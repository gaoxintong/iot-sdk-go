package device

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"iot-sdk-go/pkg/types"
	protocol "iot-sdk-go/sdk/protocol"
	request "iot-sdk-go/sdk/request"
	serializer "iot-sdk-go/sdk/serializer"
	storage "iot-sdk-go/sdk/storage"
	topics "iot-sdk-go/sdk/topics"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/imdario/mergo"
)

// Device 设备
type Device struct {
	ProductKey string
	Name       string
	Version    string
	Secret     string
	ID         int64
	Token      []byte
	Access     string
	Protocol   protocol.Protocol
	Serializer serializer.Serializer
	Topics     topics.Topics
	Storage    storage.Storage
}

// Builder 设备构造器
type Builder struct {
	ProductKey string
	Name       string
	Version    string
	Protocol   protocol.Protocol
	Serializer serializer.Serializer
	Topics     topics.Topics
	Storage    storage.Storage
}

// NewBuilder 创建设备构造器
func NewBuilder() *Builder {
	DefaultProtocol := protocol.NewMQTT()
	DefaultSerializer := serializer.NewTLV()
	DefaultTopics := topics.DefaultTopics
	DefaultStorage := &storage.LocalStorage{}
	return &Builder{
		Protocol:   DefaultProtocol,
		Serializer: DefaultSerializer,
		Topics:     DefaultTopics,
		Storage:    DefaultStorage,
	}
}

// SetProductKey 设置产品Key
func (d *Builder) SetProductKey(productKey string) *Builder {
	d.ProductKey = productKey
	return d
}

// SetDeviceName 设置设备名
func (d *Builder) SetDeviceName(name string) *Builder {
	d.Name = name
	return d
}

// SetVersion 设置版本
func (d *Builder) SetVersion(version string) *Builder {
	d.Version = version
	return d
}

// SetProtocol 设置协议
func (d *Builder) SetProtocol(protocol protocol.Protocol) *Builder {
	d.Protocol = protocol
	return d
}

// SetSerializer 设置序列化器
func (d *Builder) SetSerializer(serializer serializer.Serializer) *Builder {
	d.Serializer = serializer
	return d
}

// SetTopics 设置主题列表
func (d *Builder) SetTopics(topics topics.Topics) *Builder {
	d.Topics = topics
	return d
}

// SetStorage 设置存储
func (d *Builder) SetStorage(storage storage.Storage) *Builder {
	d.Storage = storage
	return d
}

// Build 构建设备
func (d *Builder) Build() *Device {
	return &Device{
		ProductKey: d.ProductKey,
		Name:       d.Name,
		Version:    d.Version,
		Protocol:   d.Protocol,
		Serializer: d.Serializer,
		Topics:     d.Topics,
		Storage:    d.Storage,
	}
}

// GetDeviceInfo 获取设备信息
func (d *Device) GetDeviceInfo() (*Device, error) {
	ProductKeyInter, err := d.Storage.Get("ProductKey")
	if err != nil {
		return nil, err
	}
	ProductKey, _ := types.InterfaceToString(ProductKeyInter)
	NameInter, err := d.Storage.Get("Name")
	if err != nil {
		return nil, err
	}
	Name, _ := types.InterfaceToString(NameInter)
	SecretInter, err := d.Storage.Get("Secret")
	if err != nil {
		return nil, err
	}
	Secret, _ := types.InterfaceToString(SecretInter)
	VersionInter, err := d.Storage.Get("Version")
	if err != nil {
		return nil, err
	}
	Version, _ := types.InterfaceToString(VersionInter)
	IDInter, err := d.Storage.Get("ID")
	if err != nil {
		return nil, err
	}
	IDInt, _ := types.InterfaceToInt(IDInter)
	ID := int64(IDInt)
	return &Device{
		ProductKey: ProductKey,
		Name:       Name,
		Secret:     Secret,
		Version:    Version,
		ID:         ID,
	}, nil
}

// LoadDeviceInfo 合并设备信息
func (d *Device) LoadDeviceInfo() error {
	tmp, err := d.GetDeviceInfo()
	if err != nil {
		return err
	}
	return mergo.Merge(d, tmp, mergo.WithOverride)
}

// SetDeviceInfo 设置设备信息
func (d *Device) SetDeviceInfo() error {
	storage := d.Storage
	if d.ProductKey != "" {
		if err := storage.Set("ProductKey", d.ProductKey); err != nil {
			return err
		}
	}
	if d.Name != "" {
		if err := storage.Set("Name", d.Name); err != nil {
			return err
		}
	}
	if d.Secret != "" {
		if err := storage.Set("Secret", d.Secret); err != nil {
			return err
		}
	}
	if d.Version != "" {
		if err := storage.Set("Version", d.Version); err != nil {
			return err
		}
	}
	if d.ID != 0 {
		if err := storage.Set("ID", d.ID); err != nil {
			return err
		}
	}
	if d.Token != nil {
		if err := storage.Set("Token", d.Token); err != nil {
			return err
		}
	}
	if d.Access != "" {
		if err := storage.Set("Access", d.Access); err != nil {
			return err
		}
	}
	return nil
}

// Register 注册
func (d *Device) Register() error {
	args, err := RegisterArgsFromDevice(*d)
	if err != nil {
		return err
	}
	argsStr, err := json.Marshal(args)
	if err != nil {
		return err
	}
	jsonresp, err := http.Post(d.Topics.Register, "application/json", strings.NewReader(string(argsStr)))
	if err != nil {
		return err
	}
	response := RegisterResponse{}
	body, _ := ioutil.ReadAll(jsonresp.Body)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	if err := HTTPIsOK(response); err != nil {
		return err
	}
	d.ID = response.Data.ID
	d.Secret = response.Data.Secret
	d.SetDeviceInfo()
	return nil
}

// Login 登陆
func (d *Device) Login() error {
	args, err := AuthArgsFromDevice(*d)
	if err != nil {
		return err
	}
	argsStr, err := json.Marshal(args)
	if err != nil {
		return err
	}
	jsonresp, err := http.Post(d.Topics.Login, "application/json", strings.NewReader(string(argsStr)))
	if err != nil {
		return err
	}
	response := AuthResponse{}
	body, _ := ioutil.ReadAll(jsonresp.Body)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	if err := HTTPIsOK(response); err != nil {
		return err
	}
	hexToken, err := hex.DecodeString(response.Data.AccessToken)
	if err != nil {
		return err
	}
	d.Token = hexToken
	d.Access = response.Data.AccessAddr
	d.SetDeviceInfo()
	return nil
}

// AutoLogin 自动登录
func (d *Device) AutoLogin() error {
	if d.Token == nil || d.Access == "" {
		if err := d.Register(); err != nil {
			return err
		}
	}
	return d.Login()
}

// InitProtocolClient 初始化协议客户端
func (d *Device) InitProtocolClient(opts ...interface{}) error {
	if len(opts) > 0 {
		// 用户传入配置，使用配置创建客户端
		return d.Protocol.NewClient(opts[0])
	}
	// 默认创建 MQTT 配置
	IDStr := strconv.Itoa(int(d.ID))
	TokenStr := hex.EncodeToString(d.Token) // 817aecf06c023365
	mqttOpts := map[string]interface{}{
		"Broker":    d.Access,
		"ClientID":  IDStr,
		"Username":  IDStr,
		"Password":  TokenStr,
		"KeepAlive": 30 * time.Second,
	}
	newOpts := d.Protocol.MakeOpts(mqttOpts)
	return d.Protocol.NewClient(newOpts)
}

// Publish 发布
func (d *Device) Publish(request request.Request) error {
	params := protocol.ParamsFormatter(request)
	return d.Protocol.Publish(params)
}

// Subscribe 订阅
func (d *Device) Subscribe(request request.Request) {
	params := protocol.ParamsFormatter(request)
	d.Protocol.Subscribe(params)
}

// Unsubscribe 取消订阅
func (d *Device) Unsubscribe(topics []string) {
	d.Protocol.Unsubscribe(map[string]interface{}{"topics": topics})
}

// PostProperty 发送属性
func (d *Device) PostProperty(property []interface{}) error {
	request := request.Request{}
	request.Topic = d.Topics.PostProperty
	request.Qos = 1
	request.Retained = false
	payload, err := d.Serializer.Marshal(property, serializer.PostProperty)
	if err != nil {
		return err
	}
	request.Payload = payload
	params := protocol.ParamsFormatter(request)
	return d.Protocol.Publish(params)
}

// InitOptions 初始化配置项
type InitOptions struct {
	AutoReregister           bool
	AutoRelogin              bool
	AutoReInitProtocolClient bool
}

// AutoPostProperty 自动发送属性
func (d *Device) AutoPostProperty(property []interface{}, opts ...InitOptions) error {
	finallyOpts := InitOptions{}
	if len(opts) > 0 {
		finallyOpts = opts[0]
	}
	if types.IsNil(d.Protocol.GetInstance()) {
		if err := d.AutoLogin(); err != nil {
			if finallyOpts.AutoRelogin {
				for {
					time.Sleep(5 * time.Second)
					if err := d.AutoLogin(); err == nil {
						break
					}
				}
			} else {
				return err
			}
		}
		if err := d.InitProtocolClient(); err != nil {
			if finallyOpts.AutoReInitProtocolClient {
				for {
					time.Sleep(5 * time.Second)
					if err := d.InitProtocolClient(); err == nil {
						break
					}
				}
			} else {
				return err
			}
		}
	}
	return d.PostProperty(property)
}

// OnProperty 设置属性
func (d *Device) OnProperty(callback func(property interface{})) {

}

// PostEvent 发送事件
func (d *Device) PostEvent(eventIdentifier string, property []interface{}) error {
	request := request.Request{}
	request.Topic = d.Topics.PostEvent
	request.Qos = 1
	payload, err := d.Serializer.Marshal(property, serializer.PostEvent)
	request.Payload = payload
	if err != nil {
		return err
	}
	params := protocol.ParamsFormatter(request)
	d.Protocol.Publish(params)
	return nil
}

// OnCommand 响应命令
func (d *Device) OnCommand(func(res interface{}, replyFn func(res interface{}))) {

}

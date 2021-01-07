package protocol

import (
	"errors"
	mqtt "iot-sdk-go/pkg/mqtt"
	"iot-sdk-go/pkg/types"
	"reflect"
)

// MQTT 实现
type MQTT struct {
	Client *mqtt.Client
}

// NewMQTT 创建 MQTT 对象
func NewMQTT() *MQTT {
	return &MQTT{}
}

// ParamsFormatter 参数格式化
func ParamsFormatter(s interface{}) map[string]interface{} {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	ret := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		ret[t.Field(i).Name] = v.Field(i).Interface()
	}
	return ret
}

// MakeOpts 创建配置项
func (m *MQTT) MakeOpts(params map[string]interface{}) interface{} {
	Broker, err := types.InterfaceToString(params["Broker"])
	if err != nil {
		return err
	}
	ClientID, err := types.InterfaceToString(params["ClientID"])
	if err != nil {
		return err
	}
	Username, err := types.InterfaceToString(params["Username"])
	if err != nil {
		return err
	}
	Password, err := types.InterfaceToString(params["Password"])
	if err != nil {
		return err
	}
	KeepAlive, err := types.InterfaceToDuration(params["KeepAlive"])
	if err != nil {
		return err
	}
	opts := mqtt.NewClientOptions().AddBroker("tcp://" + Broker)
	opts.SetClientID(ClientID)
	opts.SetUsername(Username)
	opts.SetPassword(Password)
	opts.SetKeepAlive(KeepAlive)
	// opts.SetDefaultPublishHandler()
	return opts
}

// NewClient 创建客户端
func (m *MQTT) NewClient(opts interface{}) error {
	typedOpts, ok := opts.(*mqtt.ClientOptions)
	if !ok {
		return errors.New("mqtt options conversion failed")
	}
	c := mqtt.NewClient(typedOpts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	m.Client = c
	return nil
}

// Publish 发布
func (m *MQTT) Publish(params map[string]interface{}) error {
	topic, err := types.InterfaceToString(params["Topic"])
	if err != nil {
		return err
	}
	qos, err := types.InterfaceToByte(params["Qos"])
	if err != nil {
		return err
	}
	retained, err := types.InterfaceToBool(params["Retained"])
	if err != nil {
		return err
	}
	payload := params["Payload"]
	if err := m.Client.Publish(topic, qos, retained, payload).Error(); err != nil {
		return err
	}
	return nil
}

// InterfaceToMqttMessageHandler 接口转函数
func InterfaceToMqttMessageHandler(v interface{}) (mqtt.MessageHandler, error) {
	switch v.(type) {
	case mqtt.MessageHandler:
		return v.(mqtt.MessageHandler), nil
	}
	return nil, errors.New("interface to Mqtt MessageHandler failed")
}

// Subscribe 订阅
func (m *MQTT) Subscribe(params map[string]interface{}) error {
	topic, err := types.InterfaceToString(params["topic"])
	if err != nil {
		return err
	}
	qos, err := types.InterfaceToByte(params["qos"])
	if err != nil {
		return err
	}
	callback, err := InterfaceToMqttMessageHandler(params["callback"])
	if err != nil {
		return err
	}
	m.Client.Subscribe(topic, qos, callback)
	return nil
}

// Unsubscribe 取消订阅
func (m *MQTT) Unsubscribe(params map[string]interface{}) error {
	topics, err := types.InterfaceToSliceString(params["topics"])
	if err != nil {
		return err
	}
	m.Client.Unsubscribe(topics...)
	return nil
}

// GetName 获取协议名
func (m *MQTT) GetName() string {
	return "mqtt"
}

// GetInstance 获取协议客户端实例
func (m *MQTT) GetInstance() interface{} {
	return m.Client
}

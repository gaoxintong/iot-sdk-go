package protocol

import (
	"iot-sdk-go/pkg/mqtt"
	"iot-sdk-go/pkg/typeconv"
	"iot-sdk-go/sdk/request"

	"github.com/pkg/errors"
)

// MQTT 实现
type MQTT struct {
	Client *mqtt.Client
}

// NewMQTT 创建 MQTT 对象
func NewMQTT() *MQTT {
	return &MQTT{}
}

// MakeOpts 创建配置项
func (m *MQTT) MakeOpts(params map[string]interface{}) (interface{}, error) {
	Broker, err := typeconv.InterfaceToString(params["Broker"])
	if err != nil {
		return nil, errors.Wrap(err, "make mqtt options failed")
	}
	ClientID, err := typeconv.InterfaceToString(params["ClientID"])
	if err != nil {
		return nil, errors.Wrap(err, "make mqtt options failed")
	}
	Username, err := typeconv.InterfaceToString(params["Username"])
	if err != nil {
		return nil, errors.Wrap(err, "make mqtt options failed")
	}
	Password, err := typeconv.InterfaceToString(params["Password"])
	if err != nil {
		return nil, errors.Wrap(err, "make mqtt options failed")
	}
	KeepAlive, err := typeconv.InterfaceToDuration(params["KeepAlive"])
	if err != nil {
		return nil, errors.Wrap(err, "make mqtt options failed")
	}
	opts := mqtt.NewClientOptions().AddBroker("tcp://" + Broker)
	opts.SetClientID(ClientID)
	opts.SetUsername(Username)
	opts.SetPassword(Password)
	opts.SetKeepAlive(KeepAlive)
	// opts.SetDefaultPublishHandler()
	return opts, nil
}

// NewClient 创建客户端
func (m *MQTT) NewClient(opts interface{}) error {
	typedOpts, ok := opts.(*mqtt.ClientOptions)
	if !ok {
		return errors.New("mqtt options conversion failed")
	}
	c := mqtt.NewClient(typedOpts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return errors.Wrap(token.Error(), "new mqtt client failed")
	}
	m.Client = c
	return nil
}

// Options 配置项
type Options struct {
	Topic    string
	Qos      byte
	Retained bool
	Payload  interface{}
	Callback func(request.Response)
}

// Publish 发布
func (m *MQTT) Publish(opts map[string]interface{}) error {
	finllyOpts, err := getOpts(opts)
	if err != nil {
		return errors.Wrap(err, "mqtt publish failed")
	}
	return m.Client.Publish(finllyOpts.Topic, finllyOpts.Qos, finllyOpts.Retained, finllyOpts.Payload).Error()
}

// InterfaceToMqttMessageHandler 接口转函数
func InterfaceToMqttMessageHandler(v interface{}) (mqtt.MessageHandler, error) {
	switch v.(type) {
	case mqtt.MessageHandler:
		return v.(mqtt.MessageHandler), nil
	}
	return nil, errors.New("interface to Mqtt MessageHandler failed")
}

// InterfaceToCallbackFn 接口转函数
func InterfaceToCallbackFn(v interface{}) (func(request.Response), error) {
	switch v.(type) {
	case func(request.Response):
		return v.(func(request.Response)), nil
	}
	return nil, errors.New("interface to callback func failed")
}

// getOpts 转换生成 MQTT 配置项
func getOpts(opts map[string]interface{}) (*Options, error) {
	topic, err := typeconv.InterfaceToString(opts["Topic"])
	if err != nil {
		return nil, err
	}
	qos, err := typeconv.InterfaceToByte(opts["Qos"])
	if err != nil {
		qos = 0
	}
	retained, err := typeconv.InterfaceToBool(opts["Retained"])
	if err != nil {
		retained = false
	}
	payload := opts["Payload"]
	callback, err := InterfaceToCallbackFn(opts["Callback"])
	if err != nil {
		callback = nil
	}
	return &Options{
		Topic:    topic,
		Qos:      qos,
		Retained: retained,
		Payload:  payload,
		Callback: callback,
	}, nil
}

// Subscribe 订阅
func (m *MQTT) Subscribe(opts map[string]interface{}) error {
	finllyOpts, err := getOpts(opts)
	if err != nil {
		return err
	}
	var cb mqtt.MessageHandler = func(c *mqtt.Client, m mqtt.Message) {
		if finllyOpts.Callback != nil {
			finllyOpts.Callback(m)
		}
	}
	return m.Client.Subscribe(finllyOpts.Topic, finllyOpts.Qos, cb).Error()
}

// Unsubscribe 取消订阅
func (m *MQTT) Unsubscribe(opts map[string]interface{}) error {
	topics, err := typeconv.InterfaceToSliceString(opts["topics"])
	if err != nil {
		return err
	}
	return m.Client.Unsubscribe(topics...).Error()
}

// GetName 获取协议名
func (m *MQTT) GetName() string {
	return "mqtt"
}

// GetInstance 获取协议客户端实例
func (m *MQTT) GetInstance() interface{} {
	return m.Client
}

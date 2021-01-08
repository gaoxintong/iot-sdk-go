package device

import (
	protocol "iot-sdk-go/sdk/protocol"
	serializer "iot-sdk-go/sdk/serializer"
	storage "iot-sdk-go/sdk/storage"
	topics "iot-sdk-go/sdk/topics"
)

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

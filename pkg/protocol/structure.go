package protocol

import (
	"iot-sdk-go/pkg/tlv"
)

type CommandEventHead struct {
	Flag        uint8
	Timestamp   uint64
	Token       [8]byte
	SubDeviceid uint16
	No          uint16
	Priority    uint16
	ParamsCount uint16
}

type Command struct {
	Head   CommandEventHead
	Params []tlv.TLV
}

type Event struct {
	Head   CommandEventHead
	Params []tlv.TLV
}

type DataHead struct {
	Flag      uint8
	Timestamp uint64
	Token     [8]byte
}

type Data struct {
	Head    DataHead
	SubData []SubData
}

type SubDataHead struct {
	SubDeviceid      uint16
	PropertyNum      uint16
	ParamsCount      uint16
	ExternalDeviceId [8]byte // 扩展设备Id为兼容网关类设备
}

type SubData struct {
	Head   SubDataHead
	Params []tlv.TLV
}

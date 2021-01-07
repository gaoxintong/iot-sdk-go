package serializer

import (
	"iot-sdk-go/pkg/tlv"
	"iot-sdk-go/pkg/types"
	"time"

	"iot-sdk-go/pkg/protocol"
)

// TLV TLV对象
type TLV struct {
	Serializer tlv.TLV
}

// NewTLV 创建TLV对象
func NewTLV() TLV {
	return TLV{}
}

// Unmarshal 反序列化
func (s TLV) Unmarshal(data []byte, t Type) (interface{}, error) {
	m := map[Type]func([]byte) (interface{}, error){
		OnCommand: UnmarshalCommand,
	}
	return m[t](data)
}

// Marshal 序列化
func (s TLV) Marshal(data []interface{}, t Type) ([]byte, error) {
	m := map[Type]func(data []interface{}) ([]byte, error){
		PostProperty: MarshalProperty,
		PostEvent:    MarshalEvent,
	}
	return m[t](data)
}

// MarshalProperty 属性序列化
func MarshalProperty(data []interface{}) ([]byte, error) {
	payloadHead := protocol.DataHead{
		Flag:      0,
		Timestamp: uint64(time.Now().Unix() * 1000),
	}
	params, err := tlv.MakeTLVs(data)
	if err != nil {
		return nil, err
	}
	// 内嵌数据
	sub := protocol.SubData{
		Head: protocol.SubDataHead{
			SubDeviceid: uint16(1),
			PropertyNum: uint16(1),
			ParamsCount: uint16(len(params)),
		},
		Params: params,
	}
	// 组装数据
	status := protocol.Data{
		Head:    payloadHead,
		SubData: []protocol.SubData{},
	}
	status.SubData = append(status.SubData, sub)
	// 转 byte
	return status.Marshal()
}

// MarshalEvent 事件序列化
func MarshalEvent(data []interface{}) ([]byte, error) {
	event := protocol.Event{}
	params, err := tlv.MakeTLVs(data)
	if err != nil {
		return nil, err
	}
	event.Params = params
	event.Head.No = 1
	event.Head.SubDeviceid = 1
	event.Head.ParamsCount = uint16(len(params))
	return event.Marshal()
}

// UnmarshalCommand 命令反序列化
func UnmarshalCommand(data []byte) (interface{}, error) {
	cmd := protocol.Command{}
	dataByte := make([]byte, len(data))
	for i, v := range data {
		v2, err := types.InterfaceToByte(v)
		if err != nil {
			return nil, err
		}
		dataByte[i] = v2
	}
	err := cmd.UnMarshal(dataByte)
	if err != nil {
		return nil, err
	}
	return cmd, nil
}

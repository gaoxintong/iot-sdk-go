package serializer

import (
	"errors"
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
func NewTLV() *TLV {
	return &TLV{}
}

// Marshal 序列化
func (t *TLV) Marshal(data interface{}) (interface{}, error) {
	v, ok := data.([]interface{})
	if ok {
		return tlv.MakeTLVs(v)
	}
	return nil, errors.New("")
}

// Unmarshal 反序列化
func (t *TLV) Unmarshal(data interface{}) (interface{}, error) {
	return nil, nil
}

// MakePropertyData 创建序列化后的属性数据
func (t *TLV) MakePropertyData(property *Property) ([]byte, error) {
	payloadHead := protocol.DataHead{
		Flag:      0,
		Timestamp: uint64(time.Now().Unix() * 1000),
	}
	params, err := t.Marshal(property.Value)
	paramsTLV, ok := params.([]tlv.TLV)
	if !ok {
		return nil, errors.New("marshal property failed")
	}
	if err != nil {
		return nil, err
	}
	// 内嵌数据
	sub := protocol.SubData{
		Head: protocol.SubDataHead{
			SubDeviceid: property.SubDeviceID,
			PropertyNum: property.PropertyID,
			ParamsCount: uint16(len(paramsTLV)),
		},
		Params: paramsTLV,
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

// MakeEventData 创建序列化后的事件数据
func (t *TLV) MakeEventData(property *Property) ([]byte, error) {
	event := protocol.Event{}
	params, err := t.Marshal(property.Value)
	if err != nil {
		return nil, err
	}
	paramsTLV, ok := params.([]tlv.TLV)
	if !ok {
		return nil, errors.New("marshal property failed")
	}
	event.Params = paramsTLV
	event.Head.No = property.PropertyID
	event.Head.SubDeviceid = property.SubDeviceID
	event.Head.ParamsCount = uint16(len(paramsTLV))
	return event.Marshal()
}

// UnmarshalCommand 命令反序列化
func (t *TLV) UnmarshalCommand(data []byte) (*Command, error) {
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
	params := map[int]interface{}{}
	for i, v := range cmd.Params {
		params[i] = v.Value
	}
	ret := &Command{
		ID:          cmd.Head.No,
		SubDeviceID: cmd.Head.SubDeviceid,
		Params:      params,
	}
	return ret, nil
}

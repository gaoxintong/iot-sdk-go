package tlv

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// 定义数据类型
const (
	TLVFLOAT64 = 1
	TLVFLOAT32 = 2
	TLVINT8    = 3
	TLVINT16   = 4
	TLVINT32   = 5
	TLVINT64   = 6
	TLVUINT8   = 7
	TLVUINT16  = 8
	TLVUINT32  = 9
	TLVUINT64  = 10
	TLVBYTES   = 11
	TLVSTRING  = 12
	TLVBOOL    = 13
)

// TLV type length value
type TLV struct {
	Tag   uint16
	Value []byte
}

func uint16ToByte(value uint16) []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, value)

	return buf.Bytes()
}

func byteToUint16(buf []byte) uint16 {
	tmpBuf := bytes.NewBuffer(buf)
	var value uint16
	binary.Read(tmpBuf, binary.BigEndian, &value)

	return value
}

// ToBinary make tlv to []byte
func (tlv *TLV) ToBinary() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, &tlv.Tag)
	binary.Write(buf, binary.BigEndian, &tlv.Value)

	return buf.Bytes()
}

// Length get tlv length
func (tlv *TLV) Length() int {
	length := int(0)
	switch tlv.Tag {
	case TLVFLOAT64:
		length = 8
	case TLVINT64:
		length = 8
	case TLVUINT64:
		length = 8
	case TLVFLOAT32:
		length = 4
	case TLVINT32:
		length = 4
	case TLVUINT32:
		length = 4
	case TLVINT16:
		length = 2
	case TLVUINT16:
		length = 2
	case TLVINT8:
		length = 1
	case TLVUINT8:
		length = 1
	case TLVBYTES:
		length = int(byteToUint16(tlv.Value[0:2]))
		length += 2
	case TLVSTRING:
		length = int(byteToUint16(tlv.Value[0:2]))
		length += 2
	default:
		length = 0
	}

	length += 2

	return length
}

// FromBinary read from binary
func (tlv *TLV) FromBinary(r io.Reader) error {
	binary.Read(r, binary.BigEndian, &tlv.Tag)
	length := uint16(0)
	switch tlv.Tag {
	case TLVFLOAT64:
		length = 8
		tlv.Value = make([]byte, length)
		binary.Read(r, binary.BigEndian, &tlv.Value)
	case TLVINT64:
		length = 8
		tlv.Value = make([]byte, length)
		binary.Read(r, binary.BigEndian, &tlv.Value)
	case TLVUINT64:
		length = 8
		tlv.Value = make([]byte, length)
		binary.Read(r, binary.BigEndian, &tlv.Value)
	case TLVFLOAT32:
		length = 4
		tlv.Value = make([]byte, length)
		binary.Read(r, binary.BigEndian, &tlv.Value)
	case TLVINT32:
		length = 4
		tlv.Value = make([]byte, length)
		binary.Read(r, binary.BigEndian, &tlv.Value)
	case TLVUINT32:
		length = 4
		tlv.Value = make([]byte, length)
		binary.Read(r, binary.BigEndian, &tlv.Value)
	case TLVINT16:
		length = 2
		tlv.Value = make([]byte, length)
		binary.Read(r, binary.BigEndian, &tlv.Value)
	case TLVUINT16:
		length = 2
		tlv.Value = make([]byte, length)
		binary.Read(r, binary.BigEndian, &tlv.Value)
	case TLVINT8:
		length = 1
		tlv.Value = make([]byte, length)
		binary.Read(r, binary.BigEndian, &tlv.Value)
	case TLVUINT8:
		length = 1
		tlv.Value = make([]byte, length)
		binary.Read(r, binary.BigEndian, &tlv.Value)
	case TLVBYTES:
		binary.Read(r, binary.BigEndian, &length)
		tlv.Value = make([]byte, length+2)
		copy(tlv.Value[0:2], uint16ToByte(length))
		binary.Read(r, binary.BigEndian, tlv.Value[2:])
	case TLVSTRING:
		binary.Read(r, binary.BigEndian, &length)
		tlv.Value = make([]byte, length+2)
		copy(tlv.Value[0:2], uint16ToByte(length))
		binary.Read(r, binary.BigEndian, tlv.Value[2:])
	default:
		return fmt.Errorf("unsuport value: %d", tlv.Tag)
	}

	return nil
}

// MakeTLV make a tlv pointer
func MakeTLV(a interface{}) (*TLV, error) {
	var tag uint16
	var length uint16
	buf := new(bytes.Buffer)
	switch a := a.(type) {
	case float64:
		tag = TLVFLOAT64
		length = 8
		binary.Write(buf, binary.BigEndian, a)
	case float32:
		tag = TLVFLOAT32
		length = 4
		binary.Write(buf, binary.BigEndian, a)
	case int8:
		tag = TLVINT8
		length = 1
		binary.Write(buf, binary.BigEndian, a)
	case int16:
		tag = TLVINT16
		length = 2
		binary.Write(buf, binary.BigEndian, a)
	case int32:
		tag = TLVINT32
		length = 4
		binary.Write(buf, binary.BigEndian, a)
	case int64:
		tag = TLVINT64
		length = 8
		binary.Write(buf, binary.BigEndian, a)
	case uint8:
		tag = TLVUINT8
		length = 1
		binary.Write(buf, binary.BigEndian, a)
	case uint16:
		tag = TLVUINT16
		length = 2
		binary.Write(buf, binary.BigEndian, a)
	case uint32:
		tag = TLVUINT32
		length = 4
		binary.Write(buf, binary.BigEndian, a)
	case uint64:
		tag = TLVUINT64
		length = 8
		binary.Write(buf, binary.BigEndian, a)
	case []byte:
		tag = TLVBYTES
		length = uint16(len(a))
		binary.Write(buf, binary.BigEndian, length)
		binary.Write(buf, binary.BigEndian, a)
	case string:
		tag = TLVSTRING
		length = uint16(len(a))
		binary.Write(buf, binary.BigEndian, length)
		binary.Write(buf, binary.BigEndian, []byte(a))
	default:
		return nil, fmt.Errorf("unsuport value: %v", a)
	}

	tlv := TLV{
		Tag:   tag,
		Value: buf.Bytes(),
	}

	if length == 0 {
		tlv.Value = []byte{}
	}

	return &tlv, nil
}

// ReadTLV read from tlv pointer
func ReadTLV(tlv *TLV) (interface{}, error) {
	tag := tlv.Tag
	length := uint16(0)
	value := tlv.Value

	buffer := bytes.NewReader(value)
	switch tag {
	case TLVFLOAT64:
		retvar := float64(0.0)
		err := binary.Read(buffer, binary.BigEndian, &retvar)
		return retvar, err
	case TLVFLOAT32:
		retvar := float32(0.0)
		err := binary.Read(buffer, binary.BigEndian, &retvar)
		return retvar, err
	case TLVINT8:
		retvar := int8(0)
		err := binary.Read(buffer, binary.BigEndian, &retvar)
		return retvar, err
	case TLVINT16:
		retvar := int16(0)
		err := binary.Read(buffer, binary.BigEndian, &retvar)
		return retvar, err
	case TLVINT32:
		retvar := int32(0)
		err := binary.Read(buffer, binary.BigEndian, &retvar)
		return retvar, err
	case TLVINT64:
		retvar := int64(0)
		err := binary.Read(buffer, binary.BigEndian, &retvar)
		return retvar, err
	case TLVUINT8:
		retvar := uint8(0)
		err := binary.Read(buffer, binary.BigEndian, &retvar)
		return retvar, err
	case TLVUINT16:
		retvar := uint16(0)
		err := binary.Read(buffer, binary.BigEndian, &retvar)
		return retvar, err
	case TLVUINT32:
		retvar := uint32(0)
		err := binary.Read(buffer, binary.BigEndian, &retvar)
		return retvar, err
	case TLVUINT64:
		retvar := uint64(0)
		err := binary.Read(buffer, binary.BigEndian, &retvar)
		return retvar, err
	case TLVBYTES:
		err := binary.Read(buffer, binary.BigEndian, &length)
		if err != nil {
			return []byte{}, err
		}
		retvar := make([]byte, length)
		err = binary.Read(buffer, binary.BigEndian, &retvar)
		return retvar, err
	case TLVSTRING:
		err := binary.Read(buffer, binary.BigEndian, &length)
		if err != nil {
			return string([]byte{}), err
		}
		retvar := make([]byte, length)
		err = binary.Read(buffer, binary.BigEndian, &retvar)
		return string(retvar), err
	default:
		return nil, errors.New("Reading TLV error ,Unkown TLV type: " + string(tag))
	}
}

// MakeTLVs ``
func MakeTLVs(a []interface{}) ([]TLV, error) {
	tlvs := []TLV{}
	for _, one := range a {
		tlv, err := MakeTLV(one)
		if err != nil {
			return nil, err
		}
		tlvs = append(tlvs, *tlv)
	}
	return tlvs, nil
}

// ReadTLVs ``
func ReadTLVs(tlvs []TLV) ([]interface{}, error) {
	values := []interface{}{}
	for _, tlv := range tlvs {
		one, err := ReadTLV(&tlv)
		if err != nil {
			return values, err
		}
		values = append(values, one)
	}
	return values, nil
}

// CastTLV cast tlv
func CastTLV(value interface{}, valueType int32) interface{} {
	switch valueType {
	case TLVFLOAT64:
		return float64(value.(float64))
	case TLVFLOAT32:
		return float32(value.(float64))
	case TLVINT8:
		return int8(value.(float64))
	case TLVINT16:
		return int16(value.(float64))
	case TLVINT32:
		return int32(value.(float64))
	case TLVINT64:
		return int64(value.(float64))
	case TLVUINT8:
		return uint8(value.(float64))
	case TLVUINT16:
		return uint16(value.(float64))
	case TLVUINT32:
		return uint32(value.(float64))
	case TLVUINT64:
		return uint64(value.(float64))
	case TLVBYTES:
		return []byte(value.(string))
	case TLVSTRING:
		return value.(string)
	default:
		return nil
	}
}

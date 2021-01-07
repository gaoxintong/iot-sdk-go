package types

import (
	"errors"
	"reflect"
	"time"
)

// InterfaceToString 接口转字符串
func InterfaceToString(v interface{}) (string, error) {
	switch v.(type) {
	case string:
		return v.(string), nil
	}
	return "", errors.New("interface to string failed")
}

// InterfaceToByte 接口转Byte
func InterfaceToByte(v interface{}) (byte, error) {
	switch v.(type) {
	case byte:
		return v.(byte), nil
	}
	return 0, errors.New("interface to byte failed")
}

// InterfaceToInt 接口转Int
func InterfaceToInt(v interface{}) (int, error) {
	switch v.(type) {
	case int:
		return v.(int), nil
	}
	return 0, errors.New("interface to int failed")
}

// InterfaceToBool 接口转Bool
func InterfaceToBool(v interface{}) (bool, error) {
	switch v.(type) {
	case bool:
		return v.(bool), nil
	}
	return false, errors.New("interface to bool failed")
}

// InterfaceToMap 接口转Map
func InterfaceToMap(v interface{}) (map[string]interface{}, error) {
	switch v.(type) {
	case map[string]interface{}:
		return v.(map[string]interface{}), nil
	}
	return nil, errors.New("interface to map failed")
}

// InterfaceToSliceString 接口转Slice
func InterfaceToSliceString(v interface{}) ([]string, error) {
	switch v.(type) {
	case []string:
		return v.([]string), nil
	}
	return nil, errors.New("interface to slice string failed")
}

// InterfaceToFunc 接口转Func
func InterfaceToFunc(v interface{}) (func(), error) {
	switch v.(type) {
	case func():
		return v.(func()), nil
	}
	return nil, errors.New("interface to bool failed")
}

// InterfaceToDuration 接口转Duration
func InterfaceToDuration(v interface{}) (time.Duration, error) {
	switch v.(type) {
	case time.Duration:
		return v.(time.Duration), nil
	}
	return 0, errors.New("interface to Duration failed")
}

// IsNil 判断空指针
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

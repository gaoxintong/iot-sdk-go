package device

import (
	"errors"
)

// RegisterArgs 设备注册参数
type RegisterArgs struct {
	ProductKey string `json:"product_key"  binding:"required"`
	DeviceCode string `json:"device_code"  binding:"required"`
	Version    string `json:"version"  binding:"required"`
}

// RegisterArgsFromDevice 从设备构建 RegisterArgs
func RegisterArgsFromDevice(device Device) (*RegisterArgs, error) {
	if device.ProductKey == "" {
		return nil, errors.New("field ProductKey cannot be empty")
	}
	if device.Name == "" {
		return nil, errors.New("field Name cannot be empty")
	}
	if device.Version == "" {
		return nil, errors.New("field Version cannot be empty")
	}
	r := &RegisterArgs{}
	r.ProductKey = device.ProductKey
	r.DeviceCode = device.Name
	r.Version = device.Version
	return r, nil
}

// RegisterResponse 注册返回数据
type RegisterResponse struct {
	Common
	Data RegisterData `json:"data"`
}

// RegisterData 注册返回数据
type RegisterData struct {
	ID         int64  `json:"device_id"`
	Secret     string `json:"device_secret"`
	Key        string `json:"device_key"`
	Identifier string `json:"device_identifier"`
}

// AuthArgs 认证参数
type AuthArgs struct {
	ID       int64  `json:"device_id" binding:"required"`
	Secret   string `json:"device_secret" binding:"required"`
	Protocol string `json:"protocol" binding:"required"`
}

// AuthArgsFromDevice 使用 Device 构建 AuthArgs
func AuthArgsFromDevice(device Device) (*AuthArgs, error) {
	if device.ID == 0 {
		return nil, errors.New("field ID cannot be empty")
	}
	if device.Secret == "" {
		return nil, errors.New("field Secret cannot be empty")
	}
	protocol := device.Protocol.GetName()
	if protocol == "" {
		return nil, errors.New("field protocol cannot be empty")
	}
	ret := &AuthArgs{}
	ret.ID = device.ID
	ret.Secret = device.Secret
	ret.Protocol = device.Protocol.GetName()
	return ret, nil
}

// Common 公用字段
type Common struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// AuthResponse 认证返回数据
type AuthResponse struct {
	Common
	Data AuthData `json:"data"`
}

// AuthData 认证返回数据
type AuthData struct {
	AccessToken string `json:"access_token"`
	AccessAddr  string `json:"access_addr"`
}

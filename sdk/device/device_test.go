package device

import (
	"fmt"
	"iot-sdk-go/sdk/topics"
	"testing"
	"time"
)

const (
	ProductKey = "491e1ba0bc0ade7bb8cdb0b14483be2b312841122ee861f8fdbf0e4a4eacff52"
	DeviceName = "qxq19900805"
	Version    = "1.0.1"
)

var device *Device

func init() {
	topics := topics.Override(topics.Topics{
		Register: "http://192.168.1.120:8088/v1/devices/registration",
		Login:    "http://192.168.1.120:8088/v1/devices/authentication",
	})
	device = NewBuilder().
		SetProductKey(ProductKey).
		SetDeviceName(DeviceName).
		SetVersion(Version).
		SetTopics(topics).
		Build()
	device.LoadDeviceInfo()
}

func TestRegister(t *testing.T) {
	err := device.Register()
	if err != nil {
		fmt.Println("err:", err.Error())
	}
}

func TestLogin(t *testing.T) {
	err := device.Login()
	if err != nil {
		fmt.Println("err:", err.Error())
	}
}

func TestPostProperty(t *testing.T) {
	// 登陆
	err := device.Login()
	if err != nil {
		t.Fatal("login error:", err.Error())
	}
	// 初始化协议客户端
	err = device.InitProtocolClient()
	if err != nil {
		t.Fatal("init client error:", err.Error())
	}
	// 发送 20 次
	count := 0
	// 发送数据
	tick := time.Tick(1 * time.Second)
	for {
		fmt.Println("post property")
		// 获取硬件数据
		var v int64 = 1
		data := []interface{}{v}
		if err := device.PostProperty(data); err != nil {
			fmt.Println("err:", err)
		}
		<-tick
		count++
		if count == 20 {
			break
		}
	}
}

func TestAutoLogin(t *testing.T) {
	err := device.AutoLogin()
	if err != nil {
		t.Fatal("login error:", err.Error())
	}
	// 发送 20 次
	count := 0
	// 发送数据
	tick := time.Tick(1 * time.Second)
	for {
		fmt.Println("post property")
		// 获取硬件数据
		var v int64 = 1
		data := []interface{}{v}
		if err := device.PostProperty(data); err != nil {
			fmt.Println("err:", err)
		}
		<-tick
		count++
		if count == 20 {
			break
		}
	}
}

func TestAutoPostProperty(t *testing.T) {
	// 发送 20 次
	count := 0
	// 发送数据
	tick := time.Tick(1 * time.Second)
	for {
		fmt.Println("post property")
		// 获取硬件数据
		var v int64 = 1
		data := []interface{}{v}
		if err := device.AutoPostProperty(data); err != nil {
			fmt.Println("err:", err)
		}
		<-tick
		count++
		if count == 20 {
			break
		}
	}
}

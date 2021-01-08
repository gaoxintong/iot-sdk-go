package device

import (
	"encoding/json"
	"fmt"
	request "iot-sdk-go/sdk/request"
	"iot-sdk-go/sdk/topics"
	"testing"
	"time"
)

const (
	ProductKey = "491e1ba0bc0ade7bb8cdb0b14483be2b312841122ee861f8fdbf0e4a4eacff52"
	DeviceName = "qxq19900805"
	Version    = "1.0.1"
)

var light *Device

func init() {
	topics := topics.Override(topics.Topics{
		Register: "http://192.168.1.101:8088/v1/devices/registration",
		Login:    "http://192.168.1.101:8088/v1/devices/authentication",
	})
	light = NewBuilder().
		SetProductKey(ProductKey).
		SetDeviceName(DeviceName).
		SetVersion(Version).
		SetTopics(topics).
		Build()
	light.LoadDeviceInfo()
}

type LightTSL struct {
	status     uint16
	brightness uint16
}

func TestRegister(t *testing.T) {
	err := light.Register()
	if err != nil {
		fmt.Println("register error:", err.Error())
	}
}

func TestLogin(t *testing.T) {
	err := light.Login()
	if err != nil {
		fmt.Println("login error:", err.Error())
	}
}

func TestPostProperty(t *testing.T) {
	// 登陆
	err := light.Login()
	if err != nil {
		t.Fatal("login error:", err.Error())
	}
	// 初始化协议客户端
	err = light.InitProtocolClient()
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
		var status uint16 = 1
		var brightness uint16 = 88
		value := []interface{}{status, brightness}
		p := Property{
			SubDeviceID: 1,
			PropertyID:  1,
			Value:       value,
		}
		if err := light.PostProperty(p); err != nil {
			fmt.Println("post property error:", err)
		}
		<-tick
		count++
		if count == 20 {
			break
		}
	}
}

func TestAutoConn(t *testing.T) {
	fmt.Println("conn12")
	err := light.AutoLogin()
	if err != nil {
		t.Fatal("login error:", err.Error())
	}
	err = light.InitProtocolClient()
	for {
	}
}

func TestAutoLogin(t *testing.T) {
	err := light.AutoLogin()
	if err != nil {
		t.Fatal("login error:", err.Error())
	}
	err = light.InitProtocolClient()
	if err != nil {
		t.Fatal("InitProtocolClient error:", err.Error())
	}
	// 发送 20 次
	count := 0
	// 发送数据
	tick := time.Tick(1 * time.Second)
	for {
		fmt.Println("post property")
		// 获取硬件数据
		var status uint16 = 1
		var brightness uint16 = 88
		value := []interface{}{status, brightness}
		p := Property{
			SubDeviceID: 1,
			PropertyID:  1,
			Value:       value,
		}
		if err := light.PostProperty(p); err != nil {
			fmt.Println("post property error:", err)
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
		var status uint16 = 1
		var brightness uint16 = 88
		value := []interface{}{status, brightness}
		p := Property{
			SubDeviceID: 1,
			PropertyID:  1,
			Value:       value,
		}
		if err := light.AutoPostProperty(p); err != nil {
			fmt.Println("Auto Post Property error:", err)
		}
		<-tick
		count++
		if count == 20 {
			break
		}
	}
}

func TestOnCommand(t *testing.T) {
	fmt.Println("test command1")
	// 自动登陆
	err := light.AutoLogin()
	if err != nil {
		t.Fatal("login error:", err.Error())
	}
	err = light.InitProtocolClient()
	if err != nil {
		t.Fatal("Init Protocol Client error:", err.Error())
	}
	// 定义亮度和状态两种命令
	adjustBrightnessCmd := Command{
		ID: 1,
		Callback: func(m map[int]interface{}) {
			subDeviceID := m[-1] // 子设备 ID
			brightness := m[0]   // 亮度
			fmt.Println("子设备 ID：", subDeviceID)
			fmt.Println("控制亮度：", brightness)
		},
	}
	switchCmd := Command{
		ID: 2,
		Callback: func(m map[int]interface{}) {
			status := m[0] // 状态
			fmt.Println("控制开关：", status)
		},
	}
	err = light.OnCommand(switchCmd, adjustBrightnessCmd)
	if err != nil {
		t.Fatal("OnCommand error:", err.Error())
	}
	for {

	}
}

func TestPublish(t *testing.T) {
	// 初始化
	light.AutoInit()
	// 创建 payload
	tsl := LightTSL{
		status:     1,
		brightness: 88,
	}
	// 转 []byte
	payload, err := json.Marshal(tsl)
	if err != nil {
		t.Fatal(err)
	}
	// 创建 request
	req := request.Request{
		Topic:   "test",
		Qos:     1,
		Payload: payload,
	}
	// 发布
	if err := light.Publish(req); err != nil {
		t.Fatal("publish error:", err)
	}
}

func TestSubscribe(t *testing.T) {
	// 初始化
	light.AutoInit()
	// 创建 request
	req := request.Request{
		Topic: "test",
		Qos:   1,
		Callback: func(resp request.Response) {
			fmt.Println(resp)
		},
	}
	// 订阅
	if err := light.Subscribe(req); err != nil {
		t.Fatal("subscribe error:", err)
	}
	// for {
	// }
	//	取消订阅
	if err := light.Unsubscribe([]string{"test"}); err != nil {
		t.Fatal("unsubscribe error:", err)
	}
}

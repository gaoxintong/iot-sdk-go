---
id: TSL-develop
title: 物模型开发
sidebar_label: 物模型开发
slug: /sdk/TSL-develop
---

物模型的开发方式让设备不用关心如何去订阅 MQTT topic，而是调用物模型相关的接口来实现属性上报、服务监听、事件上报。

## 设备属性上报

属性上报有两种方式。第一种是用户手动初始化设备的注册、登陆、创建协议客户端等流程，然后属性上报。

代码示例：

```go
// 自动初始化
err := light.AutoInit()
if err != nil {
  t.Fatal("init error:", err.Error())
}
// 发送数据
tick := time.Tick(1 * time.Second)
for {
  // 获取硬件数据
  var status uint16 = 1
  var brightness uint16 = 88
  value := []interface{}{status, brightness}
  // 设置子设备ID 和属性
  p := Property{
    SubDeviceID: 1,
    PropertyID:  1,
    Value:       value,
  }
  // 属性上报
  if err := light.PostProperty(p); err != nil {
    fmt.Println("post property error:", err)
  }
  <-tick
}
```

第二种是自动完成注册、登陆、创建协议客户端后，进行属性上报。

代码示例：

```go
tick := time.Tick(1 * time.Second)
for {
  // 获取硬件数据
  var status uint16 = 1
  var brightness uint16 = 88
  value := []interface{}{status, brightness}
  // 设置子设备ID 和属性
  p := Property{
    SubDeviceID: 1,
    PropertyID:  1,
    Value:       value,
  }
  // 自动进行属性上报
  if err := light.PostProperty(p); err != nil {
    fmt.Println("post property error:", err)
  }
  <-tick
}
```

### Property

| 属性        |          类型 | 描述      | 默认值 |
| :---------- | ------------: | :-------- | :----- |
| SubDeviceID |        uint16 | 子设备 ID | 必填   |
| PropertyID  |        uint16 | 属性 ID   | 必填   |
| Value       | []interface{} | 属性值    | 必填   |

## 属性设置

暂无

## 监听命令

可以监听一个命令或者多个命令。

```go
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
// 监听命令
err = light.OnCommand(switchCmd, adjustBrightnessCmd)
if err != nil {
  t.Fatal("OnCommand error:", err.Error())
}
```

### Command

| 属性     |                      类型 | 描述     | 默认值 |
| :------- | ------------------------: | :------- | :----- |
| ID       |                    uint16 | 命令 ID  | 必填   |
| Callback | func(map[int]interface{}) | 回调函数 | 必填   |

回调函数的参数类型是一个键值对，按照配置顺序进行排列，-1 所对应的参数是 SubDeviceID。

## 事件上报

```go
// 触发关闭事件
func closeEvent(subDeviceID int, propertyID int) {
  // 获取硬件数据
  var status uint16 = 0
  value := []interface{}{status}
  // 设置子设备ID 和属性
  p := Property{
    subDeviceID: subDeviceID,
    PropertyID:  propertyID,
    Value:       value,
  }
  // 事件上报
  if err := light.PostEvent("close", p); err != nil {
    fmt.Println("post evnet error:", err)
  }
}
```
---
id: communication-based-on-mqtt
title: 基于 MQTT 通信
sidebar_label: 基于 MQTT 通信
slug: /sdk/communication-based-on-mqtt
---

SDK 提供了与云端长链接的基础能力接口，用户可以直接使用这些接口完成自定义 Topic 相关的功能。提供的基础能力包括：发布、订阅、取消订阅。

## 发布

主要用于数据上报。

代码示例：

```go
light.AutoInit()// 初始化
// 创建 payload
tsl := LightTSL{
  status:     1,
  brightness: 88,
}
payload, err := json.Marshal(tsl)// 转 []byte
if err != nil {
  t.Fatal(err)
}
// 创建 request
req := request.Request{
  Topic:   "test",
  Payload: payload,
}
// 发布
if err := light.Publish(req); err != nil {
  t.Fatal("publish error:", err)
}
```

### Request

| 属性     |           类型 | 描述                             | 默认值 |
| :------- | -------------: | :------------------------------- | :----- |
| Topic    |         string | 主题名称。                       | 必填   |
| Qos      |           byte | 服务质量级别。                   | 0      |
| Retained |           bool | 是否保存消息。                   | false  |
| Payload  |    interface{} | 消息体，仅支持 string 或[]byte。 | 必填   |
| Callback | device.ReplyFn | 回调函数，用于订阅。             | nil    |

## 订阅

代码示例：

```go
// 初始化
light.AutoInit()
// 创建 request
req := request.Request{
  Topic: "test",
  Callback: func(resp request.Response) {
    fmt.Println(resp)
  },
}
// 订阅
if err := light.Subscribe(req); err != nil {
  t.Fatal("subscribe error:", err)
}
```

## 取消订阅

代码示例：

```go
if err := light.Unsubscribe([]string{"test"}); err != nil {
  t.Fatal("unsubscribe error:", err)
}
```

| 参数   |     类型 | 描述           | 默认值 |
| :----- | -------: | :------------- | :----- |
| topics | []string | 主题名称列表。 | 必填   |

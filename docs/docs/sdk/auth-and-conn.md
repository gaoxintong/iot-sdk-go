---
id: auth-and-conn
title: 认证与连接
sidebar_label: 认证与连接
slug: /sdk/auth-and-conn
---

## 初始化

设备的认证流程分为注册和登陆。

首先在云平台创建产品，将 ProductKey 持久化到设备本地。

设备自身具有 DeviceName 和 Version 两个属性。

通过这三个属性进行初始化设备。

接下来创建一个智能电灯。

代码示例：

```go
ProductKey := "491e1ba0bc0ade7bb8cdb0b14483be2b312841122ee861f8fdbf0e4a4eacff52"
DeviceName := "light"
Version    := "1.0.1"
topics     := topics.Override(topics.Topics{
  Register: "http://192.168.1.101:8088/v1/devices/registration",
  Login:    "http://192.168.1.101:8088/v1/devices/authentication",
})
light := New(ProductKey, DeviceName, Version, Topics(topics))
```

| 参数       |                  类型 | 描述       | 默认值                   |
| :--------- | --------------------: | :--------- | :----------------------- |
| ProductKey |                string | 产品 Key。 | 必填                     |
| DeviceName |                string | 产品名称。 | 必填                     |
| Version    |                string | 设备版本。 | 必填                     |
| Protocol   |     protocol.Protocol | 协议类型。 | protocol.MQTT            |
| Serializer | serializer.Serializer | 序列化器。 | serializer.TLV           |
| Topics     |         topics.Topics | 主题列表。 | topics.DefaultTopics     |
| Storage    |       storage.Storage | 配置存储。 | storage.LocalStorage     |
| HTTPClient |           http.Client | 配置存储。 | httpclient.DefaultClient |

## 设备注册

ProductKey、DeviceName、Version 是设备注册的三元组，通过这三个属性进行注册，如果这三项属性不正确，会注册失败。如果注册成功，会获取到 DeviceID 和 Secret，将它们挂载到 Device 实例上，并使用 Storage 进行存储。

代码示例：

```go
err := light.Register()
if err != nil {
  panic(err)
}
```

| 方法     |                                            描述 |
| :------- | ----------------------------------------------: |
| Register | 使用 ProductKey、DeviceName、Version 进行注册。 |

## 设备登陆

设备登陆需要 DeviceID 、 Secret 和 Protocol 三项参数。如果登陆成功，会获取到 Token 和 Access，将它们挂在到 Device 实例上，并使用 Storage 进行存储。

Access 是协议的地址，Token 是密码。

代码示例：

```go
err := light.Login()
if err != nil {
  panic(err)
}
```

| 方法  |       描述 |
| :---- | ---------: |
| Login | 注册接口。 |

## 自动登录

为了更方便开发，SDK 提供了自动登陆的 API。如果设备尚未注册，它会先进行注册，再进行登陆。

代码示例：

```go
err := light.AutoLogin()
if err != nil {
  panic(err)
}
```

如果设备已经注册过，那么可以从存储中加载 Token、Access 属性。

代码示例：

```go
light.LoadDeviceInfo()
```

| 方法           |                       描述 |
| :------------- | -------------------------: |
| AutoLogin      |           自动注册、登陆。 |
| LoadDeviceInfo | 从存储中加载 device 属性。 |

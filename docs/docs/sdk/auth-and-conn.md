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

device.NewBuilder 是一个 device 构造器，通过一系列配置后，调用 Build 方法创建出 device 实例。

接下来创建一个智能电灯。

代码示例：

```go
ProductKey = "491e1ba0bc0ade7bb8cdb0b14483be2b312841122ee861f8fdbf0e4a4eacff52"
DeviceName = "light"
Version    = "1.0.1"
light = NewBuilder().
  SetProductKey(ProductKey).
  SetDeviceName(DeviceName).
  SetVersion(Version).
  Build()
```

| 方法          |                                                          描述 |
| :------------ | ------------------------------------------------------------: |
| NewBuilder    | 创建 device.Builder 实例，该实例用于创建 device.Device 实例。 |
| SetProductKey |                                                设置产品 Key。 |
| SetDeviceName |                                                设置产品名称。 |
| SetVersion    |                                                设置设备版本。 |
| SetProtocol   |                                     设置协议类型，默认 MQTT。 |
| SetSerializer |                                      设置序列化器，默认 TLV。 |
| SetTopics     |                                        设置注册、登陆和主题。 |
| SetStorage    |                                      设置存储，默认本地存储。 |
| Build         |                                     构建 device.Device 实例。 |

## 设备注册

ProductKey、DeviceName、Version 是设备注册的三元组，通过这三个属性进行注册，如果这三项属性不正确，会注册失败。如果注册成功，会获取到 DeviceID 和 Secret，将它们挂在到 Device 实例上，并使用 Storage 进行存储。

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

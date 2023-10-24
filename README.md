# BAETYL GATEWAY SDK

[![Baetyl-logo](./docs/logo_with_name.png)](https://baetyl.io)

[![License](https://img.shields.io/github/license/baetyl/baetyl-gateway-sdk?color=blue)](LICENSE)
[![Stars](https://img.shields.io/github/stars/baetyl/baetyl-gateway-sdk?style=social)](Stars)

Baetyl-Gateway 是基于 [go-plugin](https://github.com/hashicorp/go-plugin) 开源框架实现的 BIE 软网关。

旨在解决工业物联网领域中面临的海量异构设备数据难以统一接入的问题，此框架致力于将各种设备使用的不同通信协议数据转化为一致的物联网标准协议，从而实现设备与物联网系统之间的互联互通，为工业生产和制造过程提供全面的数据支持。

## 架构
![Baetyl-logo](./docs/baetyl-gateway.png)

### baetyl-gateway-sdk
baetyl-gateway-sdk 负责完成具体驱动与 baetyl-gateway 主进程的基础通信的封装

对应于架构图中蓝色虚线框部分

SDK 主要提供下面三部分功能
* 驱动协议实现时可以调用对应语言的 SDK 的接口完成数据的上报同步
* 通过实现 SDK 预定义的接口，可以将驱动托管给 baetyl-gateway 主进程，由主进程负责驱动的启停等操作
* 提供对 baetyl-gateway 框架下驱动的三个配置文件 access_template.yml、models.yml、sub_devices.yml 的解析逻辑的封装

## 文件结构
```
.
├── LICENSE
├── README.md
├── demo
│   ├── java
│   ├── golang
│   ├── python
│   ├── csharp
│   └── java
├── docs
├── sdk
│   ├── README.md
│   ├── example
│   ├── java
│   ├── golang
│   ├── python
│   ├── csharp
│   └── proto
└── test
    ├── baetyl-broker
    ├── baetyl-gateway
    └── driver
        ├── custom-java
        ├── custom-golang
        ├── custom-python
        └── custom-csharp
```
./sdk
* sdk/README.md: 通用的 SDK 开发指南  
* sdk/{language}: 目录下 java、golang、python、csharp 包含具体各语言 SDK 的实现
* sdk/proto: 定义软网关宿主服务和驱动服务支持的 rpc 函数列表，具体函数定义会在下文给出介绍
* sdk/example: 给出了 baetyl-gateway 框架下最终提供给驱动的三个配置文件的示例，具体示例内容会在下文给出介绍

./demo
* demo/{language}: 目录下 java、golang、python、csharp 包含、基于对应语言 SDK 的一个自定义驱动的 Demo 的实现，Demo 实现了对模拟三个点位的采集和上报过程

./test
* test/baetyl-broker: 目录下包含一个小型的 MQTT Broker 实现的二进制程序，用于帮助驱动的开发和调试工作，具体文件见 v0.0.0 Pre Release `baetyl-broker.zip` 中
* test/baetyl-gateway: 目录下包含 baetyl 软网关的二进制程序，用于帮助驱动的开发和调试工作，具体文件见 v0.0.0 Pre Release `baetyl-gateway.zip` 中
* test/driver/custom-{language}: 提供个语言 Demo 二进制运行的配置文件，搭配上述工具可以运行软网关及开发的驱动
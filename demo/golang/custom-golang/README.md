# 自定义 Go 插件开发指南
开发驱动依赖于对应语言的 SDK，SDK 位于 baetyl/baetyl-gateway-sdk/sdk/， 开发前确保已经做好 SDK 依赖

go.mod 中

```
github.com/baetyl/baetyl-gateway-sdk/sdk/golang vx.x.x
```

一个驱动一般包含如下几个部分的功能实现
1. 驱动入口，实现 baetyl-gateway-sdk/sdk/golang/interface.go 文件中的 Driver 接口；对应示例中 driver.go
2. 驱动核心控制引擎，负责实际各子设备的连接、对各设备的采集任务、对设备的操作（召测/置数等）的封装；对应示例中 custom.go
3. 设备接入实现，根据设备接入通道信息，与设备建立连接，并提供针对设备的操作的底层实现；对应示例中 device.go
4. 采集任务的具体实现，根据采集点配置对特定点位进行采集并上报；对应示例中 work.go

## 0. 示例介绍
custom-golang 项目实现自定义驱动程序

自定义驱动模拟实现了协议采集实际设备测点的完整业务流程

### 项目构建
```shell
go mod tidy

cd demo/golang/custom-golang/cmd/custom

go build -o ../../../../../test/driver/custom-golang/custom-golang .
```

产出位于 ./test/driver/custom-golang/custom-golang

测试流程详见 [test/README.md](../../../test/README.md)	

也提供 Makefile 

```shell
# 构建当前平台的二进制，产出位于 demo/golang/custom-golang/output 下
make build 

# 构建所有平台的二进制，产出位于 demo/golang/custom-golang/output 下，平台包括
# darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 linux/arm/v7 windows/amd64
make all PLATFORMS=all

# 清理项目构建
make clean
```

## 1. 文件结构
```shell
custom-golang
├── README.md
├── cmd
│    └── custom
│        ├── Makefile
│        ├── main.go
│        └── program.yml
├── config.go
├── common.go
├── custom.go
├── device.go
├── driver.go
├── simulator.go
└── worker.go

2 directories, 11 files
```
* README.md：使用说明文件
* cmd/custom/Makefile：编译用文件，make package 可以直接生成二进制并打包入口文件生成压缩包
* cmd/custom/main.go：驱动入口文件，运行依赖 go-plugin 环境
* cmd/custom/program.yml：驱动的入口文件，格式固定（entry: xxx）其中当前路径的占位符可用 $PATH 代替
* config.go：驱动的配置数据结构定义，包含驱动名称、采集配置数据结构、接入配置数据结构
* common.go：一些公共字段和实现定义，比如实际设备点位寄存器地址信息、日志等
* custom.go：实际驱动业务逻辑控制实现，负责管理当前驱动类型设备的连接和采集任务的启动
* device.go：一个具体子设备/传感器的接入、采集和置数实现
* driver.go：实现了软网关框架定义的驱动接口，软网关grpc调用的实际实现
* simulator.go：当前自定义驱动模拟生成数据的设备模拟实现
* worker.go：根据配置要求实际的采集和上报的实现

## 2. driver.go
* driver.go 实现了 baetyl-gateway-sdk/sdk/golang/interface.go 文件中的 Driver 接口

```golang
type Driver interface {
    // GetDriverInfo 获取驱动信息
	GetDriverInfo(req *Request) (*Response, error)
    // SetConfig 配置驱动，目前只配置了驱动的配置文件路径
	SetConfig(req *Request) (*Response, error)
    // Setup 宿主进程上报接口传递，必须调用下述逻辑，其余可用户自定义
	Setup(config *BackendConfig) (*Response, error)
    // Start 驱动采集启动，用户自定义实现
	Start(req *Request) (*Response, error)
    // Restart 驱动重启，用户自定义实现
	Restart(req *Request) (*Response, error)
    // Stop 驱动停止，用户自定义实现
	Stop(req *Request) (*Response, error)

    // Get 召测，用户自定义实现
	Get(req *Request) (*Response, error)
    // Set 置数，用户自定义实现
	Set(req *Request) (*Response, error)
}
```

根据软网关实现，插件启动后，gateway/driver.go 会依次调用 Setup() SetConfig() Start() 来启动驱动  
Setup 会设置驱动名称，以及**注册上报实现**  
其中 SetConfig 配置驱动，目前只配置了驱动的配置文件路径，Request 数据示例如下：

```golang
Request {
  BrokerID: 72938,
  Req: "etc/custom",
}
```

Start 主要完成了：

1. 加载 SetConfig 中设的配置文件，配置包含3个文件，这三个文件正常由 BIE 云端框架，根据实际配置自动生成，在**部署软网关**时随服务下发
	* access_template.yml ： 存放『子设备管理』-『接入模板』中的设备点表和物模型点位映射信息
		* xxx.properties[i].visitor.custom 为设备点表信息中的采集配置，自定义协议，这里一般是一个 json 格式的字符串，比如：`{"device": "custom-dev-0","interval": 3000000000}`
		* xxx.properties[i].mapping 为物模型点位映射信息
	* models.yml : 存放『子设备管理』-『产品』中的产品测点信息
	* sub_devices.yml ： 存放『节点管理』-『子设备管理』中的驱动配置和子设备配置信息（包含具体设备的接入信息）
		* devcies[i].accessConfig.custom 为子设备的配置
		* driver 为驱动配置
2. 根据配置文件中的数据，以数组形式构建**待连接设备信息**，**待启动采集的任务信息**
	* 设备信息：一般包含设备名称、设备接入地址、接入验证的用户名密码或者证书信息等。其中设备名称来源于 sub_devices.yml 中的 devcies[i].name 字段；设备接入信息来自于 devcies[i].accessConfig 字段；
	* 采集任务信息：一般包含设备名称，采集间隔，待采集点位的配置信息（如点位ID、点位数据类型、点位名称等）。其中设备名称来源于  sub_devices.yml 中的 devcies[i].name 字段；采样间隔包含在 sub_devices.yml 中的 devcies[i].accessConfig.custom 中；待采集点位配置信息位于 access_template.yml 的 xxx.properties 中，每一个采集点的具体接入访问信息位于 access_template.yml 的 xxx.properties.Vistor.Custom 中（如果是非自定义的系统已支持的协议，可以直接在 xxx.properties.Vistor.{协议} 字段中找到相关配置）
3. 初始化自定义驱动并启动
	
## 3. custom.go
* 实现了与每个设备的连接、针对每个设备周期性采集程序的启动、操作指定设备（置数、召测）的功能

### 3.1 NewCustom
* 遍历 2 Start 中生成的**待连接设备信息**，根据每条信息建立与指定设备的连接
* 遍历 2 Start 中生成的**待启动采集的任务信息**，根据采集任务配置，构造对应的采集任务

### 3.2 Start
异步启动针对所有设备的采集任务

### 3.2 Stop
通知并等待所有采集任务停止

### 3.3 Set
设置指定设备的指定属性值  
属性值的类型根据 job 中的配置进行判断转换

### 3.4 Event
根据接收到的时间类型做响应，比如召测等

## 4. worker.go
实现周期性采集上报

### 4.1 Working
根据**待启动采集的任务信息**，进行周期性收集上报

### 4.2 Report
调用设备层暴露的获取指定点位信息的方法来进行点位采集，然后通过 gateway/driver.go Setup 中一路透传过来的上报实现来完成数据的上报

## 5. device.go
完成实际设备的连接，操作方法的封装



	
	
	
	
	
	
	
	
	
	
	
	
	
	
	


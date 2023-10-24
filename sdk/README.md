# SDK 开发指南
## 1、说明
各语言驱动开发依赖对应的 SDK 实现，SDK 位于 baetyl-gateway-sdk/sdk/， 开发前确保已经做好 SDK 依赖

* golang sdk：baetyl-gateway-sdk/sdk/golang 可直接采用源码依赖
* java sdk：baetyl-gateway-sdk/sdk/java 采用 jar 包或源码形式依赖：

```groovy
dependencies {
    // 引入本地 JAR 文件
    implementation fileTree(dir: 'libs', include: '*.jar')

    // 或引入源代码目录
    implementation project(':path/to/library')
}  
```

## 2、功能
一个 SDK 需要实现以下功能
1. 基于 baetyl-gateway-sdk/sdk/proto 下的文件生成当前语言的 grpc 操作代码
2. 定义一些后续通信 grpc 通信用的常量字段及数据结构
3. 实现 access_template.yml、models.yml、sub_devices.yml 三个文件的解析逻辑，三个文件示例：baetyl-gateway-sdk/sdk/example
4. 实现 grpc v1 标准 Health 服务
5. 实现 driver.proto 中的 Driver 服务
6. 初始化 driver.proto 中的 Report 服务的客户端
7. 启动 Health、Driver 服务

### 2.1 生成 grpc 代码
直接使用 [protoc](https://grpc.io/docs/protoc-installation/) 命令生成，或者采用语言对应的构建方式，如 java 通过 gradle 管理项目时可以通过 protobuf 插件进行 proto 对应的 grpc 文件生成

### 2.2 常量及数据结构定义
#### 2.2.1 定义 Message 数据结构

结构示例：

golang:
```go
type Message struct {
	Kind     MessageKind       `yaml:"kind" json:"kind"`
	Metadata map[string]string `yaml:"meta" json:"meta"`
	Content  []byte            `yaml:"content" json:"content"`
}
```

java:
```java
public class Message {
    private String kind;
    private Map<String, String> meta;
    private Object content;
}
```

数据结构对应的 json 示例
```json
{
	"kind": "deviceReport",
	"meta": {
		"deviceName": "custom-zx",
		"driverName": "custom-java"
	},
	"content": {
		"humidity": 68,
		"pressure": 11,
		"temperature": 49
	}
}
```

#### 2.2.2 常量
定义如下常量，java 示例

```java
    // DEFAULT_SUB_DEVICE_CONF 驱动运行所需要的三个配置文件之一
    // * 对应云端『节点管理』-『子设备管理』中的驱动配置和子设备配置信息
    // * devcies[i].accessConfig.custom为子设备的配置
    // * driver为驱动配置
    public static final String DEFAULT_SUB_DEVICE_CONF = "sub_devices.yml";
    // DEFAULT_DEVICE_MODEL_CONF 驱动运行所需要的三个配置文件之一
    // * 对应云端『子设备管理』-『产品』中的产品测点信息
    public static final String DEFAULT_DEVICE_MODEL_CONF = "models.yml";
    // DEFAULT_ACCESS_TEMPLATE_CONF 驱动运行所需要的三个配置文件之一
    // * 对应云端『子设备管理』-『接入模板』中的设备点表和物模型点位映射信息
    // * device-access-tpl.properties[i].visitor.custom为设备点表信息中的采集配置
    // * device-access-tpl.properties[i].mapping为物模型点位映射信息
    public static final String DEFAULT_ACCESS_TEMPLATE_CONF = "access_template.yml";

    public static final String TYPE_REPORT_EVENT = "report";
    public static final String KEY_DRIVER_NAME = "driverName";
    public static final String KEY_DEVICE_NAME = "deviceName";

    // MESSAGE_DEVICE_EVENT 驱动中接口收到此类型消息后，需要根据具体 event 中的内容进行相关的业务处理
    public static final String MESSAGE_DEVICE_EVENT = "deviceEvent";
    // MESSAGE_DEVICE_REPORT 采集上报。测点获取完成后进行上报的消息类型
    public static final String MESSAGE_DEVICE_REPORT = "deviceReport";
    public static final String MESSAGE_DEVICE_DESIRE = "deviceDesire";
    // MESSAGE_DEVICE_DELTA 置数消息
    public static final String MESSAGE_DEVICE_DELTA = "deviceDelta";
    // MESSAGE_DEVICE_PROPERTY_GET 驱动中接口收到此类型消息后，需要进行一次测点采集
    public static final String MESSAGE_DEVICE_PROPERTY_GET = "thing.property.get";
    public static final String MESSAGE_DEVICE_EVENT_REPORT = "thing.event.post";
    // MESSAGE_DEVICE_LIFECYCLE_REPORT 状态上报
    public static final String MESSAGE_DEVICE_LIFECYCLE_REPORT = "thing.lifecycle.post";
```

### 2.3 解析三个配置文件
文件格式为 yaml

示例 baetyl-gateway-sdk/sdk/example

数据结构可先参考 java/golang 的实现

### 2.4 【重要】实现 grpc v1 版本的 Health
基于 grpc 官方定义的 [health.proto](https://github.com/grpc/grpc/blob/master/src/proto/grpc/health/v1/health.proto) 初始化一个 health server

```protobuf
message HealthCheckRequest {
  string service = 1;
}

message HealthCheckResponse {
  enum ServingStatus {
    UNKNOWN = 0;
    SERVING = 1;
    NOT_SERVING = 2;
    SERVICE_UNKNOWN = 3;  // Used only by the Watch method.
  }
  ServingStatus status = 1;
}

service Health {
  rpc Check(HealthCheckRequest) returns (HealthCheckResponse);
  rpc Watch(HealthCheckRequest) returns (stream HealthCheckResponse);
}
```

在 check 函数加入如下逻辑：**当请求数据中的 service 字段值为 plugin 时，返回 status 为 enum 中的 SERVING 作为响应**

这用于服务端 go-plugin 通过访问 check 确定驱动插件已启动

### 2.5 实现 driver.proto 中的 Driver 服务
server driver proto :
```protobuf
message RequestArgs
{
  uint32 brokerid = 1; // setup 接口使用，接收的数据为：宿主启动的 report grpc server 所在的端口号 
  string request  = 2; 
}

message ResponseResult
{
  string data     = 1;
}

service Driver
{
  // 宿主（client） --> 驱动（server）
  rpc GetDriverInfo( RequestArgs ) returns ( ResponseResult );
  rpc SetConfig( RequestArgs ) returns ( ResponseResult );

  rpc Setup( RequestArgs ) returns ( ResponseResult );
  rpc Start( RequestArgs ) returns ( ResponseResult );
  rpc Restart( RequestArgs ) returns ( ResponseResult );
  rpc Stop( RequestArgs ) returns ( ResponseResult );

  rpc Get ( RequestArgs ) returns ( ResponseResult );
  rpc Set ( RequestArgs ) returns ( ResponseResult );
}
```

第一步：先定义对应语言的接口封装 IDriver，这个接口后续需要由具体的驱动插件实现

一般的，setup 接口的入参一般定义为 RequestArgs.request 和 基于 RequestArgs.brokerid 构建的访问 grpc report 服务的客户端的封装

示例：

* [java driver 接口](java/driversdk/src/main/java/com/baidu/bce/sdk/plugin/IDriver.java)
* [golang driver 接口](golang/plugin/interface.go)

第二步：实现 driver.proto 中的 Driver 服务，如实现类为 GrpcDriverImpl，类中包含 IDriver 类型的对象

### 2.6 初始化 driver.proto 中的 Report 服务的客户端
在实现 setup 函数逻辑时，需要根据传入的 RequestArgs 中的 brokerid，即 宿主启动的 report grpc server 所在的端口号，来初始化一个 grpc report 服务的客户端

### 2.7 启动 Health、Driver 服务
将 2.4 2.5 中初始化的 Health、Driver 服务注册到 grpc 并启动服务

**【重要】服务启动后，需要调用当前语言的标准控制台输出函数，严格按如下格式打印，其中 ${port} 部分为服务启动的端口**

```
1|1|tcp|127.0.0.1:${port}|grpc
```

各语言示例
```java
System.out.println("1|1|tcp|127.0.0.1:" + port + "|grpc");
```

```go
fmt.Printf("1|1|tcp|127.0.0.1:%d|grpc", port);
```

```python
print("1|1|tcp|127.0.0.1:xxxx|grpc")
```

## 3、日志
在驱动程序中，打印日志一定要输出到**标准控制台的 err 流中**

只有这样日志才会被宿主程序捕获到并显示出来，否则无法在宿主程序中查看插件日志

示例：java 打印日志，应使用如下方式

```java
System.err.println("xxxx");
```


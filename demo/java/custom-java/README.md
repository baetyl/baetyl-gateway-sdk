## 1.业务流程

![driver](../../../docs/drvier.png)

custom-java 项目实现自定义驱动程序

自定义驱动模拟实现了协议采集实际设备测点的完整业务流程

上图各模块对应 com.baidu.bce.driver 包下同名的类实现的功能

* Driver：实现了 SDK 中的 IDriver 接口，提供了配置加载，驱动的生命周期管理及数据上报功能功能
* Custom：驱动主程序，根据加载的配置，初始化连接各个设备的客户端，并启动等同于设备数的工作线程进行周期性采集上报
* Device：实现了与具体设备连接通信的客户端，在本示例中是初始化一个模拟器实例并建立与模拟器的连接
* Worker：根据 Custom 启动具体 Worker 时传入的设备信息和采集配置信息，周期性的通过 Device 客户端对目标设备点位进行采集并上报
* Simulator：模拟器程序实现，本示例中用于模拟一个包含三个每 10s 变化一次的点位信息的设备，实际实现驱动协议时不包含此部分，由 Device 客户端直接对接真实设备

## 2.项目构建
执行 shadowJar 构建产出（此处不要直接使用 gradle build 构建，会缺少依赖项）

生成 jar 包并复制到 baetyl-gateway/etc/custom-java 目录下: gradle 执行 copyShadowJar 任务

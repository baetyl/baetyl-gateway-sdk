package com.baidu.bce.sdk.plugin;


public interface IDriver {
    // 获取驱动信息
    String getDriverInfo(String data);

    // 配置驱动，目前只配置了驱动的配置文件路径
    String setConfig(String data);

    // 宿主进程上报接口传递，必须调用下述逻辑，其余可用户自定义
    String setup(String driver, IReport report);

    // 驱动采集启动，用户自定义实现
    String start(String data);

    // 驱动重启，用户自定义实现
    String restart(String data);

    // 驱动停止，用户自定义实现
    String stop(String data);

    // 召测，用户自定义实现
    String get(String data);

    // 置数，用户自定义实现
    String set(String data);
}

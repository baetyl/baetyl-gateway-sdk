package com.baidu.bce.sdk.plugin;

/**
 * 提供了宿主-->驱动的生命周期管理方法的定义
 */
public interface IDriver {
    /**
     * 获取驱动信息
     * @param data 根据业务需要传所需数据，一般为空
     * @return 返回当前生效的配置信息，一般使用 json 序列化
     */
    String getDriverInfo(String data);

    /**
     * 配置驱动，目前只配置了驱动的配置文件路径
     * @param data 配置文件的路径
     * @return 操作执行结果
     */
    String setConfig(String data);

    /**
     * 宿主进程上报接口传递，必须调用下述逻辑，其余可用户自定义
     * @param driver 驱动的名称
     * @param report 数据上报的实现
     * @return 操作执行结果
     */
    String setup(String driver, IReport report);

    /**
     * 驱动采集启动，用户自定义实现
     * @param data 根据业务需要传所需数据，一般为空
     * @return 操作执行结果
     */
    String start(String data);

    /**
     * 驱动重启，用户自定义实现
     * @param data 根据业务需要传所需数据，一般为空
     * @return 操作执行结果
     */
    String restart(String data);

    /**
     * 驱动停止，用户自定义实现
     * @param data 根据业务需要传所需数据，一般为空
     * @return 操作执行结果
     */
    String stop(String data);

    /**
     * 召测，用户自定义实现
     * @param data 待采集点位信息
     * @return 采集的数据
     */
    String get(String data);

    /**
     * 置数，用户自定义实现
     * @param data 置数操作的数据信息
     * @return 驱动执行结果的响应
     */
    String set(String data);
}

using System;
using Proto;

namespace DriverSDK.Plugin
{
	public interface IDriver
	{
        // 获取驱动信息
        string GetDriverInfo(string data);

        // 配置驱动，目前只配置了驱动的配置文件路径
        string SetConfig(string data);

        // 宿主进程上报接口传递，必须调用下述逻辑，其余可用户自定义
        string Setup(string driver, IReport report);

        // 驱动采集启动，用户自定义实现
        string Start(string data);

        // 驱动重启，用户自定义实现
        string Restart(string data);

        // 驱动停止，用户自定义实现
        string Stop(string data);

        // 召测，用户自定义实现
        string Get(string data);

        // 置数，用户自定义实现
        string Set(string data);
    }
}


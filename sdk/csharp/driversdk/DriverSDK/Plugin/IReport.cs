using System;

namespace DriverSDK.Plugin
{
	public interface IReport
	{
        // 驱动 --> 宿主
        string Post(string data);
        string State(string data);
    }
}


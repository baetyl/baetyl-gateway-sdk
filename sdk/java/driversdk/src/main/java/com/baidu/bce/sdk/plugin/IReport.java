package com.baidu.bce.sdk.plugin;

/**
 * 定义了驱动 --> 宿主上报消息和状态的接口
 */
public interface IReport {
    /**
     * 驱动向宿主上报信息的接口
     * @param data 自定义格式的字符串数据，包含点位采集信息
     * @return 宿主的响应
     */
    String post(String data);

    String state(String data);
}

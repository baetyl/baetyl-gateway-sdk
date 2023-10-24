package com.baidu.bce.sdk.plugin;

public interface IReport {
    // 驱动 --> 宿主
    String post(String data);

    String state(String data);
}

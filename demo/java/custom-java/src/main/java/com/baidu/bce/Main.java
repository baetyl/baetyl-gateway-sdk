package com.baidu.bce;

import com.baidu.bce.driver.Driver;
import com.baidu.bce.sdk.plugin.Serve;

public class Main {
    public static void main(String[] args) {
        new Serve(new Driver()).start();
    }
}
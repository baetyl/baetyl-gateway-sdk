package com.baidu.bce.driver;

import com.baidu.bce.models.DeviceConfig;
import com.baidu.bce.sdk.context.models.device.DeviceInfo;

public class Device {
    private DeviceConfig cfg;
    private DeviceInfo info;
    private Simulator cli;

    public Device(DeviceConfig cfg, DeviceInfo info) {
        this.cfg = cfg;
        this.info = info;
        this.cli = new Simulator(cfg.getDevice());
    }

    public Object get(int index) throws Exception {
        return this.cli.get(index);
    }

    public void set(int index, float val) {
        this.cli.set(index, val);
    }

    public DeviceInfo getInfo() {
        return info;
    }

    public void setInfo(DeviceInfo info) {
        this.info = info;
    }
}

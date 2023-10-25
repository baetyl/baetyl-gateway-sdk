package com.baidu.bce.driver;

import com.alibaba.fastjson2.JSONObject;
import com.baidu.bce.models.DeviceConfig;
import com.baidu.bce.models.DriverConfig;
import com.baidu.bce.models.Job;
import com.baidu.bce.models.Property;
import com.baidu.bce.sdk.context.Context;
import com.baidu.bce.sdk.context.models.Message;
import com.baidu.bce.sdk.context.models.access.AccessConfig;
import com.baidu.bce.sdk.context.models.access.AccessTemplate;
import com.baidu.bce.sdk.context.models.device.DeviceInfo;
import com.baidu.bce.sdk.context.models.device.DeviceProperty;
import com.baidu.bce.sdk.context.models.device.Event;
import com.baidu.bce.sdk.plugin.IDriver;
import com.baidu.bce.sdk.plugin.IReport;
import com.baidu.bce.utils.L;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Driver implements IDriver {
    private String driverName;
    private String configPath;
    private DriverConfig config;
    private IReport report;
    private Custom custom;

    @Override
    public String getDriverInfo(String data) {
        Map<String, Object> info = new HashMap<>();
        info.put("name", this.driverName);
        info.put("configPath", this.configPath);
        info.put("config", this.config);

        return JSONObject.toJSONString(info);
    }

    @Override
    public String setConfig(String data) {
        L.debug("Driver setConfig " + data);
        this.configPath = data;
        return String.format("plugin %s setConfig success", this.driverName);
    }

    @Override
    public String setup(String driver, IReport report) {
        L.debug("Driver setup " + driver + " IReport " + report);
        this.driverName = driver;
        this.report = report;
        return String.format("plugin %s setup success", driver);
    }

    @Override
    public String start(String data) {
        L.debug("Driver start " + data);
        Context ctx = new Context();
        ctx.loadYamlConfig(configPath, driverName);
        this.config = loadConfig(ctx, driverName);
        this.custom = new Custom(ctx, config, report);
        this.custom.start();
        return String.format("plugin %s start success", this.driverName);
    }

    @Override
    public String restart(String data) {
        L.debug("Driver restart " + data);
        this.custom.restart();
        return String.format("plugin %s restart success", this.driverName);
    }

    @Override
    public String stop(String data) {
        L.debug("Driver stop " + data);
        this.custom.stop();
        return String.format("plugin %s stop success", this.driverName);
    }

    @Override
    public String get(String data) {
        L.debug("Driver get " + data);

        Message msg = JSONObject.parseObject(data, Message.class);
        String driverName = msg.getMeta().get(Context.KEY_DRIVER_NAME);
        String deviceName = msg.getMeta().get(Context.KEY_DEVICE_NAME);

        if ("".equals(driverName) || "".equals(deviceName)) {
            return String.format("plugin %s get failed", this.driverName);
        }

        DeviceInfo info = this.custom.getCtx().getDevice(driverName, deviceName);

        switch (msg.getKind()) {
            case Context.MESSAGE_DEVICE_EVENT -> {
                Event event = JSONObject.parseObject(msg.getContent().toString(), Event.class);
                this.custom.event(info, event);
            }
            case Context.MESSAGE_DEVICE_PROPERTY_GET -> this.custom.propertyGet(info);
            default -> L.debug("driver get unsupported message type");
        }
        return String.format("plugin %s get success", this.driverName);
    }

    @Override
    public String set(String data) {
        L.debug("Driver set " + data);

        Message msg = JSONObject.parseObject(data, Message.class);
        String driverName = msg.getMeta().get(Context.KEY_DRIVER_NAME);
        String deviceName = msg.getMeta().get(Context.KEY_DEVICE_NAME);

        if ("".equals(driverName) || "".equals(deviceName)) {
            return String.format("plugin %s get failed", this.driverName);
        }

        DeviceInfo info = this.custom.getCtx().getDevice(driverName, deviceName);

        Map<?, ?> origData = JSONObject.parseObject(msg.getContent().toString(), Map.class);
        Map<String, Object> props = new HashMap<>();
        for (Map.Entry<?, ?> entry : origData.entrySet()) {
            props.put(entry.getKey().toString(), entry.getValue());
        }
        this.custom.set(info, props);
        return String.format("plugin %s set success", this.driverName);
    }

    private DriverConfig loadConfig(Context dm, String driverName) {
        List<DeviceConfig> devices = new ArrayList<>();
        List<Job> jobs = new ArrayList<>();

        for (DeviceInfo info : dm.getAllDevices(driverName)) {
            AccessConfig accessConfig = info.getAccessConfig();
            if (accessConfig == null) {
                continue;
            }
            DeviceConfig device = new DeviceConfig();
            device.setDevice(info.getName());
            devices.add(device);

            List<Property> jobProps = new ArrayList<>();

            AccessTemplate tpl = dm.getAccessTemplate(driverName, info.getAccessTemplate());
            if (tpl != null && tpl.getProperties() != null) {
                for (DeviceProperty prop : tpl.getProperties()) {
                    String visitor = prop.getVisitor().getCustom();
                    if (!"".equals(visitor)) {
                        Property jobProp = JSONObject.parseObject(visitor, Property.class);
                        jobProps.add(jobProp);
                    }
                }
            }

            Job job = JSONObject.parseObject(accessConfig.getCustom(), Job.class);
            job.setDevice(info.getName());
            job.setProperties(jobProps);

            jobs.add(job);
        }
        return new DriverConfig(driverName, devices, jobs);
    }
}

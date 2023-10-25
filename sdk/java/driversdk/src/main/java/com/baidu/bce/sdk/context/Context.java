package com.baidu.bce.sdk.context;

import com.baidu.bce.sdk.context.models.access.AccessTemplate;
import com.baidu.bce.sdk.context.models.device.DeviceInfo;
import com.baidu.bce.sdk.context.models.device.DeviceProperty;
import com.baidu.bce.sdk.context.models.yaml.AccessTemplateConstructor;
import com.baidu.bce.sdk.context.models.yaml.ListDevicePropertyConstructor;
import com.baidu.bce.sdk.context.models.yaml.SubDeviceYaml;
import lombok.Data;
import org.yaml.snakeyaml.DumperOptions;
import org.yaml.snakeyaml.LoaderOptions;
import org.yaml.snakeyaml.Yaml;
import org.yaml.snakeyaml.representer.Representer;

import java.io.FileInputStream;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Data
public class Context {
    public static final String DEFAULT_SUB_DEVICE_CONF = "sub_devices.yml";
    public static final String DEFAULT_DEVICE_MODEL_CONF = "models.yml";
    public static final String DEFAULT_ACCESS_TEMPLATE_CONF = "access_template.yml";

    public static final String TYPE_REPORT_EVENT = "report";
    public static final String KEY_DRIVER_NAME = "driverName";
    public static final String KEY_DEVICE_NAME = "deviceName";

    public static final String MESSAGE_DEVICE_EVENT = "deviceEvent";
    public static final String MESSAGE_DEVICE_REPORT = "deviceReport";
    public static final String MESSAGE_DEVICE_DESIRE = "deviceDesire";
    public static final String MESSAGE_DEVICE_DELTA = "deviceDelta";
    public static final String MESSAGE_DEVICE_PROPERTY_GET = "thing.property.get";
    public static final String MESSAGE_DEVICE_EVENT_REPORT = "thing.event.post";
    public static final String MESSAGE_DEVICE_LIFECYCLE_REPORT = "thing.lifecycle.post";

    private Map<String, Map<String, List<DeviceProperty>>> modelYamls; // deviceModels
    private Map<String, Map<String, AccessTemplate>> accessTemplateYamls; // accessTemplates
    private Map<String, SubDeviceYaml> subDeviceYamls;
    private Map<String, Map<String, DeviceInfo>> deviceInfos; // devices
    private Map<String, String> deviceDriverMap;

    private String driverConfigPathBase;

    public Context() {
        this.modelYamls = new HashMap<>();
        this.accessTemplateYamls = new HashMap<>();
        this.subDeviceYamls = new HashMap<>();
        this.deviceDriverMap = new HashMap<>();
        this.deviceInfos = new HashMap<>();
    }

    public void loadYamlConfig(String path, String driverName) {
        this.driverConfigPathBase = path;
        try {
            Yaml yl;
            FileInputStream fis;
            Representer representer = new Representer(new DumperOptions());
            representer.getPropertyUtils().setSkipMissingProperties(true);

            fis = new FileInputStream(Paths.get(path, DEFAULT_ACCESS_TEMPLATE_CONF).toString());
            yl = new Yaml(new AccessTemplateConstructor(new LoaderOptions()), representer);
            Map<String, AccessTemplate> aty = yl.load(fis);
            for (String name : aty.keySet()) {
                aty.get(name).setName(name);
            }
            this.accessTemplateYamls.put(driverName, aty);

            yl = new Yaml(new ListDevicePropertyConstructor(new LoaderOptions()), representer);
            fis = new FileInputStream(Paths.get(path, DEFAULT_DEVICE_MODEL_CONF).toString());
            Map<String, List<DeviceProperty>> my = yl.load(fis);
            this.modelYamls.put(driverName, my);

            yl = new Yaml(representer);
            fis = new FileInputStream(Paths.get(path, DEFAULT_SUB_DEVICE_CONF).toString());
            SubDeviceYaml sdy = yl.loadAs(fis, SubDeviceYaml.class);
            this.subDeviceYamls.put(driverName, sdy);

            Map<String, DeviceInfo> info = new HashMap<>();
            for (DeviceInfo item : sdy.getDevices()) {
                info.put(item.getName(), item);
                this.deviceDriverMap.put(item.getName(), driverName);
            }
            this.deviceInfos.put(driverName, info);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    public List<DeviceInfo> getAllDevices(String driverName) {
        List<DeviceInfo> res = new ArrayList<>();
        for (String name : this.deviceInfos.get(driverName).keySet()) {
            res.add(this.deviceInfos.get(driverName).get(name));
        }
        return res;
    }

    public DeviceInfo getDevice(String driverName, String deviceName) {
        if (this.deviceInfos.containsKey(driverName) && this.deviceInfos.get(driverName).containsKey(deviceName)) {
            return this.deviceInfos.get(driverName).get(deviceName);
        }
        return null;
    }

    public String getDriverNameByDevice(String deviceName) {
        return this.deviceDriverMap.get(deviceName);
    }

    public List<DeviceProperty> getDeviceModel(String driverName, String deviceModelName) {
        if (this.modelYamls.containsKey(driverName)) {
            return this.modelYamls.get(driverName).get(deviceModelName);
        }
        return null;
    }

    public Map<String, List<DeviceProperty>> getAllDeviceModels(String driverName) {
        return this.modelYamls.get(driverName);
    }

    public AccessTemplate getAccessTemplate(String driverName, String accessTemplateName) {
        if (this.accessTemplateYamls.containsKey(driverName)) {
            return this.accessTemplateYamls.get(driverName).get(accessTemplateName);
        }
        return null;
    }

    public Map<String, AccessTemplate> getAllAccessTemplates(String driverName) {
        return this.accessTemplateYamls.get(driverName);
    }
}

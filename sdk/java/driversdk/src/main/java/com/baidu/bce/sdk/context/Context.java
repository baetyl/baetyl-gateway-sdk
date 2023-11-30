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

/**
 * Context 是一个包含 Baetyl 设备管理上下文的类
 * 提供了加载 access_template.yml、models.yml、sub_devices.yml 三个配置文件并解析到对应数据结构类的实现
 */
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

    public static final int DeviceOffline = 0;
    public static final int DeviceOnline = 1;

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

    /**
     * 提供了加载三个 yaml 配置文件的具体实现，下面两个参数一般在实现 IDriver 接口时，由 setup 方法传入
     * @param path 配置文件所在目录
     * @param driverName 驱动的名称
     */
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

    /**
     * 获取指定名称的驱动下所有设备信息，设备信息来自对 sub_devices.yml 文件的解析
     * @param driverName 驱动名称
     * @return 设备信息列表
     */
    public List<DeviceInfo> getAllDevices(String driverName) {
        List<DeviceInfo> res = new ArrayList<>();
        for (String name : this.deviceInfos.get(driverName).keySet()) {
            res.add(this.deviceInfos.get(driverName).get(name));
        }
        return res;
    }

    /**
     * 获取指定驱动下配置的指定名称设备的信息
     * @param driverName 驱动名称
     * @param deviceName 待获取设备的名称
     * @return 设备信息
     */
    public DeviceInfo getDevice(String driverName, String deviceName) {
        if (this.deviceInfos.containsKey(driverName) && this.deviceInfos.get(driverName).containsKey(deviceName)) {
            return this.deviceInfos.get(driverName).get(deviceName);
        }
        return null;
    }

    /**
     * 根据设备名称查找对应的驱动名称
     * @param deviceName 设备名称
     * @return 驱动名称
     */
    public String getDriverNameByDevice(String deviceName) {
        return this.deviceDriverMap.get(deviceName);
    }

    /**
     * 根据驱动名称和设备产品(设备模型)名称获取设备测点列表，设备属性信息来自对 models.yml 的解析
     * @param driverName 驱动名称
     * @param deviceModelName 设备产品名称（设备模型名称/物模型名称）
     * @return 测点列表
     */
    public List<DeviceProperty> getDeviceModel(String driverName, String deviceModelName) {
        if (this.modelYamls.containsKey(driverName)) {
            return this.modelYamls.get(driverName).get(deviceModelName);
        }
        return null;
    }

    /**
     * 根据驱动名称获取所有的设备测点列表
     * @param driverName 驱动名称
     * @return 驱动名称为key，测点列表为val的map
     */
    public Map<String, List<DeviceProperty>> getAllDeviceModels(String driverName) {
        return this.modelYamls.get(driverName);
    }

    /**
     * 根据驱动名称和接入模板名称获取接入模板数据，接入模板数据来自对 access_template.yml 的解析
     * @param driverName 驱动名称
     * @param accessTemplateName 接入模板名称
     * @return 接入模板信息
     */
    public AccessTemplate getAccessTemplate(String driverName, String accessTemplateName) {
        if (this.accessTemplateYamls.containsKey(driverName)) {
            return this.accessTemplateYamls.get(driverName).get(accessTemplateName);
        }
        return null;
    }

    /**
     * 根据驱动名称获取所有的接入目标信息
     * @param driverName 驱动名称
     * @return 驱动名称为key，接入模板为val的map
     */
    public Map<String, AccessTemplate> getAllAccessTemplates(String driverName) {
        return this.accessTemplateYamls.get(driverName);
    }
}

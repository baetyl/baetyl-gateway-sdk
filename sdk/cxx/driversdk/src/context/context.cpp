#include "context.h"
#include "utils/path.h"
#include <iostream>
#include <fstream>

namespace DRIVERSDK {

    const std::string Context::DEFAULT_SUB_DEVICE_CONF = "sub_devices.yml";
    const std::string Context::DEFAULT_DEVICE_MODEL_CONF = "models.yml";
    const std::string Context::DEFAULT_ACCESS_TEMPLATE_CONF = "access_template.yml";

    const std::string Context::TYPE_REPORT_EVENT = "report";
    const std::string Context::KEY_DRIVER_NAME = "driverName";
    const std::string Context::KEY_DEVICE_NAME = "deviceName";

    const std::string Context::MESSAGE_DEVICE_EVENT = "deviceEvent";
    const std::string Context::MESSAGE_DEVICE_REPORT = "deviceReport";
    const std::string Context::MESSAGE_DEVICE_DESIRE = "deviceDesire";
    const std::string Context::MESSAGE_DEVICE_DELTA = "deviceDelta";
    const std::string Context::MESSAGE_DEVICE_PROPERTY_GET = "thing.property.get";
    const std::string Context::MESSAGE_DEVICE_EVENT_REPORT = "thing.event.post";
    const std::string Context::MESSAGE_DEVICE_LIFECYCLE_REPORT = "thing.lifecycle.post";


    void Context::loadYamlConfig(const std::string &path, const std::string &driverName) {
        driverConfigPathBase = path;
        try {
            modelYamls[driverName] = parseModelYaml(
                    PathUtils::joinPaths(driverConfigPathBase, DEFAULT_DEVICE_MODEL_CONF));
            accessTemplateYamls[driverName] = parseAccessTemplateYaml(
                    PathUtils::joinPaths(driverConfigPathBase, DEFAULT_ACCESS_TEMPLATE_CONF));
            subDeviceYamls[driverName] = parseSubDeviceYaml(
                    PathUtils::joinPaths(driverConfigPathBase, DEFAULT_SUB_DEVICE_CONF));

            std::map<std::string, DeviceInfo> info;
            for (const DeviceInfo &item: subDeviceYamls[driverName].getDevices()) {
                info[item.getName()] = item;
                deviceDriverMap[item.getName()] = driverName;
            }
            deviceInfos[driverName] = info;
        } catch (const std::exception &e) {
            std::cerr << "Error: " << e.what() << std::endl;
            return;
        }
    }

    std::vector<DeviceInfo> Context::getAllDevices(const std::string &driverName) {
        std::vector<DeviceInfo> res;
        for (const auto &item: deviceInfos[driverName]) {
            res.push_back(item.second);
        }
        return res;
    }

    DeviceInfo Context::getDevice(const std::string &driverName, const std::string &deviceName) {
        if (deviceInfos.find(driverName) != deviceInfos.end() &&
            deviceInfos[driverName].find(deviceName) != deviceInfos[driverName].end()) {
            return deviceInfos[driverName][deviceName];
        }
        return {};
    }

    std::string Context::getDriverNameByDevice(const std::string &deviceName) {
        return deviceDriverMap[deviceName];
    }

    std::vector<DeviceProperty>
    Context::getDeviceModel(const std::string &driverName, const std::string &deviceModelName) {
        if (modelYamls.find(driverName) != modelYamls.end() &&
            modelYamls[driverName].find(deviceModelName) != modelYamls[driverName].end()) {
            return modelYamls[driverName][deviceModelName];
        }
        return {};
    }

    std::map<std::string, std::vector<DeviceProperty>> Context::getAllDeviceModels(const std::string &driverName) {
        return modelYamls[driverName];
    }

    AccessTemplate Context::getAccessTemplate(const std::string &driverName, const std::string &accessTemplateName) {
        if (accessTemplateYamls.find(driverName) != accessTemplateYamls.end() &&
            accessTemplateYamls[driverName].find(accessTemplateName) != accessTemplateYamls[driverName].end()) {
            return accessTemplateYamls[driverName][accessTemplateName];
        }
        return {};
    }

    std::map<std::string, AccessTemplate> Context::getAllAccessTemplates(const std::string &driverName) {
        if (accessTemplateYamls.find(driverName) != accessTemplateYamls.end()) {
            return accessTemplateYamls[driverName];
        }
        return {};
    }

    SubDeviceYaml Context::parseSubDeviceYaml(const std::string &filePath) {
        std::ifstream file(filePath);
        YAML::Node yamlNode = YAML::Load(file);

        SubDeviceYaml sdy;
        for (const auto &item: yamlNode) {
            std::string key = item.first.as<std::string>();
            if (key == "devices") {
                const YAML::Node &devicesNode = item.second;
                std::vector<DeviceInfo> devices;
                for (const auto &dev: devicesNode) {
                    DeviceInfo info;

                    if (dev["name"]) {
                        info.setName(dev["name"].as<std::string>());
                    }
                    if (dev["version"]) {
                        info.setVersion(dev["version"].as<std::string>());
                    }
                    if (dev["deviceModel"]) {
                        info.setDeviceModel(dev["deviceModel"].as<std::string>());
                    }
                    if (dev["accessTemplate"]) {
                        info.setAccessTemplate(dev["accessTemplate"].as<std::string>());
                    }
                    if (dev["accessConfig"]) {
                        const YAML::Node &acNode = dev["accessConfig"];
                        AccessConfig ac;
                        if (acNode["custom"]) {
                            ac.setCustom(acNode["custom"].as<std::string>());
                        }
                        info.setAccessConfig(ac);
                    }
                    devices.push_back(info);
                }
                sdy.setDevices(devices);
            } else if (key == "driver") {
                sdy.setDriver(item.second.as<std::string>());
            }
        }
        return sdy;
    }

    std::map<std::string, AccessTemplate> Context::parseAccessTemplateYaml(const std::string &filePath) {
        std::ifstream file(filePath);
        YAML::Node yamlNode = YAML::Load(file);

        // 正常 access_template.yaml 配置文件根节点下只有一对 kv，所以最后直接 return
        for (const auto &item: yamlNode) {
            std::string tplName = item.first.as<std::string>();
            const YAML::Node &tplNode = item.second;

            std::map<std::string, AccessTemplate> tplMap;

            AccessTemplate tpl;
            for (const auto &tplkv: tplNode) {
                std::string key = tplkv.first.as<std::string>();

                if (key == "name") {
                    tpl.setName(tplkv.second.as<std::string>());
                } else if (key == "version") {
                    tpl.setVersion(tplkv.second.as<std::string>());
                } else if (key == "properties") {
                    const YAML::Node &properties = tplkv.second;
                    std::vector<DeviceProperty> props;
                    for (const auto &property: properties) {
                        props.push_back(parseDeviceProperty(property));
                    }
                    tpl.setProperties(props);
                } else if (key == "mappings") {
                    const YAML::Node &mappings = tplkv.second;
                    std::vector<ModelMapping> mps;
                    for (const auto &mp: mappings) {
                        ModelMapping m;
                        if (mp["attribute"]) {
                            m.setAttribute(mp["attribute"].as<std::string>());
                        }
                        if (mp["expression"]) {
                            m.setExpression(mp["expression"].as<std::string>());
                        }
                        if (mp["type"]) {
                            m.setType(mp["type"].as<std::string>());
                        }
                        if (mp["precision"]) {
                            m.setPrecision(mp["precision"].as<int>());
                        }
                        if (mp["deviation"]) {
                            m.setDeviation(mp["deviation"].as<double>());
                        }
                        if (mp["silentWin"]) {
                            m.setSilentWindow(mp["silentWin"].as<int>());
                        }
                        mps.push_back(m);
                    }
                    tpl.setMappings(mps);
                }

            }
            tplMap[tplName] = tpl;
            return tplMap;
        }

        return {};
    }

    std::map<std::string, std::vector<DeviceProperty>> Context::parseModelYaml(const std::string &filePath) {
        std::ifstream file(filePath);
        YAML::Node yamlNode = YAML::Load(file);

        // 正常 models.yaml 配置文件根节点下只有一对 kv，所以最后直接 return
        for (const auto &item: yamlNode) {
            std::string deviceName = item.first.as<std::string>();
            const YAML::Node &properties = item.second;

            std::map<std::string, std::vector<DeviceProperty>> propertyMap;

            std::vector<DeviceProperty> dpVector;
            for (const auto &property: properties) {
                dpVector.push_back(parseDeviceProperty(property));
            }

            propertyMap[deviceName] = dpVector;
            return propertyMap;
        }

        return {};
    }

    DeviceProperty Context::parseDeviceProperty(const YAML::detail::iterator_value &property) {
        DeviceProperty deviceProperty;

        if (property["name"]) {
            deviceProperty.setName(property["name"].as<std::string>());
        }
        if (property["id"]) {
            deviceProperty.setId(property["id"].as<std::string>());
        }
        if (property["type"]) {
            deviceProperty.setType(property["type"].as<std::string>());
        }
        if (property["mode"]) {
            deviceProperty.setMode(property["mode"].as<std::string>());
        }
        if (property["unit"]) {
            deviceProperty.setUnit(property["unit"].as<std::string>());
        }
        if (property["format"]) {
            deviceProperty.setName(property["format"].as<std::string>());
        }
        if (property["arrayType"]) {
            const YAML::Node &arrayTypeNode = property["arrayType"];
            ArrayType arrayType;
            if (arrayTypeNode["type"]) {
                arrayType.setType(arrayTypeNode["type"].as<std::string>());
            }
            if (arrayTypeNode["min"]) {
                arrayType.setMin(arrayTypeNode["min"].as<int>());
            }
            if (arrayTypeNode["max"]) {
                arrayType.setMax(arrayTypeNode["max"].as<int>());
            }
            if (arrayTypeNode["format"]) {
                arrayType.setFormat(arrayTypeNode["format"].as<std::string>());
            }
            deviceProperty.setArrayType(arrayType);
        }
        if (property["objectType"]) {
            const YAML::Node &objectTypeNode = property["objectType"];
            std::map<std::string, ObjectType> objectTypeMap;
            for (const auto &entry: objectTypeNode) {
                std::string objectPropertyName = entry.first.as<std::string>();
                const YAML::Node &propertyNode = entry.second;
                ObjectType objectType;
                if (propertyNode["displayName"]) {
                    objectType.setDisplayName(propertyNode["displayName"].as<std::string>());
                }
                if (propertyNode["type"]) {
                    objectType.setType(propertyNode["type"].as<std::string>());
                }
                if (propertyNode["format"]) {
                    objectType.setFormat(propertyNode["format"].as<std::string>());
                }
                objectTypeMap[objectPropertyName] = objectType;
            }
            deviceProperty.setObjectType(objectTypeMap);
        }
        if (property["objectRequired"]) {
            const YAML::Node &objectRequiredNode = property["objectRequired"];
            std::vector<std::string> objectRequiredList;
            for (const auto &entry: objectRequiredNode) {
                objectRequiredList.push_back(entry.as<std::string>());
            }
            deviceProperty.setObjectRequired(objectRequiredList);
        }
        if (property["visitor"]) {
            const YAML::Node &visitorNode = property["visitor"];
            if (visitorNode["custom"]) {
                std::string customValue = visitorNode["custom"].as<std::string>();
                PropertyVisitor visitor(customValue);
                deviceProperty.setVisitor(visitor);
            }
        }
        return deviceProperty;
    }
}

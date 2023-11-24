#pragma once

#include <map>
#include <string>
#include <vector>
#include <yaml-cpp/yaml.h>
#include "models/deviceproperty.h"
#include "models/accesstemplate.h"
#include "models/subdeviceyaml.h"

namespace DRIVERSDK {

    class Context {
    public:
        static const std::string DEFAULT_SUB_DEVICE_CONF;
        static const std::string DEFAULT_DEVICE_MODEL_CONF;
        static const std::string DEFAULT_ACCESS_TEMPLATE_CONF;

        static const std::string TYPE_REPORT_EVENT;
        static const std::string KEY_DRIVER_NAME;
        static const std::string KEY_DEVICE_NAME;

        static const std::string MESSAGE_DEVICE_EVENT;
        static const std::string MESSAGE_DEVICE_REPORT;
        static const std::string MESSAGE_DEVICE_DESIRE;
        static const std::string MESSAGE_DEVICE_DELTA;
        static const std::string MESSAGE_DEVICE_PROPERTY_GET;
        static const std::string MESSAGE_DEVICE_EVENT_REPORT;
        static const std::string MESSAGE_DEVICE_LIFECYCLE_REPORT;

        void loadYamlConfig(const std::string& path, const std::string& driverName);
        std::vector<DeviceInfo> getAllDevices(const std::string& driverName);
        DeviceInfo getDevice(const std::string& driverName, const std::string& deviceName);
        std::string getDriverNameByDevice(const std::string& deviceName);
        std::vector<DeviceProperty> getDeviceModel(const std::string& driverName, const std::string& deviceModelName);
        std::map<std::string, std::vector<DeviceProperty>> getAllDeviceModels(const std::string& driverName);
        AccessTemplate getAccessTemplate(const std::string& driverName, const std::string& accessTemplateName);
        std::map<std::string, AccessTemplate> getAllAccessTemplates(const std::string& driverName);

    private:
        std::map<std::string, std::map<std::string, std::vector<DeviceProperty>>> modelYamls;
        std::map<std::string, std::map<std::string, AccessTemplate>> accessTemplateYamls;
        std::map<std::string, SubDeviceYaml> subDeviceYamls;
        std::map<std::string, std::map<std::string, DeviceInfo>> deviceInfos;
        std::map<std::string, std::string> deviceDriverMap;

        std::string driverConfigPathBase;

        SubDeviceYaml parseSubDeviceYaml(const std::string& filePath);
        std::map<std::string, std::vector<DeviceProperty>> parseModelYaml(const std::string& filePath);
        std::map<std::string, AccessTemplate> parseAccessTemplateYaml(const std::string& filePath);
        DeviceProperty parseDeviceProperty(const YAML::detail::iterator_value& property);
    };

} // namespace DRIVERSDK


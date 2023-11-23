#include "deviceinfo.h"

namespace DRIVERSDK {
// Default constructor
    DeviceInfo::DeviceInfo() {}

// Parameterized constructor
    DeviceInfo::DeviceInfo(const std::string &deviceName, const std::string &deviceVersion,
                           const std::string &model, const std::string &templateName,
                           const AccessConfig &config)
            : name(deviceName), version(deviceVersion), deviceModel(model),
              accessTemplate(templateName), accessConfig(config) {}

// Getter and setter methods
    const std::string &DeviceInfo::getName() const {
        return name;
    }

    void DeviceInfo::setName(const std::string &deviceName) {
        name = deviceName;
    }

    const std::string &DeviceInfo::getVersion() const {
        return version;
    }

    void DeviceInfo::setVersion(const std::string &deviceVersion) {
        version = deviceVersion;
    }

    const std::string &DeviceInfo::getDeviceModel() const {
        return deviceModel;
    }

    void DeviceInfo::setDeviceModel(const std::string &model) {
        deviceModel = model;
    }

    const std::string &DeviceInfo::getAccessTemplate() const {
        return accessTemplate;
    }

    void DeviceInfo::setAccessTemplate(const std::string &templateName) {
        accessTemplate = templateName;
    }

    const AccessConfig &DeviceInfo::getAccessConfig() const {
        return accessConfig;
    }

    void DeviceInfo::setAccessConfig(const AccessConfig &config) {
        accessConfig = config;
    }
}

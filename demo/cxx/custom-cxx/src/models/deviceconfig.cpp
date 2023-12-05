#include "deviceconfig.h"

DeviceConfig::DeviceConfig() = default;

DeviceConfig::DeviceConfig(const std::string& dev)
        : device(dev) {}

const std::string& DeviceConfig::getDevice() const {
    return device;
}

void DeviceConfig::setDevice(const std::string& dev) {
    this->device = dev;
}

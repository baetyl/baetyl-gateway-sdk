#include "device.h"

// Constructor
Device::Device(const DeviceConfig &c, DRIVERSDK::DeviceInfo& inf)
        : cfg(c), info(inf), cli(c.getDevice()) {
}

// Member function implementations...
float Device::get(int index) {
    return cli.get(index);
}

void Device::set(int index, float val) {
    cli.set(index, val);
}

const DRIVERSDK::DeviceInfo &Device::getInfo() const {
    return info;
}

void Device::setInfo(const DRIVERSDK::DeviceInfo &inf) {
    this->info = inf;
}

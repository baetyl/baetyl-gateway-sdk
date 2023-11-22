#include "accessconfig.h"

namespace DRIVERSDK {

    AccessConfig::AccessConfig() {}

    AccessConfig::AccessConfig(const std::string &customValue) : custom(customValue) {}

    const std::string &AccessConfig::getCustom() const {
        return custom;
    }

    void AccessConfig::setCustom(const std::string &customValue) {
        custom = customValue;
    }

}
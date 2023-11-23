#ifndef DRIVERSDK_ACCESSCONFIG_H
#define DRIVERSDK_ACCESSCONFIG_H

#include <string>
namespace DRIVERSDK {
    class AccessConfig {
    private:
        std::string custom;

    public:
        AccessConfig();

        AccessConfig(const std::string &customValue);

        const std::string &getCustom() const;

        void setCustom(const std::string &customValue);
    };
}
#endif // DRIVERSDK_ACCESSCONFIG_H

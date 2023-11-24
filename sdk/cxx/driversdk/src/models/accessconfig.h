#pragma once

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


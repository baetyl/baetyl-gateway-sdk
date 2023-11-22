// deviceinfo.h

#ifndef DRIVERSDK_DEVICEINFO_H
#define DRIVERSDK_DEVICEINFO_H

#include <string>

#include "accessconfig.h" // Include the AccessConfig header

namespace DRIVERSDK {
    class DeviceInfo {
    private:
        std::string name;
        std::string version;
        std::string deviceModel;
        std::string accessTemplate;
        AccessConfig accessConfig;

    public:
        // Default constructor
        DeviceInfo();

        // Parameterized constructor
        DeviceInfo(const std::string &deviceName, const std::string &deviceVersion,
                   const std::string &model, const std::string &templateName,
                   const AccessConfig &config);

        // Getter and setter methods
        const std::string &getName() const;

        void setName(const std::string &deviceName);

        const std::string &getVersion() const;

        void setVersion(const std::string &deviceVersion);

        const std::string &getDeviceModel() const;

        void setDeviceModel(const std::string &model);

        const std::string &getAccessTemplate() const;

        void setAccessTemplate(const std::string &templateName);

        const AccessConfig &getAccessConfig() const;

        void setAccessConfig(const AccessConfig &config);
    };
}

#endif // DRIVERSDK_DEVICEINFO_H

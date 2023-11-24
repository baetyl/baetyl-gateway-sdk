#ifndef DRIVERSDK_ACCESSTEMPLATE_H
#define DRIVERSDK_ACCESSTEMPLATE_H

#include <string>
#include <vector>
#include "deviceproperty.h" // Include the DeviceProperty header
#include "modelmapping.h"    // Include the ModelMapping header

namespace DRIVERSDK {

    class AccessTemplate {
    private:
        std::string name;
        std::string version;
        std::vector<DeviceProperty> properties;
        std::vector<ModelMapping> mappings;

    public:
        // Default constructor
        AccessTemplate();

        // Parameterized constructor
        AccessTemplate(const std::string& templateName, const std::string& templateVersion,
                       const std::vector<DeviceProperty>& templateProperties, const std::vector<ModelMapping>& templateMappings);

        // Getter and setter methods
        const std::string& getName() const;
        void setName(const std::string& templateName);

        const std::string& getVersion() const;
        void setVersion(const std::string& templateVersion);

        const std::vector<DeviceProperty>& getProperties() const;
        void setProperties(const std::vector<DeviceProperty>& templateProperties);

        const std::vector<ModelMapping>& getMappings() const;
        void setMappings(const std::vector<ModelMapping>& templateMappings);
    };

} // namespace DRIVERSDK

#endif // DRIVERSDK_ACCESSTEMPLATE_H

#include "accesstemplate.h"

namespace DRIVERSDK {

// Default constructor
    AccessTemplate::AccessTemplate() {}

// Parameterized constructor
    AccessTemplate::AccessTemplate(const std::string& templateName, const std::string& templateVersion,
                                   const std::vector<DeviceProperty>& templateProperties, const std::vector<ModelMapping>& templateMappings)
            : name(templateName), version(templateVersion), properties(templateProperties), mappings(templateMappings) {}

// Getter and setter methods
    const std::string& AccessTemplate::getName() const {
        return name;
    }

    void AccessTemplate::setName(const std::string& templateName) {
        name = templateName;
    }

    const std::string& AccessTemplate::getVersion() const {
        return version;
    }

    void AccessTemplate::setVersion(const std::string& templateVersion) {
        version = templateVersion;
    }

    const std::vector<DeviceProperty>& AccessTemplate::getProperties() const {
        return properties;
    }

    void AccessTemplate::setProperties(const std::vector<DeviceProperty>& templateProperties) {
        properties = templateProperties;
    }

    const std::vector<ModelMapping>& AccessTemplate::getMappings() const {
        return mappings;
    }

    void AccessTemplate::setMappings(const std::vector<ModelMapping>& templateMappings) {
        mappings = templateMappings;
    }

} // namespace DRIVERSDK

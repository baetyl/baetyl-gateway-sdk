#include "accesstemplate.h"

namespace DRIVERSDK {

// Default constructor
    AccessTemplate::AccessTemplate() {}

// Parameterized constructor
    AccessTemplate::AccessTemplate(const std::string& templateName, const std::string& templateVersion,
                                   const std::list<DeviceProperty>& templateProperties, const std::list<ModelMapping>& templateMappings)
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

    const std::list<DeviceProperty>& AccessTemplate::getProperties() const {
        return properties;
    }

    void AccessTemplate::setProperties(const std::list<DeviceProperty>& templateProperties) {
        properties = templateProperties;
    }

    const std::list<ModelMapping>& AccessTemplate::getMappings() const {
        return mappings;
    }

    void AccessTemplate::setMappings(const std::list<ModelMapping>& templateMappings) {
        mappings = templateMappings;
    }

} // namespace DRIVERSDK

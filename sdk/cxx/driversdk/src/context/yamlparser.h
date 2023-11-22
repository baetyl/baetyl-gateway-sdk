#ifndef DRIVERSDK_YAMLPARSER_H
#define DRIVERSDK_YAMLPARSER_H

#include <map>
#include <string>
#include <vector>
#include "models/deviceproperty.h"

namespace DRIVERSDK {

    class YamlParser {
    public:
        static std::map<std::string, std::map<std::string, std::vector<DeviceProperty>>> parseYaml(const std::string& filePath);
    };

} // namespace DRIVERSDK

#endif // DRIVERSDK_YAMLPARSER_H

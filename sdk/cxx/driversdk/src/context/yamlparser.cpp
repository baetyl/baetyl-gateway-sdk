#include "yamlparser.h"
#include <yaml-cpp/yaml.h>
#include <fstream>

namespace DRIVERSDK {

    std::map<std::string, std::map<std::string, std::vector<DeviceProperty>>>
    YamlParser::parseYaml(const std::string &filePath) {
        std::ifstream file(filePath);
        YAML::Node yamlNode = YAML::Load(file);

        std::map<std::string, std::map<std::string, std::vector<DeviceProperty>>> modelYamls;

        for (const auto &item: yamlNode) {
            std::string deviceName = item.first.as<std::string>();
            const YAML::Node &properties = item.second;

            std::map<std::string, std::vector<DeviceProperty>> propertyMap;

            for (const auto &property: properties) {
                DeviceProperty deviceProperty;

                if (property["name"]) {
                    deviceProperty.setName(property["name"].as<std::string>());
                }
                if (property["id"]) {
                    deviceProperty.setId(property["id"].as<std::string>());
                }
                if (property["type"]) {
                    deviceProperty.setType(property["type"].as<std::string>());
                }
                if (property["mode"]) {
                    deviceProperty.setMode(property["mode"].as<std::string>());
                }
                if (property["unit"]) {
                    deviceProperty.setUnit(property["unit"].as<std::string>());
                }
                if (property["format"]) {
                    deviceProperty.setName(property["format"].as<std::string>());
                }
                if (property["arrayType"]) {
                    const YAML::Node &arrayTypeNode = property["arrayType"];
                    ArrayType arrayType;
                    if (arrayTypeNode["type"]) {
                        arrayType.setType(arrayTypeNode["type"].as<std::string>());
                    }
                    if (arrayTypeNode["min"]) {
                        arrayType.setMin(arrayTypeNode["min"].as<int>());
                    }
                    if (arrayTypeNode["max"]) {
                        arrayType.setMax(arrayTypeNode["max"].as<int>());
                    }
                    if (arrayTypeNode["format"]) {
                        arrayType.setFormat(arrayTypeNode["format"].as<std::string>());
                    }
                    deviceProperty.setArrayType(arrayType);
                }

                if (property["objectType"]) {
                    const YAML::Node &objectTypeNode = property["objectType"];
                    std::map<std::string, ObjectType> objectTypeMap;
                    for (const auto &entry: objectTypeNode) {
                        std::string objectPropertyName = entry.first.as<std::string>();
                        const YAML::Node &propertyNode = entry.second;
                        ObjectType objectType;
                        if (propertyNode["displayName"]) {
                            objectType.setDisplayName(propertyNode["displayName"].as<std::string>());
                        }
                        if (propertyNode["type"]) {
                            objectType.setType(propertyNode["type"].as<std::string>());
                        }
                        if (propertyNode["format"]) {
                            objectType.setFormat(propertyNode["format"].as<std::string>());
                        }
                        objectTypeMap[objectPropertyName] = objectType;
                    }
                    deviceProperty.setObjectType(objectTypeMap);
                }

                if (property["objectRequired"]) {
                    const YAML::Node &objectRequiredNode = property["objectRequired"];
                    std::list<std::string> objectRequiredList;
                    for (const auto &entry: objectRequiredNode) {
                        objectRequiredList.push_back(entry.as<std::string>());
                    }
                    deviceProperty.setObjectRequired(objectRequiredList);
                }

                // Handle PropertyVisitor if it exists
                if (property["visitor"]) {
                    const YAML::Node &visitorNode = property["visitor"];
                    if (visitorNode["custom"]) {
                        std::string customValue = visitorNode["custom"].as<std::string>();
                        PropertyVisitor visitor(customValue);
                        deviceProperty.setVisitor(visitor);
                    }
                }

                propertyMap[deviceProperty.getType()].push_back(deviceProperty);
            }

            modelYamls[deviceName] = propertyMap;
        }

        return modelYamls;
    }

}

#include "deviceproperty.h"

namespace DRIVERSDK {

// Default constructor
    DeviceProperty::DeviceProperty(){}

// Getter and setter methods for each member variable
    const std::string& DeviceProperty::getName() const {
        return name;
    }

    void DeviceProperty::setName(const std::string& propName) {
        name = propName;
    }

    const std::string& DeviceProperty::getId() const {
        return id;
    }

    void DeviceProperty::setId(const std::string& propId) {
        id = propId;
    }

    const std::string& DeviceProperty::getType() const {
        return type;
    }

    void DeviceProperty::setType(const std::string& propType) {
        type = propType;
    }

    const std::string& DeviceProperty::getMode() const {
        return mode;
    }

    void DeviceProperty::setMode(const std::string& propMode) {
        mode = propMode;
    }

    const std::string& DeviceProperty::getUnit() const {
        return unit;
    }

    void DeviceProperty::setUnit(const std::string& propUnit) {
        unit = propUnit;
    }

    const PropertyVisitor& DeviceProperty::getVisitor() const {
        return visitor;
    }

    void DeviceProperty::setVisitor(const PropertyVisitor& propVisitor) {
        visitor = propVisitor;
    }

    const std::string& DeviceProperty::getFormat() const {
        return format;
    }

    void DeviceProperty::setFormat(const std::string& propFormat) {
        format = propFormat;
    }

    const EnumType& DeviceProperty::getEnumType() const {
        return enumType;
    }

    void DeviceProperty::setEnumType(const EnumType& propEnumType) {
        enumType = propEnumType;
    }

    const ArrayType& DeviceProperty::getArrayType() const {
        return arrayType;
    }

    void DeviceProperty::setArrayType(const ArrayType& propArrayType) {
        arrayType = propArrayType;
    }

    const std::map<std::string, ObjectType>& DeviceProperty::getObjectType() const {
        return objectType;
    }

    void DeviceProperty::setObjectType(const std::map<std::string, ObjectType>& propObjectType) {
        objectType = propObjectType;
    }

    const std::list<std::string>& DeviceProperty::getObjectRequired() const {
        return objectRequired;
    }

    void DeviceProperty::setObjectRequired(const std::list<std::string>& propObjectRequired) {
        objectRequired = propObjectRequired;
    }

} // namespace DRIVERSDK

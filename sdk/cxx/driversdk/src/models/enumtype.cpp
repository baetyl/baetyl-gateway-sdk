#include "enumtype.h"

namespace DRIVERSDK {

// Default constructor
    EnumType::EnumType() {}

// Parameterized constructor
    EnumType::EnumType(const std::string& enumType, const std::vector<EnumValue>& enumValues)
            : type(enumType), values(enumValues) {}

// Getter and setter methods
    const std::string& EnumType::getType() const {
        return type;
    }

    void EnumType::setType(const std::string& enumType) {
        type = enumType;
    }

    const std::vector<EnumValue>& EnumType::getValues() const {
        return values;
    }

    void EnumType::setValues(const std::vector<EnumValue>& enumValues) {
        values = enumValues;
    }

} // namespace DRIVERSDK

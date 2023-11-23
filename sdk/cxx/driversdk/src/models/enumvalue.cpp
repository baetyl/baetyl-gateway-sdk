// enumvalue.cpp

#include "enumvalue.h"

namespace DRIVERSDK {

// Default constructor
    EnumValue::EnumValue() {}

// Parameterized constructor
    EnumValue::EnumValue(const std::string& enumName, const std::string& enumValue, const std::string& display)
            : name(enumName), value(enumValue), displayName(display) {}

// Getter and setter methods
    const std::string& EnumValue::getName() const {
        return name;
    }

    void EnumValue::setName(const std::string& enumName) {
        name = enumName;
    }

    const std::string& EnumValue::getValue() const {
        return value;
    }

    void EnumValue::setValue(const std::string& enumValue) {
        value = enumValue;
    }

    const std::string& EnumValue::getDisplayName() const {
        return displayName;
    }

    void EnumValue::setDisplayName(const std::string& display) {
        displayName = display;
    }

} // namespace DRIVERSDK

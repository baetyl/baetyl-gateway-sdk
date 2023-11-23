#include "objecttype.h"

namespace DRIVERSDK {

// Default constructor
    ObjectType::ObjectType() {}

// Parameterized constructor
    ObjectType::ObjectType(const std::string& display, const std::string& objType, const std::string& objFormat)
            : displayName(display), type(objType), format(objFormat) {}

// Getter and setter methods
    const std::string& ObjectType::getDisplayName() const {
        return displayName;
    }

    void ObjectType::setDisplayName(const std::string& display) {
        displayName = display;
    }

    const std::string& ObjectType::getType() const {
        return type;
    }

    void ObjectType::setType(const std::string& objType) {
        type = objType;
    }

    const std::string& ObjectType::getFormat() const {
        return format;
    }

    void ObjectType::setFormat(const std::string& objFormat) {
        format = objFormat;
    }

} // namespace DRIVERSDK

#include "propertyvisitor.h"

namespace DRIVERSDK {

// Default constructor
    PropertyVisitor::PropertyVisitor() {}

// Parameterized constructor
    PropertyVisitor::PropertyVisitor(const std::string& customValue) : custom(customValue) {}

// Getter and setter methods
    const std::string& PropertyVisitor::getCustom() const {
        return custom;
    }

    void PropertyVisitor::setCustom(const std::string& customValue) {
        custom = customValue;
    }

} // namespace DRIVERSDK

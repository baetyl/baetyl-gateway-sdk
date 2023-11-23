#ifndef DRIVERSDK_PROPERTYVISITOR_H
#define DRIVERSDK_PROPERTYVISITOR_H

#include <string>

namespace DRIVERSDK {

    class PropertyVisitor {
    private:
        std::string custom;

    public:
        // Default constructor
        PropertyVisitor();

        // Parameterized constructor
        PropertyVisitor(const std::string& customValue);

        // Getter and setter methods
        const std::string& getCustom() const;
        void setCustom(const std::string& customValue);
    };

} // namespace DRIVERSDK

#endif // DRIVERSDK_PROPERTYVISITOR_H

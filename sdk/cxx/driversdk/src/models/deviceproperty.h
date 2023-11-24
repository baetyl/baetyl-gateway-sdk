#ifndef DRIVERSDK_DEVICEPROPERTY_H
#define DRIVERSDK_DEVICEPROPERTY_H

#include <string>
#include <map>
#include <list>
#include "propertyvisitor.h" // Include the PropertyVisitor header
#include "enumtype.h"        // Include the EnumType header
#include "arraytype.h"       // Include the ArrayType header
#include "objecttype.h"      // Include the ObjectType header

namespace DRIVERSDK {

    class DeviceProperty {
    private:
        std::string name;
        std::string id;
        std::string type;
        std::string mode;
        std::string unit;
        PropertyVisitor visitor;
        std::string format;
        EnumType enumType;
        ArrayType arrayType;
        std::map<std::string, ObjectType> objectType;
        std::vector<std::string> objectRequired;

    public:
        // Default constructor
        DeviceProperty();

        // Getter and setter methods for each member variable
        const std::string& getName() const;
        void setName(const std::string& propName);

        const std::string& getId() const;
        void setId(const std::string& propId);

        const std::string& getType() const;
        void setType(const std::string& propType);

        const std::string& getMode() const;
        void setMode(const std::string& propMode);

        const std::string& getUnit() const;
        void setUnit(const std::string& propUnit);

        const PropertyVisitor& getVisitor() const;
        void setVisitor(const PropertyVisitor& propVisitor);

        const std::string& getFormat() const;
        void setFormat(const std::string& propFormat);

        const EnumType& getEnumType() const;
        void setEnumType(const EnumType& propEnumType);

        const ArrayType& getArrayType() const;
        void setArrayType(const ArrayType& propArrayType);

        const std::map<std::string, ObjectType>& getObjectType() const;
        void setObjectType(const std::map<std::string, ObjectType>& propObjectType);

        const std::vector<std::string>& getObjectRequired() const;
        void setObjectRequired(const std::vector<std::string>& propObjectRequired);

    };

} // namespace DRIVERSDK

#endif // DRIVERSDK_DEVICEPROPERTY_H

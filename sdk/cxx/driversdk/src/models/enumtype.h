// enumtype.h

#ifndef DRIVERSDK_ENUMTYPE_H
#define DRIVERSDK_ENUMTYPE_H

#include <string>
#include <vector>
#include "enumvalue.h" // Include the EnumValue header

namespace DRIVERSDK {

    class EnumType {
    private:
        std::string type;
        std::vector<EnumValue> values;

    public:
        // Default constructor
        EnumType();

        // Parameterized constructor
        EnumType(const std::string& enumType, const std::vector<EnumValue>& enumValues);

        // Getter and setter methods
        const std::string& getType() const;
        void setType(const std::string& enumType);

        const std::vector<EnumValue>& getValues() const;
        void setValues(const std::vector<EnumValue>& enumValues);
    };

} // namespace DRIVERSDK

#endif // DRIVERSDK_ENUMTYPE_H

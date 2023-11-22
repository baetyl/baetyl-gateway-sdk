#include "arraytype.h"

namespace DRIVERSDK {
    ArrayType::ArrayType() : min(0), max(0) {}

    ArrayType::ArrayType(const std::string &arrayType, int arrayMin, int arrayMax, const std::string &arrayFormat)
            : type(arrayType), min(arrayMin), max(arrayMax), format(arrayFormat) {}

    const std::string &ArrayType::getType() const {
        return type;
    }

    void ArrayType::setType(const std::string &arrayType) {
        type = arrayType;
    }

    int ArrayType::getMin() const {
        return min;
    }

    void ArrayType::setMin(int arrayMin) {
        min = arrayMin;
    }

    int ArrayType::getMax() const {
        return max;
    }

    void ArrayType::setMax(int arrayMax) {
        max = arrayMax;
    }

    const std::string &ArrayType::getFormat() const {
        return format;
    }

    void ArrayType::setFormat(const std::string &arrayFormat) {
        format = arrayFormat;
    }
}

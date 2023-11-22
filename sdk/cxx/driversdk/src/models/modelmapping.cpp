#include "modelmapping.h"

namespace DRIVERSDK {

// Default constructor
    ModelMapping::ModelMapping() : precision(0), deviation(0.0), silentWin(0) {}

// Parameterized constructor
    ModelMapping::ModelMapping(const std::string& attr, const std::string& tp, const std::string& expr,
                               int prec, double dev, int silentWindow)
            : attribute(attr), type(tp), expression(expr), precision(prec), deviation(dev), silentWin(silentWindow) {}

// Getter and setter methods
    const std::string& ModelMapping::getAttribute() const {
        return attribute;
    }

    void ModelMapping::setAttribute(const std::string& attr) {
        attribute = attr;
    }

    const std::string& ModelMapping::getType() const {
        return type;
    }

    void ModelMapping::setType(const std::string& tp) {
        type = tp;
    }

    const std::string& ModelMapping::getExpression() const {
        return expression;
    }

    void ModelMapping::setExpression(const std::string& expr) {
        expression = expr;
    }

    int ModelMapping::getPrecision() const {
        return precision;
    }

    void ModelMapping::setPrecision(int prec) {
        precision = prec;
    }

    double ModelMapping::getDeviation() const {
        return deviation;
    }

    void ModelMapping::setDeviation(double dev) {
        deviation = dev;
    }

    int ModelMapping::getSilentWindow() const {
        return silentWin;
    }

    void ModelMapping::setSilentWindow(int silentWindow) {
        silentWin = silentWindow;
    }

} // namespace DRIVERSDK

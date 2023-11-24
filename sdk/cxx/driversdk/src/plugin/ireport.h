#pragma once

#include <string>

// 定义了驱动 --> 宿主上报消息和状态的接口
class IReport {
public:
    // 驱动向宿主上报信息的接口
    virtual std::string post(const std::string& data) = 0;

    virtual std::string state(const std::string& data) = 0;

    // 虚析构函数，确保正确释放资源
    virtual ~IReport() {}
};

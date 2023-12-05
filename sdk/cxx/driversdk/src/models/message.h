#pragma once

#include <iostream>
#include <map>
#include <variant>
#include <string>
#include <vector>

namespace DRIVERSDK {

    template <typename T>
    class Message {
    public:
        // Constructors
        Message();
        Message(const std::string& kind, const std::map<std::string, std::string>& meta, const T& content);

        // Getters
        std::string getKind() const;
        std::map<std::string, std::string> getMeta() const;
        T getContent() const;

        // Setters
        void setKind(const std::string& kind);
        void setMeta(const std::map<std::string, std::string>& meta);
        void setContent(const T& content);

    private:
        std::string kind;
        std::map<std::string, std::string> meta;
        T content;
    };

} // namespace DRIVERSDK

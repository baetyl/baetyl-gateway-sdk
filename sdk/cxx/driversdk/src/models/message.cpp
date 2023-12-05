#include "message.h"

namespace DRIVERSDK {

    // Template implementation
    template <typename T>
    Message<T>::Message() {}

    template <typename T>
    Message<T>::Message(const std::string& kind, const std::map<std::string, std::string>& meta, const T& content)
            : kind(kind), meta(meta), content(content) {}

    template <typename T>
    std::string Message<T>::getKind() const {
        return kind;
    }

    template <typename T>
    std::map<std::string, std::string> Message<T>::getMeta() const {
        return meta;
    }

    template <typename T>
    T Message<T>::getContent() const {
        return content;
    }

    template <typename T>
    void Message<T>::setKind(const std::string& kind) {
        this->kind = kind;
    }

    template <typename T>
    void Message<T>::setMeta(const std::map<std::string, std::string>& meta) {
        this->meta = meta;
    }

    template <typename T>
    void Message<T>::setContent(const T& content) {
        this->content = content;
    }

    // Explicit instantiation for commonly used types
    template class Message<int>;
    template class Message<std::string>;
    template class Message<std::vector<int>>;

} // namespace DRIVERSDK

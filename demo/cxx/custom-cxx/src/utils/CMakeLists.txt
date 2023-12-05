set(UTIL_SRC_LIST
        l.cpp
)

add_library(util STATIC ${UTIL_SRC_LIST})

target_include_directories(util PUBLIC
        ${ROOT_DIR}
        ${ROOT_INCLUDE}
)
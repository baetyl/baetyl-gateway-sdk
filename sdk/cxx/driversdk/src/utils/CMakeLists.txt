set(UTIL_SRC_LIST)

add_library(util STATIC ${UTIL_SRC_LIST})

target_include_directories(util PUBLIC
        ${ROOT_DIR}
        )

target_link_libraries(util
        )

set_target_properties(util PROPERTIES
        C_VISIBILITY_PRESET hidden
        CXX_VISIBILITY_PRESET hidden
        VISIBILITY_INLINES_HIDDEN TRUE
        )
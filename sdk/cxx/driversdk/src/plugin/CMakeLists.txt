set(PLUGIN_SRC_LIST)

add_library(plugin STATIC ${PLUGIN_SRC_LIST})

target_include_directories(plugin PUBLIC
        ${ROOT_DIR}
        ${DRIVER_PROTO_INCLUDE}
        )

find_package(gRPC CONFIG REQUIRED)

target_link_libraries(plugin
        gRPC::grpc++
        )

set_target_properties(plugin PROPERTIES
        C_VISIBILITY_PRESET hidden
        CXX_VISIBILITY_PRESET hidden
        VISIBILITY_INLINES_HIDDEN TRUE
        )
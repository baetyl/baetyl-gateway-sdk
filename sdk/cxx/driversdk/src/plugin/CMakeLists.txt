find_package(Protobuf REQUIRED)
if(NOT Protobuf_FOUND)
        message(FATAL_ERROR "Protobuf not found. Please install Protobuf by running:\n"
                "  sudo apt-get install protobuf-compiler\n"
                "Or visit: https://github.com/protocolbuffers/protobuf")
endif()
message("Found Protobuf version: ${Protobuf_VERSION}")

find_package(gRPC REQUIRED)
if(NOT gRPC_FOUND)
        message(FATAL_ERROR "gRPC not found. Please install gRPC by running:\n"
                "  sudo apt-get install -y libgrpc++-dev\n"
                "Or visit: https://grpc.io/docs/quickstart/cpp/")
endif()
message("Found gRPC version: ${gRPC_VERSION}")

set(PLUGIN_SRC_LIST

)

add_library(plugin STATIC ${PLUGIN_SRC_LIST})

include_directories(${PROTOBUF_INCLUDE_DIR})

target_include_directories(plugin PUBLIC
        ${ROOT_DIR}
        ${DRIVER_PROTO_INCLUDE}
        ${_gRPC_INCLUDE_DIR}
        ${Protobuf_INCLUDE_DIRS}
        )

target_link_libraries(plugin PUBLIC
        protobuf::libprotobuf
        gRPC::grpc
        gRPC::grpc++
        )

set_target_properties(plugin PROPERTIES
        C_VISIBILITY_PRESET hidden
        CXX_VISIBILITY_PRESET hidden
        VISIBILITY_INLINES_HIDDEN TRUE
        )
find_package(Protobuf REQUIRED)
if (NOT Protobuf_FOUND)
    message(FATAL_ERROR "Protobuf not found. Please install Protobuf by running:\n"
            "  sudo apt-get install protobuf-compiler\n"
            "Or visit: https://github.com/protocolbuffers/protobuf")
endif ()
message("-- Found Protobuf version: ${Protobuf_VERSION}")

find_package(gRPC REQUIRED)
if (NOT gRPC_FOUND)
    message(FATAL_ERROR "gRPC not found. Please install gRPC by running:\n"
            "  sudo apt-get install -y libgrpc++-dev\n"
            "Or visit: https://grpc.io/docs/quickstart/cpp/")
endif ()
message("-- Found gRPC version: ${gRPC_VERSION}")

set(PLUGIN_SRC_LIST
        serve.cpp
        grpcdriver.cpp
        grpcreport.cpp
)

add_library(driverplugin STATIC ${PLUGIN_SRC_LIST})

message("-- _gRPC_INCLUDE_DIR: ${_gRPC_INCLUDE_DIR}")
message("-- PROTOBUF_INCLUDE_DIR: ${PROTOBUF_INCLUDE_DIR}")

include_directories(
        ${DRIVER_PROTO_INCLUDE}
        ${PROTOBUF_INCLUDE_DIR}
        ${_gRPC_INCLUDE_DIR}
)

target_include_directories(driverplugin PUBLIC
        ${ROOT_DIR}
        ${DRIVER_PROTO_INCLUDE}
        ${_gRPC_INCLUDE_DIR}
        ${Protobuf_INCLUDE_DIRS}
)

target_link_libraries(driverplugin PUBLIC
        protobuf::libprotobuf
        gRPC::grpc
        gRPC::grpc++
        gRPC::grpc++_reflection
)

set_target_properties(driverplugin PROPERTIES
        C_VISIBILITY_PRESET hidden
        CXX_VISIBILITY_PRESET hidden
        VISIBILITY_INLINES_HIDDEN TRUE
)
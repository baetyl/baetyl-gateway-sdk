find_package(protobuf CONFIG REQUIRED)
find_package(gRPC CONFIG REQUIRED)
find_package(Threads)

#
# Protobuf/Grpc source files
#
set(PROTO_FILES
        driver.proto
)

#
# Add Library target with protobuf sources
#
add_library(driverproto ${PROTO_FILES})

target_link_libraries(driverproto
        PUBLIC
        protobuf::libprotobuf
        gRPC::grpc
        gRPC::grpc++
)

target_include_directories(driverproto PUBLIC ${CMAKE_CURRENT_BINARY_DIR})

#
# Compile protobuf and grpc files in driverproto target to cpp
#
get_target_property(grpc_cpp_plugin_location gRPC::grpc_cpp_plugin LOCATION)
protobuf_generate(TARGET driverproto LANGUAGE cpp)
protobuf_generate(TARGET driverproto LANGUAGE grpc GENERATE_EXTENSIONS .grpc.pb.h .grpc.pb.cc PLUGIN "protoc-gen-grpc=${grpc_cpp_plugin_location}")

install(TARGETS driverproto
        EXPORT driverproto-targets
        LIBRARY DESTINATION ${CMAKE_INSTALL_LIBDIR}
)
install(DIRECTORY ${CMAKE_CURRENT_BINARY_DIR}/
        DESTINATION ${CMAKE_INSTALL_PREFIX}/include/
        FILES_MATCHING PATTERN "*.h"
)
install(CODE "
    file(GLOB include_subdirectories \"${CMAKE_INSTALL_PREFIX}/include/*\")
    foreach(dir \${include_subdirectories})
        file(GLOB_RECURSE header_files \"\${dir}/*.h\")
        if(NOT header_files)
            execute_process(COMMAND ${CMAKE_COMMAND} -E remove_directory \"\${dir}\")
        endif()
    endforeach()
")
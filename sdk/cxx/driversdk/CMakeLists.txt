cmake_minimum_required(VERSION 3.13)
project(driversdk)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_POSITION_INDEPENDENT_CODE ON)

# set root path
set(ROOT_INCLUDE "${CMAKE_SOURCE_DIR}/include")
set(ROOT_DIR "${CMAKE_SOURCE_DIR}/src")

# set 3rd path
set(SUFFIX "${CMAKE_SYSTEM_PROCESSOR}")
set(CMAKE_POSITION_INDEPENDENT_CODE ON)
set(THIRDPART_DIR "${CMAKE_SOURCE_DIR}/3rdparty")
set(3RD_FINDER "${THIRDPART_DIR}/3rdfinder")

# set proto path
set(PROTO_DIR "${CMAKE_SOURCE_DIR}/proto")

# include 3rd finder
include("${PROTO_DIR}/driver-proto.cmake")
include("${3RD_FINDER}/yaml-cpp.cmake")

include_directories(
        ${DRIVER_PROTO_INCLUDE_DIRS}
        ${YAML_CPP_INCLUDE_DIRS}
)

# add subdirectory
add_subdirectory(${ROOT_DIR}/models)
add_subdirectory(${ROOT_DIR}/utils)
add_subdirectory(${ROOT_DIR}/context)
add_subdirectory(${ROOT_DIR}/plugin)

#add_subdirectory(${ROOT_DIR}/output)

# install
install(DIRECTORY ${ROOT_DIR}/
        DESTINATION include
        FILES_MATCHING PATTERN "*.h"
)

# 安装静态链接库到指定的安装目录
install(TARGETS
        driverplugin
        driverutil
        drivermodel
        drivercontext
        ARCHIVE DESTINATION lib
)
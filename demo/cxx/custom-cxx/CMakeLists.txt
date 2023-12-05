cmake_minimum_required(VERSION 3.13)
project(custom-cxx)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_POSITION_INDEPENDENT_CODE ON)

# set root path
set(ROOT_INCLUDE "${CMAKE_SOURCE_DIR}/include")
set(ROOT_LIB_DIR "${CMAKE_SOURCE_DIR}/lib")
set(ROOT_DIR "${CMAKE_SOURCE_DIR}/src")

# set 3rd path
set(SUFFIX "${CMAKE_SYSTEM_PROCESSOR}")
set(CMAKE_POSITION_INDEPENDENT_CODE ON)

include_directories(
        ${ROOT_INCLUDE}
)

# add subdirectory
add_subdirectory(${ROOT_DIR})
add_subdirectory(${ROOT_DIR}/utils)
add_subdirectory(${ROOT_DIR}/models)
add_subdirectory(${ROOT_DIR}/driver)

# install
install(DIRECTORY ${ROOT_DIR}
        DESTINATION include
        FILES_MATCHING PATTERN "*.h"
)

# 安装静态链接库到指定的安装目录
install(TARGETS
        ARCHIVE DESTINATION lib
)
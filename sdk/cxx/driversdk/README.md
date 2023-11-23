# driversdk
driversdk 实现程序

## 构建
    cd $PROJECT_DIR
    ARCH=`uname -m`
    mkdir build-$ARCH
    cd build-$ARCH
    cmake .. -DCMAKE_BUILD_TYPE=Debug|Release
    make -j4

* 输出目录为 PROJECT_DIR/output-${CMAKE\_SYSTEM\_PROCESSOR}
* 最终输出文件为 **libdriversdk.so** 

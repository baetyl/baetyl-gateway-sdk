# custom-cxx
custom-cxx demo 实现程序

TODO: 还没测试，代码还需要完善

## 构建
先运行并安装 [cxxSDK](../../../sdk/cxx/driversdk)

```shell
ARCH=`uname -m`
rm -rf ./build-$ARCH
rm -rf ./install-$ARCH
mkdir build-$ARCH
cd build-$ARCH
cmake .. -DCMAKE_INSTALL_PREFIX=../install-$ARCH
make -j4
make install
```


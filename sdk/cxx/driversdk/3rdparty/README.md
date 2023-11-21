## 项目构建
下载 yaml-cpp 置于 ./3rdparty 目录下 

### 库版本

```
- yaml-cpp 0.8.0
```

### 编译

yaml-cpp

```
git clone git@github.com:jbeder/yaml-cpp.git
git checkout tags/0.8.0
```

```shell
cd "yaml-cpp"
ARCH=`uname -m`
rm -rf ./build-$ARCH
rm -rf ./install-$ARCH
mkdir build-$ARCH
cd build-$ARCH
cmake .. -DCMAKE_INSTALL_PREFIX=../install-$ARCH
make -j4
make install
```

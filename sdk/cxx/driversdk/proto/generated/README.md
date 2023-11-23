## 生成proto代码
### 方案1

参考 ./proto/generated/build.sh

```cmake
# 负责将 .proto 文件中的所有消息类型编译成C++
protobuf_generate(TARGET xxx LANGUAGE cpp)

# 1. 获取 grpc_cpp_plugin 插件
# 2. 通过插件生成客户端/服务端代码
get_target_property(grpc_cpp_plugin_location gRPC::grpc_cpp_plugin LOCATION)
protobuf_generate(TARGET xxx LANGUAGE grpc GENERATE_EXTENSIONS .grpc.pb.h .grpc.pb.cc PLUGIN "protoc-gen-grpc=${grpc_cpp_plugin_location}")
```

```shell
ARCH=`uname -m`
cmake .. -DCMAKE_INSTALL_PREFIX=../install-$ARCH
make -j4
make install
```

### 方案2
使用 protoc 编译器生成C代码文件。在终端中，使用以下命令

macos 安装  grpc_cpp_plugin 可以通过 brew install grpc

```shell
protoc --proto_path=./proto driver.proto --cpp_out=./proto --grpc_out=./proto --plugin=protoc-gen-grpc
```

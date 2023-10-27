## Golang SDK 说明
### 生成 *.pb.go 文件
基于 .proto 文件生成 .pb.go 文件，需要以下几个步骤

**步骤1**

安装 Protocol Buffers 编译器 protoc：

访问 Protocol Buffers 的 GitHub 仓库 [下载](https://github.com/protocolbuffers/protobuf/releases) 最新的编译器版本

解压缩下载的文件，并将 protoc 可执行文件添加到系统的可执行文件路径中

**步骤2**

安装 Go 的 Protobuf 插件 protoc-gen-go：

打开终端并运行以下命令来安装 protoc-gen-go 插件：
```shell
go get google.golang.org/protobuf/cmd/protoc-gen-go
```

**步骤3**

使用 protoc 编译器生成Go代码文件。在终端中，使用以下命令

```shell
protoc --proto_path=../../proto driver.proto --go_out=plugins=grpc:. 
```

* **protoc**:  Protocol Buffers 编译器的命令
* **--proto_path=../../proto**: 指定了 proto 文件的搜索路径。在这里，../../proto 是 proto 文件的相对路径，表示 proto 文件位于当前目录的两级父目录下的 proto 目录中。编译器将在这个路径下查找 driver.proto 文件
* **driver.proto**: 要编译的具体 protobuf 文件的名称。在这个例子中，它是 driver.proto
* **--go_out=plugins=grpc:.**: 这是生成 Go 语言代码的选项。它告诉编译器使用 gRPC 插件生成 Go 语言的代码，并将生成的文件放在当前目录 ./ 中

综合起来，此命令通知 Protocol Buffers 编译器在指定的 ../../proto 目录下查找 driver.proto 文件，然后使用 gRPC 插件生成 Go 语言的代码，并将生成的代码文件放在当前目录。这些生成的文件通常包括用于与 gRPC 通信的 Go 语言结构和函数
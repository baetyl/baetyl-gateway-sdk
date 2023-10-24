## 1.项目构建
### 1.1 生成grpc代码
拷贝 .proto 文件到 ./src/main/proto 目录下

在 sdk/java/driversdk 项目目录下执行 

```shell
gradle generateProto
```

执行命令后：
* 将生成 ./src/main/proto 目录下所有 .proto 文件的 grpc 及 java 代码
* 生成文件位于 ./src/generated 目录下

```shell
gradle clean
```

执行 clean 后会清理自动生成的代码

执行 shadowJar 构建产出（此处不要直接使用 gradle build 构建，会缺少依赖项）

执行 copyShadowJar task 会将生成的 sdk jar（默认位于 build/libs 目录） 拷贝到 demo/java/custom-java 目录

生成 jar 包并复制到 baetyl-gateway/etc/custom-java 目录下: gradle 执行 copyShadowJar 任务

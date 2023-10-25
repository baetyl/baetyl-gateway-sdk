## 1.项目构建
执行 shadowJar 构建产出（此处不要直接使用 gradle build 构建，会缺少依赖项）

生成 jar 包并复制到 baetyl-gateway/etc/custom-java 目录下: gradle 执行 copyShadowJar 任务

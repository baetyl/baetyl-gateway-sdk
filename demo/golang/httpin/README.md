# HTTP IN

插件启动后，根据配置监听对应端口，提供如下接口

```shell
GET    /health   
POST   /v1/report
```

其中 `/helath` 接口用于探测服务启动状态，正常时返回 http-code:200 , body: true

`/v1/report` 用户接收数据的上报，body为键值对格式，键为string类型，值为任意类型
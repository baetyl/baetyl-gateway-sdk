## 说明
三个配置文件在实际使用中，由 [baetyl-cloud](https://github.com/baetyl/baetyl-cloud) 创建驱动示例的时候自动生成

完整操作可以在  [BIE 控制台](https://console.bce.baidu.com/iot2/bie/device/product/list) 参考文档 [设备管理](https://cloud.baidu.com/doc/BIE/s/fkmy9wd8o) 进行实践

## 配置文件解析
### [sub_devices.yml](./sub_devices.yml)
『节点管理』-『子设备管理』中的驱动配置和子设备配置信息存放在sub_device.yml中

* devcies[i].accessConfig.custom为子设备的配置

开发时仅需关注下列字段

```yaml
devices:
  - name: custom-zx   # 待连接设备的名称
    deviceModel: custom-simulator   # 待连接设备对应的设备模型的名称，对应 models.yml 中同名的 key
    accessTemplate: custom-access-template    # 待连接设备对应的接入模板的名称，对应 access_template.yml 中同名的 key
    accessConfig: # 设备具体接入信息，所有信息包含在 custom 字段下面，字段的值由用户自行定义，后续自行在驱动实现时完成解析，如本示例，以 json 字符串格式序列化接入信息
      custom: '{"device": "custom-dev-0","interval": 3000000000}'   # 自定义的设备接入信息，本示例，json 格式提供了接入物理设备名称，采样间隔 3s 
```

### [models.yml](./models.yml)
『子设备管理』-『产品』中的产品测点信息存放在models.yml中

```yaml
custom-simulator:   # 设备模型的名称，与 sub_devices.yml 中 devcies[i] 其中一条的 `deviceModel` 字段对应
  - name: temperature   # 点位标识
    type: float32       # 点位字段数据类型，支持 bool string int16 int32 int64 float32 float64 bool 等
    mode: ro            # 点位读写类型（ro：只读；rw：读写）
```

### [access_template.yml](./access_template.yml)
『子设备管理』-『接入模板』中的设备点表和物模型点位映射信息存放在access_template.yml中

* device-access-tpl.properties[i].visitor.custom为设备点表信息中的采集配置
* device-access-tpl.properties[i].mapping为物模型点位映射信息

开发时仅需关注下列字段

```yaml
custom-access-template:   # 接入模板的名称，与 sub_devices.yml 中 devcies[i] 其中一条的 `accessTemplate` 字段对应
  properties:   # 测点信息，其中 visitor 中 custom 字段记录了这个点位在具体物理设备中访问所需的信息
    - name: temperature   # 测点名称
      id: "1"     # 在 Baetyl 设备管理系统中的测点标识
      type: float32   # 测点数据类型
      visitor:    # 测点的的具体访问信息
        custom: '{"name":"temperature","type":"float32","index":0}'   # 点位在具体物理设备中访问所需的信息，字段的值由用户自行定义，后续自行在驱动中完成解析，并依据此信息采集具体点位的值
```
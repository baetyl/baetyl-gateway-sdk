package com.baidu.bce.sdk.context.models.yaml;

import com.baidu.bce.sdk.context.models.device.DeviceInfo;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class SubDeviceYaml {
    List<DeviceInfo> devices;
    String driver;
}

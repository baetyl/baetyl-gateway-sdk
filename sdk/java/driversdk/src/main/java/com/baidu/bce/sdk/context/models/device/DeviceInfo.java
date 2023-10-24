package com.baidu.bce.sdk.context.models.device;

import com.baidu.bce.sdk.context.models.access.AccessConfig;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class DeviceInfo {
    private String name;
    private String version;
    private String deviceModel;
    private String accessTemplate;
    private AccessConfig accessConfig;
}
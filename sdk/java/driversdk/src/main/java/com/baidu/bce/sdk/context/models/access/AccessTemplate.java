package com.baidu.bce.sdk.context.models.access;

import com.baidu.bce.sdk.context.models.device.DeviceProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class AccessTemplate {
    private String name;
    private String version;
    private List<DeviceProperty> properties;
    private List<ModelMapping> mappings;
}
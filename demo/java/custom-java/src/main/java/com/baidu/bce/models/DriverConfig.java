package com.baidu.bce.models;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class DriverConfig {
    private String driverName;
    private List<DeviceConfig> devices;
    private List<Job> jobs;
}

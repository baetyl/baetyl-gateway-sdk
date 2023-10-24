package com.baidu.bce.models;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class Job {
    private String device;
    private long interval; // ns
    private List<Property> properties;
}

package com.baidu.bce.models;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class Point {
    /**
     * 点位标识符
     */
    public static final int INDEX_TEMPERATURE = 0;
    public static final int INDEX_HUMIDITY = 1;
    public static final int INDEX_PRESSURE = 2;

    private float temperature;
    private float humidity;
    private float pressure;
}
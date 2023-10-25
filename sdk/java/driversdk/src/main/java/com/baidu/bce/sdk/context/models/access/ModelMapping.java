package com.baidu.bce.sdk.context.models.access;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class ModelMapping {
    private String attribute;
    private String type;
    private String expression;
    private int precision;
    private double deviation;
    private int silentWin;
}

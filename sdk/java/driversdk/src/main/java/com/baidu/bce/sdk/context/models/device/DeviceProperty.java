package com.baidu.bce.sdk.context.models.device;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;
import java.util.Map;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class DeviceProperty {
    private String name;
    private String id;
    private String type;
    private String mode;
    private String unit;
    private PropertyVisitor visitor;
    private String format;
    private EnumType enumType;
    private ArrayType arrayType;
    private Map<String, ObjectType> objectType;
    private List<String> objectRequired;
    private Object current;
    private Object expect;
}
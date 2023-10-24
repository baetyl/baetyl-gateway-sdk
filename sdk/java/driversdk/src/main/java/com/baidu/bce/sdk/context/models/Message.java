package com.baidu.bce.sdk.context.models;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.Map;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class Message {
    private String kind;
    private Map<String, String> meta;
    private Object content;
}

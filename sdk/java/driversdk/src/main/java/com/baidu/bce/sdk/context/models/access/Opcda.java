package com.baidu.bce.sdk.context.models.access;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import java.time.Duration;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Opcda {
    private String host;
    private String clsid;
    private String programid;
    private String username;
    private String password;
    private Long interval;
}

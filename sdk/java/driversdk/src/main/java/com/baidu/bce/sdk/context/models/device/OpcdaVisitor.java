package com.baidu.bce.sdk.context.models.device;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class OpcdaVisitor {
    private String datapath;
    private String type;
}

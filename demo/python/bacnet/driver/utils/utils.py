# !/usr/bin/env python3
def bacnetType(key: int):
    typeMap = {0: "analogInput", 1: "analogOutput", 2: "analogValue", 3: "binaryInput", 4: "binaryOutput",
               5: "binaryValue"}
    return typeMap[key]

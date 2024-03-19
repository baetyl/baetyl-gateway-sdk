# !/usr/bin/env python3
from dataclasses import dataclass


@dataclass
class BacnetAccessConfig:
    ID: str = ""
    Interval: int = 0
    DeviceID: int = 0
    AddressOffset: int = 0
    Address: str = ""
    Port: int = 0

    def FormatByYamlData(self, item: dict):
        self.ID = item.get("id", "")
        self.Interval = item.get("interval", 0)
        self.DeviceID = item.get("deviceId", 0)
        self.AddressOffset = item.get("addressOffset", 0)
        self.Address = item.get("address", "")
        self.Port = item.get("port", 0)

# !/usr/bin/env python3
from dataclasses import dataclass


@dataclass
class OpcdaAccessConfig:
    Server: str
    Host: str
    Group: str
    Interval: int

    def FormatByYamlData(self, item: dict):
        self.Server = item.get("server", "")
        self.Host = item.get("host", "")
        self.Group = item.get("group", "")
        self.Interval = item.get("interval", 0)

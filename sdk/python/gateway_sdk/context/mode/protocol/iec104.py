# !/usr/bin/env python3
from dataclasses import dataclass
from typing import Optional


@dataclass
class IEC104AccessConfig:
    ID: int = 0
    Interval: float = 0
    Endpoint: str = ""
    AIOffset: int = 0
    DIOffset: int = 0
    AOOffset: int = 0
    DOOffset: int = 0

    def FormatByYamlData(self, item: dict):
        self.ID = item.get("id", 0)
        self.Interval = item.get("interval", 0)
        self.Endpoint = item.get("endpoint", "")
        self.AIOffset = item.get("aIOffset", 0)
        self.DIOffset = item.get("dIOffset", 0)
        self.AOOffset = item.get("aOOffset", 0)
        self.DOOffset = item.get("dOOffset", 0)

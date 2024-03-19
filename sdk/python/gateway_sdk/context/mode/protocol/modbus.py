# !/usr/bin/env python3
from dataclasses import dataclass
from typing import Optional
import time


@dataclass
class TCPConfig:
    Address: str = ""
    Port: int = 0

    def FormatByYamlData(self, item: dict):
        self.Address = item.get("address", "")
        self.Port = item.get("port", 0)


@dataclass
class RTUConfig:
    Port: str = ""
    BaudRate: int = 19200
    Parity: str = "E"
    DataBit: int = 8
    StopBit: int = 1

    def FormatByYamlData(self, item: dict):
        self.BaudRate = item.get("baudRate", 19200)
        self.Port = item.get("port", "")
        self.Parity = item.get("parity", "E")
        self.DataBit = item.get("dataBit", 8)
        self.StopBit = item.get("stopBit", 1)


@dataclass
class ModbusAccessConfig:
    ID: int = 0
    Interval: int = 0
    TCP: TCPConfig = None
    RTU: RTUConfig = None
    Timeout: int = 10
    IdleTimeout: int = 60

    def FormatByYamlData(self, item: dict):
        self.ID = item.get("id", 0)
        self.Interval = item.get("interval", 0)
        self.Timeout = item.get("timeout", 10)
        self.IdleTimeout = item.get("idleTimeout", 60)
        if "tcp" in item:
            config = TCPConfig()
            config.FormatByYamlData(item.get("tcp"))
            self.TCP = config
        if "rtp" in item:
            config = RTUConfig()
            config.FormatByYamlData(item.get("rtu"))
            self.RTU = config

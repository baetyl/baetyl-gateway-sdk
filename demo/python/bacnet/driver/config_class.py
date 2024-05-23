# !/usr/bin/env python3
from typing import List, Dict
from gateway_sdk.context.mode.mode import BacnetAccessConfig
from gateway_sdk.context.mode.mode import DeviceProperty as accessPorperty


class Property:
    Id: str
    Name: str
    Type: str
    Mode: str
    BacnetType: int
    BacnetAddress: int
    ApplicationTagNumber: str

    def genByConfig(self, p: accessPorperty, b: BacnetAccessConfig):
        self.Id = p.ID
        self.Name = p.Name
        self.Type = p.Type
        self.Mode = p.Mode
        self.BacnetType = b.DeviceID


class Job:
    Device: str
    DeviceId: int
    AddressOffset: int
    Properties: Dict[str, Property]
    Interval: int


class SlaveConfig:
    Device: str
    DeviceID: int
    Interval: int
    Address: str
    Port: int

    def genByConfig(self, bacnet: BacnetAccessConfig):
        self.DeviceID = bacnet.DeviceID
        self.Interval = bacnet.Interval
        self.Address = bacnet.Address
        self.Port = bacnet.Port


class Config:
    DriverName: str = ""
    Slaves: List[SlaveConfig] = None
    Jobs: list[Job] = None

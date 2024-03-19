# !/usr/bin/env python3
from dataclasses import dataclass
from typing import Optional


@dataclass
class OpcuaSecurity:
    Policy: str = ""
    Mode: str = ""

    def FormatByYamlData(self, item: dict):
        self.Policy = item.get("policy", "")
        self.Mode = item.get("mode", "")


@dataclass
class OpcuaAuth:
    Username: str = ""
    Password: str = ""

    def FormatByYamlData(self, item: dict):
        self.Username = item.get("username", "")
        self.Password = item.get("password", "")


@dataclass
class OpcuaCertificate:
    Cert: str = ""
    Key: str = ""

    def FormatByYamlData(self, item: dict):
        self.Cert = item.get("cert", "")
        self.Key = item.get("key", "")


@dataclass
class OpcuaAccessConfig:
    Interval: int = 0
    Timeout: int = 0
    Auth: OpcuaAuth = None
    Certificate: OpcuaCertificate = None
    Security: OpcuaSecurity = None
    ID: int = 0
    Endpoint: str = ""
    NsOffset: int = 0
    IDOffset: int = 0

    def FormatByYamlData(self, item: dict):
        self.ID = item.get("id", 0)
        self.Interval = item.get("interval", 0)
        self.Timeout = item.get("timeout", 0)
        self.Endpoint = item.get("endpoint", "")
        self.NsOffset = item.get("nsOffset", 10)
        self.IDOffset = item.get("IDOffset", 60)

        if "auth" in item:
            auth = OpcuaAuth()
            auth.FormatByYamlData(item.get("auth"))
            self.Auth = auth
        if "certificate" in item:
            cer = OpcuaCertificate()
            cer.FormatByYamlData(item.get("certificate"))
            self.Certificate = cer
        if "security" in item:
            sec = OpcuaSecurity()
            sec.FormatByYamlData(item.get("security"))
            self.Security = sec

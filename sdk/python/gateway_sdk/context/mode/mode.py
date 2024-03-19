# !/usr/bin/env python3
from dataclasses import dataclass
from typing import List, Dict, Any
from .protocol.bacnet import BacnetAccessConfig
from .protocol.modbus import ModbusAccessConfig
from .protocol.opcua import OpcuaAccessConfig
from .protocol.opcda import OpcdaAccessConfig
from .protocol.iec104 import IEC104AccessConfig
from .protocol.custom import CustomAccessConfig


@dataclass
class ReportProperty:
    Time: int
    Value: Any


@dataclass
class Event:
    Payload: Any
    Type: str


@dataclass
class EnumValue:
    Name: str
    Value: str
    DisplayName: str


@dataclass
class EnumType:
    Type: str = ""
    Values: List[EnumValue] = None


@dataclass
class ArrayType:
    Type: str = ""
    Min: int = 0
    Max: int = 0
    Format: str = ""


@dataclass
class ObjectType:
    DisplayName: str
    Type: str
    Format: str


@dataclass
class IEC104Visitor:
    PointNum: int = 0
    PointType: str = ""
    Type: str = ""

    def FormatByYamlData(self, item: Dict):
        self.Type = item.get("type", "")
        self.PointType = item.get("pointType", "")
        self.PointNum = item.get("pointNum", 0)


@dataclass
class ModbusVisitor:
    Function: str = ""
    Address: str = ""
    Quantity: int = 0
    Type: str = ""
    Unit: str = ""
    Scale: float = 0.0
    SwapByte: bool = False
    SwapRegister: bool = False

    def FormatByYamlData(self, item: Dict):
        self.Function = item.get("function", "")
        self.Address = item.get("address", "")
        self.Quantity = item.get("quantity", 0)
        self.Type = item.get("type", 0)
        self.Unit = item.get("unit", 0)
        self.Scale = item.get("scale", 0)
        self.SwapByte = item.get("swapByte", False)
        self.SwapRegister = item.get("swapRegister", False)


@dataclass
class OpcuaVisitor:
    NodeID: str = ""
    Type: str = ""
    NsBase: str = ""
    IDBase: str = ""
    ID: str = ""

    def FormatByYamlData(self, item: Dict):
        self.NodeID = item.get("nodeId", "")
        self.Type = item.get("type", "")
        self.IDBase = item.get("idBase", "")
        self.NsBase = item.get("nsBase", "")
        self.ID = item.get("id", "")


@dataclass
class OpcdaVisitor:
    Datapath: str = ""
    Type: str = ""

    def FormatByYamlData(self, item: Dict):
        self.Datapath = item.get("datapath", "")
        self.Type = item.get("type", "")


@dataclass
class BacnetVisitor:
    Type: str = ""
    BacnetType: int = 0
    BacnetAddress: int = 0
    ApplicationTagNumber: str = ""

    def FormatByYamlData(self, item: Dict):
        self.BacnetType = item.get("bacnetType", 0)
        self.BacnetAddress = item.get("bacnetAddress", 0)
        self.ApplicationTagNumber = item.get("applicationTagNumber", "")
        self.Type = item.get("type", "")


@dataclass
class CustomVisitor(str):
    pass


@dataclass
class PropertyVisitor:
    Modbus: ModbusVisitor = None
    Opcua: OpcuaVisitor = None
    Opcda: OpcdaVisitor = None
    Bacnet: BacnetVisitor = None
    IEC104: IEC104Visitor = None
    Custom: CustomVisitor = None

    def FormatByYamlData(self, item: dict):
        if "modbus" in item:
            modbus = ModbusVisitor()
            modbus.FormatByYamlData(item.get("modbus"))
            self.Modbus = modbus
        if "opcua" in item:
            opcua = OpcuaVisitor()
            opcua.FormatByYamlData(item.get("opcua"))
            self.Opcua = opcua
        if "opcda" in item:
            opcda = OpcdaVisitor()
            opcda.FormatByYamlData(item.get("opcda"))
            self.Opcda = opcda
        if "bacnet" in item:
            bacnet = BacnetVisitor()
            bacnet.FormatByYamlData(item.get("bacnet"))
            self.Bacnet = bacnet
        if "iec104" in item:
            iec = IEC104Visitor()
            iec.FormatByYamlData(item.get("iec104"))
            self.IEC104 = iec
        if "custom" in item:
            cus = CustomVisitor()
            cus = item.get("custom")
            self.Custom = cus


@dataclass
class AccessConfig:
    Modbus: ModbusAccessConfig = None
    Opcua: OpcuaAccessConfig = None
    Opcda: OpcdaAccessConfig = None
    Bacnet: BacnetAccessConfig = None
    IEC104: IEC104AccessConfig = None
    Custom: CustomAccessConfig = None

    def FormatByYamlData(self, item: dict):
        if "modbus" in item:
            proto = ModbusAccessConfig()
            proto.FormatByYamlData(item.get("modbus"))
            self.Modbus = proto
        if "opcua" in item:
            proto = OpcuaAccessConfig()
            proto.FormatByYamlData(item.get("opcua"))
            self.Opcua = proto
        if "opcda" in item:
            proto = OpcdaAccessConfig()
            proto.FormatByYamlData(item.get("opcda"))
            self.Opcda = proto
        if "bacnet" in item:
            proto = BacnetAccessConfig()
            proto.FormatByYamlData(item.get("bacnet"))
            self.Bacnet = proto
        if "iec104" in item:
            proto = IEC104AccessConfig()
            proto.FormatByYamlData(item.get("iec104"))
            self.IEC104 = proto
        if "custom" in item:
            proto = CustomAccessConfig()
            proto = item.get("custom")
            self.Custom = proto


@dataclass
class ModelMapping:
    Attribute: str = ""
    Type: str = ""
    Expression: str = ""
    Precision: int = 0
    Deviation: float = 0.0
    SilentWin: int = 0

    def FormatByYamlData(self, item: dict):
        self.Attribute = item.get("attribute", "")
        self.Type = item.get("type", "")
        self.Expression = item.get("expression", "")
        self.Precision = item.get("precision", 0)
        self.Deviation = item.get("deviation", 0.0)
        self.SilentWin = item.get("silentWin", 0)


@dataclass
class DeviceProperty:
    Name: str = ""
    ID: str = ""
    Type: str = ""
    Mode: str = ""
    Unit: str = ""
    Visitor: PropertyVisitor = None
    Format: str = ""
    EnumType: EnumType = EnumType()
    ArrayType: ArrayType = ArrayType()
    ObjectType: Dict[str, ObjectType] = None
    ObjectRequired: List[str] = None
    Current: Any = None
    Expect: Any = None

    def FormatByModeYamlData(self, item: dict):
        self.Name = item.get("name", "")
        self.Type = item.get("type", "")
        self.Mode = item.get("mode", "")

    def FormatByYamlData(self, item: dict):
        self.Name = item.get("name")
        self.Type = item.get("type")
        self.ID = item.get("id")
        self.Current = item.get("current")
        self.Expect = item.get("expect")
        propertyVisitor = PropertyVisitor()
        propertyVisitor.FormatByYamlData(item.get("visitor"))
        self.Visitor = propertyVisitor


@dataclass
class QOSTopic:
    QOS: int = 0
    Topic: str = ""

    def FormatByYamlData(self, item: dict):
        self.QOS = item.get("qos")
        self.Topic = item.get("topic")


@dataclass
class DeviceTopic:
    Delta: QOSTopic = None
    Report: QOSTopic = None
    Event: QOSTopic = None
    Get: QOSTopic = None
    GetResponse: QOSTopic = None
    EventReport: QOSTopic = None
    PropertyGet: QOSTopic = None
    LifecycleReport: QOSTopic = None

    def FormatByYamlData(self, item: dict):
        if "delta" in item:
            topic = QOSTopic()
            topic.FormatByYamlData(item.get("delta"))
            self.Delta = topic
        if "report" in item:
            topic = QOSTopic()
            topic.FormatByYamlData(item.get("report"))
            self.LifecycleReport = topic
        if "event" in item:
            topic = QOSTopic()
            topic.FormatByYamlData(item.get("event"))
            self.Event = topic
        if "get" in item:
            topic = QOSTopic()
            topic.FormatByYamlData(item.get("get"))
            self.Get = topic
        if "getResponse" in item:
            topic = QOSTopic()
            topic.FormatByYamlData(item.get("getResponse"))
            self.GetResponse = topic
        if "eventReport" in item:
            topic = QOSTopic()
            topic.FormatByYamlData(item.get("eventReport"))
            self.EventReport = topic
        if "propertyGet" in item:
            topic = QOSTopic()
            topic.FormatByYamlData(item.get("propertyGet"))
            self.PropertyGet = topic
        if "lifecycleReport" in item:
            topic = QOSTopic()
            topic.FormatByYamlData(item.get("lifecycleReport"))
            self.LifecycleReport = topic


@dataclass
class DeviceInfo:
    Name: str = ""
    Version: str = ""
    DeviceModel: str = ""
    AccessTemplate: str = ""
    DeviceTopic: DeviceTopic = None
    AccessConfig: AccessConfig = None

    def FormatByYamlData(self, item: dict):
        self.Name = item.get("name", "")
        self.Version = item.get("version", "")
        self.DeviceModel = item.get("deviceModel", "")
        self.AccessTemplate = item.get("accessTemplate", "")
        deviceTopic = DeviceTopic()
        deviceTopic.FormatByYamlData(item.get("deviceTopic"))
        self.DeviceTopic = deviceTopic

        accessConfig = AccessConfig()
        accessConfig.FormatByYamlData(item.get("accessConfig"))
        self.AccessConfig = accessConfig


@dataclass
class DriverConfig:
    Devices: List[DeviceInfo] = None
    Driver: str = ""

    def FormatByYamlData(self, item: dict):
        self.Devices: List[DeviceInfo] = []
        self.Driver = item.get("driver", "")
        deviceInfo = DeviceInfo()
        for it in item.get("devices"):
            deviceInfo.FormatByYamlData(it)
            self.Devices.append(deviceInfo)


@dataclass
class AccessTemplate:
    Name: str = ""
    Version: str = ""
    Properties: List[DeviceProperty] = None
    Mappings: List[ModelMapping] = None

    def FormatByYamlData(self, item: dict):
        self.Properties: List[DeviceProperty] = []

        for properties in item["properties"]:
            deviceProperty = DeviceProperty()
            deviceProperty.FormatByYamlData(properties)
            self.Properties.append(deviceProperty)

        self.Mappings: List[ModelMapping] = []
        for properties in item["mappings"]:
            modelMapping = ModelMapping()
            modelMapping.FormatByYamlData(properties)
            self.Mappings.append(modelMapping)


@dataclass
class ConfigDeviceProperty:
    PropName: str = ""
    PropVal: any = None

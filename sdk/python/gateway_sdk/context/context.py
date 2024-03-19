# !/usr/bin/env python3
from ruamel.yaml import YAML

from typing import Dict, List
from mode.mode import AccessTemplate, DeviceProperty, DeviceInfo, DriverConfig

DefaultSubDeviceConf = "/sub_devices.yml"
DefaultDeviceModelConf = "/models.yml"
DefaultAccessTemplateConf = "/access_template.yml"


def newDmCtx():
    DmCtx()


class DmCtx:
    deviceDriverMap: Dict[str, str] = {}
    devices: Dict[str, Dict[str, DeviceInfo]] = {}
    accessTemplates: Dict[str, Dict[str, AccessTemplate]] = {}
    deviceModels: Dict[str, Dict[str, List[DeviceProperty]]] = {}

    def GetDevice(self, driverName, device):
        return self.devices[driverName][device]

    def GetDriverNameByDevice(self, device):
        return self.deviceDriverMap[device]

    def GetAllDevices(self, driverName)->List[DeviceInfo]:
        deviceList: List[DeviceInfo] = [(value) for value in self.devices[driverName].values()]
        return deviceList

    def GetDeviceModel(self, driverName, device):
        return self.deviceModels[driverName].get(device.DeviceModel)

    def GetAccessTemplates(self, driverName, name) -> AccessTemplate:
        return self.accessTemplates[driverName].get(name)

    def LoadDriverConfig(self, path: str, driverName: str):
        yaml = YAML()
        # 解析model.yml
        with open(path + DefaultDeviceModelConf) as f:
            modelConfig = yaml.load(f)
        deviceMap = {}
        for model, info in modelConfig.items():
            devicePropertys = []
            for item in info:
                deviceProperty = DeviceProperty()
                deviceProperty.FormatByModeYamlData(item)
                devicePropertys.append(deviceProperty)
            deviceMap[model] = devicePropertys
        self.deviceModels[driverName] = deviceMap

        # 解析access_template
        with open(path + DefaultAccessTemplateConf) as f:
            accessConfig = yaml.load(f)

        accessTemplateMap: dict[str, AccessTemplate] = {}
        for accessName, info in accessConfig.items():
            accessTemplate = AccessTemplate()
            accessTemplate.FormatByYamlData(info)
            accessTemplateMap[accessName] = accessTemplate

        for name, access in accessTemplateMap.items():
            access.Name = name
            temp = {name: access}
            self.accessTemplates[driverName] = temp

        # 解析sub_devices
        with open(path + DefaultSubDeviceConf) as f:
            deviceConfig = yaml.load(f)

        driverConfig = DriverConfig()
        driverConfig.FormatByYamlData(deviceConfig)

        devices: dict[str, DeviceInfo] = {}
        for dev in driverConfig.Devices:
            devices[dev.Name] = dev
            self.deviceDriverMap[dev.Name] = driverName

        self.devices[driverName] = devices


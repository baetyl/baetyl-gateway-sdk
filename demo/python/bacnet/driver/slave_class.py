# !/usr/bin/env python3
import BAC0
import json
from config_class import SlaveConfig
from gateway_sdk.interface_report import IReport
from gateway_sdk.context.mode.mode import DeviceInfo
from gateway_sdk.proto.driver_pb2 import RequestArgs

SlaveOffline = 0
SlaveOnline = 1


class Slave:
    info: DeviceInfo
    cfg: SlaveConfig
    bacnet_client: BAC0.lite
    device: str
    driverName: str
    report: IReport
    fail: int
    status: int = 0

    @staticmethod
    def NewSlave(r: IReport, driverName: str, info: DeviceInfo, cfg: SlaveConfig):
        bacnet = BAC0.connect(ip="172.30.195.37/20", port=cfg.Port)
        device = bacnet.whois()
        slave = Slave()
        slave.bacnet_client = bacnet

        for value in device:
            if cfg.DeviceID == value[1]:
                slave.device = value[0]
                break
        slave.cfg = cfg
        slave.info = info
        slave.report = r
        slave.driverName = driverName
        return slave

    def UpdateStatus(self, status):
        if status == self.status:
            return
        if status == SlaveOffline:
            self.fail += 1
            if self.fail == 3:
                try:
                    self.Offline()
                except:
                    self.status = SlaveOffline
                    self.fail = 0
        elif status == SlaveOnline:
            self.Online()
            self.status = SlaveOnline

    def Online(self):
        msg = {
            "kind": "thing.lifecycle.post",
            "meta": {
                "driverName": self.driverName,
                "deviceName": self.info.Name
            },
            "content": {"Value": True}
        }
        msg_data = json.dumps(msg)
        res = self.report.Post(RequestArgs(request=msg_data))
        return res

    def Offline(self):
        msg = {
            "kind": "thing.lifecycle.post",
            "meta": {
                "driverName": self.driverName,
                "deviceName": self.info.Name
            },
            "content": {"Value": False}
        }
        msg_data = json.dumps(msg)
        res = self.report.State(RequestArgs(request=msg_data))
        return res

    def ReadValue(self, bacnetType: str, bacnetAddress: int):
        readStr = self.device + " " + bacnetType + " " + str(bacnetAddress) + " presentValue "
        try:
            value = self.bacnet_client.read(readStr)
        except:
            value = 0
        return value

    def WriteValue(self, bacnetType: int, bacnetAddress: int, val: any):
        self.bacnet_client.write(
            self.device + " " + str(bacnetType) + " " + str(bacnetAddress) + " presentValue " + str(val) + " - 1")

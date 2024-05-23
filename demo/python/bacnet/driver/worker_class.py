# !/usr/bin/env python3
import json

from config_class import Job
from slave_class import Slave
from gateway_sdk.interface_report import IReport
from utils.utils import bacnetType
from gateway_sdk.context.context import DmCtx
from gateway_sdk.proto.driver_pb2 import RequestArgs


class Worker:
    driverName: str
    report: IReport
    job: Job
    slave: Slave
    ctx: DmCtx

    @staticmethod
    def NewWorker(r: IReport, driverName: str, ctx: DmCtx, job: Job, salve: Slave):
        worker = Worker()
        worker.report = r
        worker.driverName = driverName
        worker.job = job
        worker.slave = salve
        worker.ctx = ctx
        return worker

    def Execute(self, report: bool):
        temp = {}
        for prop in self.job.Properties.values():
            bacnet_type = bacnetType(prop.BacnetType)
            value = self.slave.ReadValue(bacnet_type, prop.BacnetAddress)
            temp[prop.Name] = value
        msg_kind = "deviceReport" if report else "deviceDesire"
        msg = {
            "kind": msg_kind,
            "meta": {
                "driverName": self.driverName,
                "deviceName": self.slave.info.Name
            },
            "content": {"Value": temp}
        }
        msg_data = json.dumps(msg)
        self.report.Post(RequestArgs(request=msg_data))


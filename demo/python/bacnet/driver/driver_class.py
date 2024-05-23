# !/usr/bin/env python3

from gateway_sdk.interface_driver import IDriver
from gateway_sdk.interface_report import IReport
from gateway_sdk.context.context import DmCtx

from gateway_sdk.context.mode.mode import DeviceInfo, ConfigDeviceProperty, Event
from bacnet_class import Bacnet
from slave_class import Slave
from config_class import Config, SlaveConfig, Job, Property
from typing import List
from slave_class import SlaveOnline, SlaveOffline
from worker_class import Worker


class Driver(IDriver):
    Report: IReport
    Bacnet: Bacnet = None
    config_path: str
    DriverName: str

    @staticmethod
    def newDriver():
        return Driver()

    def GetDriverInfo(self, data):
        return "success"

    def SetConfig(self, req):
        self.config_path = req
        return "success"

    def Setup(self, driver, report):
        self.Report = report
        self.DriverName = driver
        return "success"

    def Start(self, data):
        ctx = DmCtx()
        ctx.LoadDriverConfig(self.config_path, self.DriverName)
        cfg = genConfig(ctx, self.DriverName)
        bacnet = self.NewBacnet(ctx,cfg)
        bacnet.Start()
        self.Bacnet = bacnet
        return "success"


    def Restart(self, data):
        self.Bacnet.Restart()
        return "success"

    def Stop(self, data):
        self.Bacnet.Stop()
        return "success"

    def Get(self, data):
        print(data)

    def Set(self, data):
        print(data)

    def NewBacnet(self, ctx: DmCtx, cfg: Config):

        infos: dict[str, DeviceInfo] = {}
        allDevice = ctx.GetAllDevices(self.DriverName)
        for item in allDevice:
            infos[item.Name] = item

        slaves: dict[str, Slave] = {}
        for slaveConfig in cfg.Slaves:
            if slaveConfig.Device in infos:
                slave = Slave.NewSlave(self.Report, self.DriverName,infos.get(slaveConfig.Device), slaveConfig)
                slaves[slaveConfig.Device] = slave
                slave.UpdateStatus(SlaveOnline)

        ws: dict[str, Worker] = {}
        for job in cfg.Jobs:
            slave = slaves.get(job.Device)
            ws[job.Device] = Worker.NewWorker(self.Report,self.DriverName, ctx, job, slave)

        bacnet = Bacnet.NewBacnet(ws,ctx)
        return bacnet


def genConfig(ctx: DmCtx, driverName: str) -> Config:
    config = Config()
    slaves: List[SlaveConfig] = []
    jobs: List[Job] = []

    allDevice = ctx.GetAllDevices(driverName)
    for device in allDevice:
        accessConfig = device.AccessConfig
        if accessConfig.Bacnet is None:
            raise Exception("access config bacnet is none")
        slave = SlaveConfig()
        slave.genByConfig(accessConfig.Bacnet)
        slave.Device = device.Name
        slaves.append(slave)

        jobMaps: dict[str, Property] = {}
        dev_tpl = ctx.GetAccessTemplates(driverName, device.AccessTemplate)
        if dev_tpl is None:
            raise Config
        if dev_tpl.Properties is not None:
            for prop in dev_tpl.Properties:
                if prop.Visitor.Bacnet is not None:
                    job = Property()
                    job.Id = prop.ID
                    job.Name = prop.Name
                    job.BacnetType = prop.Visitor.Bacnet.BacnetType
                    job.ApplicationTagNumber = prop.Visitor.Bacnet.ApplicationTagNumber
                    job.BacnetAddress = prop.Visitor.Bacnet.BacnetAddress + accessConfig.Bacnet.AddressOffset
                    jobMaps["%s:%s" % (str(job.BacnetType), str(job.BacnetAddress))] = job
                else:
                    raise Exception("access config Properties is none")
        job = Job()
        job.Device = device.Name
        job.Interval = accessConfig.Bacnet.Interval
        job.Properties = jobMaps
        job.DeviceId = accessConfig.Bacnet.DeviceID
        job.AddressOffset = accessConfig.Bacnet.AddressOffset
        jobs.append(job)
    config.Jobs = jobs
    config.Slaves = slaves
    config.DriverName = driverName
    return config

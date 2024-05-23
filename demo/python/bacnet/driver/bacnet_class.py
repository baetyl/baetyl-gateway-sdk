# !/usr/bin/env python3
from typing import Dict, List
from gateway_sdk.context.context import DmCtx
from gateway_sdk.context.mode.mode import DeviceInfo, ConfigDeviceProperty, Event
import threading
from utils.worke_thread import Thread
from worker_class import Worker


class Bacnet:
    ws: Dict[str, Worker] = None
    ctx: DmCtx = None
    threads: List[Thread] = []

    @staticmethod
    def NewBacnet(ws: dict[str, Worker], ctx: DmCtx):
        bacnet = Bacnet()
        bacnet.ws = ws
        bacnet.ctx=ctx
        return bacnet


    def Start(self):
        for w in self.ws.values():
            t = Thread()
            interval = w.job.Interval
            args = (interval, self.Working, (w))
            t.threading = threading.Thread(target=t.Exec, args=args)
            t.threading.start()
            self.threads.append(t)
        return

    def Restart(self):
        for t in self.threads:
            t.flag = False
            t.threading.join()

        for t in self.threads:
            t.flag = True
            t.threading.start()
        return

    def Stop(self):
        for t in self.threads:
            t.flag = False
            t.threading.join()
        return

    def Set(self, name: str, info: DeviceInfo, props: List[ConfigDeviceProperty]):
        if info.Name not in self.ws:
            raise Exception("worker not exist according to device")
        work = self.ws.get(info.Name)
        for p in props:
            for wp in work.job.Properties.values():
                if p.PropName == wp.Name:
                    work.slave.WriteValue(wp.BacnetType, wp.BacnetAddress, p.PropVal)

        return

    def Event(self, info: DeviceInfo, event: Event):
        if event.Type == "report":
            return self.PropertyGet(info,None)
        else:
            raise Exception("event type not supported yet")

    def PropertyGet(self, info: DeviceInfo, _):
        if info.Name not in self.ws:
            raise Exception("worker not exist according to device")
        self.ws.get(info.Name).Execute(False)
        return

    def Working(self, w: Worker):
        w.Execute(True)

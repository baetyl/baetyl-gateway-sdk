# !/usr/bin/env python3
from abc import ABC, abstractmethod

class IDriver(ABC):
    @abstractmethod
    def GetDriverInfo(self, data):
        pass

    @abstractmethod
    def SetConfig(self, data):
        pass

    @abstractmethod
    def Setup(self, driver, report):
        pass

    @abstractmethod
    def Start(self, data):
        pass

    @abstractmethod
    def Restart(self, data):
        pass

    @abstractmethod
    def Stop(self, data):
        pass

    @abstractmethod
    def Get(self, data):
        pass

    @abstractmethod
    def Set(self, data):
        pass
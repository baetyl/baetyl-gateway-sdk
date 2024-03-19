# !/usr/bin/env python3

from abc import ABC, abstractmethod

class IReport(ABC):
    @abstractmethod
    def Post(self, data):
        pass

    @abstractmethod
    def State(self, data):
        pass
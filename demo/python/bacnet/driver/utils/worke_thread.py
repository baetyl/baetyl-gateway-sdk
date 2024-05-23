# !/usr/bin/env python3
import threading
import sched
import time

class Thread():
    """work thread"""
    threading :threading
    flag = True

    def Exec(self, delay, func, args):
        while self.flag:
            time.sleep(delay)
            func(args)
        return

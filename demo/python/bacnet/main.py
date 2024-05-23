# !/usr/bin/env python3

from gateway_sdk.server import serve
from bacnet.driver.driver_class import Driver

if __name__ == "__main__":
    serve(Driver())

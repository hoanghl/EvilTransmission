import asyncio
import os
import os.path as osp
import sys
import traceback

import yaml
from loguru import logger

from .utils import *

## =======================================================
with open("configs.yaml", "r", encoding="utf-8") as f:
    try:
        configs = yaml.safe_load(f)
    except yaml.YAMLError as exc:
        print(exc)
        sys.exit()

HOST = configs["SERVER"]["HOST"]
PORT = configs["SERVER"]["PORT"]

MAX_SIZE_PACKET = configs["MAX_SIZE_PACKET"]


## =======================================================


class PacketType:
    LEN_FLAG = configs["LEN_FLAG"]
    LEN_SIZE = configs["LEN_SIZE"]

    def __init__(self, name, code) -> None:
        self.name = name
        self.code = code


## Define Packet Types
Data = PacketType("data", 0)
Hash = PacketType("hash", 1)
Response = PacketType("response", 2)

packets = [Data, Hash, Response]


class PacketTypes:
    names = set(x.name for x in packets)
    codes = set(x.code for x in packets)
    map_code2name = {x.code: x.name for x in packets}
    map_name2code = {x.name: x.code for x in packets}


## =======================================================
## Define Response Code
class ResponseType:
    def __init__(self, name, code) -> None:
        self.name = name
        self.code = code


InternalErr = ResponseType("internal_err", 0)
Permitted = ResponseType("permitted", 1)
NotPermitted = ResponseType("not_permitted", 2)

responses = [InternalErr, Permitted, NotPermitted]


class ResponseTypes:
    names = set(x.name for x in responses)
    codes = set(x.code for x in responses)
    map_code2name = {x.code: x.name for x in responses}
    map_name2code = {x.name: x.code for x in responses}


## =======================================================
## Define filters
packet_type_filter = int("0b00000110", 2)
response_type_filter = int("0b00011000", 2)
sucessive_filter = int("0b00000001", 2)


## =======================================================
## Define shifts
class Shifts:
    shifts = {"PacketType": 1, "ResponseType": 3}

    def __class_getitem__(cls, key):
        name = key.__class__.__name__

        if name in Shifts.shifts:
            return Shifts.shifts[name]

        name = key.__name__
        if name in Shifts.shifts:
            return Shifts.shifts[name]

        return None

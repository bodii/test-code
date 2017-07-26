#!/usr/bin/env python3
# -*- coding:utf-8 -*-


import time
import os

mod_time = os.path.getmtime('./')
print(time.ctime(mod_time))


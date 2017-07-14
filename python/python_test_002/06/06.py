#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import sys
import os

if sys.platform == 'win32':
    print('Running on a windows platform')
    command = 'cmd.exe'

if sys.platform == 'linux':
    print('Running Linux')
    command = 'uname -a'

os.system(command)

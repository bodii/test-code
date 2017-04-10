#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import os, sys


if sys.platform == 'win32':
    print('Running on a windows palatform')
    command = 'c:\\winnt\\system32\\cmd.exe'
    params = []

if sys.platform == 'linux':
    print('Running on a Linux system, identified by %s' % sys.platform)
    command = '/bin/uname'
    params = ['uname', '-a']

print(sys.platform)

print('Running %s' % command)
os.spawnv(os.P_WAIT, command, params)



#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import shutil
import os

a = open('a1.txt', 'w')
a.write('hello, python\n')
a.close()

if os.path.isfile('a1.txt'): shutil.copy('a1.txt', 'a1.backup.txt')

# 改变文件权限
os.chmod('a1.backup.txt', 777)

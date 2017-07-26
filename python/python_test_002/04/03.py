#!/usr/bin/env python3
# -*- coding:utf-8 -*-


import shutil

shutil.copy('02.py', '02.py.backup')
shutil.move('02.py.backup', '02.py.backup.new')

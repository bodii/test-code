#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import shutil
import os

shutil.copy('02.py', '02.backup.py')

# 删除文件
os.unlink('02.backup.py')

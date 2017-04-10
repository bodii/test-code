#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import os

# mkdir只能创建父级目录

try:
    os.mkdir('new_dir1/new_dir')
except Exception as e:
    print(e)
os.makedirs('new_dir1/new_dir')

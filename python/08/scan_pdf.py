#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 查找指定目录下的pdf文件 '''

import os, os.path
import re


def print_pdf(root, dirs, files):
    for file in files:
        path = os.path.join(root, file)
        path = os.path.normcase(path)
        if re.search(r'.*\.pdf', path):
            print(path)


for root, dirs, files in os.walk('.'):
    print_pdf(root, dirs, files)

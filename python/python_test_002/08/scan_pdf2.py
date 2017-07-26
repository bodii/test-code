#!/usr/bin/env python3
# -*- coding:utf-8 -*-


import os, os.path
import re



def print_pdf(arg, dir, files):
    for file in files:
        path = os.path.join(dir, file)
        path = os.path.normcase(path)
        if not re.search(r".*\.pdf", path): continue
        if re.search(r" ", path): continue
        
        print(path)


for root, dirs, files in os.walk('/home/wong/'):
    print_pdf(root, dirs, files)

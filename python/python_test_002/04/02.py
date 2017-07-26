#!/usr/bin/env python3
# -*- coding:utf-8 -*-


import time
import os

def print_dir_info(dir_path):
    for name in os.listdir(dir_path):
        full_path = os.path.join(dir_path, name)
        file_size = os.path.getsize(full_path)
        mod_time = time.ctime(os.path.getmtime(full_path))
        print("%-12s: %8d bytes, modified %s" % (name, file_size, mod_time))


if __name__ == '__main__':
    print_dir_info('./')

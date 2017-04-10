#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import os

def print_dir(dir_path):
    for name in os.listdir(dir_path):
        print(os.path.join(dir_path,name))


if __name__ == "__main__":
    print_dir("./")

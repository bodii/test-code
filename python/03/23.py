#!/usr/bin/env python3
# -*- coding:utf-8 -*-


import os

def print_path(dir_path):
    if os.path.isdir(dir_path) :
        dir_abs_path = os.path.abspath(dir_path)
        for name in os.listdir(dir_path):
            print(os.path.join(dir_abs_path, name))
    else:
        print('not\'s directory')


if __name__ == '__main__':
    print_path('./')

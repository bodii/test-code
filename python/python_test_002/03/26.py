#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import os


def print_tree(dir_path):
    """ 递归遍历目录 """
    if os.path.isdir(dir_path):
        for name in sorted(os.listdir(dir_path)):
            if os.path.isdir(os.path.join(dir_path, name)):
                print_tree(os.path.join(dir_path, name))
            elif os.path.isfile(name):
                print(os.path.abspath(os.path.join(dir_path, name)))
            else:
                print(dir_path)
    else:
        print(dir_path)


if __name__ == '__main__':
    print_tree('../')

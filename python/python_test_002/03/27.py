#!/usr/bin/env python3
# -*- coding:utf-8 -*-


import os

def print_directory_tree(dir_path):

    if os.path.isdir(dir_path):
        for name in sorted(os.listdir(dir_path)):
            path = os.path.join(dir_path, name)
            if os.path.isdir(path):
                print(path)
                print_directory_tree(path)
            elif os.path.isfile(path):
                print(path)

    else:
        print(dir_path)


if __name__ == '__main__':
    print_directory_tree('../')

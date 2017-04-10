#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import os.path

def split_fully(path):
    parent_path, name = os.path.split(path)

    if name == "":
        return (parent_path, )
    else:
        return split_fully(parent_path) + (name, )


if __name__ == "__main__":
    print(split_fully('/usr/bin/python'))


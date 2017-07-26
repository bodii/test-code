#!/usr/bin/env python3
# -*- coding:utf-8 -*-


import os
import shutil

def make_version_path(path, version):
    if version == 0:
        return path
    else:
        return path + "." + str(version)


def rotate(path, version=0):
    old_path = make_version_path(path, version)
    if not os.path.exists(old_path):
        raise IOError("'%s' doesn't exist path" % path)

    new_path = make_version_path(path, version +1)
    if os.path.exists(new_path):
        rotate(path, version +1)
    shutil.move(old_path, new_path)

def rotate_log_file(path):
    if not os.path.exists(path):
        # The file is missing, so create it.
        new_file = file(path, 'w')
        # close the new file immediately, which leaves it empty.
        del new_file

    # now rotate it.
    rotate(path)

if __name__ == '__main__':
    rotate_log_file('08.py')

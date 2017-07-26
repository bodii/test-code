#!/usr/bin/env python3
# -*- coding:utf-8 -*-


import os
import shutil

''' 
文件日志管理任务

为了给文件的新版本预留空间，这些老版本被轮换：
目前的版本web.log变成web.log.1,而web.log.1则变成web.log.2，依此类推。

用到递归函数
'''

def make_version_path(path, version):
    if version == 0 :
        # No suffix for version 0, the current version.
        return path
    else:
        # Append a suffix to indicate the older version.
        return path + "." + str(version)

def rotate(path, version=0):
    # Construct the name of the version we're rotating.
    old_path = make_version_path(path, version)

    if not os.path.exists(old_path):
        # It doesnit exist, so complain.
        raise IOError("'%s' doesn't exist" % path)
    # Construct the new version name for this file.
    new_path = make_version_path(path, version + 1)
    # Is there already a version with this name?
    if os.path.exists(new_path):
        # Yes. Rotate it out of the way first!
        rotate(path, version + 1)
    # Now we can rename the version safely.
    shutil.move(old_path, new_path)

    

if __name__ == '__main__':
    shutil.copy('08.py', '08.0.py')
    rotate('08.py')

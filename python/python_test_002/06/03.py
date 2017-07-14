#!/usr/bin/env python3
# -*- coding:utf-8 -*-


# 使用一个以上的进程
import os

pid = os.fork()
if pid == 0 : # This is the child
    print("this is the child")

else:
    print("the child is pid %d" % pid)

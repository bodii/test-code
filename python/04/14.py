#!/usr/bin/env python3
# -*- coding:utf-8 -*-


import glob

# 通配符
# 对文件名称模式中展开通配符，如 *.py
# 以列表的形式查找后目录下匹配的文件返回值

p2 = glob.glob('*2.p*')
print(p2)

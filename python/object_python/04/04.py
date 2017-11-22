#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 管理上下文和with语句 '''

import gzip
import csv

with open('subset.csv', 'w') as target:
    wtr = csv.writer( target )

    with gzip.open(path, 'r') as source:
        line_iter = (b.decode() for b in source)
        match_iter = (format_1_pat.match( line ) for line in line_iter)
        wtr.writerows( (m.groups() for m in match_iter if m is not None) )

"""
这里使用了两个上下文管理器。
外部的上下文以with open()...作为起始，使用Python中的open()函数打开一个文件，赋值给
target变量以备后用。
内部的上下文以with gzip.open(path, 'r')...为起始。gzip.opne()函数和open函数
的行为是类似的，也是打开一个文件赋值给一个上下文管理器。
当with语句结束，上下文也会相应终止并关闭引用的文件。即使当with上下文中有异常抛出，
上下文管理器也会正常终止并相应关闭所引用的文件。
"""


"""
总是使用with语句来操作文件
既然文件属于系统资源。当应用程序不再使用系统资源时，需要及时释放，这点是重要的。
with语句可以确保资源被正确释放。
"""


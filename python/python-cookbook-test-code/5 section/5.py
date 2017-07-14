#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第五章：文件与 IO
Desc: 所有程序都要处理输入和输出。这一章将涵盖处理不同类型的文件，包括文本和二
进制文件，文件编码和其他相关的内容。对文件名和目录的操作也会涉及到。

Title;       文件不存在才能写入
Issue:你想像一个文件中写入数据，但是前提必须是这个文件在文件系统上不存在。也就
是不允许覆盖已存在的文件内容。
Answer: 可以在 open() 函数中使用 x 模式来代替 w 模式的方法来解决这个问题
"""

with open('../test file/somefile', 'wt') as f:
    f.write('Hello\n')

'''
with open('../test file/somefile', 'xt') as f:
    f.writ('Hello\n')

Traceback (most recent call last):
  File python cookbook test code/5 section/5.py", line 18, in <module>
    with open('../test file/somefile', 'xt') as f:
FileExistsError: [Errno 17] File exists: '../test file/somefile
'''

'''
如果文件是二进制的，使用 xb 来代替 xt
'''


'''
这一小节演示了在写文件时通常会遇到的一个问题的完美解决方案 (不小心覆盖一
个已存在的文件)。一个替代方案是先测试这个文件是否存在，
'''
import os
if not os.path.exists('../test file/somefile'):
    with open('somefile', 'wt') as f:
        f.write('Hello\n')
else:
    print('File already exists!')

# File already exists!

'''
显而易见，使用 x 文件模式更加简单。要注意的是 x 模式是一个 Python3 对
open() 函数特有的扩展。在 Python 的旧版本或者是 Python 实现的底层 C 函数库中
都是没有这个模式的。
'''

#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第五章：文件与 IO
Desc: 所有程序都要处理输入和输出。这一章将涵盖处理不同类型的文件，包括文本和二
进制文件，文件编码和其他相关的内容。对文件名和目录的操作也会涉及到。

Title;       打印输出至文件中
Issue:你想将 print() 函数的输出重定向到一个文件中去。
Answer: 在 print() 函数中指定 file 关键字参数
"""

with open('../test file/test.txt', 'wt') as f:
    print('Hello World!', file=f)

'''
关于输出重定向到文件中就这些了。但是有一点要注意的就是文件必须是以文本模
式打开。如果文件是二进制模式的话，打印就会出错。
'''
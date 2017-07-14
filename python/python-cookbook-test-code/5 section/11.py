#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第五章：文件与 IO
Desc: 所有程序都要处理输入和输出。这一章将涵盖处理不同类型的文件，包括文本和二
进制文件，文件编码和其他相关的内容。对文件名和目录的操作也会涉及到。

Title;         文件路径名的操作
Issue:你需要使用路径名来获取文件名，目录名，绝对路径等等。
Answer: 使用 os.path 模块中的函数来操作路径名。
"""
import os
path = '../test file/data'

print(os.path.basename(path))
# data

print(os.path.dirname(path))
# ../test file

print(os.path.join('tmp', 'data', os.path.basename(path)))
# tmp\data\data

path = '~/test file/data'
print(os.path.expanduser(path))
# C:\Users\wang1/test file/data
print(os.path.splitext(path))
# ('~/test file/data', '')

'''
对于任何的文件名的操作，你都应该使用 os.path 模块，而不是使用标准字符串
操作来构造自己的代码。特别是为了可移植性考虑的时候更应如此，因为 os.path 模
块知道 Unix 和 Windows 系统之间的差异并且能够可靠地处理类似 Data/data.csv 和
Data\data.csv 这样的文件名。其次，你真的不应该浪费时间去重复造轮子。通常最
好是直接使用已经为你准备好的功能。
要注意的是 os.path 还有更多的功能在这里并没有列举出来。可以查阅官方文档
来获取更多与文件测试，符号链接等相关的函数说明。
'''


#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第五章：文件与 IO
Desc: 所有程序都要处理输入和输出。这一章将涵盖处理不同类型的文件，包括文本和二
进制文件，文件编码和其他相关的内容。对文件名和目录的操作也会涉及到。

Title;         测试文件是否存在
Issue:你想测试一个文件或目录是否存在。
Answer: 使用 os.path 模块来测试一个文件或目录是否存在。
"""

import os
print(os.path.exists('../test file/data'))
# True
print(os.path.exists('../test file/data1'))
# False

'''
你还能进一步测试这个文件时什么类型的。在下面这些测试中，如果测试的文件不
存在的时候，结果都会返回 False：
'''
print(os.path.isfile('../test file/data'))
# True

print(os.path.isdir('../test'))
# False

print(os.path.islink('../test file/'))
# False

print(os.path.realpath('../test file/'))
# D:\..\python cookbook test code\test file

print(os.path.getsize('../test file/data'))
# 1000000

print(os.path.getmtime('../test file/data'))
# 1460454662.5821352

import time
print(time.ctime(os.path.getmtime('../test file/data')))
# Tue Apr 12 17:51:02 2016

'''
使用 os.path 来进行文件测试是很简单的。在写这些脚本时，可能唯一需要注意
的就是你需要考虑文件权限的问题，特别是在获取元数据时候。
'''
'''
>>> os.path.getsize('/Users/guido/Desktop/foo.txt')
Traceback (most recent call last):
File "<stdin>", line 1, in <module>
File "/usr/local/lib/python3.3/genericpath.py", line 49, in getsize
return os.stat(filename).st_size
PermissionError: [Errno 13] Permission denied: '/Users/guido/Desktop/foo.txt'
>>>
'''
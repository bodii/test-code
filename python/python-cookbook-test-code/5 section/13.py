#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第五章：文件与 IO
Desc: 所有程序都要处理输入和输出。这一章将涵盖处理不同类型的文件，包括文本和二
进制文件，文件编码和其他相关的内容。对文件名和目录的操作也会涉及到。

Title;         获取文件夹中的文件列表
Issue:你想获取文件系统中某个目录下的所有文件列表
Answer: 使用 os.listdir() 函数来获取某个目录中的文件列表
"""

import os
names = os.listdir('../test file/')

'''
结果会返回目录中所有文件列表，包括所有文件，子目录，符号链接等等。如果你
需要通过某种方式过滤数据，可以考虑结合 os.path 库中的一些函数来使用列表推导。
'''
import os
names = [name for name in os.listdir('../test file/')
         if os.path.isfile(os.path.join('somedir', name))]

dirnames = [name for name in os.listdir('../test file/')
            if os.path.isdir(os.path.join('somedir', name))]

'''
字符串的 startswith() 和 endswith() 方法对于过滤一个目录的内容也是很有用
的。
'''
pyfiles = [name for name in os.listdir('./')
           if name.endswith('.py')]

'''
对于文件名的匹配，你可能会考虑使用 glob 或 fnmatch 模块。
'''
import glob
pyfiles = glob.glob('./*.py')

from fnmatch import fnmatch
pyfiles = [name for name in os.listdir('./')
           if fnmatch(name, '*.py')]

'''
获取目录中的列表是很容易的，但是其返回结果只是目录中实体名列表而已。如
果你还想获取其他的元信息，比如文件大小，修改时间等等，你或许还需要使用到
os.path 模块中的函数或着 os.stat() 函数来收集数据。
'''

import os
import glob
import datetime

pyfiles = glob.glob('*.py')
name_sz_date = [(name, os.path.getsize(name), os.path.getmtime(name))
                for name in pyfiles]

for name, size, mtime in name_sz_date:
    print('filename:%s  filesize:%i  filemtime:%f' % (name, size, mtime))

'''
最后还有一点要注意的就是，有时候在处理文件名编码问题时候可能会出现一些问
题。通常来讲，函数 os.listdir() 返回的实体列表会根据系统默认的文件名编码来
解码。但是有时候也会碰到一些不能正常解码的文件名。关于文件名的处理问题，在
5.14 和 5.15 小节有更详细的讲解。
'''

#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第五章：文件与 IO
Desc: 所有程序都要处理输入和输出。这一章将涵盖处理不同类型的文件，包括文本和二
进制文件，文件编码和其他相关的内容。对文件名和目录的操作也会涉及到。

Title;         忽略文件名编码
Issue:你想使用原始文件名执行文件的 I/O 操作，也就是说文件名并没有经过系统默认
编码去解码或编码过。
Answer: 默认情况下，所有的文件名都会根据 sys.getfilesystemencoding() 返回的文本
编码来编码或解码。
"""
import sys

print(sys.getfilesystemencoding())
# mbcs

'''
如果因为某种原因你想忽略这种编码，可以使用一个原始字节字符串来指定一个文
件名即可。
'''
with open('../test file/\xf1o.txt', 'w') as f:
    f.write('Spicy!')

import os

print(os.listdir('../test file/'))

print(os.listdir(b'../test file/'))
'''
#with open(b'../test file/?o.txt', 'rb') as f:
   # print(f.read())

>>> # Open file with raw filename
>>> with open(b'jalapen\xcc\x83o.txt') as f:
... print(f.read())
...
Spicy!
>>>
'''


'''
正如你所见，在最后两个操作中，当你给文件相关函数如 open() 和 os.listdir()
传递字节字符串时，文件名的处理方式会稍有不同。
7.14.3 讨论
通常来讲，你不需要担心文件名的编码和解码，普通的文件名操作应该就没问题
了。但是，有些操作系统允许用户通过偶然或恶意方式去创建名字不符合默认编码的
文件。这些文件名可能会神秘地中断那些需要处理大量文件的 Python 程序。
读取目录并通过原始未解码方式处理文件名可以有效的避免这样的问题，尽管这样
会带来一定的编程难度。
关于打印不可解码的文件名，请参考 5.15 小节
'''
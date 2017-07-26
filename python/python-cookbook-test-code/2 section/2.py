#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第二章：字符串和文本
Desc: 几乎所有有用的程序都会涉及到某些文本处理，不管是解析数据还是产生输出。这
一章将重点关注文本的操作处理，比如提取字符串，搜索，替换以及解析等。大部分
的问题都能简单的调用字符串的内建方法完成。但是，一些更为复杂的操作可能需要
正则表达式或者强大的解析器，所有这些主题我们都会详细讲解。并且在操作 Unicode
时候碰到的一些棘手的问题在这里也会被提及到。

Title;        字符串开头或结尾匹配
Issue: 你需要通过指定的文本模式去检查字符串的开头或者结尾，比如文件名后缀，URL
Scheme 等等。
Answer: 检 查 字 符 串 开 头 或 结 尾 的 一 个 简 单 方 法 是 使 用 str.startswith() 或 者 是
str.endswith() 方法。
"""

filename = '../test file/somefile.txt'
print(filename.endswith('.txt')) # True
print(filename.startswith('file:')) # False

url = 'http://www.python.org'
print(url.startswith('http:')) # True

'''
如果你想检查多种匹配可能，只需要将所有的匹配项放入到一个元组中去，然后传
给 startswith() 或者 endswith() 方法：
'''
import os
filenames = os.listdir('.')
print(filenames)
#  ['1.py', '2.py']
print([name for name in filenames if name.endswith(('.py', '.txt'))])
# ['1.py', '2.py']
print(any(name.endswith('.py') for name in filenames))  # True
print(all(name.endswith('.py') for name in filenames)) # True

from urllib.request import urlopen

def read_data(name):
    if name.startswith(('http:', 'https:', 'ftp:')):
        return urlopen(name).read()
    else:
        with open(name) as f:
            return f.read()

'''
奇怪的是，这个方法中必须要输入一个元组作为参数。如果你恰巧有一个 list 或
者 set 类型的选择项，要确保传递参数前先调用 tuple() 将其转换为元组类型
'''
choices = ['http:', 'ftp:']
url = 'http://www.python.org'
'''
print(url.startswith(choices))

Traceback (most recent call last):
File "<stdin>", line 1, in <module>
TypeError: startswith first arg must be str or a tuple of str, not list
'''
print(url.startswith(tuple(choices))) # True

'''
startswith() 和 endswith() 方法提供了一个非常方便的方式去做字符串开头和
结尾的检查。类似的操作也可以使用切片来实现，但是代码看起来没有那么优雅。
'''
filename = 'spam.txt'
print(filename[-4:]  == '.txt') # True
url = 'http://www.python.org'
print(url[:5] == 'http:' or url[:6] == 'https:' or url[:4] == 'ftp:') # True

'''
你可以能还想使用正则表达式去实现
'''
import re
url = 'http://www.python.org'
print(re.match('http:|https:|ftp:',url))
# <_sre.SRE_Match object; span=(0, 5), match='http:'>
'''
这种方式也行得通，但是对于简单的匹配实在是有点小材大用了，本节中的方法更
加简单并且运行会更快些。
'''


'''
最后提一下，当和其他操作比如普通数据聚合相结合的时候 startswith() 和
endswith() 方法是很不错的。比如：
if any(name.endswith(('.c','.h')) for name in listdir(dirname)):
    ...
'''

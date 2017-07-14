#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第二章：字符串和文本
Desc: 几乎所有有用的程序都会涉及到某些文本处理，不管是解析数据还是产生输出。这
一章将重点关注文本的操作处理，比如提取字符串，搜索，替换以及解析等。大部分
的问题都能简单的调用字符串的内建方法完成。但是，一些更为复杂的操作可能需要
正则表达式或者强大的解析器，所有这些主题我们都会详细讲解。并且在操作 Unicode
时候碰到的一些棘手的问题在这里也会被提及到。

Title;         用 Shell 通配符匹配字符串
Issue:  你想使用 Unix Shell 中常用的通配符 (比如 *.py , Dat[0-9]*.csv 等) 去匹配文
本字符串
Answer: fnmatch 模块提供了两个函数—— fnmatch() 和 fnmatchcase() ，可以用来实现
这样的匹配。
"""
from fnmatch import fnmatch, fnmatchcase
print(fnmatch('foo.txt', '*.txt')) # True
print(fnmatch('foo.txt','?oo.txt')) # True
print(fnmatch('Dat45.csv', 'Dat[0-9]*')) # True

names = ['Dat1.csv', 'Dat2.csv', 'config.ini', 'foo.py']
print([name for name in names if fnmatch(name, 'Dat*.csv')])
# ['Dat1.csv', 'Dat2.csv']

'''
fnmatch() 函数使用底层操作系统的大小写敏感规则 (不同的系统是不一样的) 来
匹配模式

>>> # On OS X (Mac)
>>> fnmatch('foo.txt', '*.TXT')
False
>>> # On Windows
>>> fnmatch('foo.txt', '*.TXT')
True

如果你对这个区别很在意，可以使用 fnmatchcase() 来代替。它完全使用你的模
式大小写匹配。比如：
>>> fnmatchcase('foo.txt', '*.TXT')
False
'''

'''
这两个函数通常会被忽略的一个特性是在处理非文件名的字符串时候它们也是很有
用的。
'''
addresses = [
    '5412 N CLARK ST',
    '1060 W ADDISON ST',
    '1039 W GRANVILLE AVE',
    '2122 N CLARK ST',
    '4802 N BROADWAY',
]
print([addr for addr in addresses if fnmatchcase(addr, '* ST')])
# ['5412 N CLARK ST', '1060 W ADDISON ST', '2122 N CLARK ST']

print([addr for addr in addresses if fnmatchcase(addr, '54[0-9][0-9] *CLARK*')])
# ['5412 N CLARK ST']

'''
fnmatch() 函数匹配能力介于简单的字符串方法和强大的正则表达式之间。如果在
数据处理操作中只需要简单的通配符就能完成的时候，这通常是一个比较合理的方案。
如果你的代码需要做文件名的匹配，最好使用 glob 模块
'''
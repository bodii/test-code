#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第八章：类与对象
Desc: 本章主要关注点的是和类定义有关的常见编程模型。包括让对象支持常见的
Python 特性、特殊方法的使用、类封装技术、继承、内存管理以及有用的设计模
式。

Title;        自定义字符串的格式化
Issue:你想通过 format() 函数和字符串方法使得一个对象能支持自定义的格式化
Answer: 为了自定义字符串的格式化，我们需要在类上面定义 format () 方法。
"""

_formats = {
    'ymd' : '{d.year}-{d.month}-{d.day}',
    'mdy' : '{d.month}/{d.day}/{d.year}',
    'dmy' : '{d.day}/{d.month}/{d.year}'
}

class Date:
    def __init__(self, year, month, day):
        self.year = year
        self.month = month
        self.day = day

    def __format__(self, code):
        if code == '':
            code = 'ymd'
        fmt = _formats[code]
        return fmt.format(d=self)

d = Date(2012, 12, 21)
print(format(d))
# 2012-12-21
print(format(d, 'mdy'))
# 12/21/2012
print('The date is {:ymd}'.format(d))
# The date is 2012-12-21
print('The date is {:mdy}'.format(d))
# The date is 12/21/2012

"""
format () 方法给 Python 的字符串格式化功能提供了一个钩子。这里需要着重
强调的是格式化代码的解析工作完全由类自己决定。因此，格式化代码可以是任何值。
"""

from datetime import date

d = date(2012, 12, 21)
print(format(d))
# 2012-12-21
print(format(d, '%A, %B, %d, %Y'))
# Friday, December, 21, 2012
print('The end is {:%d %b %Y}. Goodbye'.format(d))
# The end is 21 Dec 2012. Goodbye

"""
对于内置类型的格式化有一些标准的约定。可以参考 string 模块文档 说明。
"""
#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第八章：类与对象
Desc: 本章主要关注点的是和类定义有关的常见编程模型。包括让对象支持常见的
Python 特性、特殊方法的使用、类封装技术、继承、内存管理以及有用的设计模
式。

Title;            创建不调用 init 方法的实例
Issue:你想创建一个实例，但是希望绕过执行 init () 方法
Answer:可以通过 new () 方法创建一个未初始化的实例。
"""

class Date:
    def __init__(self, year, month, day):
        self.year = year
        self.month = month
        self.day = day

"""
下面演示如何不调用 init () 方法来创建这个 Date 实例
"""
d = Date.__new__(Date)
print(d)
# <__main__.Date object at 0x000001D618A92D30>
# print(d.year)
# AttributeError: 'Date' object has no attribute 'year'

data = {'year':2012, 'month':8, 'day':29}
for key, value in data.items():
    setattr(d, key, value)

print(d.year)
# 2012
print(d.month)
# 8
print(d.day)
# 29

"""
当我们在反序列对象或者实现某个类方法构造函数时需要绕过 init () 方法来
创建对象。例如，对于上面的 Date 来来讲，有时候你可能会像下面这样定义一个新的
构造函数 today()
"""
from time import localtime

class Date:
    def __init__(self, year, month, day):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def today(cls):
        d = cls.__new__(cls)
        t = localtime()
        d.year = t.tm_year
        d.month = t.tm_mon
        d.day = t.tm_mday
        return d

d = Date(2012, 8, 29)
b = Date.today()
print(b.day)
# 14

"""
同样，在你反序列化 JSON 数据时产生一个如下的字典对象：
"""
data = {'year':2012, 'month':8, 'day':29}
"""
如果你想将它转换成一个 Date 类型实例，可以使用上面的技术。
当你通过这种非常规方式来创建实例的时候，最好不要直接去访问底层实例字
典，除非你真的清楚所有细节。否则的话，如果这个类使用了 slots 、properties 、
descriptors 或其他高级技术的时候代码就会失效。而这时候使用 setattr() 方法会让
你的代码变得更加通用。
"""
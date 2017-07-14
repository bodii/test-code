#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第八章：类与对象
Desc: 本章主要关注点的是和类定义有关的常见编程模型。包括让对象支持常见的
Python 特性、特殊方法的使用、类封装技术、继承、内存管理以及有用的设计模
式。

Title;           在类中定义多个构造器
Issue:你想实现一个类，除了使用 init () 方法外，还有其他方式可以初始化它
Answer:为了实现多个构造器，你需要使用到类方法。
"""
import time

class Date:
    """方法一：使用类方法"""
    def __init__(self, year, month, day):
        self.year = year
        self.month = month
        self.day = day

    @classmethod
    def today(cls):
        t = time.localtime()
        print(t.tm_year, t.tm_mon, t.tm_mday)
        return cls(t.tm_year, t.tm_mon, t.tm_mday)


"""
直接调用类方法即可，下面是使用示例
"""

a = Date(2012, 12, 21)
b = Date.today()

"""
类方法的一个主要用途就是定义多个构造器。它接受一个 class 作为第一个参数
(cls)。你应该注意到了这个类被用来创建并返回最终的实例。在继承时也能工作的很
好
"""

class NewDate(Date):
    pass

c = Date.today()
d = NewDate.today()
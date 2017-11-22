#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 其他的一些抽象基类 '''

"""
 [ 迭代器的抽象基类 ]

 在for语句中使用一个可迭代的容器时， Python会隐式地创建一个迭代器。
 """

x = [1, 2, 3]
print( iter(x) )
# <list_iterator object at 0x7f6f68595c18>

x_iter = iter(x)
print( next(x_iter) )
# 1
print( next(x_iter) )
# 2
print( next(x_iter) )
# 3
# print( next(x_iter) )
# StopIteration

import collections
print( isinstance(x_iter, collections.abc.Iterator) )
# True





#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第一章：数据结构和算法
Desc: Python 提供了大量的内置数据结构，包括列表，集合以及字典。大多数情况下使
用这些数据结构是很简单的。但是，我们也会经常碰到到诸如查询，排序和过滤等等
这些普遍存在的问题。因此，这一章的目的就是讨论这些比较常见的问题和算法。另
外，我们也会给出在集合模块 collections 当中操作这些数据结构的方法。

Title:  解压序列赋值给多个变量
Issue: 现在有一个包含 N 个元素的元组或者是序列，怎样将它里面的值解压后同时赋值
给 N 个变量？
answer: 任何的序列 (或者是可迭代对象) 可以通过一个简单的赋值语句解压并赋值给多个
变量。唯一的前提就是变量的数量必须跟序列元素的数量是一样的。
"""

#tuple
p = (4, 5)
x, y = p
print(x) # 4
print(y) # 5

#list
data = ['ABCD', 11, 2503.44, (2016, 3, 26)]
name, shares, price, date = data
print(name) # ABCD
print(shares) # 11
print(price) # 2503.44
print(date) # (2016, 3, 26)

name, shares, price,(year, mon, day) = data
print(name) # ABCD
print(shares) # 11
print(price) # 2503.44
print(year) # 2016
print(mon) # 3
print(day) # 26

#str
'''
实际上，这种解压赋值可以用在任何可迭代对象上面，而不仅仅是列表或者元组。
包括字符串，文件对象，迭代器和生成器。
'''
s = 'Hello'
a, b, c, d, e = s
print(a, b, c, d, e)  # H e l l o

'''
有时候，你可能只想解压一部分，丢弃其他的值。对于这种情况 Python 并没有提
供特殊的语法。但是你可以使用任意变量名去占位，到时候丢掉这些变量就行了。
'''
data = ['ABCD', 11, 2503.44, (2016, 3, 26)]
_, shares, price, _, = data
print(shares) # 11
print(price) # 2503.44
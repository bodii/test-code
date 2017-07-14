#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第一章：数据结构和算法
Desc: Python 提供了大量的内置数据结构，包括列表，集合以及字典。大多数情况下使
用这些数据结构是很简单的。但是，我们也会经常碰到到诸如查询，排序和过滤等等
这些普遍存在的问题。因此，这一章的目的就是讨论这些比较常见的问题和算法。另
外，我们也会给出在集合模块 collections 当中操作这些数据结构的方法。

Title:       命名切片
"""

record = '....................100 .......513.25 ..........'
cost = int(record[20:23]) * float(record[31:37])
print(cost) # 51325.0

'''
命名切片
'''

SHARES = slice(20,23)
PRICE  = slice(31,37)
cost = int(record[SHARES]) * float(record[PRICE])
print(cost) # 51325.0

'''
内置的 slice() 函数创建了一个切片对象，可以被用在任何切片允许使用的地方。
'''

items = [0, 1, 2, 3, 4, 5, 6]
a = slice(2,4)
print(items[2:4]) # [2, 3]
print(items[a]) # [2, 3]
items[a] = [10, 11]
print(items) # [0, 1, 10, 11, 4, 5, 6]
del items[a]
print(items) # [0, 1, 4, 5, 6]

'''
如果你有一个切片对象 s，你可以分别调用它的 s.start , s.stop , s.step 属性来
获取更多的信息。比如：
'''

s = slice(5, 50, 2)
print(s.start) # 5
print(s.stop) # 50
print(s.step) # 2

'''
另外，你还能通过调用切片的 indices(size) 方法将它映射到一个确定大小的序
列上，这个方法返回一个三元组 (start, stop, step) ，所有值都会被合适的缩小以
满足边界限制，从而使用的时候避免出现 IndexError 异常。
'''

s1 = 'HelloWorld'
print(*a.indices(len(s1))) # 2 4 1
for i in range(*a.indices(len(s1))):
    print(s1[i]) # l  l
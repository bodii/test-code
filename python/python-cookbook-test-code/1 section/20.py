#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第一章：数据结构和算法
Desc: Python 提供了大量的内置数据结构，包括列表，集合以及字典。大多数情况下使
用这些数据结构是很简单的。但是，我们也会经常碰到到诸如查询，排序和过滤等等
这些普遍存在的问题。因此，这一章的目的就是讨论这些比较常见的问题和算法。另
外，我们也会给出在集合模块 collections 当中操作这些数据结构的方法。

Title:         合并多个字典或映射
Issue:  现在有多个字典或者映射，你想将它们从逻辑上合并为一个单一的映射后执行某些
操作，比如查找值或者检查某些键是否存在
代访问
answer:  一个非常简单扼解决方案就是使用 collections 模块中的 ChainMap 类
"""

a = {'x' : 1, 'z' : 3}
b = {'y' : 2, 'z' : 4}

'''
现在假设你必须在两个字典中执行查找操作 (比如先从 a 中找，如果找不到再在 b
中找)。
'''
from collections import ChainMap
c = ChainMap(a, b)
print(c['x']) # 1
print(c['y']) # 2
print(c['z']) # 3
print(len(c)) # 3

print(list(c.keys())) # ['x', 'z', 'y']
print(list(c.values())) # [2, 1, 3]

'''
如果出现重复键，那么第一次出现的映射值会被返回。因此，例子程序中的 c['z']
总是会返回字典 a 中对应的值，而不是 b 中对应的值。
'''

'''
对于字典的更新或删除操作总是影响的是列表中第一个字典
'''
c['z'] = 10
c['w'] = 40
del c['x']
print(a) # {'w': 40, 'z': 10}
'''
del c['y']
print(b)

    del self.maps[0][key]
KeyError: 'y'

During handling of the above exception, another exception occurred:

Traceback (most recent call last):
  File "/python cookbook test code/1 section/20.py", line 48, in <module>
    del c['y']
  File "\Python\Python35\lib\collections\__init__.py", line 914, in __delitem__
    raise KeyError('Key not found in the first mapping: {!r}'.format(key))
KeyError: "Key not found in the first mapping: 'y'"
'''

'''
ChainMap 对于编程语言中的作用范围变量 (比如 globals , locals 等) 是非常有用
的。事实上，有一些方法可以使它变得简单
'''
values = ChainMap()
values['x'] = 1
values = values.new_child()
values['x'] = 2
values = values.new_child()
values['x'] = 3
print(values)
# ChainMap({'x': 3}, {'x': 2}, {'x': 1})

print(values['x']) # 3

values = values.parents
print(values['x']) # 2

values = values.parents
print(values['x']) # 1

print(values) # ChainMap({'x': 1})

'''
作为 ChainMap 的替代，你可能会考虑使用 update() 方法将两个字典合并
'''
a = {'x' : 1, 'z' : 3}
b = {'y' : 2, 'z' : 4}
merged = dict(b)
merged.update(a)
print(merged['x']) # 1
print(merged['y']) # 2
print(merged['z']) # 3

'''
这样也能行得通，但是它需要你创建一个完全不同的字典对象 (或者是破坏现有字
典结构)。同时，如果原字典做了更新，这种改变不会反应到新的合并字典中去
'''
a['x']  = 13
print(merged['x']) # 1

'''
ChianMap 使用原来的字典，它自己不创建新的字典。
'''
a = {'x' : 1, 'z' : 3}
b = {'y' : 2, 'z' : 4}
merged = ChainMap(a, b)
print(merged['x']) # 1
a['x'] = 42
print(merged['x']) # 42

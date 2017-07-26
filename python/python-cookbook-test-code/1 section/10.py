#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第一章：数据结构和算法
Desc: Python 提供了大量的内置数据结构，包括列表，集合以及字典。大多数情况下使
用这些数据结构是很简单的。但是，我们也会经常碰到到诸如查询，排序和过滤等等
这些普遍存在的问题。因此，这一章的目的就是讨论这些比较常见的问题和算法。另
外，我们也会给出在集合模块 collections 当中操作这些数据结构的方法。

Title:      删除序列相同元素并保持顺序
Issue: 怎样在一个序列上面保持元素顺序的同时消除重复的值？
answer: 如果序列上的值都是 hashable 类型，那么可以很简单的利用集合或者生成器来解
决这个问题。
"""

def dedupe(items):
    seen = set()
    for item in items:
        if item not in seen:
            yield item
            seen.add(item)


a = [1, 5, 2, 1, 9, 1, 5, 10]
print(list(dedupe(a))) # [1, 5, 2, 9, 10]

def dedupe2(items, key=None):
    seen = set()
    for item in items:
        val = item if key is None else key(item)
        if val not in seen:
            yield item
            seen.add(val)

'''
这里的 key 参数指定了一个函数，将序列元素转换成 hashable 类型。下面是它的
用法示例
'''
a = [{'x':1, 'y':2},{'x':1, 'y':3},{'x':1, 'y':2},{'x':2, 'y':4}]
print(list(dedupe2(a,key = lambda d:(d['x'],d['y']))))
#[{'x': 1, 'y': 2}, {'x': 1, 'y': 3}, {'x': 2, 'y': 4}]
print(list(dedupe2(a,key = lambda d:d['x'])))
#[{'y': 2, 'x': 1}, {'y': 4, 'x': 2}]

'''
如果你仅仅就是想消除重复元素，通常可以简单的构造一个集合。比如：
>>> a
[1, 5, 2, 1, 9, 1, 5, 10]
>>> set(a)
{1, 2, 10, 5, 9}
>>>
'''
'''
在本节中我们使用了生成器函数让我们的函数更加通用，不仅仅是局限于列表处
理。比如，如果如果你想读取一个文件，消除重复行，你可以很容易像这样做：
'''
with open('../test file/somefile.txt','r') as f:
    for line in dedupe2(f):
        print(line)
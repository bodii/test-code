#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第一章：数据结构和算法
Desc: Python 提供了大量的内置数据结构，包括列表，集合以及字典。大多数情况下使
用这些数据结构是很简单的。但是，我们也会经常碰到到诸如查询，排序和过滤等等
这些普遍存在的问题。因此，这一章的目的就是讨论这些比较常见的问题和算法。另
外，我们也会给出在集合模块 collections 当中操作这些数据结构的方法。

Title:      字典中的键映射多个值
Issue: 怎样实现一个键对应多个值的字典 (也叫 multidict )？
answer: 一个字典就是一个键对应一个单值的映射。如果你想要一个键映射多个值，那么你
就需要将这多个值放到另外的容器中，比如列表或者集合里面。
"""

d = {
    'a' : [1, 2, 3],
    'b' : [4, 5]
}

e = {
    'a' : {1, 2, 3},
    'b' : {4, 5}
}

'''
选择使用列表还是集合取决于你的实际需求。如果你想保持元素的插入顺序就应该
使用列表，如果想去掉重复元素就使用集合（并且不关心元素的顺序问题）
'''

'''
使用 collections 模块中的 defaultdict 来构造这样的字典。
defaultdict 的一个特征是它会自动初始化每个 key 刚开始对应的值，所以你只需要
关注添加元素操作了。
'''

from collections import defaultdict

d = defaultdict(list)
d['a'].append(1)
d['a'].append(2)
d['b'].append(4)
print(d) # defaultdict(<class 'list'>, {'a': [1, 2], 'b': [4]})

d = defaultdict(set)
d['a'].add(1)
d['a'].add(2)
d['d'].add(4)
print(d) # defaultdict(<class 'set'>, {'a': {1, 2}, 'd': {4}})

'''
需要注意的是， defaultdict 会自动为将要访问的键 (就算目前字典中并不存在这
样的键) 创建映射实体。如果你并不需要这样的特性，你可以在一个普通的字典上使用
setdefault() 方法来代替。
'''
d = {}
d.setdefault('a', []).append(1)
d.setdefault('a', []).append(2)
d.setdefault('b', []).append(4)
print(d)

'''
一般来讲，创建一个多值映射字典是很简单的。但是，如果你选择自己实现的话，
那么对于值的初始化可能会有点麻烦，你可能会像下面这样来实现
'''
d = {}
for key, value in pairs:
    if key not in d:
        d[key] = []
    d[key].append(value)

'''
如果使用 defaultdict 的话代码就更加简洁了
'''
d = defaultdict(list)
for key, value in pairs:
    d[key].append(value)
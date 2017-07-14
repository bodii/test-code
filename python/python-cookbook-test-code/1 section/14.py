#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第一章：数据结构和算法
Desc: Python 提供了大量的内置数据结构，包括列表，集合以及字典。大多数情况下使
用这些数据结构是很简单的。但是，我们也会经常碰到到诸如查询，排序和过滤等等
这些普遍存在的问题。因此，这一章的目的就是讨论这些比较常见的问题和算法。另
外，我们也会给出在集合模块 collections 当中操作这些数据结构的方法。

Title:        排序不支持原生比较的对象
Issue: 你想排序类型相同的对象，但是他们不支持原生的比较操作
answer:  内置的 sorted() 函数有一个关键字参数 key ，可以传入一个 callable 对象给它，
这个 callable 对象对每个传入的对象返回一个值，这个值会被 sorted 用来排序这些
对象。
"""

class User:
    def __init__(self,user_id):
        self.user_id = user_id

    def __repr__(self):
        return 'User({})'.format(self.user_id)

def sort_notcompare():
    users = [User(23), User(3),User(99)]
    print(users) # [User(23), User(3), User(99)]
    print(sorted(users, key=lambda u:u.user_id)) # [User(3), User(23), User(99)]

sort_notcompare()

'''
另一种方式是使用 operator.attrgetter()来代替lambda函数
'''
users = [User(23), User(3),User(99)]
from operator import attrgetter
print(sorted(users, key=attrgetter('user_id')))
#[User(3), User(23), User(99)]

'''
选 择 使 用 lambda 函 数 或 者 是 attrgetter() 可 能 取 决 于 个 人 喜 好。 但 是，
attrgetter() 函数通常会运行的快点，并且还能同时允许多个字段进行比较。这
个跟 operator.itemgetter() 函数作用于字典类型很类似
by_name = sorted(users, key=attrgetter('last_name', 'first_name'))
'''

'''
同样需要注意的是，这一小节用到的技术同样适用于像 min() 和 max() 之类的函
数。
'''

print(min(users, key=attrgetter('user_id')))
# User(3)
print(max(users, key=attrgetter('user_id')))
# User(99)
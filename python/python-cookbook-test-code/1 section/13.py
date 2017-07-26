#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第一章：数据结构和算法
Desc: Python 提供了大量的内置数据结构，包括列表，集合以及字典。大多数情况下使
用这些数据结构是很简单的。但是，我们也会经常碰到到诸如查询，排序和过滤等等
这些普遍存在的问题。因此，这一章的目的就是讨论这些比较常见的问题和算法。另
外，我们也会给出在集合模块 collections 当中操作这些数据结构的方法。

Title:       通过某个关键字排序一个字典列表
Issue:  你有一个字典列表，你想根据某个或某几个字典字段来排序这个列表。
answer:  通过使用 operator 模块的 itemgetter 函数，可以非常容易的排序这样的数据结
构。假设你从数据库中检索出来网站会员信息列表，并且以下列的数据结构返回
"""

rows = [
    {'fname': 'Brian', 'lname': 'Jones', 'uid': 1003},
    {'fname': 'David', 'lname': 'Beazley', 'uid': 1002},
    {'fname': 'John', 'lname': 'Cleese', 'uid': 1001},
    {'fname': 'Big', 'lname': 'Jones', 'uid': 1004}
]

from operator import itemgetter
rows_by_fname = sorted(rows, key=itemgetter('fname'))
rows_by_uid = sorted(rows, key = itemgetter('uid'))
print(rows_by_fname)
# [{'uid': 1004, 'lname': 'Jones', 'fname': 'Big'}, {'uid': 1003, 'lname': 'Jones', 'fname': 'Brian'}, {'uid': 1002, 'lname': 'Beazley', 'fname': 'David'}, {'uid': 1001, 'lname': 'Cleese', 'fname': 'John'}]
print(rows_by_uid)
# [{'uid': 1001, 'lname': 'Cleese', 'fname': 'John'}, {'uid': 1002, 'lname': 'Beazley', 'fname': 'David'}, {'uid': 1003, 'lname': 'Jones', 'fname': 'Brian'}, {'uid': 1004, 'lname': 'Jones', 'fname': 'Big'}]

'''
itemgetter() 函数也支持多个 keys
'''
rows_by_lfname = sorted(rows, key=itemgetter('lname','fname'))
print(rows_by_lfname)
# [{'uid': 1002, 'fname': 'David', 'lname': 'Beazley'}, {'uid': 1001, 'fname': 'John', 'lname': 'Cleese'}, {'uid': 1004, 'fname': 'Big', 'lname': 'Jones'}, {'uid': 1003, 'fname': 'Brian', 'lname': 'Jones'}]
'''
itemgetter() 有时候也可以用 lambda 表达式代替，
'''
rows_by_fname = sorted(rows, key = lambda x : x['fname'])
rows_by_lfname = sorted(rows, key =lambda x : (x['fname'], x['lname']))
print(rows_by_fname)
# [{'uid': 1004, 'fname': 'Big', 'lname': 'Jones'}, {'uid': 1003, 'fname': 'Brian', 'lname': 'Jones'}, {'uid': 1002, 'fname': 'David', 'lname': 'Beazley'}, {'uid': 1001, 'fname': 'John', 'lname': 'Cleese'}]
print(rows_by_lfname)
# [{'uid': 1004, 'fname': 'Big', 'lname': 'Jones'}, {'uid': 1003, 'fname': 'Brian', 'lname': 'Jones'}, {'uid': 1002, 'fname': 'David', 'lname': 'Beazley'}, {'uid': 1001, 'fname': 'John', 'lname': 'Cleese'}]

'''
这种方案也不错。但是，使用 itemgetter() 方式会运行的稍微快点。因此，如果
你对性能要求比较高的话就使用 itemgetter() 方式
不要忘了这节中展示的技术也同样适用于 min() 和 max() 等函数
'''
print(min(rows, key=itemgetter('fname')))
# {'lname': 'Jones', 'fname': 'Big', 'uid': 1004}
print(max(rows, key=itemgetter('fname','lname')))
# {'lname': 'Cleese', 'fname': 'John', 'uid': 1001}
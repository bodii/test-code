#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第一章：数据结构和算法
Desc: Python 提供了大量的内置数据结构，包括列表，集合以及字典。大多数情况下使
用这些数据结构是很简单的。但是，我们也会经常碰到到诸如查询，排序和过滤等等
这些普遍存在的问题。因此，这一章的目的就是讨论这些比较常见的问题和算法。另
外，我们也会给出在集合模块 collections 当中操作这些数据结构的方法。

Title:         通过某个字段将记录分组
Issue:  你有一个字典或者实例的序列，然后你想根据某个特定的字段比如 date 来分组迭
代访问
answer:  内itertools.groupby() 函数对于这样的数据分组操作非常实用。
"""

rows = [
    {'address': '5412 N CLARK', 'date': '07/01/2012'},
    {'address': '5148 N CLARK', 'date': '07/04/2012'},
    {'address': '5800 E 58TH', 'date': '07/02/2012'},
    {'address': '2122 N CLARK', 'date': '07/03/2012'},
    {'address': '5645 N RAVENSWOOD', 'date': '07/02/2012'},
    {'address': '1060 W ADDISON', 'date': '07/02/2012'},
    {'address': '4801 N BROADWAY', 'date': '07/01/2012'},
    {'address': '1039 W GRANVILLE', 'date': '07/04/2012'},
]

from operator import itemgetter
from itertools import groupby

rows.sort(key=itemgetter('date'))

for date, items in groupby(rows,key=itemgetter('date')):
    print(date)
    for i in items:
        print(' ',i)

"""
07/01/2012
  {'address': '5412 N CLARK', 'date': '07/01/2012'}
  {'address': '4801 N BROADWAY', 'date': '07/01/2012'}
07/02/2012
  {'address': '5800 E 58TH', 'date': '07/02/2012'}
  {'address': '5645 N RAVENSWOOD', 'date': '07/02/2012'}
  {'address': '1060 W ADDISON', 'date': '07/02/2012'}
07/03/2012
  {'address': '2122 N CLARK', 'date': '07/03/2012'}
07/04/2012
  {'address': '5148 N CLARK', 'date': '07/04/2012'}
  {'address': '1039 W GRANVILLE', 'date': '07/04/2012'}
"""

'''
groupby() 函数扫描整个序列并且查找连续相同值 (或者根据指定 key 函数返回值
相同) 的元素序列。在每次迭代的时候，它会返回一个值和一个迭代器对象，这个迭代
器对象可以生成元素值全部等于上面那个值的组中所有对象。
一个非常重要的准备步骤是要根据指定的字段将数据排序。因为 groupby() 仅仅
检查连续的元素，如果事先并没有排序完成的话，分组函数将得不到想要的结果。
如果你仅仅只是想根据 date 字段将数据分组到一个大的数据结构中去，并且允许
随机访问，那么你最好使用 defaultdict() 来构建一个多值字典，
'''

from collections import defaultdict

rows_by_date  = defaultdict(list)
for row in rows:
    rows_by_date[row['date']].append(row)

print(rows_by_date)
# defaultdict(<class 'list'>, {'07/02/2012': [{'date': '07/02/2012', 'address': '5800 E 58TH'}, {'date': '07/02/2012', 'address': '5645 N RAVENSWOOD'}, {'date': '07/02/2012', 'address': '1060 W ADDISON'}], '07/03/2012': [{'date': '07/03/2012', 'address': '2122 N CLARK'}], '07/04/2012': [{'date': '07/04/2012', 'address': '5148 N CLARK'}, {'date': '07/04/2012', 'address': '1039 W GRANVILLE'}], '07/01/2012': [{'date': '07/01/2012', 'address': '5412 N CLARK'}, {'date': '07/01/2012', 'address': '4801 N BROADWAY'}]})

for r in rows_by_date['07/01/2012']:
    print(r)
"""
{'date': '07/01/2012', 'address': '5412 N CLARK'}
{'date': '07/01/2012', 'address': '4801 N BROADWAY'}
"""
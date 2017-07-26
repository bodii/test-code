#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第一章：数据结构和算法
Desc: Python 提供了大量的内置数据结构，包括列表，集合以及字典。大多数情况下使
用这些数据结构是很简单的。但是，我们也会经常碰到到诸如查询，排序和过滤等等
这些普遍存在的问题。因此，这一章的目的就是讨论这些比较常见的问题和算法。另
外，我们也会给出在集合模块 collections 当中操作这些数据结构的方法。

Title:    查找最大或最小的 N 个元素
Issue: 怎样从一个集合中获得最大或者最小的 N 个元素列表？
answer: heapq模块有两个函数：nlargest() 和  nsmallest() 可以完美解决这个问题
"""

import heapq

nums = [1, 8, 2, 23, 7, -4, 18, 23, 42, 37, 2]
#最大的三个数，从大到小
print(heapq.nlargest(3, nums)) # 42, 37, 23]
#最小的三个数，从小到大
print(heapq.nsmallest(3,nums)) # [-4, 1, 2]

#两个函数都能接受一个关键字参数，用于更复杂的数据结构中
portfolio = [
    {'name': 'IBM', 'shares': 100, 'price': 91.1},
    {'name': 'AAPL', 'shares': 50, 'price': 543.22},
    {'name': 'FB', 'shares': 200, 'price': 21.09},
    {'name': 'HPQ', 'shares': 35, 'price': 31.75},
    {'name': 'YHOO', 'shares': 45, 'price': 16.35},
    {'name': 'ACME', 'shares': 75, 'price': 115.65}
]
cheap = heapq.nsmallest(3, portfolio,key=lambda s:s['price'])
print(cheap)
# [{'price': 16.35, 'shares': 45, 'name': 'YHOO'}, {'price': 21.09, 'shares': 200, 'name': 'FB'}, {'price': 31.75, 'shares': 35, 'name': 'HPQ'}]
expensive = heapq.nlargest(3,portfolio,key=lambda s:s['price'])
print(expensive)
# [{'price': 543.22, 'shares': 50, 'name': 'AAPL'}, {'price': 115.65, 'shares': 75, 'name': 'ACME'}, {'price': 91.1, 'shares': 100, 'name': 'IBM'}]

'''
如果你想在一个集合中查找最小或最大的 N 个元素，并且 N 小于集合元素数量，
那么这些函数提供了很好的性能。因为在底层实现里面，首先会先将集合数据进行堆
排序后放入一个列表中
'''
nums = [1, 8, 2, 23, 7, -4, 18, 23, 42, 37, 2]
heapq.heapify(nums)
print(nums) # [-4, 2, 1, 23, 7, 2, 18, 23, 42, 37, 8]

'''
堆数据结构最重要的特征是 heap[0] 永远是最小的元素。并且剩余的元素可以很
容易的通过调用 heapq.heappop() 方法得到，该方法会先将第一个元素弹出来，然后
用下一个最小的元素来取代被弹出元素 (这种操作时间复杂度仅仅是 O(log N)，N 是
堆大小)。
'''
heapq.heappop(nums)
print(nums) # [1, 2, 2, 23, 7, 8, 18, 23, 42, 37]

heapq.heappop(nums)
print(nums) # [2, 2, 8, 23, 7, 37, 18, 23, 42]

heapq.heappop(nums)
print(nums) # [2, 7, 8, 23, 42, 37, 18, 23]
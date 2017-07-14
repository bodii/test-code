#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第一章：数据结构和算法
Desc: Python 提供了大量的内置数据结构，包括列表，集合以及字典。大多数情况下使
用这些数据结构是很简单的。但是，我们也会经常碰到到诸如查询，排序和过滤等等
这些普遍存在的问题。因此，这一章的目的就是讨论这些比较常见的问题和算法。另
外，我们也会给出在集合模块 collections 当中操作这些数据结构的方法。

Title:      字典的运算
Issue: 怎样在数据字典中执行一些计算操作 (比如求最小值、最大值、排序等等)？
answer: 为了对字典值执行计算操作，通常需要使用 zip() 函数先将键和值反转过来，可以使用sorted() 函数来排列字典数据。
"""

prices = {
    'ACME' : 45.23,
    'AAPL' : 612.78,
    'IBM' : 205.55,
    'HPQ' : 37.20,
    'FB' : 10.75
}

min_price = min(zip(prices.values(),prices.keys()))
print(min_price) # (10.75, 'FB')

max_price = max(zip(prices.values(),prices.keys()))
print(max_price) # (612.78, 'AAPL')

# 排列字典数据
prices_sorted = sorted(zip(prices.values(),prices.keys()))
print(prices_sorted) # [(10.75, 'FB'), (37.2, 'HPQ'), (45.23, 'ACME'), (205.55, 'IBM'), (612.78, 'AAPL')]

'''
执行这些计算的时候，需要注意的是 zip() 函数创建的是一个只能访问一次的迭
代器。比如，下面的代码就会产生错误：
prices_and_names = zip(prices.values(), prices.keys())
print(min(prices_and_names)) # OK
print(max(prices_and_names)) # ValueError: max() arg is an empty sequence
'''

'''
如果你在一个字典上执行普通的数学运算，你会发现它们仅仅作用于键，而不是
值。比如：
min(prices) # Returns 'AAPL'
max(prices) # Returns 'IBM'
'''

#你可以在 min() 和 max() 函数中提供 key 函数参数来获取最小值或最大值对应的
#键的信息。
print(min(prices,key = lambda k:prices[k])) # FB
print(max(prices,key = lambda k:prices[k])) # AAPL

# 但是，如果还想要得到最小值，你又得执行一次查找操作。
print(prices[min(prices,key=lambda k:prices[k])]) # 10.75

'''
前面的 zip() 函数方案通过将字典” 反转” 为 (值，键) 元组序列来解决了上述问
题。当比较两个元组的时候，值会先进行比较，然后才是键。这样的话你就能通过一
条简单的语句就能很轻松的实现在字典上的求最值和排序操作了
'''

'''
需要注意的是在计算操作中使用到了 (值，键) 对。当多个实体拥有相同的值的时
候，键会决定返回结果。比如，在执行 min() 和 max() 操作的时候，如果恰巧最小或
最大值有重复的，那么拥有最小或最大键的实体会返回：
>>> prices = { 'AAA' : 45.23, 'ZZZ': 45.23 }
>>> min(zip(prices.values(), prices.keys()))
(45.23, 'AAA')
>>> max(zip(prices.values(), prices.keys()))
(45.23, 'ZZZ')
>>>
'''
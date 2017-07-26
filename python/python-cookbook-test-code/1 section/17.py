#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第一章：数据结构和算法
Desc: Python 提供了大量的内置数据结构，包括列表，集合以及字典。大多数情况下使
用这些数据结构是很简单的。但是，我们也会经常碰到到诸如查询，排序和过滤等等
这些普遍存在的问题。因此，这一章的目的就是讨论这些比较常见的问题和算法。另
外，我们也会给出在集合模块 collections 当中操作这些数据结构的方法。

Title:         从字典中提取子集
Issue:  你想构造一个字典，它是另外一个字典的子集
代访问
answer:  最简单的方式是使用字典推导
"""

prices = {
    'ACME': 45.23,
    'AAPL': 612.78,
    'IBM': 205.55,
    'HPQ': 37.20,
    'FB': 10.75
}
p1 = {key:value for key, value in prices.items() if value >200 }
print(p1) # {'AAPL': 612.78, 'IBM': 205.55}
tech_names = {'AAPL', 'IBM', 'HPQ', 'MSFT'}
p2 = {key:value for key,value in prices.items() if key in tech_names}
print(p2) # {'HPQ': 37.2, 'AAPL': 612.78, 'IBM': 205.55}

'''
大多数情况下字典推导能做到的，通过创建一个元组序列然后把它传给 dict() 函
数也能实现。
'''
p1 = dict((key, value) for key, value in prices.items() if value > 200)
print(p1) # {'AAPL': 612.78, 'IBM': 205.55}

'''
但是，字典推导方式表意更清晰，并且实际上也会运行的更快些 (在这个例子中，
实际测试几乎比 dcit() 函数方式快整整一倍)
'''

'''
有时候完成同一件事会有多种方式。比如，第二个例子程序也可以像这样重写
'''
tech_names = { 'AAPL', 'IBM', 'HPQ', 'MSFT' }
p2 = { key:prices[key] for key in prices.keys() & tech_names}
print(p2) # {'HPQ': 37.2, 'IBM': 205.55, 'AAPL': 612.78}
'''
但是，运行时间测试结果显示这种方案大概比第一种方案慢 1.6 倍。如果对程序运
行性能要求比较高的话，需要花点时间去做计时测试。
'''

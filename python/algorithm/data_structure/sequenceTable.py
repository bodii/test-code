#!/usr/bin/env python3
# -*- coding:utf-8 -*-
# PW @ 2019-11-09 14:07:01

"""
数据结构:
顺序表
将元素顺序地存放在一块连续的存储区里，元素间的关系由它们的存储顺序自然表示。
list和tuple两种类型采用了顺序表的实现技术。

时间复杂度：
a. 尾端加入元素，时间复杂度为o(1)。
b. 非保序的加入元素(不常见),时间复杂度为o(1)。
c. 保序的元素加入，时间复杂度为o(n)。

删除元素：
a. 删除表尾元素，时间复杂度为o(1)。
b. 非保序的元素删除(不常见),时间复杂度为o(1)。
c. 保序的元素删除，时间复杂度为o(n)。
"""

from timeit import Timer

'''
测试list列表中append、insert方法执行速度
'''
def append_test():
    alist = [] 
    for i in range(10000):
        alist.append(i)


def insert_test():
    alist = [] 
    for i in range(10000):
        alist.insert(i, i)

timer1 = Timer('append_test()', 'from __main__ import append_test')
print('append', timer1.timeit(1000))

timer2 = Timer('insert_test()', 'from __main__ import insert_test')
print('insert', timer2.timeit(1000))

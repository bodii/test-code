#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第四章：迭代器与生成器
Desc: 迭代是 Python 最强大的功能之一。初看起来，你可能会简单的认为迭代只不过是
处理序列中元素的一种方法。然而，绝非仅仅就是如此，还有很多你可能不知道的，
比如创建你自己的迭代器对象，在 itertools 模块中使用有用的迭代模式，构造生成器
函数等等。这一章目的就是向你展示跟迭代有关的各种常见问题。

Title;      不同集合上元素的迭代
Issue:你想在多个对象执行相同的操作，但是这些对象在不同的容器中，你希望代码在不
失可读性的情况下避免写重复的循环。
Answer: itertools.chain() 方法可以用来简化这个任务。它接受一个可迭代对象列表作为
输入，并返回一个迭代器，有效的屏蔽掉在多个容器中迭代细节。
"""

# 例子：
from itertools import chain

a = [1, 2, 3, 4]
b = ['x', 'y', 'z']
for x in chain(a, b):
    print(x)
'''
1
2
3
4
x
y
z
'''


'''
使用 chain() 的一个常见场景是当你想对不同的集合中所有元素执行某些操作的
时候。比如：
'''
active_items = set()
inactive_items = set()
#for item in chain(active_items, inactive_items):
# 这种解决方案要比像下面这样使用两个单独的循环更加优雅，
# for item in active_items:
# for item in inactive_items:


'''
itertools.chain() 接受一个或多个可迭代对象最为输入参数。然后创建一个迭代
器，依次连续的返回每个可迭代对象中的元素。这种方式要比先将序列合并再迭代要
高效的多。比如：

for x in a + b:

for x in chain(a, b):

第一种方案中， a + b 操作会创建一个全新的序列并要求 a 和 b 的类型一致。
chian() 不会有这一步，所以如果输入序列非常大的时候会很省内存。并且当可迭代
对象类型不一样的时候 chain() 同样可以很好的工作。
'''
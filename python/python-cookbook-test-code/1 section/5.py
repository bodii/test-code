#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第一章：数据结构和算法
Desc: Python 提供了大量的内置数据结构，包括列表，集合以及字典。大多数情况下使
用这些数据结构是很简单的。但是，我们也会经常碰到到诸如查询，排序和过滤等等
这些普遍存在的问题。因此，这一章的目的就是讨论这些比较常见的问题和算法。另
外，我们也会给出在集合模块 collections 当中操作这些数据结构的方法。

Title:     实现一个优先级队列
Issue: 怎样实现一个按优先级排序的队列？并且在这个队列上面每次 pop 操作总是返回
优先级最高的那个元素
answer: 利用 heapq 模块实现了一个简单的优先级队列
"""
"""
这 一 小 节 我 们 主 要 关 注 heapq 模 块 的 使 用。 函 数 heapq.heappush() 和
heapq.heappop() 分别在队列 queue 上插入和删除第一个元素，并且队列 queue 保证
第一个元素拥有最小优先级 (1.4 节已经讨论过这个问题)。 heappop() 函数总是返回”
最小的” 的元素，这就是保证队列 pop 操作返回正确元素的关键。另外，由于 push 和
pop 操作时间复杂度为 O(log N)，其中 N 是堆的大小，因此就算是 N 很大的时候它们
运行速度也依旧很快。
在上面代码中，队列包含了一个 (-priority, index, item) 的元组。优先级为负
数的目的是使得元素按照优先级从高到低排序。这个跟普通的按优先级从低到高排序
的堆排序恰巧相反。
index 变量的作用是保证同等优先级元素的正确排序。通过保存一个不断增加的
index 下标变量，可以确保元素按照它们插入的顺序排序。而且， index 变量也在相
同优先级元素比较的时候起到重要作用。
"""
import heapq

class PriorityQueue:

    def __init__(self):
        self._queue = []
        self._index = 0

    def push(self, item, priority):
        heapq.heappush(self._queue, (-priority, self._index, item))
        self._index += 1

    def pop(self):
        return heapq.heappop(self._queue)[-1]


class Item:
    def __init__(self, name):
        self.name = name
    def __repr__(self):
        return 'Item({!r})'.format(self.name)

q = PriorityQueue()
q.push(Item('foo'), 1)
q.push(Item('bar'), 5)
q.push(Item('spam'), 4)
q.push(Item('grok'), 1)

print(q.pop()) # Item('bar')
print(q.pop()) # Item('spam')
print(q.pop()) # Item('foo')
print(q.pop()) # Item('grok')

'''
仔细观察可以发现，第一个 pop() 操作返回优先级最高的元素。另外注意到如果
两个有着相同优先级的元素 ( foo 和 grok )，pop 操作按照它们被插入到队列的顺序返
回的。
'''

'''
如果你使用元组 (priority, item) ，只要两个元素的优先级不同就能比较。但是
如果两个元素优先级一样的话，那么比较操作就会跟之前一样出错：
>>> a = (1, Item('foo'))
>>> b = (5, Item('bar'))
>>> a < b
True
>>> c = (1, Item('grok'))
>>> a < c
Traceback (most recent call last):
File "<stdin>", line 1, in <module>
TypeError: unorderable types: Item() < Item()
>>>
通过引入另外的 index 变量组成三元组 (priority, index, item) ，就能很好的
避免上面的错误，因为不可能有两个元素有相同的 index 值。Python 在做元组比较时
候，如果前面的比较以及可以确定结果了，后面的比较操作就不会发生了：
>>> a = (1, 0, Item('foo'))
>>> b = (5, 1, Item('bar'))
>>> c = (1, 2, Item('grok'))
>>> a < b
True
>>> a < c
True
>>>
如果你想在多个线程中使用同一个队列，那么你需要增加适当的锁和信号量机制。
'''
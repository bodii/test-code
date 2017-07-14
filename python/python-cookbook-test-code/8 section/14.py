#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第八章：类与对象
Desc: 本章主要关注点的是和类定义有关的常见编程模型。包括让对象支持常见的
Python 特性、特殊方法的使用、类封装技术、继承、内存管理以及有用的设计模
式。

Title;          实现自定义容器
Issue:你想实现一个自定义的类来模拟内置的容器类功能，比如列表和字典。但是你不确
定到底要实现哪些方法。
Answer:collections 定义了很多抽象基类，当你想自定义容器类的时候它们会非常有用
"""

import collections

class A(collections.Iterable):
    def __iter__(self):
        pass

"""
不过你需要实现 collections.Iterable 所有的抽象方法，否则会报错:
"""

a = A()
# TypeError: Can't instantiate abstract class A with abstract methods __iter__

#print(collections.Sequence())
# TypeError: Can't instantiate abstract class Sequence with abstract methods __getitem__, __len__

"""
下面是一个简单的示例，继承自上面 Sequence 抽象类，并且实现元素按照顺序存
储：
"""
import bisect
class SortedItems(collections.Sequence):
    def __init__(self, initial=None):
        self._items = sorted(initial) if initial is not None else []

    def __getitem__(self, index):
        return self._items[index]

    def __len__(self):
        return len(self._items)

    def add(self, itme):
        bisect.insort(self._items, itme)

items = SortedItems([5, 1, 3])
print(list(items))
# [1, 3, 5]
print(items[0], items[-1])
# 1 5
items.add(2)
print(list(items))
# [1, 2, 3, 5]

"""
可以看到，SortedItems 跟普通的序列没什么两样，支持所有常用操作，包括索引、
迭代、包含判断，甚至是切片操作。
这里面使用到了 bisect 模块，它是一个在排序列表中插入元素的高效方式。可以
保证元素插入后还保持顺序。
"""

"""
使用 collections 中的抽象基类可以确保你自定义的容器实现了所有必要的方法。
并且还能简化类型检查。你的自定义容器会满足大部分类型检查需要，
"""
items = SortedItems()

print(isinstance(items, collections.Iterable))
# True
print(isinstance(items, collections.Sequence))
# True
print(isinstance(items, collections.Container))
# True
print(isinstance(items, collections.Sized))
# True
print(isinstance(items, collections.Mapping))
# False
"""
collections 中 很 多 抽 象 类 会 为 一 些 常 见 容 器 操 作 提 供 默 认 的 实 现， 这
样 一 来 你 只 需 要 实 现 那 些 你 最 感 兴 趣 的 方 法 即 可。 假 设 你 的 类 继 承 自
collections.MutableSequence
"""
class Items(collections.MutableSequence):
    def __init__(self, initial=None):
        self._items = list(initial) if initial is not None else []

    def __getitem__(self, index):
        print('Getting:', index)
        return self._itmes[index]

    def __setitem__(self, index, value):
        print('setting:', index, value)
        self._items[index] = value

    def __delitem__(self, index):
        print('deleting:', index)
        del self._items[index]

    def insert(self, index, value):
        print('Inserting', index, value)
        self._items.insert(index, value)

    def __len__(self):
        print('Len')
        return len(self._items)

    def count(self, value):
        print('Count')
        return sum( 1 for v in self._items if v == value)

    def remove(self, value):
        print('Remove')
        count_num = sum( 1 for v in self._items if v == value)
        indexs = [i  for i in range(len(self._items)) if self._items[i] == value]
        for i in indexs:
            self.__delitem__(i)
        return count_num

    def get(self):
        return self._items





"""
如果你创建 Items 的实例，你会发现它支持几乎所有的核心列表方法 (如 append()、
remove()、count() 等)。
"""
a = Items([1, 2, 3, 2])
print(len(a))
'''
Len
3
'''
a.append(4)
print(a.count(2))
print(a.remove(4))
print(a.get())

"""
本小节只是对 Python 抽象类功能的抛砖引玉。numbers 模块提供了一个类似的跟
整数类型相关的抽象类型集合。可以参考 8.12 小节来构造更多自定义抽象基类。
"""

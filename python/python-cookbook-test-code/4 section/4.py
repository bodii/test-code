#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第四章：迭代器与生成器
Desc: 迭代是 Python 最强大的功能之一。初看起来，你可能会简单的认为迭代只不过是
处理序列中元素的一种方法。然而，绝非仅仅就是如此，还有很多你可能不知道的，
比如创建你自己的迭代器对象，在 itertools 模块中使用有用的迭代模式，构造生成器
函数等等。这一章目的就是向你展示跟迭代有关的各种常见问题。

Title;        实现迭代器协议
Issue:你想构建一个能支持迭代操作的自定义对象，并希望找到一个能实现迭代协议的简
单方法。
Answer: 目前为止，在一个对象上实现迭代最简单的方式是使用一个生成器函数。
"""

class Node:
    def __init__(self, value):
        self._value = value
        self._children = []

    def __repr__(self):
        return 'Node({!r})'.format(self._value)

    def add_child(self, node):
        self._children.append(node)

    def __iter__(self):
        return iter(self._children)

    def depth_first(self):
        yield self
        for c in self:
            yield from c.depth_first()

#Example
if __name__ == '__main__':
    root = Node(0)
    child1 = Node(1)
    child2 = Node(2)
    root.add_child(child1)
    root.add_child(child2)
    child1.add_child(Node(3))
    child1.add_child(Node(4))
    child2.add_child(Node(5))

    for ch in root.depth_first():
        print(ch)

'''
在这段代码中，depth first() 方法简单直观。它首先返回自己本身并迭代每一个
子节点并通过调用子节点的 depth first() 方法 (使用 yield from 语句) 返回对应元
素。
'''

'''
Python 的迭代协议要求一个 iter () 方法返回一个特殊的迭代器对象，这个迭
代器对象实现了 next () 方法并通过 StopIteration 异常标识迭代的完成。但是，
实现这些通常会比较繁琐。
'''
'''
下面我们演示下这种方式，如何使用一个关联迭代器类重
新实现 depth first() 方法：
'''
class Node2:
    def __init__(self, value):
        self._value = value
        self._children = []

    def __repr__(self):
        return 'Node({!r})'.format(self._value)

    def add_child(self,node):
        self._children.append(node)

    def __iter__(self):
        return iter(self._children)

    def depth_first(self):
        return DepthFirstIterator(self)

class DepthFirstIterator(object):
    '''
    Depth-first traveral
    '''
    def __init__(self, start_node):
        self._node = start_node
        self._children_iter = None
        self._child_iter = None

    def __iter__(self):
        return self

    def __next__(self):
        if self._children_iter is None:
            self._children_iter = iter(self._node)
            return self._node
        elif self._child_iter:
            try:
                nextchild = next(self._child_iter)
                return nextchild
            except StopIteration:
                self._child_iter = None
                return next(self)
        else:
            self._child_iter = next(self._children_iter)
            return next(self)


'''
DepthFirstIterator 类和上面使用生成器的版本工作原理类似，但是它写起来很
繁琐，因为迭代器必须在迭代处理过程中维护大量的状态信息。坦白来讲，没人愿意
写这么晦涩的代码。将你的迭代器定义为一个生成器后一切迎刃而解。
'''
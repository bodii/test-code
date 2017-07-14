#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第四章：迭代器与生成器
Desc: 迭代是 Python 最强大的功能之一。初看起来，你可能会简单的认为迭代只不过是
处理序列中元素的一种方法。然而，绝非仅仅就是如此，还有很多你可能不知道的，
比如创建你自己的迭代器对象，在 itertools 模块中使用有用的迭代模式，构造生成器
函数等等。这一章目的就是向你展示跟迭代有关的各种常见问题。

Title;         代理迭代
Issue: 你构建了一个自定义容器对象，里面包含有列表、元组或其他可迭代对象。你想直
接在你的这个新容器对象上执行迭代操作。
Answer: 实际上你只需要定义一个 iter () 方法，将迭代操作代理到容器内部的对象上
去。
"""

class Node:
    def __init__(self, value):
        self._value = value
        self._children = []
    def __repr__(self):
        return 'Node({!r})'.format(self._value)

    def add_child(self,node):
        self._children.append(node)
    def __iter__(self):
        return iter(self._children)

'''
在上面代码中， iter () 方法只是简单的将迭代请求传递给内部的 children
属性。
'''
#Example
if __name__ == '__main__':
    root = Node(0)
    child1 = Node(1)
    child2 = Node(2)
    root.add_child(child1)
    root.add_child(child2)
    for ch in root:
        print(ch)

'''
Python 的迭代器协议需要 iter () 方法返回一个实现了 next () 方法的迭代
器对象。如果你只是迭代遍历其他容器的内容，你无须担心底层是怎样实现的。你所
要做的只是传递迭代请求既可。
这 里 的 iter() 函 数 的 使 用 简 化 了 代 码， iter(s) 只 是 简 单 的 通 过 调 用
s. iter () 方法来返回对应的迭代器对象，就跟 len(s) 会调用 s. len () 原理
是一样的。
'''
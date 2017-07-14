#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第八章：类与对象
Desc: 本章主要关注点的是和类定义有关的常见编程模型。包括让对象支持常见的
Python 特性、特殊方法的使用、类封装技术、继承、内存管理以及有用的设计模
式。

Title;             循环引用数据结构的内存管理
Issue:你的程序创建了很多循环引用数据结构 (比如树、图、观察者模式等)，你碰到了内
存管理难题
Answer:一个简单的循环引用数据结构例子就是一个树形结构，双亲节点有指针指向孩子节
点，孩子节点又返回来指向双亲节点。这种情况下，可以考虑使用 weakref 库中的弱
引用。
"""
import weakref

class Node:
    def __init__(self, value):
        self.value = value
        self._parent = None
        self.children = []

    def __repr__(self):
        return 'Node({!r:})'.format(self.value)

    @property
    def parent(self):
        return None if self._parent is None else self._parent()

    @parent.setter
    def parent(self, node):
        self._parent = weakref.ref(node)

    def add_child(self, child):
        self.children.append(child)
        child.parent = self
"""
这种是想方式允许 parent 静默终止。
"""
root = Node('parnet')
c1 = Node('child')
root.add_child(c1)
print(c1.parent)
# Node('parnet')

del root
print(c1.parent)
# None

"""
循环引用的数据结构在 Python 中是一个很棘手的问题，因为正常的垃圾回收机制
不能适用于这种情形。
"""
class Data:
    def __del__(self):
        print('Data.__del__')

class Node:
    def __init__(self):
        self.data = Data()
        self.parent = None
        self.children = []

    def add_child(self, child):
        self.children.append(child)
        child.parent = self

"""
下面我们使用这个代码来做一些垃圾回收试验：
"""
a = Data()
del a
# Data.__del__
a = Node()
del a
# Data.__del__
a = Node()
a.add_child(Node())

"""
可以看到，最后一个的删除时打印语句没有出现。原因是 Python 的垃圾回收机制
是基于简单的引用计数。当一个对象的引用数变成 0 的时候才会立即删除掉。而对于
循环引用这个条件永远不会成立。因此，在上面例子中最后部分，父节点和孩子节点
互相拥有对方的引用，导致每个对象的引用计数都不可能变成 0。
Python 有另外的垃圾回收器来专门针对循环引用的，但是你永远不知道它什么时
候会触发。另外你还可以手动的触发它，但是代码看上去很挫：
"""
import gc
print(gc.collect())

"""
如果循环引用的对象自己还定义了自己的 del () 方法，那么会让情况变得更糟
糕。假设你像下面这样给 Node 定义自己的 del () 方法：
"""
class Node:
    def __init__(self):
        self.data = Data()
        self.parent = None
        self.children = []

    def add_child(self, child):
        self.children.append(child)
        child.parent = self

    def __del__(self):
        del self.data
        del self.parent
        del self.children
"""
这种情况下，垃圾回收永远都不会去回收这个对象的，还会导致内存泄露。如果你
试着去运行它会发现，Data. del 消息永远不会出现了, 甚至在你强制内存回收时
"""
a = Node()
a.add_child(Node())
del a
print(gc.collect())

"""
弱引用消除了引用循环的这个问题，本质来讲，弱引用就是一个对象指针，它不会
增加它的引用计数。你可以通过 weakref 来创建弱引用。
"""
import weakref
a = Node()
a_ref = weakref.ref(a)
print(a_ref)
# <weakref at 0x0000020419688638; to 'Node' at 0x000002041968DE10>

"""
为了访问弱引用所引用的对象，你可以像函数一样去调用它即可。如果那个对象还
存在就会返回它，否则就返回一个 None。由于原始对象的引用计数没有增加，那么就
可以去删除它了。
"""
del a
print(a_ref())
# None
"""
通过这里演示的弱引用技术，你会发现不再有循环引用问题了，一旦某个节点不被
使用了，垃圾回收器立即回收它。你还能参考 8.25 小节关于弱引用的另外一个例子。
"""

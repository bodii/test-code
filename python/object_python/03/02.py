#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 容器和集合 '''

"""
除了内置的容器类之外，collections模块也定义许多集合。这些容器类包括namedtuple()、deque、
ChainMap、Counter、OrderedDict 和 defaultdict。所有这些类都是基于抽象基类定义类的例子。
"""
import collections
# 显式集合是否支持一个特定方法的快捷方式 
print( isinstance({}, collections.abc.Mapping ) )
# True
print( isinstance( collections.defaultdict(int), collections.abc.Mapping) )
# True
"""
可以检查简单的dict类，看它是否遵循基本的映射协议，并支持必需的方法。
可以检查defaultdict集合，确认它也是一种映射。

内置容器的完整图谱都是用抽象基类来表示的。底层的特性包括Container \ Iterable 和 Sized。
而它们同时也是高级构造过程的一部分。它们需要一些特殊方法，分别是__contains__()\ __iter__
() 和 __len__()。

高级的特性包括以下几个部分：
1. Sequence 和 MutableSequence: 它们是list和tuple的抽象基类，具体的序列实现还包括
bytes 和 str。
2. MutableMapping: 这是dict的抽象基类，它扩展了Mapping类，但是没有内置具体的实现。
3. set 和 MutableSet：它们是frozenset 和 set 的抽象基类。

这些特性允许我们创建新的类，扩展现有的类并且能够使之与Python内置特性的集成变得清晰、流畅。
"""

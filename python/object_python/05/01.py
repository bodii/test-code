#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 创建容器和集合 '''

"""
 集合的抽象基类

collections.abc 模块提供了很多抽象基类，这些类将集合分解成许多互相独立的属性集。

Container基类要求子类实现__contains__()方法，这个特殊方法实现了in运算符。
Iterable基类要求子类实现__iter__()方法。for语句、生成器表达式和iter()函数都需要使用这
个方法。
Sized基类要求子类实现__len__()方法。len()函数需要使用这个方法，它也很稳妥地实现了__bool
__()方法，但是这个方法不是必需的。
Hashable基类要求子类实现__hash__()方法。hash()函数需要使用这个方法，如果这个方法被实现了
,就意味着当前对象是不可变的。

Sequence 和 MutableSequence类，它们是基于例如index()\count()\reverse()\extend()
和remove()方法创建的，同时也包含了这些方法的实现。
Mapping和 MutableMapping类，这两个类包含了例如keys()\items()\values()\get()和其他
的一些方法的实现。
Set 和MutableSet包含了用于set类型的比较操作和算术运算符的实现。

"""


''' 特殊方法示例 '''

def  __contains__( self, rank ):
    return any( c.rank == rank for rank in hand.cards )


''' 使用标准库的扩展 '''

"""
6个集合函数
[ import collections ] 

namedtuple()函数会创建允许包含命名属性的tuple类，我们可以使用这个函数而不是额外完整地
定义一个仅仅为属性值命名的类。

deque(注意这种不规则的拼写方式)是一个双端队列，一个类似list的集合，但是可以在任何一段
快速地进行增加和弹出操作。可以用这个类的其中一部分属性来创建常规的栈或队列。

在一些情况下，我们可以使用ChainMap，而不是合并同的映射。这是一个将多个映射连接起来的方法。

一个OrderedDict集合是会保持所有元素原始插入顺序的一种映射。

defaultDict是dict的一个子类，它内部使用一个工厂函数为所有没有值的键提供默认值。

Counter也是一个dict的子类，它可以被用来统计对象，进而创建频率表。但是，它实际上是一个更
复杂的数据结构，通常我们称为多重集合(multiset) 或者包(bag)。

import heapq.Heapq
Heapq 模块是包含了一系列将堆队列(heap queue)的属性添加到一个现有的list对象上的函数。
堆队列的不变性是指堆上的所有元素按顺序存储，这样可以对堆队列进行升序的快速检索。如果我们在
一个list结构上使用heapq方法，我们就不再需要显式地排序列表。这个方法可以带来很大的性能提升。

array模块是一种对特定值的存储方式进行优化的序列，这为包括了大量简单值的集合提供了一些类似
列表的特性。

"""


''' namedtuple函数 '''
"""
namedtuple()函数根据提供的参数创建一个新类。这个类会有一个类名，一个字段名和一对可选的
用于定义类行为的关键字。
"""
from collections import namedtuple

BlackjackCard = namedtuple( 'BackjackCard', 'rank,suit,hard,soft' )

# 一个工厂函数创建这个类的不同实例

def card( rank, suit ):
    if rank == 1:
        return BlackjackCard( 'A', suit, 1, 11 )
    elif 2 <= rank < 11:
        return BlackjackCard( str(rank), suit, rank, rank )
    elif rank == 11:
        return BlackjackCard( 'J', suit, 10, 10 )
    elif rank == 12:
        return BlackjackCard( 'Q', suit, 10, 10 )
    elif rank == 13:
        return BlackjackCard( 'K', suit, 10, 10 )


class TheNamedTuple( tuple ):
    __slots__ = ()
    _fields = {fields_name!r}

    def __new__( _cls, {arg_list} ):
        return _tuple._new__( _cls, ({arg_list}) )



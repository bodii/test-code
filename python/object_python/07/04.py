#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用标准库中的装饰器 '''
"""
标准库中有许多装饰器。例如：contextlib\ functools\ unittest\ atexit\ importlib
和reprlib模块都包含了可以作为软件设计中横切方面的经典范例的装饰器。

例如，functools库提供了total_ordering装饰器，它定义了一系列比较运算符。它用__eq__()
和__lt__()\__le__()\__gt__()或者__ge__()中的一个创建一套完整的比较运算。
"""

import functools

@functools.total_ordering
class Card():
    __slots__ = ( 'rank', 'suit' )
    
    def __new__( cls, rank, suit ):
        self = super().__new__(cls)
        self.rank = rank
        self.suit = suit
        return self

    def __eq__( self, other ):
        return self.rank == other.rank

    def __lt__( self, other ):
        return self.rank < other.rank

"""
@functools.total_ordering。这个装饰器会为我们创建未定义的方法函数。可以用这个
类创建支持所有比较运算符的对象，虽然在类中只定义了两个。
"""
c1 = Card( 3, '♠' )
c2 = Card( 3, '♥' )
print(c1 == c2)
# True
print(c1 < c2)
# False
print(c1 <= c2)
# True
print(c1 >= c2)
# True

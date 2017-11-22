#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用标准库中的mixin类 '''

"""
标准库使用了mixin类定义。有许多模块中都有这种例子，包括io\ secketserver\
urllib.request\ contextlib 和 collections.abc。

当我们基于collections.abc抽象基类自定义集合时，我们会使用mixin类确保容器的
横切方面都以一致的方式定义。最上层的集合(set、Sequence和Mapping)都是基于多
个mixin类创建的。
只看其中一行，Sequence的总结，我们可以看到它继承自Sized、Iterable和Container。
这些mixin类提供了__contains__()、__iter__()、__reversed__()、__index__()
和count()方法。
"""
import random

class Deck( list ):
    def __init__( self, size=1 ):
        super().__init__()
        self.rng = random.Random()
        for d in range(size):
            cards = [ card(r,s) for r in range(13) for s in Suits ]
            super().extend( cards )
        self._init_shuffle()

    def _init_shuffle( self ):
        self.rng.shuffle( self )


from contextlib import ContextDecorator

class TestDeck( ContextDecorator, Deck ):
    def __init__( self, size=1, seed=0 ):
        super().__init__( size=size )
        self.seed = seed
    
    def _init_shuffle( self ):
        """Don't shuffle during __init__."""
        pass
    
    def __enter__( self ):
        self.rng.seed( self.seed, version=1)
        self.rng.shuffle( self )
        return self

    def __exit__( self, exc_type, exc_value, traceback ):
        pass

"""这个例子展示了Deck如何自己生成随机的洗牌顺序。"""
for i in range(3):
    d1 = Deck(5)
    print( d1.pop(), d1.pop(), d1.pop() )

"""使用一个给定随机数种子的TestDeck, TestDeck用作上下文管理器，并生成一系列已知
顺序的牌。每次我们调用它，都会得到相同顺序的牌。"""
for i in range(3):
    with TestDeck(5, seed=0) as d2:
        print( d2.pop(), d2.pop(), d2.pop() )


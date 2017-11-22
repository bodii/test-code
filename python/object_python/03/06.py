#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' abc模块 '''

"""
创建抽象基类的核心方法定义在abc模块中。此模块中包含的ABCMeta类提供了一些有用的特性。
首先，ABCMeta类保证抽象基类不可以被实例化。但是，一个提供了所有必须实现的子类可以被实例化。
这个元类型会在执行__new__()的时候调用抽象基类中__subclasshook__()特殊方法。
如果这个方法返回NotImplemented,就会抛出一个异常，指出当前类没有实现所有必需的方法。
其次，它提供了__instancecheck__()和__subclasscheck__()的定义。这两个特殊方法实现了
isinstance()和issubclass()两个内置的函数。它们用于确保一个对象(或者类)属于正确的抽象蕨
类。同时，为了提高性能，这个类也会包括一个子类的缓存。
abc模块还包括了许多用于创建抽象方法函数的装饰器，子类中必须实现用这些装饰器定义的抽象方法。
其中最重要的是@abstractmethod装饰器。
"""

from abc import ABCMeta, abstractmethod

class AbstractBettingStrategy(metaclass=ABCMeta):
    __slots__ = ()

    @abstractmethod
    def bet( self, hand ):
        return 1

    @abstractmethod
    def record_win( self, hand ):
        pass

    @abstractmethod
    def record_loss( self, hand ):
        pass

    @abstractmethod
    def __subclasshook__( cls, subclass ):
        if cls is Hand:
            if ( any('bet' in B.__dict__ for B in subclass.__mro__) 
            and any('record_win' in B.__dict__ for B in subclass.__mro__)
            and any('record_loss' in B.__dict__ for B in subclass.__mro__)
            ):
                return True
        return NotImplemented

"""
这个类用ABCMeta类作为它的元类型;同时，它也实现了用于检查类型完整性的__subclasshook__()
方法。这些方法提供了作为一个抽象基类应有的核心特性。

这个抽象基类使用 @abstractmethod 装饰器定义了3个抽象方法，任何试图实现这个抽象类的子类
都必须实现这3个方法。

__subclasshook__方法需要这3个方法都被实现。这样的做法似乎有一些过于严格了，因为一个很简单
的下注方法不应该提供用来计算输赢的方法。

__subclasshook__依赖于Python内置的两个类特性：__dict__和__mro__
__dict__特性用于存储类中定义的方法名和特性名，这个特性中存储了类的主体。

__mro__特性中记录了解析方法的顺序，这个特性记录了当前类层次结构的顺序。由于Python中允许多
继承，因此有可能有许多基类，那么这些基类的顺序同时也决定了解析方法名的优先级。
"""

# test 继承基类 AbstractBettingStrategy
class Simple( AbstractBettingStrategy ):
    def bet( self, hand ):
        return 1

    def record_win( self, hand ):
        pass

    def record_loss( self, hand ):
        pass


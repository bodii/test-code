#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 向类中添加方法函数 '''

def memento( class_ ):
    def memento( self ):
        return '{0.__class__.__qualname__}({0!r})'.format(self)
    class_.memento = memento
    return class_
"""
这个装饰器包含了一个即将被插入到类中的方法函数。
"""

@memento
class SomeClass:
    def __init__( self, value ):
        self.value = value

    def __repr__( self ):
        return '{0.value}'.format(self)
"""
装饰器向类中插入了一个新方法---memento()。但是，这种做法有一些缺点。
1. 我们不能通过重载memento()方法函数的实现来处理特殊情况，它是在定义之后内嵌到类中的。
2. 我们很难扩展装饰器函数。如果我们想要扩展功能或者处理特殊情况，我们必须将它升级为可调
用对象。我们准备升级到一个可调用对象，那么我们应该完全放弃这种方法，转而使用一个mixin类
添加方法。
"""

# 下面是添加一个标准方法的mixin类
class Memento:
    def memento( self ):
        return '{0.__class__.__qualname__}({0!r})'.format(self)

# 使用这个Memento类定义一个类
class SomeClass2( Memento ):
    def __inif__( self, value ):
        self.value = value 

    def __repr__( self ):
        return '{0.value}'.format(self)
"""
这个mixin提供了一个新方法---memento(),这是一个mixin类的经典用法。通过扩展Memento
 mixin类添加功能更容易，另外，我们可以重载memento()方法函数，用于处理特殊情况。
"""


#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' __getattribute__()方法 '''

"""
__getattribute__()方法提供了对属性更底层的一些操作。默认的实现逻辑是先从内部的__dict__
（或__slots__）中查找已有的属性。如果属性没有找到则调用__getattr__()函数。如果值是
一个修饰符，对修饰符进行处理。否则，返回当前值即可。

 通过重写这个方法，可以达到以下目的：
1. 可以有效阻止属性访问。在这个方法中，抛出异常而非返回值。相比于在代码中仅仅使用下划线
(_)开头来把一个名字标记为私有的方法，这种方法使得属性的封装更透彻。

2. 可仿照__getattr__()函数的工作方式来创建新属性。在这种情况下，可以绕过__getattribute
__()的实现逻辑。

3. 可以使得属性执行单独或不同的任务。但这样会降低程序的可读性和可维护性，这是个很糟糕的想法。

4. 可以改变修饰符的行为。虽然技术上可行，改变修饰符的行为却是个糟糕的想法。


当实现__getattribute__()方法时，将阻止任何内部属性访问函数体，这一点很重要。如果试图
获取self.name的值，会导致无限递归。

获得__getattribute__()方法中的属性值，必须显式调用 object基类中的方法，像：
object.__getattribute__(self, name)
通过使用__getattribute__()方法阻止对内部__dict__属性的访问来实现不呆变。
"""

# 下面类定义隐藏了所有名称以下划线（_）为开头的属性。
class BlackJackCard3:
    ''' Abstract Superclass '''
    def __init__( self, rank, suit, hard, soft ):
        super().__setattr__( 'rank', rank )
        super().__setattr__( 'suit', suit )
        super().__setattr__( 'hard', hard )
        super().__setattr__( 'soft', soft )

    def __setattr__( self, name, value ):
        if name in self.__dict__:
            raise AttributeError( 'Cannot set {name}'.format(name = name ) )
        raise AttributeError( "'{__class__.__name__}' has no attribute '{name}'".format(
            __class__ = self.__class__,
            name = name
        ) )

    def __getattribute__( self, name ):
        if name.startswitch('_'): raise AttributeError
        return object.__getattribute__( self, name )


c = BlackJackCard3( 'A', '♠', 1, 11 )
# c.rank = 12
"""
一般情况下，不会轻易使用__getattribute__()。该函数的默认实现非常复杂，大多数情况下，
使用特性或改变__getattr__()函数的行为就足以满足需求了。
"""
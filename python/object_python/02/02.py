#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用特殊方法完成属性访问 '''

"""
3个用于属性访问的标准函数：__getattr__()、__setattr__()和__delattr__()。
还可以用__dir__()函数来查看属性的名称。

__setattr__()函数用于属性的创建和赋值。
__getattr__()函数可以用来做两件事。首先，如果属性已经被赋值，__getattr__()则不会
被调用，直接返回属性值即可。其次，如果属性没有被赋值，那么将使用__getattr__()函数的
返回值。如果找不到相关属性，要记得抛出AttributeError异常。
__delattr__()函数用于删除属性
__dir__()函数用于返回属性名称列表

__getattr__()函数只是复杂逻辑中的一个小步骤而已，仅当属性未知的情况下它才会被使用。如果
属性已知，这个函数将不会使用。__setattr__()函数和__delattr__()函数没有内部的处理过程，
也没有和其他函数逻辑有交互。

    扩展类。通过重写__setattr__()和__delattr__()函数使得它几乎是不可变的。也可以使用
    __slots__替换内部的__dict__对象

    封装类。提供对象（或对象集合）属性访问的代理实现。这可能需要完全重写和属性相关的那3个
    函数

    创建类并提供和特性功能一样的函数。使用这些方法来对特性逻辑集中处理。

    创建延迟计算属性，仅当需要时才触发计算过程。对于一些属性，它的值可能来自文件、数据库
    或网络。这是__getattr__()函数的常见用法。

    创建主动计算属性，其他属性更新时会相应地更新主动计算属性的值，这是通过重写__setattr
    __()函数实现的。
"""


''' 使用__slots__创建不可变对象 '''
class BlackJackCard:
    ''' Abstract Superclass '''
    __slots__ = ( 'rank', 'suit', 'hard', 'soft' )
    def __init__( self, rank, suit, hard, soft ):
        super().__setattr__( 'rank', rank )
        super().__setattr__( 'suit', suit )
        super().__setattr__( 'hard', hard )
        super().__setattr__( 'soft', soft )

    def __str__( self ):
        return '{0.rank} {0.suit}'.format( self )

    def __setattr__( self, name, value ):
        raise AttributeError( 
            "'{__class__.__name__}' has no attribue '{name}'".format(
                __class__ = self.__class__,
                name = name
            )
        )

""" 
如果一个属性是不允许被赋值或创建的，就被称为不可变。

有3处修改：
1. 把__slots__设为唯一被允许操作的属性。这会使得对象内部的__dict__对象不再有效并阻止对
其他属性的访问

2. 在__setattr__()函数中，代码逻辑仅仅是抛出异常。

3. 在__init__()函数中，调用了基类中__setattr__()实现，为了确保当类没有包含有效的__setattr__()
函数时，属性依然可被正确赋值。

如果需要，也可以像这样绕过不可变对象：
object.__setattr__(c, 'bad', 5)

如果不喜欢一个类的不可变性，可以修改它并移除重定义过的__setattr__()函数。一个类似的例子是：
__hash__()，不可变对象的目的是能够返回一致的值而非阻止写糟糕的代码。
"""

''' __slots__的主要目的是通过限制属性的数量来节约内存。 '''

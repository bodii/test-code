#!/usr/bin/env python3 
# -*- coding:utf-8 -*-

''' 循环引用和weakref模块 '''
"""
    如果我们需要循环引用，但是又希望将清理资源的代码写在__del__()中，这时候我们可以使用
    弱引用。循环引用的一个常用场景是互相引用： 一个父类中包含一个集合，集合中的每一个实例
    也包含一个指向父类的引用。如果一个Player对象中包含多个Hand实例，那么在每个Hand对象
    中都包括一个指向对应的Player类的引用可能会更方便。
    默认的对象间的引用可以被称为强引用，但是，叫直接引用可能更好。Python的引用计数机制会
    直接使用它们，而且如果引用计数无法删除这些对象的话，垃圾回收机器也能及时发现。它们是不
    可忽略的对象。
"""

# a = B()
# 变量a直接引用B类的一个对象。此时B的引用计数至少是1,因为a变量包含了一个指向它的引用。
""" 
想要找个弱引用相关的对象需要两个步骤。一个弱引用会调用x.parent(),这个函数将弱引用作为一个
可调用对象来查找它真正的父对象。这个过程让引用计数器得以归零，垃圾回收器可以回收引用的对象，
但是不回收这个弱引用。

weakref定义了一系列使用了弱引用而没有使用强引用的集合。它让我们可以创建一种特殊的字典类型，
当这种字典的对象没有用时，可以保证被垃圾回收。
"""

import weakref
class Parent2:
    def __init__( self, *children ):
        self.children = list(children)
        for child in self.children:
            child.parent = weakref.ref(self)

    def __del__( self ):
        print( 'Removing {__class__.__name__} {id:d}'.format(
            __class__ = self.__class__,
            id = id(self)
        ) )

''' 我们将child中的parent引用改为一个weakref对象的引用 '''

"""
p = self.parent()
if p is not None:
    # process p, the Parent instance
else:
    # the parent instance was garbage collected.

"""

class Child:
    def __del__( self ):
        print( 'Removing {__class__.__name__} {id:d}'.format(__class__= self.__class__,
        id= id(self)))

p = Parent2( Child(), Child() )
del p
"""
Removing Parent2 140392197220224
Removing Child 140392197220280
Removing Child 140392220907728
"""

''' 当一个weakref引用变成死引用时（因为引用被销毁了）， 我们有3个可能的方案
    1. 重新创建引用对象，或重新从数据库中加载
    2. 当垃圾回收器在低内存情况下错误地删除了一些对象时，使用warnings模块记录调试信息。
    3. 忽略这个问题。

    通常，weakref引用变成死引用是因为响应的对象已经被删除了。例如，变量的作用域已经执行结束，
    一个没有用的命令空间，应用程序正在关闭。对于这个原因，通常我们会采取第3种响应方法。因为试
    图创建这个引用的对象时很可能马上就会被删除。
 '''

#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 抽象基类设计的一致性 '''


"""
抽象基类

抽象基类的核心定义在一个名为abc的模块中。模块中包括了创建抽象基类需要的修饰符和元类型。

 [ 抽象基类的特性 ]
 1. 抽象意味着这些类中不包括我们需要的所有方法的定义。为了让它成为一个真正有用的子类，
 我们需要提供一些方法定义。

 2. 基类意味着其他类会把它当作基类来使用。
 
 3. 抽象类本身提供了一些方法的定义。更重要的是，抽象基类为缺失的方法函数提供了方法签名。
 子类必须提供正确的方法来创建符合抽象类定义的接口的具体类。

 [ 抽象类的设计初衷 ]

 1. 我们可以用它们为Python的内部类和我们的程序中的自定义类定义一组一致的基类。

 2. 我们可以用它创建一些通用的、可重用的抽象。

 3. 我们可以用它们来支持适当的类检查，以确定一个类的功能。这让我们的自定义类可以与
 库中的类更好地协作。为了正确地实现类检查，对于像"容器"和“数值”这种概念，我们最好
 可以有正式的定义。


 "使用抽象类，你就可以保证一个给定的类会有所有我们预期的行为。"
 """

 # 创建抽象类
 # 第一种方法
 import collections.abc
 class SomeApplicationClass( collections.abc ):
     pass

"""
一个函数是一个Callable类的具体例子。抽象基类是定义了__call__()方法的类。
"""

# 第二种方法

def some_method( self, other ):
    assertisinstance( other, collections.abc.Iterator )

"""
some_method()方法要求other参数是Iterator的一个子类。如果other参数无法通过这个测试，
我们会得到一个异常。另外一种更好的方法是在if语句中抛出一个TypeError。
"""

# 第三种方法
try:
    some_obj.some_method( another )
exceptAttributeError:
    warning.warn( '{0!r} not an Iterator, found {0.__class__.__bases__!r}'.format( another ) )
    raise


"""
asssertisinsatnce( some_argument, collections.abc.Container ), 
'{0!r} not a Container'.format(some_argument)

当有问题发生时，这段代码会抛出一个AssertionError异常。这样做的优点是用很短的代码就达到
了我们的目的。但是，这样的用法有两个缺点：assert语句有可能不会抛出异常，所以显式地抛出一
个TypeError异常可能会更好一些

如下 一个更好的例子：
if not isinstance(some_argument, collections.abc.Container):
    raise TypeError( '{0!r} not a Container'.format(some_argeument))


策略：
try:
    found = value in some_argument
exceptTypeError:
    if not isinstance(some_argument, collections.abc.Container):
        warnings.warn('{0!r} not a Container'.format(some_argument))
    raise

"""

def test(n):
    return n * n

print( isinstance(test, collections.abc.Callable) )
# Ture

print(abs(3))
# 3
print( isinstance(abs, collections.abc.Callable) )
# True

"""
所有的函数都认为自已属于Callable类。


#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第九章：元编程
Desc: 软件开发领域中最经典的口头禅就是“don’t repeat yourself”。也就是说，任何时
候当你的程序中存在高度重复 (或者是通过剪切复制) 的代码时，都应该想想是否有更
好的解决方案。在 Python 当中，通常都可以通过元编程来解决这类问题。简而言之，
元编程就是关于创建操作源代码 (比如修改、生成或包装原来的代码) 的函数和类。主
要技术是使用装饰器、类装饰器和元类。不过还有一些其他技术，包括签名对象、使
用 exec() 执行代码以及对内部函数和类的反射技术等。本章的主要目的是向大家介绍
这些元编程技术，并且给出实例来演示它们是怎样定制化你的源代码行为的。

Title;               避免重复的属性方法
Issue:你在类中需要重复的定义一些执行相同逻辑的属性方法，比如进行类型检查，怎样
去简化这些重复代码呢？
重载。但是你不确定应该怎样去实现（或者到底行得通不）。
Answer:
"""
class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age

    @property
    def name(self):
        return self._name

    @name.stter
    def name(self, value):
        if not isinstance(value, str):
            raise TypeError('name must be a string')
        self._name = value

    @property
    def age(self):
        return self._age

    @age.setter
    def age(self, value):
        if not isinstance(value, int):
            raise TypeError('age must be a integer')
        self._age = value

"""
可以看到，为了实现属性值的类型检查我们写了很多的重复代码。只要你以后看到
类似这样的代码，你都应该想办法去简化它。一个可行的方法是创建一个函数用来定
义属性并返回它。
"""
def typed_property(name, expected_type):
    storage_name = "_" + name

    @property
    def prop(self):
        return getattr(self, storage_name)

    @prop.setter
    def prop(self, value):
        if not isinstance(value, expected_type):
            raise TypeError('{} must ba a {}'.format(value, expected_type))
        setattr(self, storage_name, value)

    return prop

class Person:
    name = typed_property('name', str)
    age = typed_property('age', int)

    def __init__(self, name, age):
        self.name = name
        self.age = age

"""
本节我们演示内部函数或者闭包的一个重要特性，它们很像一个宏。例子中的函
数 typed property() 看上去有点难理解，其实它所做的仅仅就是为你生成属性并返
回这个属性对象。因此，当在一个类中使用它的时候，效果跟将它里面的代码放到
类定义中去是一样的。尽管属性的 getter 和 setter 方法访问了本地变量如 name ,
expected type 以及 storate name ，这个很正常，这些变量的值会保存在闭包当中。
我们还可以使用 functools.partial() 来稍稍改变下这个例子，很有趣。例如，
你可以像下面这样：
"""
from functools import partial

String = partial(typed_property, expected_type=str)
Integer = partial(typed_property, expected_type=int)

class Person:
    name = String('name')
    age = Integer('age')

    def __init__(self, name, age):
        self.name = name
        self.age = age

"""
其实你可以发现，这里的代码跟 8.13 小节中的类型系统描述器代码有些相似。
"""
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

Title;           定义有可选参数的元类
Issue:你想定义一个元类，允许类定义时提供可选参数，这样可以控制或配置类型的创建
过程。
Answer:‘‘metaclass‘‘关键字参数来指定特定的元类
"""
from abc import ABCMeta, abstractmethod

class IStream(metaclass=ABCMeta):

    @abstractmethod
    def read(self, maxsize=None):
        pass

    @abstractmethod
    def write(self, data):
        pass

"""
为了使元类支持这些关键字参数，你必须确保在__prepare__() ,__new__() 和
__init__() 方法中都使用强制关键字参数。就像下面这样：
"""
class MyMeta(type):

    @classmethod
    def __prepare__(cls, name, bases, *, debug=False, synchronize=False):
        pass
        return super().__prepare__(name,bases)

    def __new__(cls, name, bases, ns, *, debug=False, synchronize=False):
        pass
        return super().__new__(cls, name, bases, ns)

    def __init__(self, name, bases, ns, *, debug=False, synchronize=False):
        pass
        super().__init__(name, bases, ns)

"""
给一个元类添加可选关键字参数需要你完全弄懂类创建的所有步骤，因为这些参数
会被传递给每一个相关的方法。 prepare () 方法在所有类定义开始执行前首先被调
用，用来创建类命名空间。通常来讲，这个方法只是简单的返回一个字典或其他映射
对象。 new () 方法被用来实例化最终的类对象。它在类的主体被执行完后开始执
行。 init () 方法最后被调用，用来执行其他的一些初始化工作。
当我们构造元类的时候，通常只需要定义一个 new () 或 init () 方法，但不
是两个都定义。但是，如果需要接受其他的关键字参数的话，这两个方法就要同时提
供，并且都要提供对应的参数签名。默认的 prepare () 方法接受任意的关键字参
数，但是会忽略它们，所以只有当这些额外的参数可能会影响到类命名空间的创建时
你才需要去定义 prepare () 方法。
通过使用强制关键字参数，在类的创建过程中我们必须通过关键字来指定这些参
数。
"""

"""
使用关键字参数配置一个元类还可以视作对类变量的一种替代方式。
"""
class Spam(metaclass=MyMeta):
    debug = True
    synchronize = True
    pass

"""
将这些属性定义为参数的好处在于它们不会污染类的名称空间，这些属性仅仅只从
属于类的创建阶段，而不是类中的语句执行阶段。另外，它们在 prepare () 方法中
是可以被访问的，因为这个方法会在所有类主体执行前被执行。但是类变量只能在元
类的 new () 和 init () 方法中可见。
"""


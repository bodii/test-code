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

Title;              将装饰器定义为类的一部分
Issue:你想在类中定义装饰器，并将其作用在其他函数或方法上。
Answer:，装饰器要被定义成类方法并且你必须显式的使用父类名去调用它。
"""
from functools import wraps

class A:
    def decorator1(self, func):
        @wraps(func)
        def wrapper(*args, **kwargs):
            print('Decorator 1')
            return func(*args, **kwargs)
        return wrapper

    @classmethod
    def decorator2(cls, func):
        @wraps(func)
        def wrapper(*args, **kwargs):
            print('Decorator 2')
            return func(*args, **kwargs)
        return wrapper

a = A()
@a.decorator1
def spam():
    pass

@A.decorator2
def grok():
    pass

"""
仔细观察可以发现一个是实例调用，一个是类调用。
"""


"""
在类中定义装饰器初看上去好像很奇怪，但是在标准库中有很多这样的例子。特别
的，@property 装饰器实际上是一个类，它里面定义了三个方法 getter(), setter(),
deleter() , 每一个方法都是一个装饰器。
"""
class Person:
    first_name = property()

    @first_name.getter
    def first_name(self):
        return self._first_name

    @first_name.setter
    def first_name(self, value):
        if not isinstance(value, str):
            raise TypeError('Expected a string')
        self._first_name = value


"""
它为什么要这么定义的主要原因是各种不同的装饰器方法会在关联的 property 实
例上操作它的状态。因此，任何时候只要你碰到需要在装饰器中记录或绑定信息，那
么这不失为一种可行方法。
在类中定义装饰器有个难理解的地方就是对于额外参数 self 或 cls 的正确使用。
尽管最外层的装饰器函数比如 decorator1() 或 decorator2() 需要提供一个 self 或
cls 参数，但是在两个装饰器内部被创建的 wrapper() 函数并不需要包含这个 self 参
数。你唯一需要这个参数是在你确实要访问包装器中这个实例的某些部分的时候。其
他情况下都不用去管它。
对于类里面定义的包装器还有一点比较难理解，就是在涉及到继承的时候。例如，
假设你想让在 A 中定义的装饰器作用在子类 B 中。你需要像下面这样写：
"""
class B(A):
    @A.decorator2
    def bar(self):
        pass

"""
也就是说，装饰器要被定义成类方法并且你必须显式的使用父类名去调用它。你不
能使用 @B.decorator2 ，因为在方法定义时，这个类 B 还没有被创建。
"""


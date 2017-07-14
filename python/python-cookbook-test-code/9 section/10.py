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

Title;               为类和静态方法提供装饰器
Issue:你想给类或静态方法提供装饰器。
Answer:给类或静态方法提供装饰器是很简单的，不过要确保装饰器在 @classmethod 或
@staticmethod 之前。
"""
import time
from functools import wraps

def timethis(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        start = time.time()
        r = func(*args, **kwargs)
        end = time.time()
        print(end-start)
        return r
    return wrapper

class Spam:
    @timethis
    def instance_method(self, n):
        print(self, n)
        while n > 0:
            n -= 1

    @classmethod
    @timethis
    def class_method(cls, n):
        print(cls, n)
        while n > 0:
            n -= 1

    @staticmethod
    @timethis
    def static_method(n):
        print(n)
        while n > 0:
            n -= 1

"""
装饰后的类和静态方法可正常工作，只不过增加了额外的计时功能：
"""
s = Spam()
s.instance_method(1000000)
# <__main__.Spam object at 0x000002A9B5F9E3C8> 1000000
# 0.12602758407592773

Spam.class_method(1000000)
'''
<class '__main__.Spam'> 1000000
0.1371612548828125
'''
Spam.static_method(1000000)
'''
1000000
0.12651586532592773
'''

"""
如果你把装饰器的顺序写错了就会出错。例如，假设你像下面这样写：
"""
class Spam:

    @timethis
    @staticmethod
    def static_method(n):
        print(n)
        while n > 0:
            n -=1

s = Spam()
#s.static_method(1000000)
#TypeError: 'staticmethod' object is not callable
"""
问题在于 @classmethod 和 @staticmethod 实际上并不会创建可直接调用的对象，
而是创建特殊的描述器对象 (参考 8.9 小节)。因此当你试着在其他装饰器中将它们当
做函数来使用时就会出错。确保这种装饰器出现在装饰器链中的第一个位置可以修复
这个问题。
当我们在抽象基类中定义类方法和静态方法 (参考 8.12 小节) 时，这里讲到的知识
就很有用了。例如，如果你想定义一个抽象类方法，可以使用类似下面的代码：
"""
from abc import ABCMeta, abstractmethod
class A(metaclass=ABCMeta):
    @classmethod
    @abstractmethod
    def method(cls):
        pass
"""
在这段代码中，@classmethod 跟 @abstractmethod 两者的顺序是有讲究的，如果
你调换它们的顺序就会出错。
"""
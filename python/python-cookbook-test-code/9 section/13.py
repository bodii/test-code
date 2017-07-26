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

Title;            使用元类控制实例的创建
Issue:你想通过改变实例创建方式来实现单例、缓存或其他类似的特性。
Answer:如果你想自定义这个步骤，你可以定义一个元类并自己实现 call () 方法
"""
class Spam():

    def __init__(self, name):
        self.name = name

a = Spam("Guido")
b = Spam('Diana')

"""
如果你想自定义这个步骤，你可以定义一个元类并自己实现__call__() 方法。
"""

class NoInstances(type):
    def __call__(self, *args, **kwargs):
        raise TypeError("Can't instantiate directly")

class Spam(metaclass=NoInstances): # 使用‘‘metaclass‘‘关键字参数来指定特定的元类
    @staticmethod
    def grok(x):
        print('Spam.grok')
"""
这样的话，用户只能调用这个类的静态方法，而不能使用通常的方法来创建它的实
例。
"""
Spam.grok(42)
# Spam.grok

# s = Spam()
# TypeError: Can't instantiate directly

"""
现在，假如你想实现单例模式（只能创建唯一实例的类），实现起来也很简单
"""
class Singleton(type):

    def __init__(self, *args, **kwargs):
        self.__instance = None
        print(self.__instance)
        super().__init__(*args, **kwargs)

    def __call__(self, *args, **kwargs):
        if self.__instance is None:
            self.__initance = super().__call__(*args, **kwargs)
            return self.__instance
        else:
            return self.__instance

class Spam(metaclass=Singleton):

    def __init__(self):
        print('Creating Spam')

"""
那么 Spam 类就只能创建唯一的实例了
"""
a = Spam()
'''
None
Creating Spam
'''
b = Spam()
# Creating Spam
print(a is b)
# True
c = Spam()
# Creating Spam
print(a is c)
# True

"""
最后，假设你想创建 8.25 小节中那样的缓存实例。下面我们可以通过元类来实现
"""
import weakref

class Cached(type):

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.__cache = weakref.WeakValueDictionary()

    def __call__(self, *args):
        if args in self.__cache:
            print(self.__cache[args])
        else:
            obj = super().__call__(*args)
            self.__cache[args] = obj
            return obj

class Spam(metaclass=Cached):

    def __init__(self, name):
        print('Creating Spam({!r})'.format(name))
        self.name = name

a = Spam('Guido')
# Creating Spam('Guido')
b = Spam('Diana')
# Creating Spam('Diana')
c = Spam('Guido')

print(a is b)
# False
print(a is c)
# True

"""
利用元类实现多种实例创建模式通常要比不使用元类的方式优雅得多。
假设你不使用元类，你可能需要将类隐藏在某些工厂函数后面。比如为了实现一个
单例，你你可能会像下面这样写：
"""

class _Spam:

    def __init__(self):
        print('Creating Spam')

_spam_instance = None

def Spam():

    global _spam_instance

    if _spam_instance is not None:
        return _spam_instance
    else:
        _spam_instance = _Spam()
        return _spam_instance

"""
尽管使用元类可能会涉及到比较高级点的技术，但是它的代码看起来会更加简洁舒
服，而且也更加直观。
更多关于创建缓存实例、弱引用等内容，请参考 8.25 小节。
"""
Spam()
# Creating Spam
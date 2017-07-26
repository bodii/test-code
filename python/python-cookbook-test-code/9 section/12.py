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

Title;           使用装饰器扩充类的功能
Issue:你想通过反省或者重写类定义的某部分来修改它的行为，但是你又不希望使用继承
或元类的方式。
Answer:类装饰器
"""

def log_getattribute(cls):

    orig_getattribute = cls.__getattribute__

    def new_getattribute(self, name):
        print('getting:', name)
        return orig_getattribute(self, name)

    cls.__getattribute__ = new_getattribute
    return cls

@log_getattribute
class A:

    def __init__(self, x):
        self.x = x

    def spam(self):
        pass

a = A(42)
print(a.x)
'''
getting: x
42
'''
a.spam()
# getting: spam

"""
类装饰器通常可以作为其他高级技术比如混入或元类的一种非常简洁的替代方案。
比如，上面示例中的另外一种实现使用到继承：
"""
class LoggedGetattribute:

    def __getattribute__(self, name):
        print('getting:', name)
        return super().__getattribute__(name)

class A(LoggedGetattribute):

    def __init__(self, x):
        self.x = x

    def spam(self):
        pass

"""
这种方案也行得通，但是为了去理解它，你就必须知道方法调用顺序、super() 以
及其它 8.7 小节介绍的继承知识。某种程度上来讲，类装饰器方案就显得更加直观，并
且它不会引入新的继承体系。它的运行速度也更快一些，因为他并不依赖 super() 函
数。
如果你系想在一个类上面使用多个类装饰器，那么就需要注意下顺序问题。例如，
一个装饰器 A 会将其装饰的方法完整替换成另一种实现，而另一个装饰器 B 只是简单
的在其装饰的方法中添加点额外逻辑。那么这时候装饰器 A 就需要放在装饰器 B 的前
面。
你还可以回顾一下 8.13 小节另外一个关于类装饰器的有用的例子。
"""


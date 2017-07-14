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

Title;               将装饰器定义为类
Issue:你想使用一个装饰器去包装函数，但是希望返回一个可调用的实例。你需要让你的
装饰器可以同时工作在类定义的内部和外部。
Answer:为了将装饰器定义成一个实例，你需要确保它实现了 call () 和 get () 方法
"""
import types
from functools import wraps

class Profiled:
    def __init__(self, func):
        wraps(func)(self)
        self.ncalls = 0

    def __call__(self, *args, **kwargs):
        self.ncalls += 1
        return self.__wrapped__(*args, **kwargs)

    def __get__(self, instance, cls):
        if instance is None:
            return self
        else:
            return types.MethodType(self, instance)

"""
你可以将它当做一个普通的装饰器来使用，在类里面或外面都可以：
"""
@Profiled
def add(x, y):
    return x + y

class Spam:
    @Profiled
    def bar(self, x):
        print(self, x)

print(add(2, 3))
# 5
print(add(4, 5))
# 9
s = Spam()
s.bar(1)
# <__main__.Spam object at 0x000001E9D3572E80> 1
s.bar(2)
# <__main__.Spam object at 0x000001A719CC2E80> 2
s.bar(3)
# <__main__.Spam object at 0x00000241F6FB2E48> 3
print(Spam.bar.ncalls)
# 3

"""
将装饰器定义成类通常是很简单的。但是这里还是有一些细节需要解释下，特别是
当你想将它作用在实例方法上的时候。
首先，使用 functools.wraps() 函数的作用跟之前还是一样，将被包装函数的元
信息复制到可调用实例中去。
其次，通常很容易会忽视上面的 get () 方法。如果你忽略它，保持其他代码不
变再次运行，你会发现当你去调用被装饰实例方法时出现很奇怪的问题。例如：
>>> s = Spam()
>>> s.bar(3)
Traceback (most recent call last):
...
TypeError: bar() missing 1 required positional argument: 'x'
"""

"""
出错原因是当方法函数在一个类中被查找时，它们的 get () 方法依据描述器协
议被调用，在 8.9 小节已经讲述过描述器协议了。在这里， get () 的目的是创建一
个绑定方法对象 (最终会给这个方法传递 self 参数)。下面是一个例子来演示底层原理：
"""
s = Spam()
def grok(self, x):
    pass

print(grok.__get__(s, Spam))
# <bound method grok of <__main__.Spam object at 0x0000020688E12EB8>>

"""
get () 方法是为了确保绑定方法对象能被正确的创建。 type.MethodType() 手
动创建一个绑定方法来使用。只有当实例被使用的时候绑定方法才会被创建。如果这
个方法是在类上面来访问，那么 get () 中的 instance 参数会被设置成 None 并直接
返回 Profiled 实例本身。这样的话我们就可以提取它的 ncalls 属性了。
如果你想避免一些混乱，也可以考虑另外一个使用闭包和 nonlocal 变量实现的装
饰器，这个在 9.5 小节有讲到。
"""
def profiled(func):
    ncalls = 0
    @wraps(func)
    def wrapper(*args, **kwargs):
        nonlocal ncalls
        ncalls +=1
        return func(*args, **kwargs)
    wrapper.ncalls = lambda:ncalls
    return wrapper

@profiled
def add(x, y):
    return x + y
"""
这个方式跟之前的效果几乎一样，除了对于 ncalls 的访问现在是通过一个被绑定
为属性的函数来实现，例如：
"""
print(add(2, 3))
# 5
print(add(4, 5))
# 9
print(add.ncalls())
# 2



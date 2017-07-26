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

Title;            解除一个装饰器
Issue:一个装饰器已经作用在一个函数上，你想撤销它，直接访问原始的未包装的那个函
数。
Answer:假设装饰器是通过 @wraps (参考 9.2 小节) 来实现的，那么你可以通过访问
__wrapped__属性来访问原始函数：
"""

#@somedecorator
def add(x, y):
    return x + y
orig_add = add
print(orig_add(3, 4))
"""
直接访问未包装的原始函数在调试、内省和其他函数操作时是很有用的。但是我们
这里的方案仅仅适用于在包装器中正确使用了 @wraps 或者直接设置了 wrapped 属
性的情况。
如果有多个包装器，那么访问 wrapped 属性的行为是不可预知的，应该避免这
样做。在 Python3.3 中，它会略过所有的包装层，比如，假如你有如下的代码
"""
from functools import wraps

def decorator1(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        print('Decorator 1')
        return func(*args, **kwargs)
    return wrapper

def decorator2(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        print('Decorator 2')
        return func(*args, **kwargs)
    return wrapper

@decorator1
@decorator2
def add(x, y):
    return x + y

print(add(2, 3))
print(add.__wrapped__(2, 3))

"""
最后要说的是，并不是所有的装饰器都使用了 @wraps ，因此这里的方案并不全部
适用。特别的，内置的装饰器 @staticmethod 和 @classmethod 就没有遵循这个约定
(它们把原始函数存储在属性 func 中)。
"""
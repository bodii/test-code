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

Title;            创建装饰器时保留函数元信息
Issue:你写了一个装饰器作用在某个函数上，但是这个函数的重要的元信息比如名字、文
档字符串、注解和参数签名都丢失了。
Answer:任何时候你定义装饰器的时候，都应该使用 functools 库中的 @wraps 装饰器来注
解底层包装函数。
"""
import time
from functools import wraps

def timethis(func):

    @wraps(func)
    def wrapper(*args, **kwargs):
        start = time.time()
        result = func(*args, **kwargs)
        end = time.time()
        print(func.__name__, end -start)
        return result
    return wrapper

"""
下面我们使用这个被包装后的函数并检查它的元信息：
"""
@timethis
def countdown(n:int):
    '''
    Counts down
    '''
    while n > 0:
        n -= 1

countdown(100000)
print(countdown.__name__)
# countdown
print(countdown.__doc__)
#     Counts down
print(countdown.__annotations__)
# {'n': <class 'int'>}

"""
在编写装饰器的时候复制元信息是一个非常重要的部分。如果你忘记了使用 @wrap
，那么你会发现被装饰函数丢失了所有有用的信息。比如如果忽略 @wrap 后的效果是
下面这样的：
>>> countdown.__name__
'wrapper'
>>> countdown.__doc__
>>> countdown.__annotations__
{}
@wraps 有一个重要特征是它能让你通过属性 wrapped 直接访问被包装函数。
"""
countdown.__wrapped__(100000)

"""
__wrapped__ 属性还能让被装饰函数正确暴露底层的参数签名信息。例如：
"""
from inspect import signature
print(signature(countdown))
# (n:int)

"""
一个很普遍的问题是怎样让装饰器去直接复制原始函数的参数签名信息，如果想自
己手动实现的话需要做大量的工作，最好就简单的使用 wrapped 装饰器。通过底层
的 wrapped 属性访问到函数签名信息。更多关于签名的内容可以参考 9.16 小节。
"""

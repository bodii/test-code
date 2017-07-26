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

Title;            定义一个带参数的装饰器
Issue:你想定义一个可以接受参数的装饰器
Answer:定义一个接受参数的包装器看上去比较复杂主要是因为底层的调用序列。
"""
from functools import wraps
import logging

def logged(level, name=None, message=None):
    def decorate(func):
        logname = name if name else func.__module__
        log = logging.getLogger(logname)
        logmsg = message if message else func.__name__

        @wraps(func)
        def wrapper(*args, **kwargs):
            log.log(level, logmsg)
            return func(*args, **kwargs)
        return wrapper
    return decorate

@logged(logging.DEBUG)
def add(x, y):
    return x + y

@logged(logging.CRITICAL, 'example')
def spam():
    print('Spam!')
"""
初看起来，这种实现看上去很复杂，但是核心思想很简单。最外层的函数 logged()
接受参数并将它们作用在内部的装饰器函数上面。内层的函数 decorate() 接受一个函
数作为参数，然后在函数上面放置一个包装器。这里的关键点是包装器是可以使用传
递给 logged() 的参数的。
"""

#@decorator(x, y, z)
def func(a, b):
    pass

"""
装饰器处理过程跟下面的调用是等效的;
"""
def func(a, b):
    pass
#func = decorator(x, y, z)(func)
"""
decorator(x, y, z) 的返回结果必须是一个可调用对象，它接受一个函数作为参
数并包装它，可以参考 9.7 小节中另外一个可接受参数的包装器例子。
"""
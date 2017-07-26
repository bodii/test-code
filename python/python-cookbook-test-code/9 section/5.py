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

Title;            可自定义属性的装饰器
Issue:你想写一个装饰器来包装一个函数，并且允许用户提供参数在运行时控制装饰器行
为。
Answer:引入一个访问函数，使用 nolocal 来修改内部变量。然后这个访问函数被作为一
个属性赋值给包装函数。
"""

from functools import wraps, partial
import logging

def attach_wrapper(obj, func=None):
    if func is None:
        return partial(attach_wrapper, obj)
    setattr(obj, func.__name__, func)
    return func

def logged(level, name=None, message=None):
    def decorate(func):
        logname = name if name else func.__module__
        log = logging.getLogger(logname)
        logmsg = message if message else func.__name__

        @wraps(func)
        def wrapper(*args, **kwargs):
            log.log(level, logmsg)
            return func(*args, **kwargs)

        @attach_wrapper(wrapper)
        def set_level(newlevel):
            nonlocal level
            level = newlevel

        @attach_wrapper(wrapper)
        def set_message(newmsg):
            nonlocal logmsg
            logmsg = newmsg

        return wrapper
    return decorate

@logged(logging.DEBUG)
def add(x, y):
    return x + y

@logged(logging.CRITICAL, 'example')
def spam():
    print('Spam!')

"""
下面是交互环境下的使用例子：
"""
logging.basicConfig(level=logging.DEBUG)
print(add(2, 3))
'''
5
DEBUG:__main__:add
'''
add.set_message('Add called')
print(add(2, 3))
'''
DEBUG:__main__:Add called
5
'''
add.set_level(logging.WARNING)
print(add(2, 3))
'''
WARNING:__main__:Add called
5
'''

"""
这一小节的关键点在于访问函数 (如 set message() 和 set level() )，它们被作
为属性赋给包装器。每个访问函数允许使用 nonlocal 来修改函数内部的变量。
还有一个令人吃惊的地方是访问函数会在多层装饰器间传播 (如果你的装饰器都使
用了 @functools.wraps 注解)。例如，假设你引入另外一个装饰器，比如 9.2 小节中
的 @timethis ，像下面这样：

@timethis
@logged(logging.DEBUG)
def countdown(n):
    while n > 0:
        n -= 1

你会发现访问函数依旧有效：
>>> countdown(10000000)
DEBUG:__main__:countdown
countdown 0.8198461532592773
>>> countdown.set_level(logging.WARNING)
>>> countdown.set_message("Counting down to zero")
>>> countdown(10000000)
WARNING:__main__:Counting down to zero
countdown 0.8225970268249512
"""
"""
你还会发现即使装饰器像下面这样以相反的方向排放，效果也是一样的：
"""
import time
@logged(logging.DEBUG)
@timethis
def countdown(n):
    while n > 0:
        n -= 1

"""
还能通过使用 lambda 表达式代码来让访问函数的返回不同的设定值：
"""
@attach_wrapper(wrapper)
def get_level():
    return get_level
wrapper.get_level = lambda: get_level

"""
一个比较难理解的地方就是对于访问函数的首次使用。例如，你可能会考虑另外一
个方法直接访问函数的属性，
"""
@wraps(func)
def wrapper(*args, **kwargs):
    wrapper.log.log(wrapper.level, wrapper.logmsg)
    return func(*args, **kwargs)

wrapper.level = level
wrapper.logmsg = logmsg
wrapper.log = log
"""
这个方法也可能正常工作，但前提是它必须是最外层的装饰器才行。如果它的上面
还有另外的装饰器 (比如上面提到的 @timethis 例子)，那么它会隐藏底层属性，使得
修改它们没有任何作用。而通过使用访问函数就能避免这样的局限性。
最后提一点，这一小节的方案也可以作为 9.9 小节中装饰器类的另一种实现方法。
"""

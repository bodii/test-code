#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 带参数的装饰器 '''

"""
有时会希望传递更复杂的参数给装饰器。做法是修改封装的函数

例如：
@decorator(arg)
def func():
    pass

上面装饰器的用法是下在代码的一种简化版本：
def func():
    pass

func = decorator(arg)(func)

两个例子都做了3件事：
1. 定义了一个函数， func
2. 将抽象的装饰器应用于它的参数上，创建了一个具体的装饰器，decorator（arg）
3. 将具体的装饰器应用于函数上，创建了被装饰版本的函数，decorator(arg)(func)。

这意味着带参的装饰器需要间接地创建最后的函数。再修改调试装饰器为下面：

@debug('log_name')
def some_function( arg ):
    pass
"""

import functools

def decorator( config ):
    def concrete_decorator( function ):
        @functools.wraps( function )
        def wrapped( *args, **kw ):
            return function( *args, **kw )
        return wrapped
    return concrete_decorator

"""
装饰器定义（def decorator(config))规定了使用它时需要提供的参数。主体部分是具体
的装饰器，它会作为整个装饰器的返回值。将被应用于函数上的是具体的装饰器(def concrete
_decorator(function))。然后，和之前看到的简单装饰器类似。它创建并返回了一个封装好的
函数(def wrapped( *args, **kw) ),下面是带有记录器名称的调试装饰器。
"""
import logging
import sys
import os

def debug_named( log_name ):
    def concrete_decorator( function ):
        @functools.wraps( function )
        def wrapped( *args, **kw ):
            log = logging.getLogger( log_name )
            log.debug( '%s( %r, %r)', function.__name__, args, kw, )
            result = function( *args, **kw )
            log.debug('%s = %r', function.__name__, result)
            return result
        return wrapped
    return concrete_decorator

# 装饰器会添加许多调试信息。它们直接输出到名为recursion的日志中：
@debug_named('recursion')
def ackermann( m, n ):
    if m == 0: return n+1
    elif m > 0 and n == 0: return ackermann( m-1, 1 )
    elif m > 0 and n > 0 : return ackermann( m-1, ackermann( m, n-1))

logging.basicConfig(
    filename = os.path.dirname(__file__) + '/recursion', 
    datefmt='[%d/%b/%Y %H:%M:%S]',
    level=logging.DEBUG
    )


ackermann(2, 4)
#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 写一个简单的函数装饰器 '''

"""
一个decorator（装饰器）函数是一种用于返回新函数的函数（或者可调用对象）。最简单的只
需要一个参数：将被装饰的函数。装饰器的返回值是一个被包装的函数。基本上，额外的功能会
加到原始功能之前或者之后，这是函数中两个现成的连接点。

当定义一个装饰器时，我们会想确保装饰过的函数有原始函数的函数名和docstring。这些将被
用于写装饰函数的属性应该由装饰器来设定。用functools.wraps来写装饰器简化了所需要做
的工作，因为它已经为我们处理这些事情。

为了展示两个可以插入新功能的地方，可以创建一个调试跟踪装饰器，它会将一个函数的参数和
返回值写入日志。这个装饰器会在调用函数前和调用函数后分别插入新的功能。


import logging

logging.debug( "function(", args, kw, ")" )

result = some_function( *args, **kw )
logging.debug( 'result = ', result )
return result
这段代码展示了如何用新的处理逻辑封装原来的函数。
"""
import functools
import logging
import sys

def debug( function ):
    @functools.wraps( function )
    def logged_function( *args, **kw ):
        logging.debug( '%s( %r, %r )', function.__name__, args, kw, )
        result = function( *args, **kw )
        logging.debug( '%s = %r', function.__name__, result )
        return result
    return logged_function


@debug
def ackermann( m, n ):
    if m == 0: return n + 1
    elif m > 0 and n == 0: return ackermann( m-1, 1 )
    elif m > 0 and n > 0: return ackermann( m-1, ackermann(m, n-1) )

"""
这段代码包装了ackermann()函数，它用logging模块将调试信息写入到根记录中。会用
下面的方式配置日志记录器
"""
logging.basicConfig(stream=sys.stderr, level=logging.DEBUG)
print( ackermann(2, 4) )
"""
加入 logging.basicConfig后
logging.basicConfig是配置logging的输出
DEBUG:root:ackermann( (2, 4), {} )
DEBUG:root:ackermann( (2, 3), {} )
DEBUG:root:ackermann( (2, 2), {} )
DEBUG:root:ackermann( (2, 1), {} )
DEBUG:root:ackermann( (2, 0), {} )
DEBUG:root:ackermann( (1, 1), {} )
DEBUG:root:ackermann( (1, 0), {} )
DEBUG:root:ackermann( (0, 1), {} )
DEBUG:root:ackermann = 2
DEBUG:root:ackermann = 2
DEBUG:root:ackermann( (0, 2), {} )
DEBUG:root:ackermann = 3
DEBUG:root:ackermann = 3
DEBUG:root:ackermann = 3
DEBUG:root:ackermann( (1, 3), {} )
DEBUG:root:ackermann( (1, 2), {} )
DEBUG:root:ackermann( (1, 1), {} )
DEBUG:root:ackermann( (1, 0), {} )
DEBUG:root:ackermann( (0, 1), {} )
DEBUG:root:ackermann = 2
DEBUG:root:ackermann = 2
DEBUG:root:ackermann( (0, 2), {} )
DEBUG:root:ackermann = 3
DEBUG:root:ackermann = 3
DEBUG:root:ackermann( (0, 3), {} )
DEBUG:root:ackermann = 4
DEBUG:root:ackermann = 4
DEBUG:root:ackermann( (0, 4), {} )
DEBUG:root:ackermann = 5
DEBUG:root:ackermann = 5
DEBUG:root:ackermann = 5
DEBUG:root:ackermann( (1, 5), {} )
DEBUG:root:ackermann( (1, 4), {} )
DEBUG:root:ackermann( (1, 3), {} )
DEBUG:root:ackermann( (1, 2), {} )
DEBUG:root:ackermann( (1, 1), {} )
DEBUG:root:ackermann( (1, 0), {} )
DEBUG:root:ackermann( (0, 1), {} )
DEBUG:root:ackermann = 2
DEBUG:root:ackermann = 2
DEBUG:root:ackermann( (0, 2), {} )
DEBUG:root:ackermann = 3
DEBUG:root:ackermann = 3
DEBUG:root:ackermann( (0, 3), {} )
DEBUG:root:ackermann = 4
DEBUG:root:ackermann = 4
DEBUG:root:ackermann( (0, 4), {} )
DEBUG:root:ackermann = 5
DEBUG:root:ackermann = 5
DEBUG:root:ackermann( (0, 5), {} )
DEBUG:root:ackermann = 6
DEBUG:root:ackermann = 6
DEBUG:root:ackermann( (0, 6), {} )
DEBUG:root:ackermann = 7
DEBUG:root:ackermann = 7
DEBUG:root:ackermann = 7
DEBUG:root:ackermann( (1, 7), {} )
DEBUG:root:ackermann( (1, 6), {} )
DEBUG:root:ackermann( (1, 5), {} )
DEBUG:root:ackermann( (1, 4), {} )
DEBUG:root:ackermann( (1, 3), {} )
DEBUG:root:ackermann( (1, 2), {} )
DEBUG:root:ackermann( (1, 1), {} )
DEBUG:root:ackermann( (1, 0), {} )
DEBUG:root:ackermann( (0, 1), {} )
DEBUG:root:ackermann = 2
DEBUG:root:ackermann = 2
DEBUG:root:ackermann( (0, 2), {} )
DEBUG:root:ackermann = 3
DEBUG:root:ackermann = 3
DEBUG:root:ackermann( (0, 3), {} )
DEBUG:root:ackermann = 4
DEBUG:root:ackermann = 4
DEBUG:root:ackermann( (0, 4), {} )
DEBUG:root:ackermann = 5
DEBUG:root:ackermann = 5
DEBUG:root:ackermann( (0, 5), {} )
DEBUG:root:ackermann = 6
DEBUG:root:ackermann = 6
DEBUG:root:ackermann( (0, 6), {} )
DEBUG:root:ackermann = 7
DEBUG:root:ackermann = 7
DEBUG:root:ackermann( (0, 7), {} )
DEBUG:root:ackermann = 8
DEBUG:root:ackermann = 8
DEBUG:root:ackermann( (0, 8), {} )
DEBUG:root:ackermann = 9
DEBUG:root:ackermann = 9
DEBUG:root:ackermann = 9
DEBUG:root:ackermann( (1, 9), {} )
DEBUG:root:ackermann( (1, 8), {} )
DEBUG:root:ackermann( (1, 7), {} )
DEBUG:root:ackermann( (1, 6), {} )
DEBUG:root:ackermann( (1, 5), {} )
DEBUG:root:ackermann( (1, 4), {} )
DEBUG:root:ackermann( (1, 3), {} )
DEBUG:root:ackermann( (1, 2), {} )
DEBUG:root:ackermann( (1, 1), {} )
DEBUG:root:ackermann( (1, 0), {} )
DEBUG:root:ackermann( (0, 1), {} )
DEBUG:root:ackermann = 2
DEBUG:root:ackermann = 2
DEBUG:root:ackermann( (0, 2), {} )
DEBUG:root:ackermann = 3
DEBUG:root:ackermann = 3
DEBUG:root:ackermann( (0, 3), {} )
DEBUG:root:ackermann = 4
DEBUG:root:ackermann = 4
DEBUG:root:ackermann( (0, 4), {} )
DEBUG:root:ackermann = 5
DEBUG:root:ackermann = 5
DEBUG:root:ackermann( (0, 5), {} )
DEBUG:root:ackermann = 6
DEBUG:root:ackermann = 6
DEBUG:root:ackermann( (0, 6), {} )
DEBUG:root:ackermann = 7
DEBUG:root:ackermann = 7
DEBUG:root:ackermann( (0, 7), {} )
DEBUG:root:ackermann = 8
DEBUG:root:ackermann = 8
DEBUG:root:ackermann( (0, 8), {} )
DEBUG:root:ackermann = 9
DEBUG:root:ackermann = 9
DEBUG:root:ackermann( (0, 9), {} )
DEBUG:root:ackermann = 10
DEBUG:root:ackermann = 10
DEBUG:root:ackermann( (0, 10), {} )
DEBUG:root:ackermann = 11
DEBUG:root:ackermann = 11
DEBUG:root:ackermann = 11
"""
"""
不加：
11
"""
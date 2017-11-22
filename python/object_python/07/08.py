#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 创建独立的日志记录器 '''

# 装饰器为每个函数创建一个独立的日志记录器
import functools
import logging
import sys

def debug2( function ):
    @functools.wraps( function )
    def logged_function( *args, **kw ):
        log = logging.getLogger( function.__name__ )
        log.debug( 'call( %r, %r )', args, kw, )
        result = function( *args, **kw )
        log.debug( 'result %s', result )
        return result
    return logged_function

@debug2
def ackermann( m, n ):
    if m == 0: return n + 1
    elif m > 0 and n == 0: return ackermann( m-1, 1 )
    elif m > 0 and n > 0: return ackermann( m-1, ackermann(m, n-1) )

logging.basicConfig(stream=sys.stderr, level=logging.DEBUG)

print( ackermann(2, 4) )
"""
DEBUG:ackermann:call( (2, 4), {} )
DEBUG:ackermann:call( (2, 3), {} )
DEBUG:ackermann:call( (2, 2), {} )
DEBUG:ackermann:call( (2, 1), {} )
DEBUG:ackermann:call( (2, 0), {} )
DEBUG:ackermann:call( (1, 1), {} )
DEBUG:ackermann:call( (1, 0), {} )
DEBUG:ackermann:call( (0, 1), {} )
DEBUG:ackermann:result 2
DEBUG:ackermann:result 2
DEBUG:ackermann:call( (0, 2), {} )
DEBUG:ackermann:result 3
DEBUG:ackermann:result 3
DEBUG:ackermann:result 3
DEBUG:ackermann:call( (1, 3), {} )
DEBUG:ackermann:call( (1, 2), {} )
DEBUG:ackermann:call( (1, 1), {} )
DEBUG:ackermann:call( (1, 0), {} )
DEBUG:ackermann:call( (0, 1), {} )
DEBUG:ackermann:result 2
DEBUG:ackermann:result 2
DEBUG:ackermann:call( (0, 2), {} )
DEBUG:ackermann:result 3
DEBUG:ackermann:result 3
DEBUG:ackermann:call( (0, 3), {} )
DEBUG:ackermann:result 4
DEBUG:ackermann:result 4
DEBUG:ackermann:call( (0, 4), {} )
DEBUG:ackermann:result 5
DEBUG:ackermann:result 5
DEBUG:ackermann:result 5
DEBUG:ackermann:call( (1, 5), {} )
DEBUG:ackermann:call( (1, 4), {} )
DEBUG:ackermann:call( (1, 3), {} )
DEBUG:ackermann:call( (1, 2), {} )
DEBUG:ackermann:call( (1, 1), {} )
DEBUG:ackermann:call( (1, 0), {} )
DEBUG:ackermann:call( (0, 1), {} )
DEBUG:ackermann:result 2
DEBUG:ackermann:result 2
DEBUG:ackermann:call( (0, 2), {} )
DEBUG:ackermann:result 3
DEBUG:ackermann:result 3
DEBUG:ackermann:call( (0, 3), {} )
DEBUG:ackermann:result 4
DEBUG:ackermann:result 4
DEBUG:ackermann:call( (0, 4), {} )
DEBUG:ackermann:result 5
DEBUG:ackermann:result 5
DEBUG:ackermann:call( (0, 5), {} )
DEBUG:ackermann:result 6
DEBUG:ackermann:result 6
DEBUG:ackermann:call( (0, 6), {} )
DEBUG:ackermann:result 7
DEBUG:ackermann:result 7
DEBUG:ackermann:result 7
DEBUG:ackermann:call( (1, 7), {} )
DEBUG:ackermann:call( (1, 6), {} )
DEBUG:ackermann:call( (1, 5), {} )
DEBUG:ackermann:call( (1, 4), {} )
DEBUG:ackermann:call( (1, 3), {} )
DEBUG:ackermann:call( (1, 2), {} )
DEBUG:ackermann:call( (1, 1), {} )
DEBUG:ackermann:call( (1, 0), {} )
DEBUG:ackermann:call( (0, 1), {} )
DEBUG:ackermann:result 2
DEBUG:ackermann:result 2
DEBUG:ackermann:call( (0, 2), {} )
DEBUG:ackermann:result 3
DEBUG:ackermann:result 3
DEBUG:ackermann:call( (0, 3), {} )
DEBUG:ackermann:result 4
DEBUG:ackermann:result 4
DEBUG:ackermann:call( (0, 4), {} )
DEBUG:ackermann:result 5
DEBUG:ackermann:result 5
DEBUG:ackermann:call( (0, 5), {} )
DEBUG:ackermann:result 6
DEBUG:ackermann:result 6
DEBUG:ackermann:call( (0, 6), {} )
DEBUG:ackermann:result 7
DEBUG:ackermann:result 7
DEBUG:ackermann:call( (0, 7), {} )
DEBUG:ackermann:result 8
DEBUG:ackermann:result 8
DEBUG:ackermann:call( (0, 8), {} )
DEBUG:ackermann:result 9
DEBUG:ackermann:result 9
DEBUG:ackermann:result 9
DEBUG:ackermann:call( (1, 9), {} )
DEBUG:ackermann:call( (1, 8), {} )
DEBUG:ackermann:call( (1, 7), {} )
DEBUG:ackermann:call( (1, 6), {} )
DEBUG:ackermann:call( (1, 5), {} )
DEBUG:ackermann:call( (1, 4), {} )
DEBUG:ackermann:call( (1, 3), {} )
DEBUG:ackermann:call( (1, 2), {} )
DEBUG:ackermann:call( (1, 1), {} )
DEBUG:ackermann:call( (1, 0), {} )
DEBUG:ackermann:call( (0, 1), {} )
DEBUG:ackermann:result 2
DEBUG:ackermann:result 2
DEBUG:ackermann:call( (0, 2), {} )
DEBUG:ackermann:result 3
DEBUG:ackermann:result 3
DEBUG:ackermann:call( (0, 3), {} )
DEBUG:ackermann:result 4
DEBUG:ackermann:result 4
DEBUG:ackermann:call( (0, 4), {} )
DEBUG:ackermann:result 5
DEBUG:ackermann:result 5
DEBUG:ackermann:call( (0, 5), {} )
DEBUG:ackermann:result 6
DEBUG:ackermann:result 6
DEBUG:ackermann:call( (0, 6), {} )
DEBUG:ackermann:result 7
DEBUG:ackermann:result 7
DEBUG:ackermann:call( (0, 7), {} )
DEBUG:ackermann:result 8
DEBUG:ackermann:result 8
DEBUG:ackermann:call( (0, 8), {} )
DEBUG:ackermann:result 9
DEBUG:ackermann:result 9
DEBUG:ackermann:call( (0, 9), {} )
DEBUG:ackermann:result 10
DEBUG:ackermann:result 10
DEBUG:ackermann:call( (0, 10), {} )
DEBUG:ackermann:result 11
DEBUG:ackermann:result 11
DEBUG:ackermann:result 11
11
"""

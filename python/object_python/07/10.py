#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 创建方法函数装饰器 '''

"""
一个类中方法函数的装饰器和一个单独的函数的装饰器是一样的，只是在不同的上下文中使用。
这种上下文所带来的一个轻微的后果是必须经常显式地声明self变量。

方法函数装饰器的一个应用是追踪对象状态的改变。

当设计一个有状态的类时，任何setter方法都会带来状态的改变。这些setter方法通常会用
@property装饰器，这样它们就可以被当作简单的属性来使用。如果我们这么做，我们可以添加
一个@audit装饰器，这样就能合理地追踪所有的改变。
"""
import functools
import logging

def audit( mithod ):
    @functools.wraps(method)
    def wrapper( self, *args, **kw ):
        audit_log = logging.getLogger( 'audit' )
        before = repr( self )
        try: 
            result = method( self, *args, **kw )
            after = repr(self)
        except Exception as e:
            audit_log.exception(
                '%s before %s\n after %s',
                method.__qualname__, 
                before,
                after
            )
            raise
        audit_log.info(
            '%s before %s\n after %s',
            method.__qualname__, 
            before,
            after
        )
        return result
    return wrapper
"""
我们创建了一个对象被修改前的文本表示。然后，调用原始的方法函数。如果有异常，会生成一个
包含异常信息的追踪日志。否则，会在日志中生成一个INFO条目，它包含方法的全名、修改前的
文本表示和修改后的文本表示。
"""

class Hand:
    def __init__( self, *cards ):
        self._cards = list(cards)
    
    @audit
    def __iadd__( self, card ):
        self._cards.append(card)
        return self

    def __repr__( self ):
        cards = ', '.join( map(str, self._cards) )
        return '{__class__.__name__}({cards})'.format(
            __class__=self.__class__,
            cards = cards
        )
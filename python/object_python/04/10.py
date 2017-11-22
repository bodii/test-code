#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 可调用对象的设计要素和折中方案 '''

"""
在设计一个可调用对象时，需要考虑以下几点：
1. 首先是API。如果一个对象使用函数式接口更好一些，那么使用可调用对象是合理的。通过使用
collections.abc.Callable来确保可调用API被正确地创建，并且可以让读代码的人很明确地了
解类的目的。
2. 其次是函数的状态。Python中的普通函数没有迟滞性————没有保存的状态，而可调用对象可以
保存状态，记忆化设计模式很好地应用了有状态的可调用对象。
"""
import collections
# 把一个普通函数转换为可调用对象：
def x( args ):
    body

class X( collections.abc.Callable ):
    def __call__( self, args ):
        body

x = X()



''' 上下文管理器的设计要素和折中方案 '''

"""
上下文通常用于获取/释放，打开/关闭和加锁和解锁这类的操作对。大多数操作是与文件的I/O相关的，
而且Python中的大多数文件操作对象已经是不错的上下文管理器了。
对于包含了多个处理步骤，而每步又包含了括号的逻辑来说，上下文管理器总是需要的。特别是对于最终
需要调用close()方法的逻辑，应该包含在上下文管理器中。
Python中的一些类库提供了打开/关闭操作，可其中的对象并不是上下文对象。比如对于shelve模块，
可以在操作shelve文件时使用contextllib.closing()上下文 。
"""

import contextlib

with contextlib.closing( myClass() ) as my_object:
    process( my_object )


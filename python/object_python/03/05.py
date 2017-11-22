#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 上下文和上下文管理器 '''

"""
一个上下文管理是和with语句一起使用的。

# 如：
with function(arg) as context:
    process( context )
"""

"""
一个常用的上下文管理器是一个文件。当打开文件时，我们应该创建一个会自动关闭文件的上下文
管理器。所以，我们几乎应该总是以下面的方式操作文件：
"""
with open('some file') as the_file:
    process(the_file)

"""
在with语句中的代码执行完成后，Python可以保证文件一定会正确地关闭。Contextlib模块提供
了一些用于创建上下文管理器的工具。这个库没有提供任何抽象基类，而是提供了装饰器和contextlib.
ContextDecorator基类，其中装饰器会将简单的函数转成上下文管理器，
contextlib.ContextDecorator基类可以被扩展并创建一个上下文管理器类。
"""

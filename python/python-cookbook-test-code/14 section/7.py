#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十四章：测试、调试和异常
Description: 试验还是很棒的，但是调试？就没那么有趣了。事实是，在 Python 测试代码之前
没有编译器来分析你的代码，因此使的测试成为开发的一个重要部分。本章的目标是
讨论一些关于测试、调试和异常处理的常见问题。但是并不是为测试驱动开发或者单
元测试模块做一个简要的介绍。因此，笔者假定读者熟悉测试概念。

Title:              捕获所有异常
Issue: 怎样捕获代码中的所有异常？
Answer: 想要捕获所有的异常，可以直接捕获 Exception 即可
try:
    ...
except Exception as e:
    ...
    log('Reason:', e)
"""
"""

捕获所有异常通常是由于程序员在某些复杂操作中并不能记住所有可能的异常。如
果你不是很细心的人，这也是编写不易调试代码的一个简单方法。
正因如此，如果你选择捕获所有异常，那么在某个地方（比如日志文件、打印异常
到屏幕）打印确切原因就比较重要了。如果你没有这样做，有时候你看到异常打印时
可能摸不着头脑，就像下面这样
"""

def parse_int(s):
    try:
       n = int(v)
    except Exception:
        print("Gouldn't parse")


parse_int('n/a')
# Gouldn't parse
parse_int('42')
# Gouldn't parse
parse_int(42)
# Gouldn't parse

"""
这时候你就会挠头想：“这咋回事啊？”假如你像下面这样重写这个函数
"""
def parse_int(s):
    try:
        n = int(v)
    except Exception as e:
        print("Couldn't parse")
        print('Reason:', e)

"""
这时候你能获取如下输出，指明了有个编程错误：
"""
parse_int('42')
# Reason: name 'v' is not defined

"""
很明显，你应该尽可能将异常处理器定义的精准一些。不过，要是你必须捕获所有
异常，确保打印正确的诊断信息或将异常传播出去，这样不会丢失掉异常。
"""

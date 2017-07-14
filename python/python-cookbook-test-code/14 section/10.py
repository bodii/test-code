#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十四章：测试、调试和异常
Description: 试验还是很棒的，但是调试？就没那么有趣了。事实是，在 Python 测试代码之前
没有编译器来分析你的代码，因此使的测试成为开发的一个重要部分。本章的目标是
讨论一些关于测试、调试和异常处理的常见问题。但是并不是为测试驱动开发或者单
元测试模块做一个简要的介绍。因此，笔者假定读者熟悉测试概念。

Title:              重新抛出被捕获的异常
Issue: 你在一个 except 块中捕获了一个异常，现在想重新抛出它。
Answer: 简单的使用一个单独的 rasie 语句即可
"""
def example():
    try:
        int('N/A')
    except ValueError:
        print("Didn't work")
        raise


#example()
# ValueError: invalid literal for int() with base 10: 'N/A'

"""
这个问题通常是当你需要在捕获异常后执行某个操作（比如记录日志、清理等），
但是之后想将异常传播下去。一个很常见的用法是在捕获所有异常的处理器中
try:
    ...
except Exception as e:
    ...
    raise
"""
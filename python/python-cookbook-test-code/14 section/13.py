#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十四章：测试、调试和异常
Description: 试验还是很棒的，但是调试？就没那么有趣了。事实是，在 Python 测试代码之前
没有编译器来分析你的代码，因此使的测试成为开发的一个重要部分。本章的目标是
讨论一些关于测试、调试和异常处理的常见问题。但是并不是为测试驱动开发或者单
元测试模块做一个简要的介绍。因此，笔者假定读者熟悉测试概念。

Title:           给你的程序做性能测试
Issue: 你想测试你的程序运行所花费的时间并做性能测试。
Answer: 如果你只是简单的想测试下你的程序整体花费的时间，通常使用 Unix 时间函数就
行了
"""
"""
比如：
bash % time python3 someprogram.py
real 0m13.937s
user 0m12.162s
sys 0m0.098s
bash %

如果你还需要一个程序各个细节的详细报告，可以使用 cProfile 模块
bash % python3 -m cProfile someprogram.py
859647 function calls in 16.016 CPU seconds
Ordered by: standard name
ncalls tottime percall cumtime percall filename:lineno(function)
263169 0.080 0.000 0.080 0.000 someprogram.py:16(frange)
513 0.001 0.000 0.002 0.000 someprogram.py:30(generate_mandel)
262656 0.194 0.000 15.295 0.000 someprogram.py:32(<genexpr>)
1 0.036 0.036 16.077 16.077 someprogram.py:4(<module>)
262144 15.021 0.000 15.021 0.000 someprogram.py:4(in_mandelbrot)
1 0.000 0.000 0.000 0.000 os.py:746(urandom)
1 0.000 0.000 0.000 0.000 png.py:1056(_readable)
1 0.000 0.000 0.000 0.000 png.py:1073(Reader)
1 0.227 0.227 0.438 0.438 png.py:163(<module>)
512 0.010 0.000 0.010 0.000 png.py:200(group)
...
bash %


不过通常情况是介于这两个极端之间。比如你已经知道代码运行时在少数几个函数
中花费了绝大部分时间。对于这些函数的性能测试，可以使用一个简单的装饰器：
import time
from functools import wraps
def timethis(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        start = time.perf_counter()
        r = func(*args, **kwargs)
        end = time.perf_counter()
        print('{}.{} :{}'.format(func.__module__, func.__name__, end - start))
        return r
    return wrapper
"""

import time
from functools import wraps
def timethis(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        start = time.perf_counter()
        r = func(*args, **kwargs)
        end = time.perf_counter()
        print('{}.{} :{}'.format(func.__module__, func.__name__, end - start))
        return r
    return wrapper


"""要使用这个装饰器，只需要将其放置在你要进行性能测试的函数定义前即可，比
如："""
@timethis
def countdown(n):
    while n > 0:
        n -= 1

countdown(10000000)
# __main__.countdown :0.6399373729273985

"""
要测试某个代码块运行时间，你可以定义一个上下文管理器，例如：
"""
from contextlib import contextmanager

@contextmanager
def timeblock(label):
    start = time.perf_counter()
    try:
        yield
    finally:
        end = time.perf_counter()
        print('{} :{}'.format(label, end - start))

"""
下面是使用这个上下文管理器的例子:
"""
with timeblock('counting'):
    n = 10000000
    while n > 0:
        n -= 1

"""
对于测试很小的代码片段运行性能，使用 timeit 模块会很方便，例如：
"""
from timeit import timeit
timeit('math.sqrt(2)', 'import math')
timeit('sqrt(2)', 'from math import sqrt')

"""
timeit 会执行第一个参数中语句 100 万次并计算运行时间。第二个参数是运行测
试之前配置环境。如果你想改变循环执行次数，可以像下面这样设置 number 参数的值
"""
timeit('math.sqrt(2)', 'import math', number=10000000)
timeit('sqrt(2)', 'from math import sqrt', number=10000000)

"""
当 执 行 性 能 测 试 的 时 候， 需 要 注 意 的 是 你 获 取 的 结 果 都 是 近 似 值。
time.perf counter() 函数会在给定平台上获取最高精度的计时值。不过，它仍然
还是基于时钟时间，很多因素会影响到它的精确度，比如机器负载。如果你对于执行
时间更感兴趣，使用 time.process_time() 来代替它。例如：
"""
from functools import wraps
def timethis(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        start = time.process_time()
        r = func(*args, **kwargs)
        end = time.process_time()
        print('{}.{}: {}'.format(func.__module__, func.__name__, end - start))
        return r
    return wrapper

"""
最后， 如果你想进行更深入的性能分析，那么你需要详细阅读 time 、 timeit 和其
他相关模块的文档。这样你可以理解和平台相关的差异以及一些其他陷阱。还可以参
考 13.13 小节中相关的一个创建计时器类的例子
"""
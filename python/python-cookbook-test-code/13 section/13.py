#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十三章：脚本编程与系统管理
Description: 许多人使用 Python 作为一个 shell 脚本的替代，用来实现常用系统任务的自动化，
如文件的操作，系统的配置等。本章的主要目标是描述光宇编写脚本时候经常遇到的
一些功能。例如，解析命令行选项、获取有用的系统配置数据等等。第 5 章也包含了
与文件和目录相关的一般信息。

Title:           实现一个计时器
Issue:你想记录程序执行多个任务所花费的时间
Answer: time 模块包含很多函数来执行跟时间有关的函数。
"""
import time

class Timer:
    def __init__(self, func=time.perf_counter()):
        self.elapsed = 0.0
        self._func = func
        self._start = None

    def start(self):
        if self._start is not None:
            raise RuntimeError('Already started')
        self._start = self._func()

    def stop(self):
        if self._start is None:
            raise RuntimeError('Not started')
        end = self._func()
        self.elapsed += end - self._start
        self._start = None

    def reset(self):
        self.elapsed = 0.0

    @property
    def running(self):
        return self._start is not None

    def __enter__(self):
        self.start()
        return self

    def __exit__(self, *args):
        self.stop()

"""这个类定义了一个可以被用户根据需要启动、停止和重置的计时器。它会在
elapsed 属性中记录整个消耗时间。下面是一个例子来演示怎样使用它："""

def countdown(n):
    while n > 0:
        n -= 1

t = Timer()
t.start()
countdown(1000000)
t.stop()
print(t.elapsed)

with t:
    countdown(1000000)

print(t.elapsed)
with Timer() as t2:
    countdown(10000000)
print(t2.elapsed)

"""
本节提供了一个简单而实用的类来实现时间记录以及耗时计算。同时也是对使用
with 语句以及上下文管理器协议的一个很好的演示。

在 计 时 中 要 考 虑 一 个 底 层 的 时 间 函 数 问 题。 一 般 来 说， 使 用 time.time()
或 time.clock() 计 算 的 时 间 精 度 因 操 作 系 统 的 不 同 会 有 所 不 同。 而 使 用
time.perf counter() 函数可以确保使用系统上面最精确的计时器。
上述代码中由 Timer 类记录的时间是钟表时间，并包含了所有休眠时间。如果你
只想计算该进程所花费的 CPU 时间，应该使用 time.process time() 来代替
"""
t = Timer(time.process_time())
with t:
    countdown(10000000)
print(t.elapsed)

"""
time.perf counter() 和 time.process time() 都会返回小数形式的秒数时间。实
际的时间值没有任何意义，为了得到有意义的结果，你得执行两次函数然后计算它们
的差值。
更多关于计时和性能分析的例子请参考 14.13 小节。
"""
#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第九章：元编程
Desc: 软件开发领域中最经典的口头禅就是“don’t repeat yourself”。也就是说，任何时
候当你的程序中存在高度重复 (或者是通过剪切复制) 的代码时，都应该想想是否有更
好的解决方案。在 Python 当中，通常都可以通过元编程来解决这类问题。简而言之，
元编程就是关于创建操作源代码 (比如修改、生成或包装原来的代码) 的函数和类。主
要技术是使用装饰器、类装饰器和元类。不过还有一些其他技术，包括签名对象、使
用 exec() 执行代码以及对内部函数和类的反射技术等。本章的主要目的是向大家介绍
这些元编程技术，并且给出实例来演示它们是怎样定制化你的源代码行为的。

Title;               定义上下文管理器的简单方法
Issue:你想自己去实现一个新的上下文管理器，以便使用 with 语句。
Answer:实现一个新的上下文管理器的最简单的方法就是使用 contexlib 模块中的
@contextmanager 装饰器。
"""
import time
from contextlib import contextmanager

@contextmanager
def timethis(label):
    start = time.time()
    try:
        yield
    finally:
        end = time.time()
        print('{}: {}'.format(label, end - start))

with timethis('counting'):
    n = 10000000
    while n > 0:
        n -= 1

# counting: 1.9710750579833984

"""
在函数 timethis() 中，yield 之前的代码会在上下文管理器中作为 enter ()
方法执行，所有在 yield 之后的代码会作为 exit () 方法执行。如果出现了异常，
异常会在 yield 语句那里抛出。
下面是一个更加高级一点的上下文管理器，实现了列表对象上的某种事务：
"""
@contextmanager
def list_transaction(orig_list):
    working = list(orig_list)
    yield working
    orig_list[:] = working
"""
这段代码的作用是任何对列表的修改只有当所有代码运行完成并且不出现异常的情
况下才会生效。
"""
itmes = [1, 2, 3]
with list_transaction(itmes) as working:
    working.append(4)
    working.append(5)

print(itmes)
# [1, 2, 3, 4, 5]
with list_transaction(itmes) as working:
    working.append(6)
    working.append(7)
    raise RuntimeError('oops')
# RuntimeError: oops

"""
通常情况下，如果要写一个上下文管理器，你需要定义一个类，里面包含一个
__enter__() 和一个__exit__() 方法，如下所示：
"""
class timethis:
    def __init__(self, label):
        self.label = label

    def __enter__(self):
        self.start = time.time()

    def __exit__(self, exc_ty, exc_val, exc_td):
        end = time.time()
        print('{}: {}'.format(self.label, end - self.start))

"""
尽管这个也不难写，但是相比较写一个简单的使用 @contextmanager 注解的函数
而言还是稍显乏味。
@contextmanager 应该仅仅用来写自包含的上下文管理函数。如果你有一些对
象 (比如一个文件、网络连接或锁)，需要支持 with 语句，那么你就需要单独实现
__enter__() 方法和__exit__() 方法。
"""
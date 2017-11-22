#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用pickle进行转储和加载 '''

"""
pick 模块是Python内部的一种模式，用来完成对象的持久化
Python标准库中是这样描述pickle的：
pickle模块可以将一个复杂的对象转换为一个字节数组并且使用相同的内部结构将字节流转换为一个
对象。将这些字节流写入文件或许是最常见的场景，但也可能输出到网络进行传输或是数据库。

pickle模块用很多方式完成了与Python的轻量集成。例如，一个类的__reduce__()和__reduce_ex
__()方法用于提供对pickle处理过程的支持。
"""


import pickle
import os.path

travel = {'a': 1, 'b': 2}

with open(os.path.join(os.path.dirname(__file__) , 'travel_blog.p'), 'wb') as target:
    pickle.dump( travel, target )


with open(os.path.join(os.path.dirname(__file__), 'travel_blog.p'), 'rb') as source:
    copy = pickle.load( source )

print(copy)

"""
由于pickle数据是使用字节写入的，文件必须以”rb“模式打开。pickle对象将被正确地绑定于适当的类字义。
底层的字节流不是用来直接读的。必须经过适当调整才具备可读性，但设计它的初衷不是为了像YAML一样可读。
"""



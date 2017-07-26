#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第四章：迭代器与生成器
Desc: 迭代是 Python 最强大的功能之一。初看起来，你可能会简单的认为迭代只不过是
处理序列中元素的一种方法。然而，绝非仅仅就是如此，还有很多你可能不知道的，
比如创建你自己的迭代器对象，在 itertools 模块中使用有用的迭代模式，构造生成器
函数等等。这一章目的就是向你展示跟迭代有关的各种常见问题。

Title;        使用生成器创建新的迭代模式
Issue: 你想实现一个自定义迭代模式，跟普通的内置函数比如 range() , reversed() 不
一样。
Answer: 如果你想实现一种新的迭代模式，使用一个生成器函数来定义它。
去。
"""

'''
下面是一个生产某个范围内浮点数的生成器：
'''

def frange(start, stop, increment):
    x = start
    while x < stop:
        yield x
        x += increment

'''
为了使用这个函数，你可以用 for 循环迭代它或者使用其他接受一个可迭代对象的
函数 (比如 sum() , list() 等)。
'''
for n in frange(0, 4, 0.5):
    print(n)

'''
0
0.5
1.0
1.5
2.0
2.5
3.0
3.5
'''
print(list(frange(0, 1, 0.125)))
# [0, 0.125, 0.25, 0.375, 0.5, 0.625, 0.75, 0.875]

'''
一个函数中需要有一个 yield 语句即可将其转换为一个生成器。跟普通函数不同
的是，生成器只能用于迭代操作。
'''

def countdown(n):
    print('Starting to count from', n)
    while n > 0:
        yield n
        n -=1
    print('Done!')

c = countdown(3)
print(c)
# <generator object countdown at 0x0000029CC12D4CA8>

print(next(c))
'''
Starting to count from 3
3
'''

print(next(c))
# 2

print(next(c))
# 1
#print(next(c))
'''
Done!
Traceback (most recent call last):
File "<stdin>", line 1, in <module>
StopIteration
'''

'''
一个生成器函数主要特征是它只会回应在迭代中使用到的 next 操作。一旦生成器
函数返回退出，迭代终止。我们在迭代中通常使用的 for 语句会自动处理这些细节，所
以你无需担心。
'''

def frange(start, stop, step=None):
    try:
        if step is None:
            step = 1
        while start < stop:
            yield start
            start += step
    except StopIteration:
        pass

print([d for d in frange(0, 5)])
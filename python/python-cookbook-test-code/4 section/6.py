#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第四章：迭代器与生成器
Desc: 迭代是 Python 最强大的功能之一。初看起来，你可能会简单的认为迭代只不过是
处理序列中元素的一种方法。然而，绝非仅仅就是如此，还有很多你可能不知道的，
比如创建你自己的迭代器对象，在 itertools 模块中使用有用的迭代模式，构造生成器
函数等等。这一章目的就是向你展示跟迭代有关的各种常见问题。

Title;      带有外部状态的生成器函数
Issue:你想定义一个生成器函数，但是它会调用某个你想暴露给用户使用的外部状态值。
Answer: 如果你想让你的生成器暴露外部状态给用户，别忘了你可以简单的将它实现为一个
类，然后把生成器函数放到 iter () 方法中过去。
"""

from collections import deque

class linehistory:
    def __init__(self, lines, histlen=3):
        self.lines = lines
        self.history = deque(maxlen=histlen)

    def __iter__(self):
        for lineno, line in enumerate(self.lines, 1):
            self.history.append((lineno, line))
            yield line

    def clear(self):
        self.history.clear()

'''
为了使用这个类，你可以将它当做是一个普通的生成器函数。然而，由于可以创建
一个实例对象，于是你可以访问内部属性值，比如 history 属性或者是 clear() 方法。
代码示例如下：
'''
with open('../test file/somefile.txt') as f:
    lines = linehistory(f)
    for line in lines:
        if 'python' in line:
            for lineno, hline in lines.history:
                print('{} : {}'.format(lineno, hline), end='')

'''
1 : this is python test file 1 somefile.
1 : this is python test file 1 somefile.
2 : this is python test file 2 somefile.
1 : this is python test file 1 somefile.
2 : this is python test file 2 somefile.
3 : this is python test file 3 somefile.
2 : this is python test file 2 somefile.
3 : this is python test file 3 somefile.
4 : this is python test file 4 somefile.
3 : this is python test file 3 somefile.
4 : this is python test file 4 somefile.
5 : this is python test file 5 somefile.
4 : this is python test file 4 somefile.
5 : this is python test file 5 somefile.
6 : this is python test file 6 somefile.
5 : this is python test file 5 somefile.
6 : this is python test file 6 somefile.
7 : this is python test file 7 somefile.
6 : this is python test file 6 somefile.
7 : this is python test file 7 somefile.
8 : this is python test file 8 somefile.
7 : this is python test file 7 somefile.
8 : this is python test file 8 somefile.
9 : this is python test file 9 somefile.
8 : this is python test file 8 somefile.
9 : this is python test file 9 somefile.
10 : this is python test file 10 somefile.
9 : this is python test file 9 somefile.
10 : this is python test file 10 somefile.
11 : this is python test file 11 somefile.
10 : this is python test file 10 somefile.
11 : this is python test file 11 somefile.
12 : this is python test file 12 somefile.
11 : this is python test file 11 somefile.
12 : this is python test file 12 somefile.
13 : this is python test file 13 somefile
'''

'''
关于生成器，很容易掉进函数无所不能的陷阱。如果生成器函数需要跟你的程序
其他部分打交道的话 (比如暴露属性值，允许通过方法调用来控制等等)，可能会导致
你的代码异常的复杂。如果是这种情况的话，可以考虑使用上面介绍的定义类的方式。
在 iter () 方法中定义你的生成器不会改变你任何的算法逻辑。由于它是类的一部
分，所以允许你定义各种属性和方法来供用户使用。
一个需要注意的小地方是，如果你在迭代操作时不使用 for 循环语句，那么你得先
调用 iter() 函数。比如：
'''

f = open('../test file/somefile.txt')
lines = linehistory(f)
#print(next(lines))

'''
Traceback (most recent call last):
  File "python cookbook test code/4 section/6.py", line 95, in <module>
    next(lines)
TypeError: 'linehistory' object is not an iterator
'''
line = iter(lines)
print(next(line))
# 13 : this is python test file 13 somefile. this is python test file 1 somefile.
#
print(next(line))
# this is python test file 2 somefile.
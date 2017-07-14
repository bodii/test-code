#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第八章：类与对象
Desc: 本章主要关注点的是和类定义有关的常见编程模型。包括让对象支持常见的
Python 特性、特殊方法的使用、类封装技术、继承、内存管理以及有用的设计模
式。

Title;       改变对象的字符串显示
Issue:你想改变对象实例的打印或显示输出，让它们更具可读性。
Answer: 要改变一个实例的字符串表示，可重新定义它的 str () 和 repr () 方法。
"""
class Pair:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def __repr__(self):
        return 'Pair({0.x!r}, {0.y!r})'.format(self)

    def __str__(self):
        return '({0.x!r}, {0.y!r})'.format(self)

"""
__repr__ () 方法返回一个实例的代码表示形式，通常用来重新构造这个实例。内
置的 repr() 函数返回这个字符串，跟我们使用交互式解释器显示的值是一样的。
__str__ () 方法将实例转换为一个字符串，使用 str() 或 print() 函数会输出这个字符
串。
"""

p = Pair(3, 4)
print(p)
# (3, 4)

p = Pair(3, 4)
print('p is {0!r}'.format(p))
# p is Pair(3, 4)
print('p is {0}'.format(p))
# p is (3, 4)

"""
自定义 repr () 和 str () 通常是很好的习惯，因为它能简化调试和实例输出。
例如，如果仅仅只是打印输出或日志输出某个实例，那么程序员会看到实例更加详细
与有用的信息。
repr () 生成的文本字符串标准做法是需要让 eval(repr(x)) == x 为真。如果
实在不能这样子做，应该创建一个有用的文本表示，并使用 < 和 > 括起来。
"""
f = open('../test file/file.dat')
"""
如果__str__ () 没有被定义，那么就会使用 __repr__ () 来代替输出。
上面的 format() 方法的使用看上去很有趣，格式化代码 {0.x} 对应的是第 1 个参
数的 x 属性。因此，在下面的函数中，0 实际上指的就是 self 本身：
"""
"""
def __repr__(self):
    return 'Pair({0.x!r}, {0.y!r}.format(self))'

作为这种实现的一个替代，你也可以使用 % 操作符
def __repr__(self):
    return 'Pair(%r, %r) % (self.x, self.y))'
"""
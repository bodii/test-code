#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第七章：函数
Desc: 使用 def 语句定义函数是所有程序的基础。本章的目标是讲解一些更加高级和不
常见的函数定义与使用模式。涉及到的内容包括默认参数、任意数量参数、强制关键
字参数、注解和闭包。另外，一些高级的控制流和利用回调函数传递数据的技术在这
里也会讲解到。

Title;      访问闭包中定义的变量
Issue:你想要扩展函数中的某个闭包，允许它能访问和修改函数的内部变量。
Answer: 通常来讲，闭包的内部变量对于外界来讲是完全隐藏的。但是，你可以通过编写访
问函数并将其作为函数属性绑定到闭包上来实现这个目的。
"""

def sample():
    n = 0
    def func():
        print('n=', n)

    def get_n():
        return n

    def set_n(value):
        nonlocal n
        n = value

    func.get_n = get_n
    func.set_n = set_n
    return func

f = sample()
f()
# n= 0
f.set_n(10)
f()
# n= 10
print(f.get_n())
# 10

"""
为了说明清楚它如何工作的，有两点需要解释一下。首先，nonlocal 声明可以让
我们编写函数来修改内部变量的值。其次，函数属性允许我们用一种很简单的方式将
访问方法绑定到闭包函数上，这个跟实例方法很像 (尽管并没有定义任何类)。
还可以进一步的扩展，让闭包模拟类的实例。你要做的仅仅是复制上面的内部函数
到一个字典实例中并返回它即可。
"""
import sys
class ClosureInstance:
    def __init__(self, locals=None):
        if locals is None:
            locals = sys._getframe(1).f_locals

        self.__dict__.update(
            (key, value) for key, value in locals.items()
            if callable(value)
        )

    def __len__(self):
        return self.__dict__['__len__']()


def Stack():
    items = []
    def push(item):
        items.append(item)

    def pop():
        return items.pop()

    def __len__():
        return len(items)

    return ClosureInstance()

"""
下面是一个交互式会话来演示它是如何工作的
"""
s = Stack()
print(s)
# <__main__.ClosureInstance object at 0x000002D6FAB534A8>
print(s.push(10))
# None
print(s.push(20))
# None
print(s.push('hello'))
print(len(s))
# 3
print(s.pop())
# hello
print(s.pop())
# 20
print(s.pop())
# 10



class Stack2:
    def __init__(self):
        self.items = []

    def push(self, item):
        self.items.append(item)

    def pop(self):
        return self.itmes.pop()

    def __len__(self):
        return len(self.itmes)

"""
如果这样做，你会得到类似如下的结果
"""
from timeit import timeit
s = Stack()
print(timeit('s.push(1);s.pop()', 'from __main__ import s'))

s = Stack2()
print(timeit('s.push(1);s.pop()', 'from __main__ import s'))

"""
结果显示，闭包的方案运行起来要快大概 8%，大部分原因是因为对实例变量的简
化访问，闭包更快是因为不会涉及到额外的 self 变量。
Raymond Hettinger 对于这个问题设计出了更加难以理解的改进方案。不过，你得
考虑下是否真的需要在你代码中这样做，而且它只是真实类的一个奇怪的替换而已，
例如，类的主要特性如继承、属性、描述器或类方法都是不能用的。并且你要做一些其
他的工作才能让一些特殊方法生效 (比如上面 ClosureInstance 中重写过的 len ()
实现。)
最后，你可能还会让其他阅读你代码的人感到疑惑，为什么它看起来不像一个普通
的类定义呢？(当然，他们也想知道为什么它运行起来会更快)。尽管如此，这对于怎样
访问闭包的内部变量也不失为一个有趣的例子。
总体上讲，在配置的时候给闭包添加方法会有更多的实用功能，比如你需要重置内
部状态、刷新缓冲区、清除缓存或其他的反馈机制的时候。
"""
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

Title;              以编程方式定义类
Issue:你在写一段代码，最终需要创建一个新的类对象。你考虑将类的定义源代码以字符
串的形式发布出去。并且使用函数比如 exec() 来执行它，但是你想寻找一个更加优雅
的解决方案。
Answer:你可以使用函数 types.new class() 来初始化新的类对象。你需要做的只是提供
类的名字、父类元组、关键字参数，以及一个用成员变量填充类字典的回调函数。
"""
def __init__(self, name, shares, price):
    self.name = name
    self.shares = shares
    self.price = price

def cost(self):
    return self.shares * self.price

cls_dict = {
    '__init__' : __init__,
    'cost' : cost,
}

import types

Stock = types.new_class('Stock', (), {}, lambda ns:ns.update(cls_dict))
Stock.__module__ = __name__

"""
这种方式会构建一个普通的类对象，并且按照你的期望工作：
"""
s = Stock('ACME', 50, 91.1)
print(s)
# <__main__.Stock object at 0x0000022497162E48>
print(s.cost())
# 4555.0

"""
这 种 方 法 中， 一 个 比 较 难 理 解 的 地 方 是 在 调 用 完 types.new_class() 对
Stock. module 的赋值。每次当一个类被定义后，它的 module 属性包含定义
它的模块名。这个名字用于生成_repr_() 方法的输出。它同样也被用于很多库，比
如 pickle 。因此，为了让你创建的类是“正确”的，你需要确保这个属性也设置正确
了。
如果你想创建的类需要一个不同的元类，可以通过 types.new_class() 第三个参
数传递给它。
"""
import abc
Stock = types.new_class('Stock', (), {'metaclass':abc.ABCMeta},lambda ns:ns.update(cls_dict))

Stock.__module__ = __name__
print(Stock)
#<class '__main__.Stock'>
print(type(Stock))
# <class 'abc.ABCMeta'>
"""
第三个参数还可以包含其他的关键字参数。比如，一个类的定义如下：
"""
# class Spam(Base, debug=True, typecheck=False):
#     pass

"""
那么可以将其翻译成如下的 new_class() 调用形式
"""
#Spam = types.new_class('Spam', (Base,), {'debug':True, 'typecheck':False},lambda ns:ns.update(cls_dict))

"""
new class() 第四个参数最神秘，它是一个用来接受类命名空间的映射对象的函
数。通常这是一个普通的字典，但是它实际上是 prepare () 方法返回的任意对象，
这个在 9.14 小节已经介绍过了。这个函数需要使用上面演示的 update() 方法给命名
空间增加内容。
"""

"""
很 多 时 候 如 果 能 构 造 新 的 类 对 象 是 很 有 用 的。 有 个 很 熟 悉 的 例 子 是 调 用
collections.namedtuple() 函数，
"""
import collections
Stock = collections.namedtuple('Stock', ['name', 'Shares', 'price'])
print(Stock)
# <class '__main__.Stock'>

"""
namedtuple() 使用 exec() 而不是上面介绍的技术。但是，下面通过一个简单的变
化，我们直接创建一个类：
"""
import operator
import types
import sys

def named_tuple(classname, fieldnames):
    cls_dict = {
        name: property(operator.itemgetter(n))
        for n, name in enumerate(fieldnames)
    }

    def __new__(cls, *args):
        if len(args) != len(fieldnames):
            raise TypeError('Expected {} arguments'. format(len(fieldnames)))
        return tuple.__new__(cls, args)

    cls_dict['__new__'] = __new__

    cls = types.new_class(classname, (tuple,), {}, lambda ns:ns.update(cls_dict))

    cls.__module__ = sys._getframe(1).f_globals['__name__']
    return cls

"""
这段代码的最后部分使用了一个所谓的” 框架魔法”，通过调用 sys. getframe()
来获取调用者的模块名。另外一个框架魔法例子在 2.15 小节中有介绍过。
下面的例子演示了前面的代码是如何工作的：
"""
Point = named_tuple('Point', ['x', 'y'])
print(Point)
# <class '__main__.Point'>
p = Point(4, 5)
print(len(p))
# 2
print(p.x)
# 4
print(p.y)
# 5
# p.x = 2
# AttributeError: can't set attribute

print('%s %s' % p)
# 4 5

"""
这项技术一个很重要的方面是它对于元类的正确使用。你可能像通过直接实例化一
个元类来直接创建一个类：

Stock = type('Stock', (), cls_dict)
这种方法的问题在于它忽略了一些关键步骤，比如对于元类中 prepare () 方法
的调用。通过使用 types.new class() ，你可以保证所有的必要初始化步骤都能得到
执行。比如，types.new class() 第四个参数的回调函数接受 prepare () 方法返回
的映射对象。

如果你仅仅只是想执行准备步骤，可以使用 types.prepare class() 。例如：
"""
import types
metaclass, kwargs, ns = types.prepare_class('Stock', (), {'metaclass':type})

"""
它会查找合适的元类并调用它的 prepare () 方法。然后这个元类保存它的关键
字参数，准备命名空间后被返回。
"""


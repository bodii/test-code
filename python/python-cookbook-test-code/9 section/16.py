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

Title;            *args 和 **kwargs 的强制参数签名
Issue:你有一个函数或方法，它使用 *args 和 **kwargs 作为参数，这样使得它比较通用，
但有时候你想检查传递进来的参数是不是某个你想要的类型。
Answer:对任何涉及到操作函数调用签名的问题，你都应该使用 inspect 模块中的签名特
性。我们最主要关注两个类：Signature 和 Parameter 。
"""
from inspect import Signature, Parameter

parms = [
    Parameter('x', Parameter.POSITIONAL_OR_KEYWORD),
    Parameter('y', Parameter.POSITIONAL_OR_KEYWORD, default=42),
    Parameter('z', Parameter.KEYWORD_ONLY, default=None)
]

sig = Signature(parms)
print(sig)
# (x, y=42, *, z=None)

"""
一旦你有了一个签名对象，你就可以使用它的 bind() 方法很容易的将它绑定到
*args 和 **kwargs 上去。下面是一个简单的演示：
"""
def func(*args, **kwargs):
    bound_values = sig.bind(*args, **kwargs)
    for name, value in bound_values.arguments.items():
        print(name, value)

func(1, 2, z=3)
'''
x 1
y 2
z 3
'''

func(1)
# x 1

func(1, z=3)
'''
x 1
z 3
'''
func(y=2, x=1)
'''
x 1
y 2
'''
#func(1, 2, 3, 4)
# TypeError: too many positional arguments
#func(y=2)
# TypeError: missing a required argument: 'x'
#func(1, y=2, x=3)
# TypeError: multiple values for argument 'x'

"""
可以看出来，通过将签名和传递的参数绑定起来，可以强制函数调用遵循特定的规
则，比如必填、默认、重复等等。
下面是一个强制函数签名更具体的例子。在代码中，我们在基类中先定义了一个非
常通用的 init () 方法，然后我们强制所有的子类必须提供一个特定的参数签名。
"""

def make_sig(*names):
    parms = [
        Parameter(name, Parameter.POSITIONAL_OR_KEYWORD)
        for name in names
    ]
    return Signature(parms)

class Structure:
    __signature__ = make_sig()

    def __init__(self, *args, **kwargs):
        bound_values = self.__signature__.bind(*args, **kwargs)
        for name, value in bound_values.arguments.items():
            setattr(self, name, value)

class Stock(Structure):
    __signature__ = make_sig('name', 'shares', 'price')

class Point(Structure):
    __signature__ = make_sig('x', 'y')

import inspect
print(inspect.signature(Stock))
# (name, shares, price)
s1 = Stock('ACME', 100, 490.1)
#s2 = Stock('ACME', 100)
# TypeError: missing a required argument: 'price'

# s3 = Stock('ACME', 100, 490.1, shares=50)
# TypeError: multiple values for argument 'shares'

"""
在我们需要构建通用函数库、编写装饰器或实现代理的时候，对于 *args 和
**kwargs 的使用是很普遍的。但是，这样的函数有一个缺点就是当你想要实现自
己的参数检验时，代码就会笨拙混乱。在 8.11 小节里面有这样一个例子。这时候我们
可以通过一个签名对象来简化它。
在最后的一个方案实例中，我们还可以通过使用自定义元类来创建签名对象。下面
演示怎样来实现：
"""
from inspect import Signature, Parameter

def make_sig(*names):
    params = [
        Parameter(name, Parameter.POSITIONAL_OR_KEYWORD)
        for name in names
    ]
    return Signature(params)

class StructureMeta(type):
    def __new__(cls, clsname, bases, clsdict):
        clsdict['__signature__'] = make_sig(*clsdict.get('_fields', []))
        return super().__new__(cls, clsname, bases, clsdict)

class Structure(metaclass=StructureMeta):
    _fields = []
    def __init__(self, *args, **kwargs):
        bound_values = self.__signature__.bind(*args, **kwargs)
        for name, value in bound_values.arguments.items():
            setattr(self, name, value)

class Stock(Structure):
    _fields = ['name', 'shares', 'price']

class Point(Structure):
    _fields = ['x', 'y']

"""
当我们自定义签名的时候，将签名存储在特定的属性 signature 中通常是很有
用的。这样的话，在使用 inspect 模块执行内省的代码就能发现签名并将它作为调用
约定。
"""
import inspect
print(inspect.signature(Stock))
# (name, shares, price)
print(inspect.signature(Point))
# (x, y)
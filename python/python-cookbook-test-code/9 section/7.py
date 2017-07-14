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

Title;             利用装饰器强制函数上的类型检查
Issue:作为某种编程规约，你想在对函数参数进行强制类型检查。
Answer:inspect.signature() 函数
"""
"""
演示实际代码前，先说明我们的目标：能对函数参数类型进行断言，类似下面这
样:
@typeassert(int, int)
def add(x, y):
    return x + y
add(2, 3)
# 5
add(2, 'hello')
# TypeError: Argument y must be <class 'int'>
"""

"""下面是使用装饰器技术来实现 @typeassert ："""

from inspect import signature
from functools import wraps

def typeassert(*ty_args, **ty_kwargs):
    def decorate(func):
        if not __debug__:
            return func

        sig = signature(func)
        bound_types = sig.bind_partial(*ty_args, **ty_kwargs).arguments

        @wraps(func)
        def wrapper(*args, **kwargs):
            bound_values = sig.bind(*args, **kwargs)
            for name, value in bound_values.arguments.items():
                if name in bound_types:
                    if not isinstance(value, bound_types[name]):
                        raise TypeError(
                            'Argument {} must be {}'.format(name,bound_types[name])
                        )
            return func(*args, **kwargs)
        return wrapper
    return decorate

"""
可以看出这个装饰器非常灵活，既可以指定所有参数类型，也可以只指定部分。并
且可以通过位置或关键字来指定参数类型。下面是使用示例：
"""
@typeassert(int,int, z=int)
def spam(x, y, z=42):
    print(x, y, z)

spam(1, 2, 3)
# 1 2 3

#spam(1, 'hello', 3)
# TypeError: Argument y must be <class 'int'>
"""
这节是高级装饰器示例，引入了很多重要的概念。
首先，装饰器只会在函数定义时被调用一次。有时候你去掉装饰器的功能，那么你
只需要简单的返回被装饰函数即可。下面的代码中，如果全局变量　 debug 被设置
成了 False(当你使用 -O 或 -OO 参数的优化模式执行程序时)，那么就直接返回未修改
过的函数本身：
def decorate(func):
    if not __debug__:
        return func

其 次， 这 里 还 对 被 包 装 函 数 的 参 数 签 名 进 行 了 检 查， 我 们 使 用 了
inspect.signature() 函数。简单来讲，它运行你提取一个可调用对象的参数签名
信息。
"""
from inspect import signature
def spam(x, y, z=42):
    pass

sig = signature(spam)
print(sig)
# (x, y, z=42)
print(sig.parameters)
# OrderedDict([('x', <Parameter "x">), ('y', <Parameter "y">), ('z', <Parameter "z=42">)])
print(sig.parameters['z'].name)
# z
print(sig.parameters['z'].default)
# 42
print(sig.parameters['z'].kind)
# POSITIONAL_OR_KEYWORD

"""
装饰器的开始部分，我们使用了 bind partial() 方法来执行从指定类型到名称的
部分绑定。下面是例子演示：
"""
bound_types = sig.bind_partial(int, z=int)
print(bound_types)
# <BoundArguments (x=<class 'int'>, z=<class 'int'>)>
"""
在这个部分绑定中，你可以注意到缺失的参数被忽略了 (比如并没有对 y 进行绑
定)。不过最重要的是创建了一个有序字典 bound types.arguments 。这个字典会将参
数名以函数签名中相同顺序映射到指定的类型值上面去。在我们的装饰器例子中，这
个映射包含了我们要强制指定的类型断言。
在 装 饰 器 创 建 的 实 际 包 装 函 数 中 使 用 到 了 sig.bind() 方 法。 bind() 跟
bind partial() 类似，但是它不允许忽略任何参数。因此有了下面的结果：

bound_values = sig.bind(1, 2, 3)
print(bound_values.arguments)
# OrderedDict([('x', 1), ('y', 2), ('z', 3)])
for name, value in bound_values.arguments.items():
    if name in bound_types.arguments:
        if not isinstance(value, bound_types.arguments[name]):
            raise TypeError()
"""
"""
不过这个方案还有点小瑕疵，它对于有默认值的参数并不适用。比如下面的代码可
以正常工作，尽管 items 的类型是错误的
"""
@typeassert(int, list)
def bar(x, items=None):
    if items is None:
        items = []
    items.append(x)
    return items

print(bar(2))
# [2]

#print(bar(2, 3))
# TypeError: Argument items must be <class 'list'>
print(bar(4, [1, 2, 3]))
# [1, 2, 3, 4]

"""
最后一点是关于适用装饰器参数和函数注解之间的争论。例如，为什么不像下面这
样写一个装饰器来查找函数中的注解呢？
@typeassert
def spam(x:int, y, z:int = 42):
    print(x,y,z)
一个可能的原因是如果使用了函数参数注解，那么就被限制了。如果注解被用来做
类型检查就不能做其他事情了。而且 @typeassert 不能再用于使用注解做其他事情的
函数了。而使用上面的装饰器参数灵活性大多了，也更加通用。
可以在 PEP 362 以及 inspect 模块中找到更多关于函数参数对象的信息。在 9.16
小节还有另外一个例子。
"""
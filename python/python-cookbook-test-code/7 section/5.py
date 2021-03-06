#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第七章：函数
Desc: 使用 def 语句定义函数是所有程序的基础。本章的目标是讲解一些更加高级和不
常见的函数定义与使用模式。涉及到的内容包括默认参数、任意数量参数、强制关键
字参数、注解和闭包。另外，一些高级的控制流和利用回调函数传递数据的技术在这
里也会讲解到。

Title;      定义有默认参数的函数
Issue:你想定义一个函数或者方法，它的一个或多个参数是可选的并且有一个默认值。
Answer: 定义一个有可选参数的函数是非常简单的，直接在函数定义中给参数指定一个默认
值，并放到参数列表最后就行了。
"""

def spam(a, b=42):
    print(a, b)

spam(1)
# 1 42
spam(1, 2)
# 1 2

"""
如果默认参数是一个可修改的容器比如一个列表、集合或者字典，可以使用 None
作为默认值，
"""
def spam(a, b=None):
    if b is None:
        b = []
"""
如果你并不想提供一个默认值，而是想仅仅测试下某个默认参数是不是有传递进
来
"""
_no_value = object()
def spam(a, b=_no_value):
    if b is _no_value:
        print('No b value supplied')

spam(1)
# No b value supplied
spam(1, 2)
spam(1, None)

"""
仔细观察可以发现到传递一个 None 值和不传值两种情况是有差别的。
"""

x = 42
def spam(a, b=x):
    print(a, b)

spam(1)
# 1 42
x = 23
spam(1)
#  1 42

"""
注意到当我们改变 x 的值的时候对默认参数值并没有影响，这是因为在函数定义的
时候就已经确定了它的默认值了。

其次，默认参数的值应该是不可变的对象，比如 None、True、False、数字或字符
串。特别的，千万不要像下面这样写代码：
"""
def spam(a, b=[]):
    pass

"""
如果你这么做了，当默认值在其他地方被修改后你将会遇到各种麻烦。这些修改会
影响到下次调用这个函数时的默认值。
"""
def spam(a, b=[]):
    print(b)
    return b

x = spam(1)
# []
print(x)
# []

x.append(99)
x.append('Yow!')
print(x)
# [99, 'Yow!']
spam(1)
# [99, 'Yow!']

"""
这种结果应该不是你想要的。为了避免这种情况的发生，最好是将默认值设为
None，然后在函数里面检查它，前面的例子就是这样做的。
在测试 None 值时使用 is 操作符是很重要的，也是这种方案的关键点。有时候大
家会犯下下面这样的错误
"""
def spam(a, b=None):
    if not b:
        b = []

"""
这么写的问题在于尽管 None 值确实是被当成 False，但是还有其他的对象 (比如长
度为 0 的字符串、列表、元组、字典等) 都会被当做 False。因此，上面的代码会误将
一些其他输入也当成是没有输入。
"""
spam(1)
x = []
print(spam(1, x))
# None
print(spam(1, 0))
# None
print(spam(1, ''))
# None

"""
最后一个问题比较微妙，那就是一个函数需要测试某个可选参数是否被使用者传递
进来。这时候需要小心的是你不能用某个默认值比如 None、0 或者 False 值来测试用
户提供的值 (因为这些值都是合法的值，是可能被用户传递进来的)。因此，你需要其
他的解决方案了。
为了 解 决 这 个 问 题， 你 可 以创建 一个 独一无 二的 私有对 象实例， 就像上 面
的 no value 变量那样。在函数里面，你可以通过检查被传递参数值跟这个实例是否
一样来判断。这里的思路是用户不可能去传递这个 no value 实例作为输入。因此，这
里通过检查这个值就能确定某个参数是否被传递进来了。
这里对 object() 的使用看上去有点不太常见。object 是 python 中所有类的基类。
你可以创建 object 类的实例，但是这些实例没什么实际用处，因为它并没有任何有用
的方法，也没有哦任何实例数据 (因为它没有任何的实例字典，你甚至都不能设置任何
属性值)。你唯一能做的就是测试同一性。这个刚好符合我的要求，因为我在函数中就
只是需要一个同一性的测试而已。
"""
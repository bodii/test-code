#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第七章：函数
Desc: 使用 def 语句定义函数是所有程序的基础。本章的目标是讲解一些更加高级和不
常见的函数定义与使用模式。涉及到的内容包括默认参数、任意数量参数、强制关键
字参数、注解和闭包。另外，一些高级的控制流和利用回调函数传递数据的技术在这
里也会讲解到。

Title;       可接受任意数量参数的函数
Issue: 你想构造一个可接受任意数量参数的函数。
Answer: 为了能让一个函数接受任意数量的位置参数，可以使用一个 * 参数。
"""

def avg(first, *rest):
    return (first + sum(rest)) / (1 + len(rest))

print(avg(1, 2))
# 1.5
print(avg(1, 2, 3, 4))
#  2.5
"""
在这个例子中，rest 是由所有其他位置参数组成的元组。然后我们在代码中把它当
成了一个序列来进行后续的计算。
"""
"""
为了接受任意数量的关键字参数，使用一个以 ** 开头的参数。
"""
import html

def make_element(name, value, **attrs):
    keyvals = [' %s="%s"' % item for item in attrs.items()]
    attr_str = "".join(keyvals)
    element = '<{name}{attrs}>{value}</{name}>'.format(
        name = name,
        attrs = attr_str,
        value = html.escape(value)
    )
    return element

make_element('item', 'Albatross', size='large', quantity=6)
make_element('p', '<spam>')

"""
在这里，attrs 是一个包含所有被传入进来的关键字参数的字典。
如果你还希望某个函数能同时接受任意数量的位置参数和关键字参数，可以同时使
用 * 和 **。
"""
def anyargs(*args, **kwargs):
    print(args)
    print(kwargs)
"""
使用这个函数时，所有位置参数会被放到 args 元组中，所有关键字参数会被放到
字典 kwargs 中。
"""
"""
一个 * 参数只能出现在函数定义中最后一个位置参数后面，而**参数只能出现在
最后一个参数。有一点要注意的是，在 * 参数后面仍然可以定义其他参数。
"""
def a(x, *args, y):
    pass
def b(x, *args, y, **kwargs):
    pass
"""
这种参数就是我们所说的强制关键字参数，
"""

#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第七章：函数
Desc: 使用 def 语句定义函数是所有程序的基础。本章的目标是讲解一些更加高级和不
常见的函数定义与使用模式。涉及到的内容包括默认参数、任意数量参数、强制关键
字参数、注解和闭包。另外，一些高级的控制流和利用回调函数传递数据的技术在这
里也会讲解到。

Title;      定义匿名或内联函数
Issue:你想为 sort() 操作创建一个很短的回调函数，但又不想用 def 去写一个单行函
数，而是希望通过某个快捷方式以内联方式来创建这个函数。
Answer: 当一些函数很简单，仅仅只是计算一个表达式的值的时候，就可以使用 lambda 表
达式来代替了。
"""

add = lambda x, y: x + y
print(add(2, 3))
# 5
print(add('hello', ' world'))
# hello world

"""
这里使用的 lambda 表达式跟下面的效果是一样的
"""
def add(x, y):
    return x + y
print(add(2, 3))
# 5

"""
lambda 表达式典型的使用场景是排序或数据 reduce
"""
names = [
    'David Beazley', 'Brian Jones',
    'Raymond Hettinger', 'Ned Batchelder'
]

print(sorted(names, key=lambda name:name.split()[-1].lower()))
# ['Ned Batchelder', 'David Beazley', 'Raymond Hettinger', 'Brian Jones']

"""
尽管 lambda 表达式允许你定义简单函数，但是它的使用是有限制的。你只能指定
单个表达式，它的值就是最后的返回值。也就是说不能包含其他的语言特性了，包括
多个语句、条件表达式、迭代以及异常处理等等。
你可以不使用 lambda 表达式就能编写大部分 python 代码。但是，当有人编写大
量计算表达式值的短小函数或者需要用户提供回调函数的程序的时候，你就会看到
lambda 表达式的身影了。
"""
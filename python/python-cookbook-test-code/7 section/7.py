#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第七章：函数
Desc: 使用 def 语句定义函数是所有程序的基础。本章的目标是讲解一些更加高级和不
常见的函数定义与使用模式。涉及到的内容包括默认参数、任意数量参数、强制关键
字参数、注解和闭包。另外，一些高级的控制流和利用回调函数传递数据的技术在这
里也会讲解到。

Title;      匿名函数捕获变量值
Issue:你用 lambda 定义了一个匿名函数，并想在定义时捕获到某些变量的值
Answer: lambda
"""
x = 10
a = lambda y: x + y
x = 20
b = lambda y: x + y

"""
现在我问你，a(10) 和 b(10) 返回的结果是什么？如果你认为结果是 20 和 30，那
么你就错了
"""
print(a(10))
# 30
print(b(10))
# 30

"""
这其中的奥妙在于 lambda 表达式中的 x 是一个自由变量，在运行时绑定值，而不
是定义时就绑定，这跟函数的默认值参数定义是不同的。因此，在调用这个 lambda 表
达式的时候，x 的值是执行时的值。
"""
x = 15
print(a(10))
# 25

x = 3
print(a(10))
# 13

"""
如果你想让某个匿名函数在定义时就捕获到值，可以将那个参数值定义成默认参数
即可，
"""
x = 10
a = lambda y, x= x: x + y
x = 20
b = lambda y, x=x: x + y
print(a(10))
# 20
print(b(10))
# 30

"""
在这里列出来的问题是新手很容易犯的错误，有些新手可能会不恰当的 lambda 表
达式。比如，通过在一个循环或列表推导中创建一个 lambda 表达式列表，并期望函数
能在定义时就记住每次的迭代值。
"""

funcs = [lambda x: x+n for n in range(5)]
for f in funcs:
    print(f(0))
'''
4
4
4
4
4
'''
"""
但是实际效果是运行是 n 的值为迭代的最后一个值。现在我们用另一种方式修改
一下
"""
funcs = [lambda x, n=n: x + n for n in range(5)]
for f in funcs:
    print(f(0))
'''
0
1
2
3
4
'''

"""
通过使用函数默认值参数形式，lambda 函数在定义时就能绑定到值。
"""

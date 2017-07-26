#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第三章：数字日期和时间
Desc: 在 Python 中执行整数和浮点数的数学运算时很简单的。尽管如此，如果你需要执
行分数、数组或者是日期和时间的运算的话，就得做更多的工作了。本章集中讨论的
就是这些主题。

Title;         分数运算
Issue: 你进入时间机器，突然发现你正在做小学家庭作业，并涉及到分数计算问题。或者
你可能需要写代码去计算在你的木工工厂中的测量值
Answer: fractions 模块可以被用来执行包含分数的数学运算
"""

from fractions import Fraction

a = Fraction(5, 4)
b = Fraction(7, 16)
print(a + b)
# 27/16
print(a * b)
# 35/64

c = a * b
print(c.numerator)
# 35
print(c.denominator)
# 64

print(float(c))
# 0.546875

print(c.limit_denominator(8))
# 4/7

x = 3.75
y = Fraction(*x.as_integer_ratio())
print(y)
# 15/4

'''
在大多数程序中一般不会出现分数的计算问题，但是有时候还是需要用到的。比
如，在一个允许接受分数形式的测试单位并以分数形式执行运算的程序中，直接使用
分数可以减少手动转换为小数或浮点数的工作。
'''
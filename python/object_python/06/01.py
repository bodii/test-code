#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 创建数值类型 '''

"""
通过扩展numbers模块的基本抽象类来创建新的数值类型。

numbers包提供了大量的数值类型，它们都实现了numbers.Number。另外，fractions
和decimal模块提供了可扩展的数值类型：fractions.Fraction和decimal.Decimal。

Python定义了以下的抽象类以及它们关的实现：
complex实现了numbers.Complex。
float实现了numbers.Real。
fractions.Fractions 实现了numbers.Rational
int实现了numbers.Integral
此外，还有decimal.Decimal，虽然看起来像float，但它不是numbers.Real的子类。
"""

print( (3*5*7*11) / (11*13*17*23*29) )
# 0.0007123135264946712

"""
从原则上来说，在数塔中，越接近塔底的数，无穷大除数越小。不同的数字都定义了各自的
无穷大阶数，可以证明，各自的无穷大阶数大小是不同的。原则上，浮点数的表示比整数表
示包含了更多的数字。实际上，一个64位的浮点数和64位的整数包含了相同数量的不同的数
字。
数值类型的定义中，包含了一系列不同类型之间的转换。

complex: 它无法转换到任何其他类型。一个complex值可被分解为real和imag部分，
它们都是float。

float：它可以被显式转换到任何类型，包括decimal.Decimal。算术运算符无法隐式
地将float值转换为Decimal

Fractions.Fraction：它可以被转换到除了decimal.Decimal之外的其他任何类型。
转为decimal包括两部分操作：（1）转为float;(2)转为decimal.Decimal。所求得
的是近似值。

int: 可以被转换成其他任何类型。

decimal: 可以转换为其他任何类型，但算术运算符不会隐式完成转换过程。
"""

''' 决定使用哪种类型 '''
"""
1. 复杂类型： 一旦涉及复杂的数学操作，将会使用complex、float和cmath模块。
可能根本不会使用Fraction或者Decimal。可是，没有任何理由要给数值类型强加限
制，大多数数字都可以被转换为复杂类型。

2. 货币类型： 对于有关货币的操作，必须使用Decimal。

3. 位运算： 当涉及位和字节的计算，总会使用int类型。

4. 常规情况： 除上述之外的其他情况。对于大多数常规的数学运算来说，int、float
和Fraction都可以互转的。
"""

#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 数值类型 '''

"""
当我闪需要创建新的数值类型或者扩展现有的数值类型时，需要用到numbers模块。
这个模块包含了Python内置数值类型的抽象定义。
"""

import numbers

print( isinstance( 12, numbers.Number) )
# True
print( isinstance( 355/113, numbers.Number) )
# True

"""
整数和浮点数都是numbers.numbers抽象基类的子类。
这个基类的子类包括nubmers.complex \ numbers.Real \ numbers.Rational 和 nubmers.Integral。
这些定义和数学上对不同数字的分类是一致的。
但，decimal.Decimal类和这些类不太一样。
"""
import decimal

print( issubclass(decimal.Decimal, numbers.Number) )
# True

print( issubclass(decimal.Decimal, numbers.Integral) )
# False

print( issubclass(decimal.Decimal, numbers.Real) )
# False

print( issubclass(decimal.Decimal, numbers.Complex) )
# False

print( issubclass( decimal.Decimal, numbers.Rational) )
# False

"""
对于Decimal与所有内置的具体数值类型都没有关系，这一点并不用感到奇怪。对于numbers.Rational
的具体实现，请参见fractions模式。
"""


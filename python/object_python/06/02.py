#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 算术运算符的特殊方法 '''

# 如：
print( 0.14.__rsub__(7) )
# 6.86

"""
方法                           运算符     描述
object.__add__(self, other)     +        加法
object.__sub__(self, other)     -        减法
object.__rsub__(self, other)    -        右边减左边
object.__mul__(self, other)     *        乘法
object.__truediv__(self, other) /        除法
object.__floordiv__(slef,other) //       除法取整
object.__mod__(self, other)     %        取余
object.__divmod__(self,other)   divmod() 返回一tuple(取余，取余后的余数)
object.__pow__(self,other[, modulo])  pow()函数和**

# 一元运算符和函数
object.__neg__(self)            -   # 例如： 8.00.__neg__()  =  -8.00
object.__pos__(self)            +
object.__abs__(self)
object.__complex__(self)         complex()
object.__int__(self)             int()
object.__float__(self)
object.__round__(self[, n])      round()       
object.__trunc__(self[, n])      math.trunc()
object.__ceil__(self[, n])       math.ceill()
object.__floor__(self[, n])      math.floor()
"""

# 定义一个简单的追踪函数，来看一下其中的部分细节
def trace( frame, event, arg ):
    if frame.f_code.co_name.startswith('__'):
        print( frame.f_code.co_name, frame.f_code.co_filename, event )

import sys

sys.settrace(trace)
"""
一旦完成安装，系统中任何执行代码都会经过trace()函数。我们会过滤出所有特殊方法名
的事件。
"""

class noisyfloat( float ):
    def __add__( self, other ):
        print( self, '+', other )
        return super().__add__( other )

    def __radd__( self, other ):
        print( self, 'r+', other )
        return super().__add__( other )

"""
这个类只重载了运算符中的两个特殊方法。当执行noisyfloat值的加法时，会看到打印出
的运算符相关的统计信息。
"""

x = noisyfloat( 2 )
print( x + 3 )
# __add__ /home/wong/coding/python/object_python/06/02.py call
# 2.0 + 3
# 5.0

print( 2 + x )
# __radd__ /home/wong/coding/python/object_python/06/02.py call
# 2.0 r+ 2
# 4.0

print( x + 2.3 )
# __add__ /home/wong/coding/python/object_python/06/02.py call
# 2.0 + 2.3
# 4.3

print( 2.3 + x )
# __radd__ /home/wong/coding/python/object_python/06/02.py call
# 2.0 r+ 2.3
# 4.3
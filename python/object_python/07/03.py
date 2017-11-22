#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 一些类的设计原则 '''

"""
当定义类时，我们的属性和方法有下面3种来源：
1. class语句。
2. 类级的装饰器。
3. mixin类和最后一个基类。

"""

''' 使用内置的装饰器 '''
"""
内置装饰器：@property \ @classmethod 和 @staticmethod
@property装饰器将一个方法函数转换成描述器。我们用这种简单属性语法来定义方法函数。
当将@property装饰器用于方法上时，也会额外创建一对属性，它们可以用于创建setter和
deleter属性。
@classmethod 和@staticmethod装饰器将一个方法函数转换成一个类级函数。被装饰的
方法现在可以用类调用，而不是对象。对于静态方法，没有显式的类引用。
"""
import math

class Angle( float ):
    __slots__ = ( '_degrees', )

    @staticmethod
    def from_radians( value ):
        return Angle( 180*value/math.pi )

    def __new__( cls, value ):
        self = super().__new__(cls)
        self._degrees = value
        return self

    @property
    def radians( self ):
        return math.pi * self._degress / 180
    
    @property
    def degrees( self ):
        return self._degrees


b = Angle.from_radians(.227)
print(b.degrees)
# 13.006141949469686

"""
静态方法实际上是一个函数，因为它与self实例变量无关。它的优点是它的语法直接与类绑定，
用Angle.from_radians比调用一个名为angle_from_radians的函数更为直观。使用这些
装饰器可以确保实现的正确性和一致性。
"""

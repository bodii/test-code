#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用数据修饰符 '''

"""
数据修饰符使得设计变得更复杂了，因为它的接口是受限制的。它必须包含__get__()方法，并且只能包
含 __set__()方法和__delete__()方法。与接口相关的限制：可以包含以上方法的1～3个，不能包含
其他方法。引入额外的方法将意味着Python不能把这个类正确地识别为一个数据修饰符。
"""

class Unit:
    conversion = 1.0
    def __get__( self, instance, owner ):
        return instance.kph * self.conversion

    def __set__(  self, instance, value ):
        instance.kph = value / self.conversion
    """
    以上的类通过简单的乘除运算实现了标准单位和非标准单位的互换。
    """

class Knots( Unit ):
    conversion = 0.5399568

class MPH( Unit ):
    conversion = 0.62137119

class KPH( Unit ):
    def __get__( self, instance, owner ):
        return instance._kph

    def __set__( self, instance, value ):
        instance._kph = value 


class Measurement:
    kph = KPH()
    knots = Knots()
    mph = MPH()

    def __init__( self, kph=None, mph=None, knots=None ):
        if kph: self.kph = kph
        elif mph: self.mph = mph
        elif knots: self.knots = knots
        else:
            raise TypeError

    def __str__( self ):
        return 'rate: {0.kph} kph = {0.mph} mph = {0.knots} konts'.format(self)


m2 = Measurement( knots=5.9 )
print( str(m2) )
# rate: 10.92680006993152 kph = 6.789598762345432 mph = 5.9 konts

print( m2.kph )
# 10.92680006993152
print( m2.mph )
# 6.789598762345432

m3 = Measurement( kph=5.2 )
print( str(m3) )
# rate: 5.2 kph = 3.231130188 mph = 2.8077753600000004 konts

"""
在Python中，建议把所有属性公有，这意味着如下：

1. 此处应当有很好的文档说明。
2. 此处应能够正确地反映出对象的状态，它们不应当是暂时性的或者临时变量。
3. 在少数情形下，属性包含了容易产生歧义的值，可使用单下划线（_）来标记为”不在接口定义范围内“，
以此表明它并不是真正意义上的私有。

私有属性是令人厌烦的。封装，并不会因为某种编程语言缺乏一种复杂的私有机制而被破坏，它只会被糟糕
的设计所破坏。
"""
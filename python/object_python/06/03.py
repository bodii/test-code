#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 创建一个数字类 '''

import numbers
import math
import sys

class FixedPoint( numbers.Rational ):
    __slots__ = ( 'value', 'scale', 'default_format' )

    def __new__( cls, value, scale=100 ):
        self = super().__new__( cls )
        if isinstance( value, FixedPoint ):
            self.value = value.value
            self.scale = value.scale
        elif isinstance( value, int ):
            self.value = value
            self.scale = scale
        elif isinstance( value, float ):
            self.value = int(scale*value+.5)
            self.scale = scale
        else:
            raise TypeError
        digits = int( math.log10(scale) )
        self.default_format = '{{0:.{digits}f}}'.format(digits=digits)
        return self

    def __str__( self ):
        return self.__format__( self.default_format )

    def __repr__( self ):
        return '{__class__.__name__:s}({value:d}, scale={scale:d})'.format(
            __class__ = self.__class__,
            value=self.value,
            scale=self.scale
        )

    def __format__( self, specification ):
        if specification == '': specification = self.default_format
        return specification.format( self.value / self.scale )

    def numerator( self ):
        return self.value

    def denominator( self ):
        return self.scale

    """
    FixedPoint类继承自numbers.Rational。接下来会包含两个整数值scale和value，以及
    分数的一般定义。这将需要定义大量的特殊方法，因为初始化是针对不可变对象的，因此会重写
    __new__()方法而非__init__()方法。为阻止添加属性，限制了slots的数量。初始化包括：
    1. 如果赋值的对象类型为FixedPoint，将复制内部属性并创建新的FixedPoint对象，就是
    对原对象进行克隆。尽管它将有唯一的ID，但它将有相同的哈希值用来比较对象是否相等，使得
    克隆对象和原对象基本没区别。
    2. 如果赋值的对象类型为整数或有理数(int或float)，它们用于为value和scale属性赋值。
    3. 可以加入对decimal.Decimal和fractions.Fraction的处理逻辑以及字符串解析。
    """
    def __add__( self, other ):
        if not isinstance(other, FixedPoint):
            new_scale = self.scale
            new_value = self.value + other * self.scale
        else:
            new_scale = max(self.scale, other.scale)
            new_value = (self.value * (new_scale // self.scale)
            + other.value * (new_scale // other.scale))
        return FixedPoint( int(new_value), int(new_scale))

    def __sub__( self, other ):
        if not isinstance(other, FixedPoint):
            new_scale = self.scale
            new_value = self.value - other * self.new_scale
        else:
            new_scale = max(self.scale, other.scale)
            new_value = (self.value * (new_scale // self.scale)
            - other.value * (new_scale//other.scale))
        return FixedPoint( int(new_value), scale=new_scale)

    def __mul__( self, other ):
        if not isinstance(other, FixedPoint):
            new_scale = self.scale
            new_value = self.value
        else:
            new_scale = self.scale * other.scale
            new_value = self.value * other.value
        return FixedPoint( int(new_value), scale=new_scale)

    def __truediv__( self, other ):
        if not isinstance( other, FixedPoint ):
            new_value = int(self.value / other)
        else:
            new_value = int( self.value / (other.value/other.scale))
        return FixedPoint( new_value, scale = self.scale)

    def __floordiv__( self, other ):
        if not isinstance(other, FixedPoint):
            new_value = int(self.value // other)
        else:
            new_value = int(self.value // (other.value / other.scale))
        return FixedPoint(new_value, scale=self.scale)

    def __mod__( self, other ):
        if not isinstance(other, FixedPoint):
            new_value = (self.value/self.scale) % other
        else:
            new_value = self.value % (other.value / other.scale)
        return FixedPoint(new_value, scale=self.scale)

    def __pow__( slef, other ):
        if not isinstance(other, FixedPoint):
            new_value = (self.value/self.scale) ** other
        else:
            new_value = (self.value/self.scale) ** (other.value/other.scale)
        return FixedPoint(int(new_value) * self.scale, scale=self.scale)

    ''' 一元算术运算符 '''
    def __abs__( self ):
        return FixedPoint( abs(self.value), self.scale)

    def __float__( self ):
        return self.value / self.scale

    def __init__( self ):
        return int( self.value / self.scale)

    def __trunc__( self ):
        return FixedPoint(math.trunc(self.value/self.scale), self.scale)

    def __ceil__( self ):
        return FixedPoint( math.ceil(self.value/self.scale), self.scale)

    def __floor__( self ):
        return FixedPoint( math.floor(self.value/self.scale), self.scale)

    def __round__( slef, ndigits ):
        return FixedPoint( round(self.value/self.scale, ndigits=0), self.scale)

    def __neg__( self ):
        return FixedPoint( -self.value, self.scale )

    def __pos__( self ):
        return self

    ''' 反向运算 '''

    def __radd__( self, other ):
        if not isinstance(other, FixedPoint):
            new_scale = self.scale
            new_value = other * self.scale + self.value
        else:
            new_scale = max(self.scale, other.scale)
            new_value = (other.value * (new_scale // other.scale)
            + self.value *(new_scale // self.scale))
        return FixedPoint( int(new_value), scale=new_scale)

    def __rsub__( self, other ):
        if not isinstance(other, FixedPoint):
            new_scale = self.new_scale
            new_value = other*self.scale - self.value
        else:
            new_scale = max(self.scale, other.scale)
            new_value = (other.value * (new_scale//other.scale)
            - self.value*(new_scale//self.scale))
        return FixedPoint( int(new_value), scale=new_scale)

    def __rmul__( self, other ):
        if not isinstance(other, FixedPoint):
            new_scale = self.scale
            new_value = other*self.value
        else:
            new_scale = self.scale * other.scale
            new_value = other.value * self.value
        return FixedPoint( int(new_value), scale=new_scale)

    def __rtruediv__( self, other ):
        if not isinstance(other, FixedPoint):
            new_value = self.scale  * int(other/(self.value/self.scale))
        else:
            new_value = int((other.value/other/scale) / self.value)
        return FixedPoint( new_value, scale=self.scale)

    def __rfloordiv__( self, other ):
        if not isinstance(other, FixedPoint):
            new_value = self.scale*int(other// (self.value/self.scale))
        else:
            new_value = int((other.value/other.scale) // self.value)
        return FixedPoint(new_value, scale=self.scale)

    def __rmod__( self, other ):
        if not isinstance(other, FixedPoint):
            new_value = other % (self.value / self.scale)
        else:
            new_value = (other.value/other.scale) % (self.value/self.scale)
        return FixedPoint( new_value, scale=self.scale)

    def __rpow__( self, other ):
        if not isinstance( other, FixedPoint):
            new_value = other ** (self.value/ self.scale)
        else:
            new_value = (other.value/other.scale) ** (self.value/self.scale)
        return FixedPoint( int(new_value)*self.scale, scale=self.scale)

    ''' 比较运算符 '''

    def __eq__( self, other ):
        if isinstance( other, FixedPoint):
            if self.scale == other.scale:
                return self.value == other.value
            else:
                return self.value*other.scale // self.scale == other.value
        else:
            return abs(self.value/self.scale - float(other)) < .5/self.new_scale

    def __ne__( self, other ):
        return not (self == other)

    def __le__( self, other ):
        return self.value/self.scale <= float(other)

    def __lt__( self, other ):
        return self.value/self.scale < float(other)

    def __ge__( self, other ):
        return self.value/self.scale >= float(other)

    def __gt__( self, other ):
        return self.value/self.scale > float(other)

    ''' 计算一个数字的哈希值 '''
    def __hash__( self ):
        p = sys.hash_info.modulus
        m, n = self.value, self.scale

        while m % p == n % p == 0:
            m, n = m // p, n // p
        if n % p == 0:
            hash_ = sys.hash_info.inf
        else:
            hash_ = (abs(m) %p) * pow(n, p-2, p) % p
        if m < 0:
            hash_ = - hash_
        if hash_ == -1:
            hash_ = -2
        return hash_

    ''' 设计更有用的四舍五入方法 '''
    def round_to( self, new_scale ):
        f = new_scale / self.scale
        return FixedPoint( int(self.value*f+.5), scale=new_scale)

    ''' 实现其他的特殊方法 '''
    """
    方法                               运算符
    object.__lshift__(self, other)      <<
    object.__rshift__(self, other)      >>
    object.__and__(self, other)          &
    object.__xor__(self, other)          ^
    object.__or__(self, other)           |
    object.__rlshift__(self, other)      <<
    object.__rrshift__(self, other)      >>
    object.__rand__(self, other)         &
    object.__rxor__(self, other)         ^
    object.__ror__(self, other)          |
    object.__invert__(self)              <<
    """

    ''' 原地运算符 '''
    """
    方法                               运算符
    object.__iadd__(self, other)        +=
    object.__isub__(self, other)        -=
    object.__imul__(self, other)        *=
    object.__itruediv__(self, other)    /=
    object.__ifloordiv__(self,other)    //=
    object.__imod__(self, other)        %=
    object.__ipow__(self, other[, modulo]) **=
    object.__ilshift__(self,other)       <<=
    object.__irshift__(self, other)      >>=
    object.__iand__(self, other)         &=
    object.__ixor__(self, other)         ^=
    object.__ior__(self, other)          |=
    """

    """
    例如：
    def __iadd__(self, aCard):
        self._cards.append(aCard)
        return self

    这样一来，就可以使用如下代码完成发牌：
    player_hand += deck.pop()
    """

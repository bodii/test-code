#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用tuple子类创建不可变对象 '''

"""
我们也可以通过让Card特性成为tuple类的子类并重写__getattr__()函数来实现一个不可变对象。
这样一来，我们将把对__getattr__(name)的访问转换为对self[index]的访问。
self[index]被实现为__getitem__(index)
"""

class BlackJackCard2( tuple ):
    def __new__( cls, rank, suit, hard, soft ):
        return super().__new__( cls, (rank, suit, hard, soft ) )

    def __getattr__( self, name ):
        return self[
        {'rank': 0, 'suit': 1, 'hard': 2, 'soft': 3}[name]
    ]

    def __setattr__( self, name, value ):
        raise AttributeError

d = BlackJackCard2( 'A', '♠', 1, 11 )
print( d.rank ) # A
print( d.suit ) # ♠
# d.bad = 2 
# AttributeError


''' 主动计算的属性 '''


class RateTimeDistance( dict ):
    def __init__( self, *args, **kw ):
        super().__init__( *args, **kw)
        self._solve()

    def __getattr__( self, name ):
        return self.get( name, None )

    def __setattr__( self, name, value ):
        self[name] = value
        self._solve()

    def __dir__( self ):
        return list( self.keys() )

    def _solve( self ):
        if self.rate is not None and self.time is not None:
            self['distance'] = self.rate * self.time

        elif self.rate is not None and self.distance is not None:
            self['time'] = self.distance / self.rate
        
        elif self.time is not None and self.distance is not None:
            self['rate'] = self.distance / self.time

rtd = RateTimeDistance(rate=6.3, time=8.25, distance=None)
print( 'Rate={rate}, Time={time}, Distance={distance}'.format( **rtd ) )
# Rate=6.3, Time=8.25, Distance=51.975

"""
dict 类型使用__init__()方法完成字典值的填充，然后判断是否提供了足够的初始化数据。它使用
了__setattr__()函数来为字典添加新项，每当属性的赋值操发生时就会调用_solve()函数。
在__getattr__()函数中，使用None来标识属性值的缺失。对于未赋值的属性，可以使用None标记为
缺失的值，这样会强制对这个值进行查找。例如，属性值来自用户输入或网络传输的数据，只有一个变量
值为None而其他变量都有值。
"""

rtd.time = None
rtd.rate = 6.1
print( 'Rate={rate}, Time={time}, Distance={distance}'.format( **rtd ) )
# Rate=6.1, Time=8.25, Distance=50.324999999999996
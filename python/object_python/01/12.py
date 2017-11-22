#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 实现同一个类的对象比较 '''

class BlackJackCard:
    def __init__( self, rank, suit ):
        self.rank = rank
        self.suit = suit

    def __lt__( self, other ):
        if not isinstance( other, BlackJackCard ): return NotImplemented
        return self.rank < other.rank

    def __le__( self, other ):
        # 隐式的类型检查使用try:语句块。使用try：语句块的优点是可以避免重复的类名称。
        try:
            return self.rank <= other.rank
        except AttributeError:
            return NotImplemented

    def __gt__( self, other ):
        # 显式的类型检查 调用 isinstance
        if not isinstance( other, BlackJackCard ): return NotImplemented
        return self.rank > other.rank

    def __ge__( self, other ):
        if not isinstance( other, BlackJackCard ): return NotImplemented
        return self.rank >= other.rank

    def __eq__( self, other ):
        if not isinstance( other, BlackJackCard ): return NotImplemented
        return self.rank == other.rank and self.suit == other.suit

    def __ne__( self, other ):
        if not isinstance( other, BlackJackCard ): return NotImplemented
        return self.rank != other.rank and self.suit != other.suit

    def __str__( self ):
        return '{rank}{suit}'.format( **self.__dict__ )




two = BlackJackCard( 2, '♠' )
three = BlackJackCard( 3, '♠' )
two_c = BlackJackCard( 2, '♣' )

print( two == two_c ) # False

print( two.rank == two_c.rank ) # True

print( two < three ) # True
 
print( two_c < three ) # True



''' 类实例与一个int值进行比较 '''
# print( two < 2 ) # TypeError

print( two > 2 ) # TypeError: unorderable types: BlackJackCard() > int()
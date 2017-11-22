#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 比较运算符方法 '''
"""
x < y   调用   x.__lt__(y)
x <= y  调用   x.__le__(y)
x == y  调用   x.__eq__(y)
x != y  调用   x.__ne__(y)
x > y   调用   x.__gt__(y)
x >= y  调用   x.__ge__(y)
"""

class BlackJackCard_p:
    def __init__( self, rank, suit ):
        self.rank = rank
        self.suit = suit
    
    def __lt__( self, other ):
        print( 'Compare {0} < {1}'.format(self, other))
        return self.rank < other.rank

    def __eq__( self, other ):
        print( 'Compare {0} == {1}'.format(self, other))
        return self.rank == other.rank

    def __str__( self ):
        return '{rank}{suit}'.format( **self.__dict__ )


two = BlackJackCard_p( 2, '♠' )
three = BlackJackCard_p( 3, '♠' )

print( two < three )  # Compare 2♠ < 3♠ True

print( two > three )  # Compare 3♠ < 2♠ False

print( two == three )  # Compare 2♠ == 3♣ False

# print( two <= three )


two_c = BlackJackCard_p( 2, '♣' )

print( two == two_c ) # Compare 2♠ == 2♣ True
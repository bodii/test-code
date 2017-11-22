#!/usr/bin/env python3
# -*- coding:utf-8 -*-


class Rectangle:
    def area( self ):
        return self.length * self.width


r = Rectangle()
r.length, r.width = 13, 8
print(r.area())

''' 在基类中实现__init__()方法 '''
"""
通过实现__init__（）方法来初始化一对象。每当创建一个对象，Python会先创建一个
空对象，然后调用该对象的__init__()函数。这个方法提供了对象内部变量以及其他一
些一次性过程的初始化操作。

'♠ ♣ ♥ ♦♤ ♧ ♡ ♢'

"""

class Card:
    def __init__( self, rank, suit ):
        self.suit = suit
        self.rank = rank
        self.hanrd, self.soft = self._points()

class NumberCard( Card ):
    def _points( self ):
        return int(self.rank), int(self.rank)

class AceCard( Card ):
    def _points( self ):
        return 1, 11

class FaceCard( Card ):
    def _points( self ):
        return 10, 10


cards = [ AceCard('A', '♠'), AceCard('A', '♣'), AceCard('A', '♥'), AceCard('A', '♦') ];

class Suit:
    def __init__( self, name, symbol ):
        self.name = name
        self.symbol = symbol

Club, Diamond, Heart, Spade = Suit('Club', '♠'), Suit('Diamond', '♦'),
Suit('Heart', '♥'), Suit('Spade', '♣')

cards = [ AceCard('A',Spade), NumberCard('2', Spade), NumberCard('3', Spade), ]
#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 全局模块和模块项 '''

__all__ = ['Deck', 'Shoe']

class Suit:
    def __init__( self, suit ):
        pass

suits = [ Suit('♣'), Suit('♦'), Suit('♥'), Suit('♠') ]

class Card:
    pass

def card( rank, suit ):
    pass

class Deck:
    pass

class Shoe( Deck ):
    pass
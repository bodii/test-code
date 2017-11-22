#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 包含边界测试 '''

from  unittest import TestCase
import random

Suits = '♠', '♣', '♥', '♦'

# 工厂函数
class TestCardFactory( TestCase ):
    def test_rank1_should_createAceCard( self ):
        c = card( 1, Suits[1])
        self.assertIsInstance(c, AceCard)
    
    def test_rank2_should_createCard( self ):
        c = card(2, Suits[3])
        self.assertIsInstance(c, Card)
    
    def test_rank10_should_createCard( self ):
        c = card(10, Suits[2])
        self.assertIsInstance( c, Card )

    def test_rank10_should_createFaceCard( self ):
        c = card(11, Suits[0])
        self.assertIsInstance(c, Card)

    def test_rank13_should_createFaceCard( self ):
        c = card(13, Suits[1])
        self.assertIsInstance(c, Card)

    def test_otherRank_should_exception( self ):
        with self.assertRaises(LogicError):
            c = card(14, Suits[3])
        
        with self.assertRaises( LogicError ):
            c = card(0, Suits[3])



class Deck3( list ):
    def __init__( self, size=1, 
        random = random.Random(),
        card_factory=card):
        super().__init__()
        self.rng = random
        for d in range(size):
            super().extend(
                card_factory(r, s) for r in range(1, 13) for s in Suits
            )
        self.rng.shuffle( self )

    def deal( self ):
        try:
            return self.pop(0)
        except IndexError:
            raise DeckEmpty()


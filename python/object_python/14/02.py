#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 最小化依赖 '''

class Card:
    def __init__( self, rank, suit, hard=None, soft=None ):
        self.rank = rank
        self.suit = suit
        self.hard = hard or int(rank)
        self.soft = soft or int(rank)

    def __str__( self ):
        return '{0.rank!s}{0.suit!s}'.format(self)

class AceCard( Card ):
    def __init__( self, rank, suit ):
        super().__init__( rank, suit, 1, 11 )

class FaceCard( Card ):
    def __init__( self, rank, suit ):
        super().__init__( rank, suit, 10, 10 )


import random
# 可以（错误）定义一个Deck类，包含一些相关依赖
Suits = '♠', '♣', '♥', '♦'
# print( Suits )
class Deck1( list ):
    def __init__( self, size=1 ):
        super().__init__()
        self.rng = random.Random()
        for d in range(size):
            for s in Suits:
                cards = (
                    [AceCard(1, s)]
                    + [Card(r, s) for r in range(2, 12)]
                    + [FaceCard(r, s) for r in range(12, 14)]
                )
                super().extend( cards )
        self.rng.shuffle( self )


"""
以上设计有两点缺陷。首先，它与Card类层次结构中的3个类是紧耦合的。无法将Deck从Card
中分离出来做单元测试。其次，它依赖于随机数生成器，不是一个可重复的测试。
一方面，Card是一个非常简单的类，可以很容易将Deck以及所包含的Card一起测试。
另一方面，可能会想重用扑克牌或者皮纳克尔牌，它们在21点游戏中有不同的行为。
理想情况下，可以使用Deck独立于任何Card的实现。如果做到了这一点，就能够在不依赖Card
实现的前提下，对Deck进行测试，也可以对任何Card和Deck的组合情况进行测试。
"""
# 以下实现使用了工厂函数，是一种比较好的解耦方式
def card( rank, suit ):
    if rank == 1: return AceCard(rank, suit)
    elif 2 <= rank < 11: return Card(rank, suit)
    elif 11 <= rank < 14: return FaceCard(rank, suit)
    else: raise Exception( 'LogicError' )


"""
card()函数会基于传入的rank值完成Card子类的创建。这样一来，Deck类就可以使用这个函数
来完成创建Card实例的过程。我们通过插入一个中间函数将两个类的定义分离。
还胡其他用于将Card类从Deck类解耦的方式。可以对工厂函数进行重构，将它变成Deck类的方法。
也可以通过使用类级别的属性，或初始化方法参数的方式将类名独立出来进行绑定。
"""
# 在初始化方法中使用了复杂的绑定来代替工厂函数
class Deck2( list ):
    def __init__( self, size=1, random =random.Random(),
        ace_class=AceCard, card_class = Card, face_class=FaceCard
        ):
        super().__init__()
        self.rng = random
        for d in range(size):
            for s in Suits:
                cards = (
                    [ace_class(1, s)]
                    + [card_class(r, s) for r in range(2, 12)]
                    + [face_class(r, s) for r in range(12, 14)]
                )
                super().extend(cards)
        self.rng.shuffle( self )

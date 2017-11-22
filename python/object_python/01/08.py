#!/usr/bin/env python3
# -*- coding:utf-8 -*-

class Hand:
    def __init__( self, dealer_card, *cards ):
        self.dealer_card = dealer_card
        self.cards = list(cards)

    def __str__( self ):
        return ', '.join( map(str, self.cards) )

    def __repr__( self ):
        return "{__class__.__name__}({dealer_card!r}, {_cards_str})".format(
    __class__=self.__class__, _cards_str= ', '.join( map(repr, self.cards) ),
    **self.__dict__)

    def __eq__( self, other ):
        return self.cards == other.cards and self.dealer_card == other.dealer_card

    __hash__ = None
    # 这是一个包含适当的相等性比较的可变对象（__hash__是None）

# 下面是不可变的Hand版本
import sys
import defaultdict

class FrozenHand( Hand ):
    def __init__( self, *args, **kw ):
        if len(args) == 1 and isinstance(args[0], Hand):
            # Clone a hand
            other = args[0]
            self.dealer_card = other.dealer_card
            self.cards = other.cards
        else:
            # Build a fresh hand
            super().__init__( *args, **kw )

    def __hash__( self ):
        h = 0
        for c in self.cards:
            h = (h + hash(c)) % sys.hash_info.modulus
        return h 

""" 不变的版本中有一个构造函数，从另外一个Hand类创建一个Hand类。同时，还定义了一个
__hash__()方法，用sys.hash_info.modulus的值来计算cards的哈希值。大多数情况下，
这种基于模计算复合对象哈希值的方法能够满足我们的要求 """

# 使用这些类
stats = defaultdict(int) 
print(stats)
# 初始化一个整数数据字典，也可以用collections.Counter对象作为这个字典
d = Deck()
# Hand类冻结后，我们就可以将它用作字典的键，用这个键对应的值来统计实际的出牌次数
h  = Hand( d.pop(), d.pop(), d.pop() )
h_f = FrozenHand( h )
stats[h_f] += 1

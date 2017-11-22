#!/usr/bin/env python3
# -*- coding:utf-8 -*-


''' 属性访问、特性和修饰符 '''

# 属性

class Generic:
    pass

g = Generic()
g.attribute = 'value'
print( g.attribute )
# print( g.unset )
# AttributeError: 'Generic' object has no attribute 'unset'

del g.attribute
# print(g.attribute)
# AttributeError: 'Generic' object has no attribute 'attribute'


''' 另一种更好的办法是从types.SimpleNamepace类创建实例。此时不需要额外定义一个新类，就能实现同样
的功能。 '''
import types
n = types.SimpleNamespace()

n.attribute = 'value'
print( n.attribute )
# value

del n.attribute
# print( n.attribute )
# AttributeError: 'types.SimpleNamespace' object has no attribute 'attribute'

"""
使用SimpleNamespace类的做法会略有不同。object类的实例不允许创建新属性，因为它缺少Python内
部用来存储属性值的__dict__结构。
"""


''' 特性和__init__()方法 '''
"""
def split( self, deck ):
    assert self.cards[0].rank == self.cards[1].rank
    ''' [ assert 用法 ]
        1、assert语句用来声明某个条件是真的。
        2、如果你非常确信某个你使用的列表中至少有一个元素，而你想要检验这一点，
        并且在它非真的时候引发一个错误，那么assert语句是应用在这种情形下的理想
        语句。
        3、当assert语句失败的时候，会引发一AssertionError。
    '''
    try:
        self.split_count
        raise CannotResplit
    except AttributeError:
        h0 = Hand( self.dealer_card, self.cards[0], deck.pop() )
        h1 = Hand( self.dealer_card, self.cards[1], deck.pop() )
        h0.split_count = h1.split_count = 1
        return h0, h1
"""


''' 创建特性 '''

"""
    特性是一个函数，看起来（在语法上）就是一个简单的属性。我们可以获取、设置和删除特性
    值，正如我们可以获取、设置和删除属性值。这里有一个重要的区别：特性是一个函数，而且
    可以被调用，而不仅仅是用于存储的对象的引用。
    除了复杂程度，特性和属性的另一个区别在于，我们不能轻易地为已有对象添加新特性。但是
    默认情况下，我们可以很容易地给对象添加新属性。在这一点上，特性和属性有很大的区别。

    可以用两种方法来创建特性。我们可以使用 [ @property ]修饰符或者使用 [ property() ]
    函数。

    两种特性的设计模式：
    1. 主动计算（Eager Calculation）：每当更新特性值时，其他相关特性值都会立即被重新计算。
    2. 延迟计算（Lazy Calculation)：仅当访问特性时，才会触发计算过程。
"""

class Hand:
    def __str__( self ):
        return ', '.join( map(str, self.card) )

    def __repr__( self ):
        return '{__class__.__name__}({dealer_card!r}, {_cards_str})'.format(
        __class__ = self.__class__,
        _cards_str = ', '.join( map(repr, self.card) ),
        **self.__dict__
    )

class Hand_Lazy( Hand ):
    def __init__( self, dealer_card, *cards ):
        self.dealer_card = dealer_card
        self._cards = lsit(cards)

    @property
    def total( self ):
        delta_soft = max( c.soft - c.hard for c in self._cards )
        hard_total = sum( c.hard for c in self._cards )
        if hard_total + delta_soft <= 21: return hard_total + delta_soft
        return hard_total

    @property
    def card( self ):
        return self._cards
    
    @card.setter
    def card( self, aCard ):
        self._cards.append( aCard )

    @card.deleter
    def card( self ):
        self._cards.pop(-1)

    """
    Hand_Lazy类使用了一个Cards对象的集合来初始化Hand对象。其中total特性被定义为
    一个方法，仅当被调用时才会计算总值。另外，也定义了一些其他特性来更新手中的纸牌。
    card特性可以用来获取、设置或删除手中的牌。
    """


''' 主动计算特性 '''


class Hand_Eager( Hand ):
    def __init__( self, dealer_card, *cards ):
        self.dealer_card = dealer_card
        self.total = 0
        self._delta_soft = 0
        self._hard_total = 0
        self._cards = list()
        for c in cards:
            self.card = c

    @property
    def card( self ):
        return self._cards

    @card.setter
    def card( self, aCard ):
        self._cards.append(aCard)
        self._delta_soft = max(aCard.soft - aCard.hard, self._delta_soft)
        self._hard_total += aCard.hard
        self._set_total()

    @card.deleter
    def card( self ):
        removed = self._cards.pop(-1)
        self._hard_total -= removed.hard
        # Issue: was this the only ace?
        self._delta_soft = max( c.soft-c.hard for c in self._cards)
        self._set_total()

    def _set_total( self ):
        if self._hard_total + self._delta_soft <= 21:
            self.total = self._hard_total + self._delta_soft
        else:
            self.total = self._hard_total


''' setter和deleter特性 '''


"""
由于setter（和deleter）特性是基于getter属性创建的，因此先要使用如下代码定义一个getter
特性

@property
def card( self ):
    return self._cards

@card.setter
def card( self, aCard ):
    self._cards.append(aCard)

@card.deleter
def card( self ):
    self._cards.pop(-1)
"""

    
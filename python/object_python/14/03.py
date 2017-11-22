#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 创建简单的单元测试 '''

"""
有3种关于测试方法命名的方式：
1. 可以使用_should_将命名分为两个部分，例如StateUnderTest_should_Expected
Behavior, 包含了状态和结果。我们将重点关注这种方式。
2. 可以使用由when_和_should_组成的包含两个部分的名称，例如when_StateUnderTest_
should_ExpectedBehavior,也是包含了状态和结果，但是使用了更多的命名语法。
3. 可以使用一个包含了3个部分的名称，UnitOfWork_StateUndeerTest_ExpectedBehavior
它结合了要测试的工作单元进行说明，在读测试日志时会很有用。
"""


# 可以通过对unittest模块进行配置来使用不同的模式对测试方法进行查找

import unittest

class Card:
    def __init__( self, rank, suit, hard=None, soft=None ):
        self.rank = rank
        self.suit = suit
        self.hard = hard or int(rank)
        self.soft = soft or int(rank)

    def __str__( self ):
        return '{0.rank!s}{0.suit!s}'.format(self)

class TestCard( unittest.TestCase ):
    def setUp( self ):
        self.three_clubs = Card(3, '♣')
    
    def test_should_returnStr( self ):
        self.assertEqual( '3♣', str(self.three_clubs) )
    
    def test_should_getAttrValues( self ):
        self.assertEqual( 3, self.three_clubs.rank )
        self.assertEqual( '♣', self.three_clubs.suit )
        self.assertEqual( 3, slef.three_clubs.hard )
        self.assertEqual( 3, self.three_clubs.soft )


t = TestCard()
t.setUp()
t.test_should_returnStr()
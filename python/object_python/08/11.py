#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 转储和加载CSV '''

"""
CSV模块将简单的list和dict实例进行编码和解码，存入CSV模式。
一个CSV文件中每列的内容只是纯文本。

在namedtuple实例和CSV文件的行记录之间的映射是不难的，每行表示了一个不同的namedtuple
"""

from collections import namedtuple

GameStat = namedtuple( 'GameStat', 'player,bet,rounds,final' )

def gamestat_iter( player, betting, limit=100 ):
    for sample in range(30):
        b = Blackjack( player(), betting() )
        b.until_broke_or_rounds(limit)
        yield GameStat( 
            player.__name__, 
            betting.__name__,
            b.rounds, 
            b.betting.stake
            )

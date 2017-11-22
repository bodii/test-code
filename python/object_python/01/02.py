#!/usr/bin/env python3
# -*- coding:utf-8 -*-

# 用来生成Card子类对象的一个工厂函数

def card( rank, suit ):
    if rank == 1 : return AceCard( 'A', suit )
    elif 2 <= rank <= 11 : return NumberCard( str(rank), suit )
    elif 11 <= rank <14:
        name = { 11: 'J', 12: 'Q', 13: 'K' }[rank]
        return FaceCard( name, suit )
    else:
        raise Exception( 'Rank out of range' )


deck = [ card(rank, suit) 
        for rank in range(1, 14)
            for suit in (Club, Diamond, Heart, Spade) ]



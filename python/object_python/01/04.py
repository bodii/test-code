#!/usr/bin/env python3
# -*- coding:utf-8 -*-

from functools import partial

part_class = {
    1: partial(AceCard, 'A'),
    11: partial(FaceCard, 'J'),
    12: partial(FaceCard, 'Q'),
    13: partial(FaceCard, 'K'),  
}.get(rank, partial(NumberCard, str(rank)))
return part_class( suit )


class Card:
    def __init__( self, rank, suit, hard, soft ):
        self.rank = rank
        self.suit = suit
        self.hard = hard
        self.soft = soft


class NumberCard( card ):
    def __init__( self, rank, suit ):
        super().__init__( str(rank), suit, rank, rank )

class AceCard( Card ):
    def __init__( self, rank, suit ):
        super().__init__( 'A', suit, 1, 11 )

class FaceCard( Card ):
    def __init__( self, rank, suit ):
        super().__init__( {11: 'J', 12: 'Q', 13: 'K'}[rank], suit, 10, 10 )


def card10( rank, suit ):
    if rank == 1: return AceCard( rank, suit )
    elif 2 <= rank < 11: return NumberCard( rank, suit )
    elif 11 <= rank < 14: return FaceCard( rank, suit )
    else:
        raise Exception( 'Rank out of range' )


#!/usr/bin/env python3
# -*- coding:utf-8 -*-

class Card:
    insure = False
    def __init__( self, rank, suit, hard, soft ):
        self.rank  = rank
        self.suit = suit
        self.hard = hard
        self.soft = soft

    def __repr__( self ):
        return "{__class__.__name__}(suit={suit!r}, rank={rank!r})" .format(
    __class__ = self.__class__, **self.__dict__)

    def __str__( self ):
        return "{rank}{suit}".format(**self.__dict__)


class AceCard( Card ):
    def __init__( self, rank, suit ):
        super().__init__( 'A', suit, 1, 11 )

class FaceCard( Card ):
    def __init__( self, rank, suit ):
        super().__init__( { 11: 'J', 12: 'Q', 13: 'K' }[rank], suit, 10, 10 )


c1 = AceCard( 1, '♣' )
c2 = AceCard( 1, '♣' )

print( id(c1), id(c2) )
print( c1 is c2 )
print( hash(c1), hash(c2) )
print( id(c1) /16, id(c2) /16 )
print( c1 == c2 )

print( set([c1, c2]) )

print( c1.__class__.__name__ )
print( c1.__dict__ )
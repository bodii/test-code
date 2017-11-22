#!/usr/bin/env python3
# -*- coding:utf-8 -*-

class Card2:
    insure = False

    def __init__( self, rank, suit, hard, soft ):
        self.rank = rank
        self.suit = suit
        self.hard = hard
        self.soft = soft

    def __repr__( self ):
        return "{__class__.__name__}(suit={suit!r}, rank={rank!r})".format(
    __class__= self.__class__, **self.__dict__)

    def __str__( self ):
        return "{rank}{suit}".format(**self.__dict__)

    def __eq__( self, other ):
        return self.suit == other.suit and self.rank == other.rank

    def __hash__( self ):
        return hash(self.suit) ^ hash(self.rank)


class AceCard( Card2 ):
    insure = True

    def __init__( self, rank, suit ):
        super().__init__( 'A', suit, 1, 11 )


c1 = AceCard( 1, '♣' )
c2 = AceCard( 2, '♣' )

print( set([c1, c2]) )



class Card3:
    insure = False

    def __init__( self, rank, suit, hard, soft ):
        self.rank = rank
        self.suit = suit
        self.hard = hard
        self.soft = soft
    
    def __repr__( self ):
        return "{__class__.__name__}(suit={suit!r}, rank={rank!r})".format(
    __class__= self.__class__, **self.__dict__)

    def __str__( self ):
        return "{rank}{suit}".format( **self.__dic__ )

    def __eq__( self, other ):
        return self.suit == other.suit and self.rank == other.rank

    __hash__ = None

class AceCard3( Card3 ):
    insure = True

    def __init__( self, rank, suit ):
        super().__init__( 'A', suit, 1, 11 )


c1 = AceCard3( 1, '♣')
c2 = AceCard3( 1, '♣' )

print( id(c1), id(c2) )
# print( hash(c1), hash(c2) ) 
# 因为__hash__被设为None，所以这些用Card3生成的对象不可以被哈希，
# 也就无法通过hash()函数提供哈希值了。

print( c1 == c2 )

print( set([c1, c2]) )
# 

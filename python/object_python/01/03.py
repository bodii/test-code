#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用映射和类来简化设计 '''

def card4( rank, suit ):
    class_ = {1: AceCard, 11: FaceCard, 12: FaceCard, 
    13: FaceCard}.get(rank, NumberCard)
    return class_( rank, suit )


''' 也可以使用一个defaultdict类 '''
defaultdict( lambda: NumberCard, {1: AceCard, 11: FaceCard, 
12: FaceCard, 13: FaceCard})


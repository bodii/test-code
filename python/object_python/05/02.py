#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' deque类 '''


"""
一个list对象旨在为容器内的任何位置元素的修改提供一致的性能。但是，有一些操作会损失性能。最值得
注意的是，任何在list头部的操作(list.insert(0, item)) 或者list.pop(0))都会损失一些性能，
因为列表的大小改变了，同时所有元素的位置也改变了。

deque --- 一个双向队列 ---旨在为列表中的第1个和最后一个元素提供一致的性能。它的追加和弹出操
作会比内置的list对象快。
"""

# 用timeit来比较用list和deque洗牌的性能

from timeit import timeit 
from collections import deque

# l = timeit( 'random.shuffle(x)', """
# import random
# x=list(range(1*10))
# """
# )
# print( 'list:', l )
# list: 11.97796523300076

de = timeit('random.shuffle(d)', """
import random
from collections import deque
d = deque(range(1*10))
"""
)

print( 'deque:', de )
# deque: 12.345927018999646

from collections import deque
import random

class Deck( deque ):
    def __init__( self, size=1 ):
        super().__init__()
        for d in range(size):
            cards = [ card(r, s) for r in range(13) for s in Suits ]
            super().extend( cards )
        random.shuffle( self )

l = timeit('x.pop()', "x = list(range(100000)", number=100000)
# l = 0.03230..
de = timeit('x.pop()', """from collections import deque;
x=deque(range(100000))""", number=100000)
# de = 0.01350..



#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 记忆化与缓存 '''

"""
所谓的记忆化功能，即缓存上一次的计算结果为了下次可以重用。为了提高性能，可考虑通过使用更多
的内存而尽量避免计算。
一个普通的函数通常不会缓存上次的执行结果。函数通常是无状态的，然而一个可调用对象可以是有状态
的，它可以缓存上次的执行结果。
"""

import collections

class Power5( collections.abc.Callable ):
    def __init__( self ):
        self.memo = {}

    def __call__( self, x, n ):
        if (x, n) not in self.memo:
            if n == 0:
                self.memo[x, n] = 1
            elif n % 2 == 1:
                self.memo[x, n] = self.__call__(x, n-1) * x
            elif n % 2 == 0:
                t = self.__call__(x, n / 2)
                self.memo[x, n] = t * t
            else:
                raise Exception('Logic Error')

            return self.memo[x, n]

pow5 = Power5()
"""
修改算法， 加入self.memo缓存
如果x^n的值之前计算过，再次访问时将会从缓存取而非重新计算。这就是之前说的性能提升的地方。
"""

''' 使用functools完成记忆化 '''

from functools import lru_cache

@lru_cache(None)
def pow6( x, n ):
    if n == 0: return 1
    elif n % 2 == 1:
        return pow6( x, n-1 ) * x
    else:
        t = pow6( x, n/2)
        return t * t 

"""
lru_cache （Least Recently Used, LRU）最近最少使用的缓存。
它会把之前的请求放入缓存。请求在缓存中的大小是有限的。存入最新的请求，移除最近最少使用的
请求，正是LRU缓存机制的逻辑。
"""

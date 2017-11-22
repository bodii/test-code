#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 用__iter__()创建迭代器 '''

"""
尽管创建一个Iterator对象不是非常复杂，但是通常不需要这么做。创建生成器函数简单得多。对于一个封装
的集合，应该总是简单地把__iter__()方法的行为委托给内部的集合。

def __iter__( self ):
    return iter(self._list)
"""


''' 创建一种新的映射 '''
"""
Python内置了dict映射，在库中也有许多映射类型。除了collections模块对dict的扩展(defaultdic、
Counter和ChainMap)之外，库中还有一些模块包含了类似于映射的结构。

shelve模块是其他映射的一个重要示例。

mailbox 和 email.message模块中的类为邮箱提供了一个类似于dict的接口，这个接口被用于管理本地邮件。
"""

from collections import Counter
import math

class StatsCounter( Counter ):
    @property
    def mean( self ):
        sum0 = sum( v for k, v in self.items() )
        sum1 = sum( k * v for k, v in self.items() )
        return sum1 / sum0

    @property
    def stdev( self ):
        sum0 = sum( v for k, v in self.items() )
        sum1 = sum( k * v for k, v in self.items() )
        sum2 = sum( k * k * v for k, v in self.items() )
        return math.sqrt( sum0 * sum2 - sum1 * sum1) / sum0


sc = StatsCounter( [2, 4, 4, 4, 5, 5, 7, 9] )
print( sc.mean )
# 5.0
print( sc.stdev )
# 2.0

print( sc.most_common(1) )
# [(4, 3)]

print( list(sorted(sc.elements())) )
# [2, 4, 4, 4, 5, 5, 7, 9]

"""
most_common()的结果是一个包含两个元素的元组，其中一个是众数（4），另一个是这个值出现的次数（3）
。我们可能想获取前3个众数。通过执行类似sc.most_common(3)， 就能获取出现频率最高的几个值。
elements()方法按原始数据中所有元素的顺序重建列表。

"""
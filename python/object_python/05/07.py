#!/usr/bin/env python3
# -*- coding:utf-8 -*-

"""
实现MutableSequence类必须实现的方法：__getitem__\__setitem__\__delitem__\__len__\insert
\append\reverse\extend\pop\remove和__iadd__。

每个方法中必须实现的逻辑：
__getitem__: 无，因为不涉及状态的改变。
__setitem__: 这个方法改变一个元素的状态。我们需要从每个总和中减去原本元素的值，然后再将新元素的值
累加进总和中。
__delitem__: 这个方法会删除一个元素。我们需要从总和中移除被删除元素的值。
__len__: 无， 因为也不涉及状态的改变。
insert: 由于这个方法插入一个新元素。因为我们需要将这个元素累加进总和中。
append: 由于这个方法插入一个新元素。因为我们需要将这个元素累加进总和中。
reverse: 无，因为不会影响平均值和标准差的计算。
extend: 这个方法会添加许多元素，例如__init__,所以在扩展list之前，我们需要处理每个新加入的元素。
pop: 这个方法会删除一个元素。我们需要从总和中移除对应元素
remove: 这个方法也会删除一个元素。我们需要从总和中移除对应元素
__iadd__: 这个方法实现了 += 增量赋值语句，它和extend关键字完全相同。

"和常量" ： 变量包含了一种特定的和，并且这些和在类的状态被改变时仍然与类保持恒定的关系
(如下面StatsList2 的 sum0、sum1、sum2)
"""
import math

class StatsList2(list):
    """Eager Stats."""
    def __init__( self, *args, **kw ):
        self.sum0 = 0 # len( self )
        self.sum1 = 0 # sum( self )
        self.sum2 = 0 # sum( x**2 for x in self )
        super().__init__( *args, **kw )
        for x in self:
            self._new(x)
    
    def __setitem__(self, key, value ):
        super().__setitem__(key, value)
        self._new(value)

    def __delitem__(self, key ):
        super().__delitem__(key)
        self._rmv(self[key])

    def _new( self, value ):
        self.sum0 += 1
        self.sum1 += value
        self.sum2 += value * value 

    def _rmv( self, value ):
        self.sum0 -= 1
        self.sum1 -= value
        self.sum2 -= value * value

    def insert( self, index, value ):
        super().insert(index, value)
        self._new(value)

    def pop( self, index=0 ):
        value = super().pop(index)
        self._rmv(value)
        return value
    
    @property
    def mean( self ):
        return self.sum1 / self.sum0

    @property
    def stdev( self ):
        return math.sqrt( self.sum0 * self.sum2 -self.sum1 * self.sum1) / self.sum0


sl2 = StatsList2( [2, 4, 3, 4, 5, 5, 7, 9, 10])
print( sl2.sum0, sl2.sum1, sl2.sum2 )
# 9 49 325

sl2[2] = 4
print( sl2 )
# [2, 4, 4, 4, 5, 5, 7, 9, 10]
print( sl2.sum0, sl2.sum1, sl2.sum2 )
# 10 53 341
del sl2[-1] 
print( sl2.sum0, sl2.sum1, sl2.sum2 )
# 9 44 260
sl2.insert( 0, -1 )
print( sl2.sum0, sl2.sum1, sl2.sum2 )
# 10 43 261
print( sl2.pop() )
# -1
print( sl2.sum0, sl2.sum1, sl2.sum2 )
# 9 44 260

print( sl2.mean )
# 4.888888888888889
print( sl2.stdev )
# 2.23330569358242
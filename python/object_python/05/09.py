#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 实现__getitem__()\ __setitem__() 和__delitem__() '''
"""
当我们实现__getitem__()\ __setitem__() 和__delitem__()方法时，需要接受两种参数：int和slice。
"""

def __setitem__( self, index, value ):
    if isinstance(index, slice):
        start, stop, step = index.indices(len(self))
        olds = [ self[i] for i in range(start, stop, step) ]
        super().__setitem__( index, value )
        for x in olds:
            self._rmv(x)

        for x in value:
            self._new(x)

    else:
        old = self(index)
        super().__setitem__( index, value )
        self._rmv(old)
        self._new(value)

def __delitem__( self, index ):
    if isinstance(index, slice):
        start, stop, step = index.indices(len(self))
        olds = [ self[i] for i in range(start, stop, step) ]
        super().__delitem__( index )
        for x in olds:
            self._rmv(x)

    else:
        super().__delitem__(index)
        self._rmv(self[index])



''' 封装list和委托 '''

# 设计一个只支持append和__getitem__的类
import math

class StatsList3:
    def __init__( self ):
        self._list = list()
        self.sum0 = 0 # len(self), sometimes called 'N'
        self.sum1 = 0 # sum(self)
        self.sum2 = 0 # sum(x**2 for x in self)

    def append( self, value ):
        self._list.append(value)
        self.sum0 += 1
        self.sum1 += value
        self.sum2 += value * value

    def __getitem__( self, index ):
        return self._list.__getitem__( index )

    @property
    def mean( self ):
        return self.sum1 / self.sum0
    
    @property
    def stdev( self ):
        return math.sqrt( self.sum0 * self.sum2 - self.sum1 * self.sum1) /self.sum0

"""
这个类中的_list对象是Python内置的list类。这个列表总是初始化为空列表。由于append()是唯一的更新
列表的方式，我们可以很容易地维护不同种类的和。
可以直接将__getitem__()委托给内部的_list对象，而不用去关心参数和结果。
"""

sl3 = StatsList3()
for data in 2, 4, 4, 4, 5, 5, 7, 9:
    sl3.append(data)

print( sl3.mean )
# 5.0

print( sl3.stdev )
# 2.0

"""
我们创建一个空列表，然后将元素添加到列表中。由于每次有元素添加到列表中时，都会更新“和常量”，
因此可以快速地算出平均值和标准差。

没有让类变成可迭代的，没有定义__iter__()。
"""

print(sl3[0])
# 2

for x in sl3:
    print(x)

# 2
# 4
# 4
# 4
# 5
# 5
# 7
# 9

# print(len(sl3))
# TypeError: object of type 'StatsList3' has no len()
"""
可以添加一个__len__()方法并将它委托给内部的_list对象。可能也会想将__hash__设为None，
但是需要很小心，因为这是一个可变对象。

可以定义__contains__()并且将真正的工作委托给内部的_list对象。这样一来，就可以创建一个极简单
的容器，但是它仍然提供了一个容器所应具有的底层特性。
"""
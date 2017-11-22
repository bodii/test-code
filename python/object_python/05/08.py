#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' __getitem__ __setitem__ __delitem__ 和 slice 操作 '''


"""
了解slice操作

序列包括了两种不同类型的索引：
1. a[i]: 这是一个简单的整数索引
2. a[i:j] 或者 a[i:j:k]: 这些使用了start:stop:step值的slice表达式。slice表达式有7种不同的
重载。

当我们做一些类似于seq[:-1]的操作时，我们就是在写slice表达式。底层的__getitem__()方法会接受一个
slice对象作为参数，而不是一个简单的整数。

一个slice对象包含3个属性：start、stop、step。同时，它也有一个叫作indices()的函数，当上述任何
属性缺失时，这个函数会正确地计算出缺失属性的值。
"""
class Explore( list ):
    def __getitem__( self, index ):
        print( index, index.indices(len(self)) )
        return super().__getitem__( index )

x = Explore( 'abcdefg' )
print( x[:] ) 
# slice(None, None, None) (0, 7, 1)
# ['a', 'b', 'c', 'd', 'e', 'f', 'g']

print( x[:-1] )
# slice(None, -1, None) (0, 6, 1)
# ['a', 'b', 'c', 'd', 'e', 'f']

print( x[1:] )
# slice(1, None, None) (1, 7, 1)
# ['b', 'c', 'd', 'e', 'f', 'g']

print( x[::2] )
# slice(None, None, 2) (0, 7, 2)
# ['a', 'c', 'e', 'g']
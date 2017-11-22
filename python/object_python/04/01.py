#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 可调用对象和上下文的使用 '''

''' 使用ABC可调用对象进行设计 '''

"""
在Python中有两种创建可调用对象的方式：
1. 使用def语句创建一个函数。
2. 通过创建继承自collections.abc.Callable类的实例。

也可以将一个变量赋值为lambda表达式。
一个lambda表达式是一个小的匿名函数，其中只包含一个表达式语句。我们不倾向于将lambda表达式
保存在变量中，因为当使用了一个类似函数的、并未使用def语句定义可调用对象，这种做法就会带来
困惑。
"""

import collections.abc

class Power1( collections.abc.Callable ):
    def __call__( self, x, n ):
        p = 1
        for i in range(n):
            p *= x
        return p


pow1 = Power1()
"""
以上的可调用对象包含了3部分：
1. 类继承自abc.Callable
2. 定义了__call__()方法。
3. 创建了类的实例pow1()

像使用其他函数那样，现在可以使用刚刚定义的pow1()函数了。如下：
"""
print( pow1( 2, 0 ) )
# 1
print( pow1( 2, 1 ) )
# 2
print( pow1( 2, 2 ) )
# 4
print( pow1(2, 10 ) )
# 1024


"""
class Power2( collections.abc.Callable ):
    def __call_( self, x, n ):
        p = 1
        for i in range(n):
            p *= x
        return p 
        
        以上的类定义有一个错误并且不符合可调用基类的定义规则。
        

pow2 = Power2()

# methods __call__
"""



''' 提高性能 '''

"""
首先，考虑使用更好的算法。然后，考虑带有记忆的算法，包含缓存。因此，函数变得有状态了，
这是可调用对象的亮点。
"""

"""
如果只是简单地使用一个固定的值做乘法运算，我们会使用"快速幂"算法。它使用了以下3个
基本逻辑来计算x^n。

1. 如果 n =0: x^n = 1,结果返回1。
2. 如果 n 是奇数并且n mod 2 = 1, 结果为x^n-1 * x。这里包含了x^n-1的递归运算。这里
仍做了一次乘法运算但并非是一个优化的点。
3. 如果是偶数并且 n mod = 0, 结果为x^n/2 * x^n/2。 这里包含了x^n/2的递归计算。这里
把乘法计算的数量减少了一半。

以下是一个递归的、可调用对象的实现：
"""
class Power4( collections.abc.Callable ):
    def __call__( self, x, n ):
        if n == 0 : return 1
        elif n % 2 == 1:
            return self.__call__(x, n-1) * x
        else: # n % 2 == 0
            t = self.__call__(x, n / 2)
            return t * t

pow4 = Power4()
print(pow4(2, 2))
# 4

import timeit
iterative = timeit.timeit( 'pow1(2, 1024)', """
import collections.abc
class Power1( collections.abc.Callable ):
    def __call__( self, x, n ):
        p = 1
        for i in range(n):
            p *= n

        return p

pow1 = Power1()
""", number=100000
# 通过定义number=100000来让测试更快地结束。
)

print( 'Iterative', iterative )
# Iterative 59.55600157199996
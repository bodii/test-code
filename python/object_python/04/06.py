#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 定义__enter__()和__exit__()方法 '''

"""
上下文管理器的定义包含两个特殊方法：__enter__()和__exit__()。 with语句使用它们进行上下文 
的进入和退出。
"""

import random

class knownSequence:
    def __init__( self, seed = 0 ):
        self.seed = 0

    def __enter__( self ):
        self.was = random.getstate()
        random.seed( self.seed, version=1 )
        return self
    def __exit__( self, exc_type, exc_value, traceback ):
        random.setstate( self.was )

"""
上面定义了所需的__enter__()和__exit__()方法。__enter__()方法会保存随机模块上次的状态
并将种子重置为设定的值。__exit__()方法用于恢复随机数生成器之前的状态。
注意，__enter__()方法返回了self.这点对于被添加到其他类定义中的mixin上下文管理器来说是常
见的。
__exit__()方法的参数值在正常情况下会被赋值为None。除非我们有特殊的异常处理需要，我们通常会
忽略参数值。
"""

print( tuple(random.randint(-1, 36) for i in range(5)) )
# (27, 19, 2, 22, 36)

with knownSequence():
    print( tuple(random.randint(-1, 36) for i in range(5)) )

# (23, 25, 1, 15, 31)

print( tuple(random.randint(-1, 36) for i in range(5)) )
# (29, 34, 17, 0, 15)

with knownSequence():
    print( tuple(random.randint(-1, 36) for i in range(5)) )
# (23, 25, 1, 15, 31)
"""
每次创建一个KnownSequence的实例，修改了random模块的实现。在with语句的上下文中，会得到
一串固定的随机数。而在上下文之外，由于随机种子被复原了，因此会得到随机的数值。

种子在上下文中被固定了。
"""
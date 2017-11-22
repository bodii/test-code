#!/usr/bin/env python3 
# -*- coding:utf-8 -*-

'''
 __del__()

 CPython的实现中，对象会包括一个引用计数器。当对象被赋值给一个变量时，这个计数器会递增;、
 当变量被删除时，这个计数器会递减。当引用计数器的值为0时，表示我们的程序不再需要这个对象
 并且可以销毁这个对象。对于简单对象，当执行删除对象的操作时会调用__del__()方法。

 对于包含循环引用的复杂对象，引用计数器有可能永远不会归零，这样就很难让__del__()被调用。
  '''

class Noisy:
    def __del__( self ):
        print( 'Removeing {0}'.format( id(self)) )


x = Noisy()
del x

# 浅复制
ln = [ Noisy(), Noisy() ]
ln2 = ln[:]
del ln
del ln2

a = b = Noisy()
c = [ Noisy() ] * 2

#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 上下文管理器工厂 '''

"""
可以创建一个上下文管理器类来作为应用程序对象的工厂。这样的设计使得耦合降低，而且无需在应用程序
编写过多有关上下文管理器功能的逻辑。
"""
import random 
class Deterministic_Deck:
    def __init__( self, *args, **kw ):
        self.args = args
        self.kw = kw

    def __enter__( self ):
        self.was = random.getstate()
        random.seed( 0, version=1 )
        return Deck( *self.args, **self.kw )

    def __exit__( self, exc_type, exc_value, traceback ):
        random.setstate( self.was ) 

"""
以上的上下文管理器类存了参数值，为了用于创建Deck。
__enter__()方法存了旧的随机数状态，然后将随机数模块的逻辑改为：产生固定的随机值，用来创建和
洗牌。
注意__enter__()方法返回了一个新的Deck对象，用于with语句的上下文中，并使用with的as语句为其
赋值。
"""

""" 以下是使用这个上下文管理器工厂的方式"""
with Deterministic_Deck( size=6 ) as deck:
    h = Hand( deck.pop(), deck.pop(), deck.pop() )
    
    
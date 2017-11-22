#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 设计更高级的复合命令 '''

class Command:
    def set_config( self, config ):
        self.__dict__.update(config.__dict__)

    config = property(fset=set_config)
    # property 返回属性的属性
    # 等同于@property，不过它是设置
    '''
    class C:
        def __init__( self ):
            self._x = None

        def getx( slef ):
            return self._x

        def setx( self, value ):
            self._x = value

        def delx( self ):
            del self._x

        x = property(getx, setx, delx, "I'm  the 'x' property.")

    # 调用 
    c = C()
    c.x 将调用getter,
    c.x = value 将调用setter
    del c.x 将调用删除器


    等同于
    class C:
        def __init__( self ):
            self._x = None

        @property
        def x( self ):
            """I'm the 'x' property."""
            return self._x

        @x.setter
        def x( self, value ):
            self._x = value

        @x.deleter
        def x( self ):
            del self._x
    '''

    def run( self ):
        pass


class Command_Sequence( Command ):
    sequence = []
    def __init__( self ):
        slef._sequence = [ class_() for class_ in self.sequence ]

    def set_config( self, config ):
        for step in self._sequence:
            step.config = config

    config = property( fset=set_config )
    
    def run( self ):
        for step in self._sequence:
            step.run()

# 基于两个Command子类创建的另一个Command子类
class Simulate_and_Analyze( Command_Sequence ):
    sequence = [Simulate_Command, Analyze_Command]


import argparse

parser = argparse.ArgumentParser()

parser.add_argument('command', action='store', default='simulate',
choices=['simulate', 'analyze', 'simulate_analyze'])

args = parser.parse_args()
print( args )


class ForAllBets_Simulate( Command ):
    def run( self ):
        for bet_class in 'Flat', 'Martingale', 'OneThreeTwoSix':
            self.betting_rule = bet_class
            self.outputfile = 'p3_c16_simulation7_{0}.dat'.format(bet_class)
            sim = Simulate_Command()
            sim.conifg = self
            sim.run()
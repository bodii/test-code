#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 为控制、调试、审计和安全创建专门的日志 '''

"""
日志有很多不同种类，我们会关注下面的4种
1. 错误和控制（Errors and Contrrol）

2. 调试（Debugging）

3. 设计（Audit）

4. 安全（Security）
"""

from collections import Counter
import logging 

# 日志记录器
class Main:
    def __init__( self ):
        self.balance = Counter()
        self.log = logging.getLogger( self.__class__.__qualname__ )

    def run( self ):
        self.log.info( 'Start' )

        # Some processing
        self.balance['count'] += 1
        self.balance['balance'] += 3.14

        self.log.info( 'Counts {0}'.format( self.balance ) )
        for k in self.balance:
            self.log.info('{0:.<16s} {1:n}'.format(k, self.balance[k]))


def logged(class_):
    class_.logger = logging.getLogger(class_.__qualname__)
    return class_


@logged
class OneThreeTwoSix( BettingStrategy ):
    def __init__( self ):
        self.wins = 0

    def _state( self ):
        return dict( wins = self.wins)

    def bet( self ):
        bet = {0: 1, 1: 3, 2: 2, 3: 6}[slef.wins % 4]
        self.logger.debug( 'Bet {1}; based on {0}'.format(self._state(), bet) )

    def record_win( self ):
        self.wins += 1
        self.logger.debug( 'Win: {0}'.format(self._state()) )

    def record_loss( self ):
        self.wins = 0
        self.logger.debug( 'Loss: {0}'.format(self._state()) )
    

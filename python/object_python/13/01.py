#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 测试、调试、部署和维护 '''

''' Loggin和Warning模块 '''

''' 创建基本日志 '''

"""
创建日志有两个必要的步骤：
1. 通过logging.getLogger()函数获得logging.Logger实例
2. 用获得的Logger创建消息。有许多用于创建不同重要性级别消息的方法，例如warn()、info()
debug()、error()和fatal()。

以上不提供任何输出。只有当我们需要查看输出的时候，才会使用第3个步骤。有一些日志是为调试准备的，
所以并不总是希望看到这种日志。还有一个可选的步骤是配置logging模块的handlers、filters和
formatters,可以用logging.basicConfig()函数完成这些配置。
"""

import logging
import sys

logging.basicConfig( stream=sys.stderr, level=logging.DEBUG )


class Player:
    def __init__( self, bet, strategy, stake ):
        self.logger = logging.getLogger( self.__class__.__qualname__ )
        self.logger.debug( 'int bet {' + bet + ', strategy ' + strategy + ', stake' + stake +'.' )

def logged( class_ ):
    class_.logger = logging.getLogger( class_.__qualname__ )
    return class_


@logged
class Player2:
    def __init__( self, bet, strategy, stake ):
        self.logger.debug( 'int bet ' + bet + ', strategy ' + strategy + ', stake ' + stake +'.' )

p = Player2('1', '3', '5')

  

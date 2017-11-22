#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 创建审计和安全日志 '''

"""
审计和安全日志通常会在两个处理程序中重复：主控制处理程序和用来检查审计与安全的文件
处理程序，这意味着我们需要做下面几件事：
1. 为审计和安全定义额外的日志处理器
2. 为这个日志处理器定义多个处理程序
3. 可选地为审计处理程序定义不同的格式

"""
import logging 
import yaml
import os.path
from logging import config

current_path = os.path.dirname(__file__) + os.path.sep
filename = current_path + 'test_config2.yaml'


# 创建包含审计功能的类的装饰器
def audited( class_ ):
    class_.logger = logging.getLogger( class_.__qualname__ )
    class_.audit = logging.getLogger( 'audit.' + class_.__qualname__  )
    return class_


@audited
class Table:
    def bet( self, bet, amount ):
        self.audit.info('Bet {0} Amount {1}'.format(bet, amount))

if __name__ == '__main__':
    with open(filename) as file:
        logging.config.dictConfig(yaml.load(file))
    t = Table
    t.bet('1', '2')
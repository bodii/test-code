#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 扩展日志等级 '''

"""
logging模块预定义了5个重要性级别。每个级别都有一个（或者两个）带有级别号码的全局变量。
重要性级别代表的是从调试信息（很少重要到需要显示出来)到关键的或是致命的错误（总是非常
重要）的一个可选范围

日志模块变量               值

DEBUG                    10
INFO                     20
WARNING 或WARN           30
ERROR                    40
CRITICAL或者FATAL         50


我们可以像下面这样定义详细信息的新级别：
logging.addLevelName(15, 'VERBOSE')
logging.VERBOSE = 15

可以通过Logger.log()方法使用新级别，这个方法接受一个级别数字作为参数
self.logger.log( logging.VERBOSE, 'Some Message' )
"""

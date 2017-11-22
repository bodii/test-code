#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 发送日志消息到远程的进程 '''

"""
共享日志解决方案会使用multiprocessing的共享日志。

建立多进程应用程序需要3个步骤：
1. 首先，创建共享队列对象，这样一来日志消息志就可以筛选消息
2. 其次， 创建从队列中获取日志的记录的消费者进程
3. 最后，我们会创建用于处理真正应用程序工作的生产者进程池，这些进程会将记录写入
共享队列。
"""

# 消费者进程的定义

import collections
import logging
import multiprocessing
import logging.config
import yaml
import os.path

BASE_PATH = os.path.dirname(__file__)
filename = os.path.join(BASE_PATH ,'consumer_config.yaml')
with open(filename) as file:
    consumer_config = file

class Log_Consumer_1(multiprocessing.process):
    """In effect, an instance of QueueListener."""
    def __init__( self, queue ):
        self.source = queue
        super().__init__()
        logging.config.dictConfig(yaml.load(consumer_config))
        self.combined = logging.getLogger(
            'combined.' + self.__class__.__qualname__
        )
        self.log = logging.getLogger(self.__class__.__qualname__)
        slef.counts = collections.Counter()
    
    def run( self ):
        self.log.info( 'Consumer Started' )
        while True:
            log_record = self.source.get()
            if log_record == None: break
            self.combined.handle( log_record )
            words = log_record.getMessage().split()
            self.count[words[1]] += 1
        self.log.info( 'Consumer Finished' )
        self.log.info( self.counts )


# 日志生产者
class Log_Producer(multiprocessing.process):
    handler_class = logging.handlers.QueueHandler
    def __init__( self, proc_id, queue ):
        self.proc_id = proc_id
        self.destination = queue
        super().__init__()
        self.log = logging.getLogger(
            '{0}.{1}'.format(self.__class__.__qualname__, self.proc_id)
        )
        self.log.handlers = [ self.handler_class(self.destination) ]
        self.log.setLevel( logging.INFO )

    def run( self ):
        self.log.info('Producer {0} started'.format(self.proc_id))
        for i in range(100):
            self.log.info('Producer {:d} Message {:d}'.format(self.proc_id, i))
        self.log.info('Producer {0} Finished'.format(self.proc_id))


# 创建队列
queue = multiprocessing.queues(100)
# 启动消费者进程
consumer = Log_Consumer_1(queue)
consumer.start()
# 启动生产者进程数组
producers = []
for i in range(10):
    proc = Log_Producer(i, queue)
    proc.start()
    producers.append( proc )

# 10个并发的生产者将会让队列溢出。每个生产者都会收到一些队列异常，这些异常就
# 意味着消自丢失了。
# 正确地结束空这个处理过程
for p in producers:
    p.join()
queue.put( None )
consumer.join()
"""
首先，我们等待每个生产者进程结束，然后加入父进程。其次，我们将一个哨兵对象插入队列中，
这样消费者进程就会正确结束。最后，我们等待消费进程结束并加入进程。
"""

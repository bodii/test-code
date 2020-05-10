#!/usr/bin/env python3
# -*- encoding=utf-8 -*-

from threading import Thread, Event
from multiprocessing import cpu_count
from typing import List
from task import Task
from queue import ThreadSafeQueue

'''
任务处理线程
'''
class ProcessThread(Thread):

    def __init__(self, task_queue, *args, **kwargs):
        Thread.__init__(self, *args, **kwargs)
        # 线程停止标记
        self.dismiss_flag = Event()
        # 任务队列
        self.task_queue = task_queue
        self.args = args
        self.kwargs = kwargs

    def run(self):
        while True:
            # 判断线程是否被要求停止
            if self.dismiss_flag.is_set():
                break

            task_q = self.task_queue.pop()
            if not isinstance(task_q, Task):
                continue
            result = task_q.callable(*task_q.args, **task_q.kwargs)

    def dismiss(self):
        self.dismiss_flag.set()

    def stop(self):
        self.dismiss()


'''
任务类型错误异常类
'''
class TaskTypeErrorException(Exception):
    pass

'''
线程池类
'''
class ThreadPool:
    def __init__(self, size=0):
        if not size:
            # 线程池的大小为CPU核数的2倍
            size = cpu_count() * 2
        # 线程池
        self.pool = ThreadSafeQueue(size)
        # 任务队列
        self.task_queue = ThreadSafeQueue()

        # 将任务队列放在线程池中
        for i in range(size):
            self.pool.put(ProcessThread(task_queue=self.task_queue))

    def start(self):
        """
        开始线程池中的处理程序
        """
        for i in range(self.pool.size()):
            thread = self.pool.get(i)
            thread.start()

    def join(self):
        """
        停止线程池中的线程处理程序
        """
        for i in range(self.pool.size()):
            thread = self.pool.get(i)
            thread.stop()
        
        while self.pool.size():
            thread = self.pool.pop()
            thread.join()

    def put(self, item: Task):
        """
        向线程池提交任务

        :param:item: 任务元素\n
        """
        if not isinstance(item, Task):
            raise TaskTypeErrorException()
        self.task_queue.put(item)

    def batch_put(self, item_list: List):
        """
        向线程池添加一组任务

        :param:item_list: 任务列表\n
        """
        if not isinstance(item_list, list):
            item_list = list(item_list)

        for item in item_list:
            self.put(item)

    def size(self):
        """
        获取线程池中任务的数量

        :return: 任务的数量
        """
        return self.pool.size()
    
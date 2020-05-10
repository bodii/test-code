#!/usr/bin/env python3
# -*- encoding=utf-8 -*-

'''
测试自定义线程池
'''

import time
from task import Task
from pool import ThreadPool


class SimpleTask(Task):
    def __init__(self, callable):
        super().__init__(callable)
    pass

def process():
    time.sleep(1)
    print("This is a SimpleTask callable function.1")
    time.sleep(1)
    print("This is a SimpleTask callable function.2")

def test():
    # 1. 初始化一个线程池
    test_pool = ThreadPool()
    test_pool.start()

    # 2. 生成一系列的任务
    for i in range(10):
        simple_task = SimpleTask(process)
        # 3. 往线程池提交任务执行
        test_pool.put(simple_task)


if __name__ == '__main__':
    test()
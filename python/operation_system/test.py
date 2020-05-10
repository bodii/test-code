#!/usr/bin/env python3
# -*- encoding=utf-8 -*-

'''
测试自定义线程池
'''

import time
from task import Task, AsyncTask
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

def test_acync_task():
    def async_process() -> int:
        num = 0
        for i in range(100):
            num += i
        
        return num


    # 1. 初始化一个线程池
    test_pool = ThreadPool()
    test_pool.start()

    # 2. 生成一系列的任务
    for i in range(10):
        async_task = AsyncTask(func=async_process)
        # 3. 往线程池提交任务执行
        test_pool.put(async_task)
        result = async_task.get_result()
        print(f'get result: {result}')


def test_acync_task2():
    def async_process() -> int:
        num = 0
        for i in range(100):
            num += i
        # 等待一下
        time.sleep(1) 

        return num


    # 1. 初始化一个线程池
    test_pool = ThreadPool()
    test_pool.start()

    # 2. 生成一系列的任务
    for i in range(10):
        async_task = AsyncTask(func=async_process)
        # 3. 往线程池提交任务执行
        test_pool.put(async_task)
        print(f'get result in timestamp: {int(time.time())}')
        result = async_task.get_result()
        print(f'get result in timestamp: {int(time.time())}, result: {result}')


def test_acync_task3():
    def async_process() -> int:
        num = 0
        for i in range(100):
            num += i

        return num


    # 1. 初始化一个线程池
    test_pool = ThreadPool()
    test_pool.start()

    # 2. 生成一系列的任务
    for i in range(10):
        async_task = AsyncTask(func=async_process)
        # 3. 往线程池提交任务执行
        test_pool.put(async_task)
        print(f'get result in timestamp: {int(time.time())}')
        # 等待一下
        time.sleep(5) 
        result = async_task.get_result()
        print(f'get result in timestamp: {int(time.time())}, result: {result}')

if __name__ == '__main__':
    # test()
    # test_acync_task()
    # test_acync_task2()
    test_acync_task3()
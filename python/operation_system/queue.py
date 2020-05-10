#!/usr/bin/env python3
#-*- encoding=utf-8 -*- 


from threading import Lock, Condition, Thread
import time
from typing import List, Any

from task import Task, AsyncTask

class ThreadSafeQueueException(Exception):
    pass

'''
线程安全的队列
'''
class ThreadSafeQueue:
    def __init__(self, max_size=0):
        self.queue = []
        self.max_size = max_size
        # 互斥锁
        self.lock = Lock()
        # 条件锁
        self.condition = Condition()
    
    def size(self):
        """
        获取当前队列元素的数量
        """
        # 加锁
        self.lock.acquire() 
        # 获取元素列表的长度
        size = len(self.queue)
        # 解锁
        self.lock.release()

        return size

    def put(self, item: Task):
        """
        往队列里放入元素
        """
        if self.max_size != 0 and self.size() > self.max_size:
            return ThreadSafeQueueException()

        # 加锁
        self.lock.acquire() 
        # 添加元素
        self.queue.append(item)
        # 解锁
        self.lock.release()

        self.condition.acquire()
        self.condition.notify()
        self.condition.release()

    def batch_put(self, item_list: List[Task]):
        """
        往队列内添加元素列表
        """
        if not isinstance(item_list, list):
            item_list = list(item_list)
        for item in item_list:
            self.put(item)

    def pop(self, block: bool=False, timeout: int=None)-> Any: 
        """
        从队列中取出元素

        :param block: 是否需要阻塞等待\n
        :param timeout: 等待的时间\n
        :return: 取出的元素或None\n
        """
        if self.size() == 0:
            # 如果需要阻塞等待
            if block:
                self.condition.acquire()
                self.condition.wait(timeout)
                self.condition.release()
            else:
                return None
       
        # 如果还有元素，则
        # 加锁
        self.lock.acquire()
        item = None
        if len(self.queue) > 0:
            # 取出元素
            item = self.queue.pop()
        # 解锁
        self.lock.release()
        # 返回元素
        return item


    def get(self, index: int)-> Any:
        """
        获取列表中指定索引的元素

        :param: index 元素列表的索引\n
        :return: 获取的元素或None\n
        """
        if self.size() == 0:
            return None

        self.lock.acquire()
        item = self.queue[index] or None
        self.lock.release()

        return item


if __name__ == '__main__':

    queue = ThreadSafeQueue(max_size=100)
    def producer():
        while True:
            queue.put(1)
            time.sleep(1)

    def consumer():
        while True:
            print('get item from queue: %d' % queue.pop(block=True, timeout=2))
            time.sleep(1)

    thread1 = Thread(target=producer)
    thread2 = Thread(target=consumer)
    thread1.start()
    thread2.start()
    thread1.join()
    thread2.join()
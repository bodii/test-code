#!/usr/bin/env python3
# -*- coding:utf-8 -*-
# PW @ 2019-11-12 22:59:43

# 队列
class Queue:
    __container = []

    def __init__(self):
        pass

    def enqueue(self, item):
        """
        入队
        """
        self.__container.append(item)


    def dequeue(self):
        """
        出队
        """
        return self.__container.pop(0)


    def size(self) -> int:
        """
        队的长度
        """
        return len(self.__container)


    def is_empty(self) -> bool:
        """
        队列是否为空
        """
        return self.size() <= 0


if __name__ == '__main__':
    queue = Queue()
    print("is empty: ", queue.is_empty())
    queue.enqueue(10)
    queue.enqueue(20)
    queue.enqueue(30)
    print("length: ", queue.size())
    print(queue.dequeue())
    print(queue.dequeue())
    print(queue.dequeue())
    print("is empty: ", queue.is_empty())



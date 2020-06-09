#!/usr/bin/env python3
# -*- coding:utf-8 -*-
# PW @ 2020-06-09 13:49:34

from typing import List
from threading import Thread, Condition, Lock, ThreadError
from random import Random, randrange
from time import sleep

class Request:
    def __init__(self, name: str):
        self.name = name

    def getName(self):
        return self.name

    def __str__(self):
        return '[ Request %s ]' % (self.name)
    
    __repr__ = __str__


class RequestQueue:
    def __init__(self):
        self.queue = []
        self.cv = Condition()

    def getRequest(self) -> Request:
        self.cv.acquire()
        while 1 > len(self.queue):
            self.cv.wait()
        self.cv.release()
        return self.queue.pop()

    def putRequest(self, request: Request):
        self.cv.acquire()
        self.queue.append(request)
        self.cv.notify_all()
        self.cv.release()


class ServerThread(Thread):
    def __init__(self, requestQueue: RequestQueue, name: str, seed: int):
        Thread.__init__(self,name=name)
        self.requestQueue = requestQueue
        self.seed = seed

    def run(self):
        for i in range(1, 1000):
            try:
                request = self.requestQueue.getRequest()
                print(self.getName() + ' handles ' + str(request))
                sleep(Random(self.seed))
            except Exception as e:
                pass


class ClientThread(Thread):
    def __init__(self, requestQueue: RequestQueue, name: str, seed: int):
        Thread.__init__(self,name=name)
        self.requestQueue = requestQueue
        self.seed = seed

    def run(self):
        for i in range(1, 1000):
            request = Request('No.' + str(i))
            self.requestQueue.putRequest(request)
            print(self.getName() + ' handles ' + str(request))

            try:
                sleep(Random(self.seed))
            except Exception as e:
                pass


if __name__ == '__main__':
    requestQueue = RequestQueue()
    ClientThread(requestQueue, 'Alice', 0.357642).start()
    ServerThread(requestQueue, 'Bobby', 0.642234).start()
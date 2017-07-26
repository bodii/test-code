#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十二章：并发编程
Description: 对于并发编程, Python 有多种长期支持的方法, 包括多线程, 调用子进程, 以及各种
各样的关于生成器函数的技巧. 这一章将会给出并发编程各种方面的技巧, 包括通用的
多线程技术以及并行计算的实现方法.
像经验丰富的程序员所知道的那样, 大家担心并发的程序有潜在的危险. 因此, 本章
的主要目标之一是给出更加可信赖和易调试的代码.

Title:            判断线程是否已经启动
Issue: 你已经启动了一个线程，但是你想知道它是不是真的已经开始运行了
Answer: 线程的一个关键特性是每个线程都是独立运行且状态不可预测。如果程序中的其
他线程需要通过判断某个线程的状态来确定自己下一步的操作，这时线程同步问题就
会变得非常棘手。为了解决这些问题，我们需要使用 threading 库中的 Event 对象。
Event 对象包含一个可由线程设置的信号标志，它允许线程等待某些事件的发生。在
初始情况下， event 对象中的信号标志被设置为假。如果有线程等待一个 event 对象，
而这个 event 对象的标志为假，那么这个线程将会被一直阻塞直至该标志为真。一个
线程如果将一个 event 对象的信号标志设置为真，它将唤醒所有等待这个 event 对象的
线程。如果一个线程等待一个已经被设置为真的 event 对象，那么它将忽略这个事件，
继续执行。
"""

from threading import Thread, Event
import time

def countdown(n, started_evt):
    print('countdown starting')
    started_evt.set()
    while n > 0:
        print('T-minus', n)
        n -= 1
        time.sleep(5)

started_evt = Event()
print('Launching countdown')
t = Thread(target=countdown, args=(10, started_evt))
t.start()
started_evt.wait()
print('countdown is running')

"""
当你执行这段代码，“ countdown is running”总是显示在“ countdown starting”之
后显示。这是由于使用 event 来协调线程，使得主线程要等到 countdown() 函数输出
启动信息后，才能继续执行。
"""

"""
event 对象最好单次使用，就是说，你创建一个 event 对象，让某个线程等待这个
对象，一旦这个对象被设置为真，你就应该丢弃它。尽管可以通过 clear() 方法来重
置 event 对象，但是很难确保安全地清理 event 对象并对它重新赋值。很可能会发生错
过事件、死锁或者其他问题（特别是，你无法保证重置 event 对象的代码会在线程再
次等待这个 event 对象之前执行）。如果一个线程需要不停地重复使用 event 对象，你
最好使用 Condition 对象来代替。下面的代码使用 Condition 对象实现了一个周期定
时器，每当定时器超时的时候，其他线程都可以监测到
"""
import threading
import time

class PeriodicTimer:
    def __init__(self, interval):
        self._interval = interval
        self._flag = 0
        self._cv = threading.Condition()

    def start(self):
        t = threading.Thread(target=self.run)
        t.daemon = True

        t.start()

    def run(self):
        '''
        Run the timer and notify waiting threads after each interval
        :return:
        '''
        while True:
            time.sleep(self._interval)
            with self._cv:
                self._flag ^= 1
                self._cv.notify_all()

    def wait_for_tick(self):
        '''
        Wait for the next tick of the timer
        :return:
        '''
        with self._cv:
            last_flag = self._flag
            while last_flag == self._flag:
                self._cv.wait()

ptimer = PeriodicTimer(5)
ptimer.start()

def countdown(nticks):
    while nticks > 0:
        ptimer.wait_for_tick()
        print('T-minus', nticks)
        nticks -= 1

def countup(last):
    n = 0
    while n < last:
        ptimer.wait_for_tick()
        print('Counting', n)
        n += 1

threading.Thread(target=countdown, args=(10, )).start()
threading.Thread(target=countup, args=(5, )).start()

"""
event 对象的一个重要特点是当它被设置为真时会唤醒所有等待它的线程。如果你
只想唤醒单个线程，最好是使用信号量或者 Condition 对象来替代。考虑一下这段使
用信号量实现的代码
"""
def worker(n, sema):
    sema.acquire()
    print('Working', n)

sema = threading.Semaphore(0)
nworkers = 10
for n in range(nworkers):
    t = threading.Thread(target=worker, args=(n, sema,))
    t.start()

"""
运行上边的代码将会启动一个线程池，但是并没有什么事情发生。这是因为所有的
线程都在等待获取信号量。每次信号量被释放，只有一个线程会被唤醒并执行，示例
如下：
"""
sema.release()
# Working 0
sema.release()
# Working 1

"""
编写涉及到大量的线程间同步问题的代码会让你痛不欲生。比较合适的方式是使用
队列来进行线程间通信或者每个把线程当作一个 Actor，利用 Actor 模型来控制并发。
下一节将会介绍到队列，而 Actor 模型将在 12.10 节介绍。
"""
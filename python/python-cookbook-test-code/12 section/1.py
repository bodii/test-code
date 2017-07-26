#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十二章：并发编程
Description: 对于并发编程, Python 有多种长期支持的方法, 包括多线程, 调用子进程, 以及各种
各样的关于生成器函数的技巧. 这一章将会给出并发编程各种方面的技巧, 包括通用的
多线程技术以及并行计算的实现方法.
像经验丰富的程序员所知道的那样, 大家担心并发的程序有潜在的危险. 因此, 本章
的主要目标之一是给出更加可信赖和易调试的代码.

Title:            启动与停止线程
Issue: 你要为需要并发执行的代码创建/销毁线程
Answer: threading 库可以在单独的线程中执行任何的在 Python 中可以调用的对象。你可
以创建一个 Thread 对象并将你要执行的对象以 target 参数的形式提供给该对象。
"""
import time
def countdown(n):
    while n > 0:
        print('T-minus', n)
        n -= 1
        time.sleep(5)

from threading import Thread
t = Thread(target=countdown, args=(10, ))
t.start()
'''
T-minus 10
T-minus 9
T-minus 8
T-minus 7
T-minus 6
T-minus 5
T-minus 4
T-minus 3
T-minus 2
T-minus 1
'''

"""
当你创建好一个线程对象后，该对象并不会立即执行，除非你调用它的 start()
方法（当你调用 start() 方法时，它会调用你传递进来的函数，并把你传递进来的参
数传递给该函数）。 Python 中的线程会在一个单独的系统级线程中执行（比如说一个
POSIX 线程或者一个 Windows 线程），这些线程将由操作系统来全权管理。线程一旦
启动，将独立执行直到目标函数返回。你可以查询一个线程对象的状态，看它是否还
在执行：
"""

if t.is_alive():
    print('Still running')
else:
    print('Completed')

"""你也可以将一个线程加入到当前线程，并等待它终止："""
t.join()
"""
Python 解释器在所有线程都终止后才继续执行代码剩余的部分。对于需要长时间
运行的线程或者需要一直运行的后台任务，你应当考虑使用后台线程。
"""
t = Thread(target=countdown, args=(10, ), daemon=True)
t.start()

"""
后台线程无法等待，不过，这些线程会在主线程终止时自动销毁。除了如上所示的
两个操作，并没有太多可以对线程做的事情。你无法结束一个线程，无法给它发送信
号，无法调整它的调度，也无法执行其他高级操作。如果需要这些特性，你需要自己
添加。比如说，如果你需要终止线程，那么这个线程必须通过编程在某个特定点轮询
来退出。你可以像下边这样把线程放入一个类中
"""
class CountdownTask:
    def __init__(self):
        self._running = True

    def terminate(self):
        self._running = False

    def run(self, n):
        while self._running and n > 0:
            print('T-minus', n)
            n -= 1
            time.sleep(5)

c = CountdownTask()
t = Thread(target=c.run, args=(10, ))
t.start()
c.terminate()
t.join()

"""
如果线程执行一些像 I/O 这样的阻塞操作，那么通过轮询来终止线程将使得线程
之间的协调变得非常棘手。比如，如果一个线程一直阻塞在一个 I/O 操作上，它就永
远无法返回，也就无法检查自己是否已经被结束了。要正确处理这些问题，你需要利
用超时循环来小心操作线程。例子如下
"""
import socket
class IOTask:
    def terminate(self):
        self._running = False

    def run(self, sock):
        sock.settimeout(5)
        while self._running:
            try:
                data = sock.recv(8192)
                break
            except socket.timeout:
                continue
        return

"""
由于全局解释锁（GIL）的原因， Python 的线程被限制到同一时刻只允许一个线程
执行这样一个执行模型。所以， Python 的线程更适用于处理 I/O 和其他需要并发执行
的阻塞操作（比如等待 I/O、等待从数据库获取数据等等），而不是需要多处理器并行
的计算密集型任务。
有时你会看到下边这种通过继承 Thread 类来实现的线程
"""
from threading import Thread
class CountdownThread(Thread):
    def __init__(self, n):
        super().__init__()
        self.n = 0

    def run(self):
        while self.n > 0:
            print('T-minus', self.n)
            self.n -=1
            time.sleep(5)

c = CountdownThread(5)
c.start()

"""
尽管这样也可以工作，但这使得你的代码依赖于 threading 库，所以你的这些代
码只能在线程上下文中使用。上文所写的那些代码、函数都是与 threading 库无关的，
这样就使得这些代码可以被用在其他的上下文中，可能与线程有关，也可能与线程无
关。比如，你可以通过 multiprocessing 模块在一个单独的进程中执行你的代码：
"""
import multiprocessing
c = CountdownTask(5)
p = multiprocessing.Process(target = c.run)
p.start()
"""
再次重申，这段代码仅适用于 CountdownTask 类是以独立于实际的并发手段（多
线程、多进程等等）实现的情况。
"""
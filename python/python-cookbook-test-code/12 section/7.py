#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十二章：并发编程
Description: 对于并发编程, Python 有多种长期支持的方法, 包括多线程, 调用子进程, 以及各种
各样的关于生成器函数的技巧. 这一章将会给出并发编程各种方面的技巧, 包括通用的
多线程技术以及并行计算的实现方法.
像经验丰富的程序员所知道的那样, 大家担心并发的程序有潜在的危险. 因此, 本章
的主要目标之一是给出更加可信赖和易调试的代码.

Title:            创建一个线程池
Issue: 你创建一个工作者线程池，用来相应客户端请求或执行其他的工作。
Answer: concurrent.futures 函数库有一个 ThreadPoolExecutor 类可以被用来完成这个
任务。
"""

"""
下面是一个简单的 TCP 服务器，使用了一个线程池来响应客户端：
"""
from socket import AF_INET, SOCK_STREAM, socket
from concurrent.futures import ThreadPoolExecutor

def echo_client(sock, client_addr):
    '''
    Handle a client connection
    :param sock:
    :param client_addr:
    :return:
    '''
    print('Got connection from', client_addr)
    while True:
        msg = sock.recv(65536)
        if not msg:
            break
        sock.sendall(msg)
    print('Client closed connection')
    sock.close()

def echo_server(addr):
    pool = ThreadPoolExecutor(128)
    sock = socket(AF_INET, SOCK_STREAM)
    sock.bind(addr)
    sock.listen(5)
    while True:
        client_sock, client_addr = sock.accept()
        pool.submit(echo_client, client_sock, client_addr)

echo_server(('', 15000))

"""
如果你想手动创建你自己的线程池，通常可以使用一个 Queue 来轻松实现。下面
是一个稍微不同但是手动实现的例子：
"""

from threading import Thread
from queue import Queue

def echo_client(q):
    sock, client_addr = q.get()
    print('Got connection from', client_addr)
    while True:
        msg = sock.recv(65536)
        if not msg:
            break
        sock.sendall(msg)
    print('Client closed connection')
    sock.close()

def echo_server(addr, nworkers):
    q = Queue
    for n in range(nworkers):
        t = Thread(target=echo_client, args=(q, ))
        t.daemon = True
        t.start()

    sock = socket(AF_INET, SOCK_STREAM)
    sock.bind(addr)
    sock.listen(5)
    while True:
        client_sock, client_addr = sock.accept()
        q.put((client_sock, client_addr))

echo_server(('', 15000), 128)

"""
使用 ThreadPoolExecutor 相对于手动实现的一个好处在于它使得任务提交者更方
便的从被调用函数中获取返回值。
"""

from concurrent.futures import ThreadPoolExecutor
import urllib.request
def fetch_url(url):
    u = urllib.request.urlopen(url)
    data = u.read()
    return data

pool = ThreadPoolExecutor(10)
a = pool.submit(fetch_url, 'http://www.python.org')
b = pool.submit(fetch_url, 'http://www.pypy.org')

x = a.result()
y = b.result()

"""
例子中返回的 handle 对象会帮你处理所有的阻塞与协作，然后从工作线程中返回
数据给你。特别的， a.result() 操作会阻塞进程直到对应的函数执行完成并返回一个
结果。
通常来讲，你应该避免编写线程数量可以无限制增长的程序。
例如，看看下面这个服务器：
"""
from threading import Thread
from socket import socket, AF_INET, SOCK_STREAM

def echo_client(sock, client_addr):
    print('Got connection from', client_addr)
    while True:
        msg = sock.recv(65536)
        if not msg:
            break
        sock.sendall(msg)
    print('Client closed connection')
    sock.close()

def echo_server(addr, nworkers):
    sock = socket(AF_INET, SOCK_STREAM)
    sock.bind(addr)
    sock.listen(5)
    while True:
         client_sock, client_addr = sock.accept()
         t = Thread(target=echo_client, args=(client_sock,client_addr))
         t.daemon = True
         t.start()

echo_server(('', 15000))

"""
尽管这个也可以工作，但是它不能抵御有人试图通过创建大量线程让你服务器资源
枯竭而崩溃的攻击行为。通过使用预先初始化的线程池，你可以设置同时运行线程的
上限数量。
你可能会关心创建大量线程会有什么后果。现代操作系统可以很轻松的创建几千个
线程的线程池。甚至，同时几千个线程等待工作并不会对其他代码产生性能影响。当
然了，如果所有线程同时被唤醒并立即在 CPU 上执行，那就不同了——特别是有了全
局解释器锁 GIL。通常，你应该只在 I/O 处理相关代码中使用线程池。
创建大的线程池的一个可能需要关注的问题是内存的使用。例如，如果你在 OS X
系统上面创建 2000 个线程，系统显示 Python 进程使用了超过 9GB 的虚拟内存。不
过，这个计算通常是有误差的。当创建一个线程时，操作系统会预留一个虚拟内存区
域来放置线程的执行栈（通常是 8MB 大小）。但是这个内存只有一小片段被实际映射
到真实内存中。因此， Python 进程使用到的真实内存其实很小（比如，对于 2000 个
线程来讲，只使用到了 70MB 的真实内存，而不是 9GB）。如果你担心虚拟内存大小，
可以使用 threading.stack size() 函数来降低它。
"""
import threading
threading.stack_size(65536)

"""
如果你加上这条语句并再次运行前面的创建 2000 个线程试验，你会发现 Python
进程只使用到了大概 210MB 的虚拟内存，而真实内存使用量没有变。注意线程栈大小
必须至少为 32768 字节，通常是系统内存页大小（4096、 8192 等）的整数倍。
"""

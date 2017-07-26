#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十一章：网络与 WEB 编程
Desc: 本章是关于在网络应用和分布式应用中使用的各种主题。主题划分为使用 Python
编写客户端程序来访问已有的服务，以及使用 Python 实现网络服务端程序。也给出了
一些常见的技术，用于编写涉及协同或通信的的代码。

Title;             在不同的 Python 解释器之间交互
Issue:你在不同的机器上面运行着多个 Python 解释器实例，并希望能够在这些解释器之
间通过消息来交换数据。
Answer:通过使用 multiprocessing.connection 模块可以很容易的实现解释器之间的通
信。
"""
from multiprocessing.connection import Listener
import traceback

def echo_client(conn):
    try:
        while True:
            msg = conn.recv()
            conn.send(msg)
    except EOFError:
        print('Connection Closed')

def echo_sever(address, authkey):
    serv = Listener(address, authkey = authkey)
    while True:
        try:
            client = serv.accept()
            echo_client(client)

        except Exception:
            traceback.print_exc()

echo_sever(('', 25000), authkey=b'peekaboo')

'''
from multiprocessing.connection import Client

c = Client(('localhost', 25000), authkey=b'peekaboo')
c.send('hello')
print(c.recv())
# hello
c.send(42)
print(c.recv())
# 42
c.send([1, 2, 3, 4, 5])
print(c.recv())
# [1, 2, 3, 4, 5]
'''

"""
跟底层 socket 不同的是，每个消息会完整保存（每一个通过 send() 发送的对象
能通过 recv() 来完整接受）。另外，所有对象会通过 pickle 序列化。因此，任何兼容
pickle 的对象都能在此连接上面被发送和接受。

目前有很多用来实现各种消息传输的包和函数库，比如 ZeroMQ、 Celery 等。你还
有另外一种选择就是自己在底层 socket 基础之上来实现一个消息传输层。但是你想要
简单一点的方案，那么这时候 multiprocessing.connection 就派上用场了。仅仅使
用一些简单的语句即可实现多个解释器之间的消息通信。
如果你的解释器运行在同一台机器上面，那么你可以使用另外的通信机制，比如
Unix 域套接字或者是 Windows 命名管道。要想使用 UNIX 域套接字来创建一个连接，
只需简单的将地址改写一个文件名即可：
s = Listener('/tmp/myconn', authkey=b'peekaboo')
要想使用 Windows 命名管道来创建连接，只需像下面这样使用一个文件名：
s = Listener(r'\\.\pipe\myconn', authkey=b'peekaboo')
一个通用准则是，你不要使用 multiprocessing 来实现一个对外的公共服务。
Client() 和 Listener() 中的 authkey 参数用来认证发起连接的终端用户。如果密钥
不对会产生一个异常。此外，该模块最适合用来建立长连接（而不是大量的短连接），
例如，两个解释器之间启动后就开始建立连接并在处理某个问题过程中会一直保持连
接状态。
如果你需要对底层连接做更多的控制，比如需要支持超时、非阻塞 I/O 或其他类
似的特性，你最好使用另外的库或者是在高层 socket 上来实现这些特性。
"""

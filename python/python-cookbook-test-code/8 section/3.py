#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第八章：类与对象
Desc: 本章主要关注点的是和类定义有关的常见编程模型。包括让对象支持常见的
Python 特性、特殊方法的使用、类封装技术、继承、内存管理以及有用的设计模
式。

Title;        让对象支持上下文管理协议
Issue:你想让你的对象支持上下文管理协议 (with 语句)。
Answer: 为了让一个对象兼容 with 语句，你需要实现 enter () 和 exit () 方法。
"""

from socket import socket, AF_INET, SOCK_STREAM

class LazyConnection:
    def __init__(self, address, family=AF_INET, type=SOCK_STREAM):
        self.address = address
        self.family = family
        self.type = type
        self.scok = None

    def __enter__(self):
        if self.sock is not None:
            raise RuntimeError('Already connected')
        self.sock = socket(self.family, self.type)
        self.sock.connect(self.address)
        return self.sock

    def __exit__(self, exc_ty, exc_val, tb):
        self.sock.close()
        self.sock = None

"""
这个类的关键特点在于它表示了一个网络连接，但是初始化的时候并不会做任何事
情 (比如它并没有建立一个连接)。连接的建立和关闭是使用 with 语句自动完成的，
"""
from functools import partial

conn = LazyConnection(('www.python.ort',80))
with conn as s:
    s.send(b'GET /index.html HTTP/1.0\r\n')
    s.send(b'Host:www.python.org\r\n')
    s.send(b'\r\n')
    resp = b''.join(iter(partial(s.recv, 8192), b''))

"""
编写上下文管理器的主要原理是你的代码会放到 with 语句块中执行。当出现 with
语句的时候，对象的 enter () 方法被触发，它返回的值 (如果有的话) 会被赋值给
as 声明的变量。然后，with 语句块里面的代码开始执行。最后， exit () 方法被触
发进行清理工作。
不管 with 代码块中发生什么，上面的控制流都会执行完，就算代码块中发生了异
常也是一样的。事实上， exit () 方法的第三个参数包含了异常类型、异常值和追溯
信息 (如果有的话)。 exit () 方法能自己决定怎样利用这个异常信息，或者忽略它
并返回一个 None 值。如果 exit () 返回 True ，那么异常会被清空，就好像什么都
没发生一样， with 语句后面的程序继续在正常执行。
还有一个细节问题就是 LazyConnection 类是否允许多个 with 语句来嵌套使用连
接。很显然，上面的定义中一次只能允许一个 socket 连接，如果正在使用一个 socket
的时候又重复使用 with 语句，就会产生一个异常了。不过你可以像下面这样修改下上
面的实现来解决这个问题
"""

class LazyConnection:
    def __init__(self, address, family=AF_INET, type=SOCK_STREAM):
        self.address = address
        self.family = family
        self.type = type
        self.connections = []

    def __enter__(self):
        sock = socket(self.family, self.type)
        sock.connect(self.address)
        self.connctions.append(sock)
        return sock

    def __exit__(self, exc_ty, exc_val, tb):
        self.connections.pop().colse()

from functools import partial
conn = LazyConnection(('www.python.org', 80))
with conn as s1:
    pass
    with conn as s2:
        pass

"""
在第二个版本中，LazyConnection 类可以被看做是某个连接工厂。在内部，一个
列表被用来构造一个栈。每次 enter () 方法执行的时候，它复制创建一个新的连接
并将其加入到栈里面。 exit () 方法简单的从栈中弹出最后一个连接并关闭它。这
里稍微有点难理解，不过它能允许嵌套使用 with 语句创建多个连接，就如上面演示的
那样。
在需要管理一些资源比如文件、网络连接和锁的编程环境中，使用上下文管理器是
很普遍的。这些资源的一个主要特征是它们必须被手动的关闭或释放来确保程序的正
确运行。例如，如果你请求了一个锁，那么你必须确保之后释放了它，否则就可能产
生死锁。通过实现 enter () 和 exit () 方法并使用 with 语句可以很容易的避免
这些问题，因为 exit () 方法可以让你无需担心这些了。
在 contextmanager 模块中有一个标准的上下文管理方案模板，可参考 9.22 小节。
同时在 12.6 小节中还有一个对本节示例程序的线程安全的修改版。
"""

#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十一章：网络与 WEB 编程
Desc: 本章是关于在网络应用和分布式应用中使用的各种主题。主题划分为使用 Python
编写客户端程序来访问已有的服务，以及使用 Python 实现网络服务端程序。也给出了
一些常见的技术，用于编写涉及协同或通信的的代码。

Title;             进程间传递 Socket 文件描述符
Issue:你有多个 Python 解释器进程在同时运行，你想将某个打开的文件描述符从一个解
释器传递给另外一个。比如，假设有个服务器进程相应连接请求，但是实际的相应逻
辑是在另一个解释器中执行的。
Answer:为了在多个进程中传递文件描述符，你首先需要将它们连接到一起。在 Unix 机器
上，你可能需要使用 Unix 域套接字，而在 windows 上面你需要使用命名管道。不过
你无需真的需要去操作这些底层，通常使用 multiprocessing 模块来创建这样的连接
会更容易一些。
一 旦 一 个 连 接 被 创 建， 你 可 以 使 用 multiprocessing.reduction 中 的
send handle() 和 recv handle() 函数在不同的处理器直接传递文件描述符。
"""

import multiprocessing
from multiprocessing.reduction import recv_handle, send_handle
import socket

def worker(in_p, out_p):
    out_p.close()
    while True:
        fd = recv_handle(in_p)
        print('GHILD: GOT FD', fd)
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM, fileno=fd) as s:
            while True:
                msg = s.recv(1024)
                if not msg:
                    break
                print('GHILD:RECV {!r}'.format(msg))
                s.send(msg)

def server(address, in_p, out_p, worker_pid):
    in_p.close()
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, True)
    s.bind(address)
    s.listen(1)
    while True:
        client, addr = s.accept()
        print('SERVER: Got connection from', addr)
        send_handle(out_p, client.fileno(), worker_pid)
        client.close()

if __name__ == '__main__':
    c1, c2 = multiprocessing.Pipe()
    worker_p = multiprocessing.Process(target=worker, args=(c1, c2))
    worker_p.start()

    server_p = multiprocessing.Process(target=server,
                                       args=(('', 15000), c1, c2, worker_p.pid))
    server_p.start()
    c1.close()
    c2.close()

"""
在这个例子中，两个进程被创建并通过一个 multiprocessing 管道连接起来。服
务器进程打开一个 socket 并等待客户端连接请求。工作进程仅仅使用 recv handle()
在管道上面等待接收一个文件描述符。当服务器接收到一个连接，它将产生的 socket
文件描述符通过 send handle() 传递给工作进程。工作进程接收到 socket 后向客户端
回应数据，然后此次连接关闭。
如果你使用 Telnet 或类似工具连接到服务器，下面是一个演示例子：
bash % python3 passfd.py SERVER: Got connection from (‘127.0.0.1’, 55543)
CHILD: GOT FD 7 CHILD: RECV b’Hellorn’ CHILD: RECV b’Worldrn’
此例最重要的部分是服务器接收到的客户端 socket 实际上被另外一个不同的进程
处理。服务器仅仅只是将其转手并关闭此连接，然后等待下一个连接。

对于大部分程序员来讲在不同进程之间传递文件描述符好像没什么必要。但是，有
时候它是构建一个可扩展系统的很有用的工具。例如，在一个多核机器上面，你可以
有多个 Python 解释器实例，将文件描述符传递给其它解释器来实现负载均衡。
send handle() 和 recv handle() 函数只能够用于 multiprocessing 连接。使用
它们来代替管道的使用（参考 11.7 节），只要你使用的是 Unix 域套接字或 Windows
管道。例如，你可以让服务器和工作者各自以单独的程序来启动。
"""
# servermp.py
from multiprocessing.connection import Listener
from multiprocessing.reduction import send_handle
import socket

def server(work_address, port):
    work_serv = Listener(work_address, authkey=b'peekaboo')
    worker = work_serv.accept()
    worker_pid = worker.recv()

    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, True)
    s.bind(('', port))
    s.listen(1)
    while True:
        client, addr = s.accept()
        print('SERVER: Got connection from', addr)

        send_handle(worker, client.fileno(), worker_pid)
        client.close()

if __name__ == '__main__':
    import sys
    if len(sys.argv) != 3:
        print('Usage: server.py server_address port', file=sys.stderr)
        raise SystemExit(1)
    server(sys.argv[1], int(sys.argv[2]))

"""
运行这个服务器，只需要执行 python3 servermp.py /tmp/servconn 15000 ，下面是
相应的工作者代码：
"""
# workermp.py

from multiprocessing.connection import Client
from multiprocessing.reduction import recv_handle
import os
from socket import socket, AF_INET, SOCK_STREAM

def worker(server_address):
    serv = Client(server_address, authkey=b'peekaboo')
    serv.send(os.getpid())
    while True:
        fd = recv_handle(serv)
        print('WORKEY: GOT FD', fd)
        with socket(AF_INET, SOCK_STREAM, fileno=fd) as client:
            while True:
                msg = client.recv(1024)
                if not msg:
                    break
                print('WORKEY: RECV {!r}'.format(msg))
                client.send(msg)

if __name__ == '__main__':
    import sys
    if len(sys.argv) != 2:
        print('Usage: worker.py server_address', file=sys.stderr)
        raise SystemExit(1)
    worker(sys.argv[1])

"""
要运行工作者，执行执行命令 python3 workermp.py /tmp/servconn . 效果跟使用
Pipe() 例子是完全一样的。文件描述符的传递会涉及到 UNIX 域套接字的创建和套接
字的 sendmsg() 方法。不过这种技术并不常见，下面是使用套接字来传递描述符的另
外一种实现：
"""
# server.py
import socket
import struct

def send_fd(sock, fd):
    sock.sendmsg([b'x'],
                 [(socket.SOL_SOCKET, socket.SCM_RIGHTS, struct.pack('i', fd))])
    ack = sock.recv(2)
    assert sck == b'OK'

def server(work_address, port):
    work_serv = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    work_serv.bind(work_address)
    work_serv.listen(1)
    worker, addr = work_serv.accept()

    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, True)
    s.bind(('', port))
    s.listen(1)
    while True:
        client, addr = s.accept()
        print('SERVER: Got connection from', addr)
        send_fd(worker, client.fileno())
        client.close()

if __name__ == '__main__':
    import sys
    if len(sys.argv) != 3:
        print('Usage: server.py server_address port', file=sys.stderr)
        raise SystemExit(1)
    server(sys.argv[1], int(sys.argv[2]))

"""
下面是使用套接字的工作者实现
"""
# worker.py
import socket
import struct

def recv_fd(sock):
    msg. ancdata, flags, addr = sock.recvmsg(1, socket.CMSG_LEN(struct.calcsize('i')))
    cmsg_level, cmsg_type, cmsg_data = ancdata[0]
    assert cmsg_level == socket.SOL_SOCKET and cmsg_type == socket.SCM_RIGHTS
    sock.sendall(b'OK')
    return struct.unpack('i', cmsg_data)[0]

def worker(server_address):
    serv = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    serv.connect(server_address)
    while True:
        fd = recv_fd(serv)
        print('WORKER: GOT FD', fd)
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM, fileno=fd) as client:
            while True:
                msg = client.recv(1024)
                if not msg:
                    break
                print('WORKER: RECV {!r}'.format(msg))
                client.send(msg)

if __name__ == '__main__':
    import sys
    if len(sys.argv) != 2:
        print('Usage: worker.py server_address', file=sys.stderr)
        raise SystemExit(1)
    worker(sys.argv[1])

"""
如果你想在你的程序中传递文件描述符，建议你参阅其他一些更加高级的文
档，比如 Unix Network Programming by W. Richard Stevens (Prentice Hall,
1990) . 在 Windows 上 传 递 文 件 描 述 符 跟 Unix 是 不 一 样 的， 建 议 你 研 究 下
multiprocessing.reduction 中的源代码看看其工作原理。
"""

#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十一章：网络与 WEB 编程
Desc: 本章是关于在网络应用和分布式应用中使用的各种主题。主题划分为使用 Python
编写客户端程序来访问已有的服务，以及使用 Python 实现网络服务端程序。也给出了
一些常见的技术，用于编写涉及协同或通信的的代码。

Title;               创建 TCP 服务器
Issue:你想实现一个服务器，通过 TCP 协议和客户端通信。
Answer: 创建一个 TCP 服务器的一个简单方法是使用 socketserver 库。
"""
from socketserver import BaseRequestHandler, TCPServer

class EchoHandler(BaseRequestHandler):
    def handle(self):
        print('Got conection from', self.client_address)
        while True:
            msg = self.request.recv(8192)
            if not msg:
                break
            self.request.send(msg)

# if __name__ == '__main__':
#     serv = TCPServer(('',20000),EchoHandler)
#     serv.serve_forever()

"""
在这段代码中，你定义了一个特殊的处理类，实现了一个 handle() 方法，用来为
客户端连接服务。 request 属性是客户端 socket， client address 有客户端地址。为
了测试这个服务器，运行它并打开另外一个 Python 进程连接这个服务器：
"""
# from socket import socket, AF_INET, SOCK_STREAM
# s = socket(AF_INET, SOCK_STREAM)
# s.connect(('localhost',20000))
# s.send(b'Hello')
# s.recv(8192)

"""
很 多 时 候， 可 以 很 容 易 的 定 义 一 个 不 同 的 处 理 器。 下 面 是 一 个 使 用
StreamRequestHandler 基类将一个类文件接口放置在底层 socket 上的例子
"""
from socketserver import StreamRequestHandler, TCPServer
class EchoHandler(StreamRequestHandler):
    def handle(self):
        print('Got connection from', self.client_address)
        for line in self.rfile:
            self.wfile.write(line)

# if __name__ == '__main__':
#     serv = TCPServer(('', 20000), EchoHandler)
#     serv.serve_forever()

"""
socketserver 可以让我们很容易的创建简单的 TCP 服务器。但是，你需要注意的
是，默认情况下这种服务器是单线程的，一次只能为一个客户端连接服务。如果你想
处理多个客户端，可以初始化一个 ForkingTCPServer 或者是 ThreadingTCPServer 对
象。
"""

from socketserver import ThreadingTCPServer

# if __name__ == '__main__':
#     serv = ThreadingTCPServer(('', 20000), EchoHandler)
#     serv.serve_forever()

"""
使用 fork 或线程服务器有个潜在问题就是它们会为每个客户端连接创建一个新的
进程或线程。由于客户端连接数是没有限制的，因此一个恶意的黑客可以同时发送大
量的连接让你的服务器奔溃。
如果你担心这个问题，你可以创建一个预先分配大小的工作线程池或进程池。你先
创建一个普通的非线程服务器，然后在一个线程池中使用 serve forever() 方法来启
动它们。
"""
# if __name__ == '__main__':
#     from threading import Thread
#     NWORKERS = 16
#     serv = TCPServer(('', 20000), EchoHandler)
#     for n in range(NWORKERS):
#         t = Thread(target=serv.serve_forever)
#         t.daemon = True
#         t.start()
#     serv.serve_forever()

"""
一 般 来 讲， 一 个 TCPServer 在 实 例 化 的 时 候 会 绑 定 并 激 活 相 应 的 socket
。不过，有时候你想通过设置某些选项去调整底下的 socket‘ ，可以设置参数
bind and activate=False 。如下：
"""

# if __name__ == '__main__':
#     serv = TCPServer(('', 20000), EchoHandler, bind_and_activate=False)
#     import socket
#     serv.socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, True)
#     serv.server_bind()
#     serv.server_activate()
#     serv.serve_forever()

"""
上面的 socket 选项是一个非常普遍的配置项，它允许服务器重新绑定一个之前使
用过的端口号。由于要被经常使用到，它被放置到类变量中，可以直接在 TCPServer
上面设置。在实例化服务器的时候去设置它的值，
"""
# if __name__ == '__main__':
#     TCPServer.allow_reuse_address = True
#     serv = TCPServer(('', 20000), EchoHandler)
#     serv.serve_forever()

"""
在上面示例中，我们演示了两种不同的处理器基类（BaseRequestHandler 和
StreamRequestHandler ）。 StreamRequestHandler 更加灵活点，能通过设置其他
的类变量来支持一些新的特性。
"""
import socket
class EchoHandler(StreamRequestHandler):
    timeout = 5
    rbufsize = -1
    wbufsize = 0
    disable_nagle_algorithm = False
    def handle(self):
        print('Got connect from', self.client_address)
        try:
            for line in self.rfile:
                self.wfile.write(line)
        except socket.timeout:
            print('Timed out!')

"""
最后，还需要注意的是巨大部分 Python 的高层网络模块（比如 HTTP、 XMLRPC 等）都是建立在 socketserver 功能之上。也就是说，直接使用 socket 库来实现
服务器也并不是很难。下面是一个使用 socket 直接编程实现的一个服务器简单例子
"""
from socket import socket, AF_INET, SOCK_STREAM

def echo_handler(address, client_sock):
    print('Got connection from {}'. format(address))
    while True:
        msg = client_sock.recv(8192)
        if not msg:
            break
        client_sock.sendall(msg)
    client_sock.close()

def echo_server(address, backlog=5):
    sock = socket(AF_INET, SOCK_STREAM)
    sock.bind(address)
    sock.listen(backlog)
    while True:
        client_sock, client_addr = sock.accept()
        echo_handler(client_addr, client_sock)

if __name__ == '__main__':
    echo_server(('', 20000))
#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十一章：网络与 WEB 编程
Desc: 本章是关于在网络应用和分布式应用中使用的各种主题。主题划分为使用 Python
编写客户端程序来访问已有的服务，以及使用 Python 实现网络服务端程序。也给出了
一些常见的技术，用于编写涉及协同或通信的的代码。

Title;             实现远程方法调用
Issue:你想在一个消息传输层如 sockets 、 multiprocessing connections 或 ZeroMQ 的
基础之上实现一个简单的远程过程调用（RPC）。
Answer:将函数请求、参数和返回值使用 pickle 编码后，在不同的解释器直接传送 pickle 字
节字符串，可以很容易的实现 RPC。
"""
import pickle
class RPCHandler:
    def __init__(self):
        self._functions = {}

    def register_function(self, func):
        self._functions[func.__name__] = func

    def handle_connection(self, connection):
        try:
            while True:
                func_name, args, kwargs = pickle.loads(connection.recv())
                try:
                    r = self._functions[func_name](*args, **kwargs)
                    connection.send(pickle.dumps(r))
                except Exception as e:
                    connection.send(pickle.dumps(e))

        except EOFError:
            pass

"""
要使用这个处理器，你需要将它加入到一个消息服务器中。你有很多种选择，但是
使用 multiprocessing 库是最简单的。下面是一个 RPC 服务器例子：
"""
from multiprocessing.connection import Listener
from threading import Thread

def rpc_server(handler, address, authkey):
    sock = Listener(address, authkey = authkey)
    while True:
        client = sock.accept()
        t = Thread(target=handler.handle_connection, args=(client,))
        t.daemon = True
        t. start()

def add(x, y):
    return x + y

def sub(x, y):
    return x - y

# handler = RPCHandler()
# handler.register_function(add)
# handler.register_function(sub)
#
# rpc_server(handler, ('localhost', 17000), authkey=b'peekaboo')

"""
为了从一个远程客户端访问服务器，你需要创建一个对应的用来传送请求的 RPC
代理类。
"""

import pickle

class RPCProxy:
    def __init__(self, connection):
        self._connection = connection

    def __getattr__(self, name):
        def do_rpc(*args, **kwargs):
            self._connection.send(pickle.dumps((name, args, kwargs)))
            result = pickle.loads(self._connection.recv())
            if isinstance(result, Exception):
                raise result
            return result
        return do_rpc

"""
要使用这个代理类，你需要将其包装到一个服务器的连接上面，例如：
"""
from multiprocessing.connection import Client
c = Client(('localhost', 17000), authkey=b'peekaboo')
proxy = RPCProxy(c)
print(proxy.add(2, 3))
# 5
print(proxy.sub([1, 2], 4))
# TypeError: unsupported operand type(s) for -: 'list' and 'int'

"""
要注意的是很多消息层（比如 multiprocessing ）已经使用 pickle 序列化了数据。
如果是这样的话，对 pickle.dumps() 和 pickle.loads() 的调用要去掉。

RPCHandler 和 RPCProxy 的基本思路是很比较简单的。如果一个客户端想要调用
一个远程函数，比如 foo(1, 2, z=3) , 代理类创建一个包含了函数名和参数的元组
('foo', (1, 2), {'z': 3}) 。这个元组被 pickle 序列化后通过网络连接发生出去。
这一步在 RPCProxy 的 getattr () 方法返回的 do rpc() 闭包中完成。服务器接收
后通过 pickle 反序列化消息，查找函数名看看是否已经注册过，然后执行相应的函
数。执行结果 (或异常) 被 pickle 序列化后返回发送给客户端。我们的实例需要依赖
multiprocessing 进行通信。不过，这种方式可以适用于其他任何消息系统。例如，
如果你想在 ZeroMQ 之上实习 RPC，仅仅只需要将连接对象换成合适的 ZeroMQ 的
socket 对象即可。
由于底层需要依赖 pickle，那么安全问题就需要考虑了（因为一个聪明的黑客可以
创建特定的消息，能够让任意函数通过 pickle 反序列化后被执行）。因此你永远不要
允许来自不信任或未认证的客户端的 RPC。特别是你绝对不要允许来自 Internet 的任
意机器的访问，这种只能在内部被使用，位于防火墙后面并且不要对外暴露。
作为 pickle 的替代，你也许可以考虑使用 JSON、 XML 或一些其他的编码格式
来序列化消息。例如，本机实例可以很容易的改写成 JSON 编码方案。还需要将
pickle.loads() 和 pickle.dumps() 替换成 json.loads() 和 json.dumps() 即可：
"""

import json

# class RPCHandler:
#     def __init__(self):
#         self._functions = {}
#
#     def register_function(self, func):
#         self._functions[func.__name__] = func
#
#     def handle_connection(self, connection):
#         try:
#             while True;
#             func_name, args, kwargs = json.loads(connection.recv())
#             try:
#                 r = self._functions[func_name](*args, **kwargs)
#                 connection.send(json.dumps(r))
#             except Exception as e:
#                 connection.send(json.dumps(e))
#
#         except EOFError:
#             pass
#
# import json
# class RPCProxy:
#     def __init__(self, connection):
#         self._connection = connection
#
#     def __getattr__(self, name):
#         def do_rpc(*args, **kwargs):
#             self._connection.send(json.dumps((name, args, kwargs)))
#             result = json.loads(self._connection.recv())
#             return result
#         return do_rpc

"""
实现 RPC 的一个比较复杂的问题是如何去处理异常。至少，当方法产生异常时服
务器不应该奔溃。因此，返回给客户端的异常所代表的含义就要好好设计了。如果你
使用 pickle，异常对象实例在客户端能被反序列化并抛出。如果你使用其他的协议，那
得想想另外的方法了。不过至少，你应该在响应中返回异常字符串。我们在 JSON 的
例子中就是使用的这种方式。
对 于 其 他 的 RPC 实 现 例 子， 我 推 荐 你 看 看 在 XML-RPC 中 使 用 的
SimpleXMLRPCServer 和 ServerProxy 的实现，也就是 11.6 小节中的内容。
"""

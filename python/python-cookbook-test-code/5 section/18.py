#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第五章：文件与 IO
Desc: 所有程序都要处理输入和输出。这一章将涵盖处理不同类型的文件，包括文本和二
进制文件，文件编码和其他相关的内容。对文件名和目录的操作也会涉及到。

Title;           将文件描述符包装成文件对象
Issue:你有一个对应于操作系统上一个已打开的 I/O 通道 (比如文件、管道、套接字等)
的整型文件描述符，你想将它包装成一个更高层的 Python 文件对象。
Answer: 一个文件描述符和一个打开的普通文件是不一样的。文件描述符仅仅是一个由操作
系统指定的整数，用来指代某个系统的 I/O 通道。如果你碰巧有这么一个文件描述符，
你可以通过使用 open() 函数来将其包装为一个 Python 的文件对象。你仅仅只需要使
用这个整数值的文件描述符作为第一个参数来代替文件名即可。
"""

import os
fd = os.open('../test file/somefile.txt', os.O_WRONLY | os.O_CREAT)
f = open(fd, 'wt')
f.write('hello world\n')
f.close()

'''
当高层的文件对象被关闭或者破坏的时候，底层的文件描述符也会被关闭。如果这
个并不是你想要的结果，你可以给 open() 函数传递一个可选的 colsefd=False 。
'''
f = open(fd, 'wt', closefd=False)

'''
在 Unix 系统中，这种包装文件描述符的技术可以很方便的将一个类文件接口作用
于一个以不同方式打开的 I/O 通道上，如管道、套接字等。举例来讲，下面是一个操
作管道的例子：
'''

from socket import socket, AF_INET, SOCK_STREAM

def echo_client(client_sock, addr):
    print('Got connection from', addr)

    client_in = open(client_sock.fileno(), 'rt', encoding='latin-1',closefd=False)
    client_out = open(client_sock.fileno(), 'wt', encoding='latin-1',closefd=False)
    for line in client_in:
        client_out.write(line)
        client_out.flush()

    client_sock.clost()

def echo_server(address):
    sock = socket(AF_INET, SOCK_STREAM)
    sock.bind(address)
    sock.listen(1)
    while True:
        client, addr = sock.accept()
        echo_client(client, addr)

'''
需要重点强调的一点是，上面的例子仅仅是为了演示内置的 open() 函数的一个特
性，并且也只适用于基于 Unix 的系统。如果你想将一个类文件接口作用在一个套接字
并希望你的代码可以跨平台，请使用套接字对象的 makefile() 方法。但是如果不考虑
可移植性的话，那上面的解决方案会比使用 makefile() 性能更好一点。
你也可以使用这种技术来构造一个别名，允许以不同于第一次打开文件的方式使用
它。例如，下面演示如何创建一个文件对象，它允许你输出二进制数据到标准输出 (通
常以文本模式打开)：
'''

import sys
bstdout = open(sys.stdout.fileno(), 'wb', closefd=False)
bstdout.write(b'Hello World\n')
bstdout.flush()

'''
尽管可以将一个已存在的文件描述符包装成一个正常的文件对象，但是要注意的是
并不是所有的文件模式都被支持，并且某些类型的文件描述符可能会有副作用 (特别是
涉及到错误处理、文件结尾条件等等的时候)。在不同的操作系统上这种行为也是不一
样，特别的，上面的例子都不能在非 Unix 系统上运行。我说了这么多，意思就是让你
充分测试自己的实现代码，确保它能按照期望工作。
'''
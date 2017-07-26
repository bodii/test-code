#!/usr/bin/env python3
# -*- coding:utf-8 -*-

'''
在调用os.fork 之后最常做的一件事情是立即调用os.execl以运行其他的程序。
os.execl是一个用新程序代替正在运行的程序的指令，所以它使得调用程序中止，
并且一个新的程序出现在它的位置（指示：UNIX系统使用fork和exec方法运行所有的
程序）
'''

import os

pid = os.fork()

# fork and exec together
print('second test')
if pid == 0 : # This is the child
    print("this is the child")
    print("I'm going to exec another program now")
    os.execl('/bin/cat', 'cat', '/etc/host.conf')
else:
    print("the child is pid %d" % pid)
os.wait()

'''
os.wait函数通知Python让父进程什么都不做，一直等到子进程返回。了解这个函数的
工作原理很有用的，因为它只在像Linux这样的UNIX和类UNIX的平台上能够正常工作。
Windows也有启动新进程的机制。
为了使启动新程序这样普通的任务变得更容易，Python提供了单独的一族函数，合并
了在类UNIX系统上的os.fork 和os.exec，并且能够在Windows上完成类似的工作。当
只想要启动一个新程序时，可以使用os.spawn函数族。因为这些函数的名称很相似，
所以被称作是一个族，但是每个函数的行为都稍有不同。
在类UNIX的系统上，os.spawn族包括spawnl,spawnle,spawnlp,spawnlpe,spawnv,
spawnve,spawnvp 和spawnvpe。在Windows系统上，spawn族只包括spawnl，spawnle，
spawnv和spawnve。
v表示传递一个列表（实际上v代表vector）作为参数。这允许在不同的实例中运行
差别很大的命令，而不需要对程序进行任何修改。1变量只需要一个简单的参数列表。
使用e时，需要传递一个包括名称和值的字典作为新创建的程序的环境，而不是使用
当前的环境。
p变体使用环境字典中PATH键的值来寻找程序。P变体只能用在类UNIX平台上。这意味
着在Windows中，程序必须有一个完全限定的路径，使其可以用于os.spawn调用，否
则必须自己搜索路径。
'''

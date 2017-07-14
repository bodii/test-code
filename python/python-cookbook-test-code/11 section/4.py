#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十一章：网络与 WEB 编程
Desc: 本章是关于在网络应用和分布式应用中使用的各种主题。主题划分为使用 Python
编写客户端程序来访问已有的服务，以及使用 Python 实现网络服务端程序。也给出了
一些常见的技术，用于编写涉及协同或通信的的代码。

Title;              通过 CIDR 地址生成对应的 IP 地址集
Issue:你有一个 CIDR 网络地址比如“ 123.45.67.89/27”，你想将其转换成它所代表的所
有 IP （比如，“ 123.45.67.64” , “ 123.45.67.65” , …, “ 123.45.67.95” )）
Answer: 可以使用 ipaddress 模块很容易的实现这样的计算。
"""
import ipaddress
net = ipaddress.ip_network('123.45.67.64/27')
print(net)
# 123.45.67.64/27
for a in net:
    print(a)

'''
123.45.67.64
123.45.67.65
123.45.67.66
123.45.67.67
123.45.67.68
123.45.67.69
123.45.67.70
123.45.67.71
123.45.67.72
123.45.67.73
123.45.67.74
123.45.67.75
123.45.67.76
123.45.67.77
123.45.67.78
123.45.67.79
123.45.67.80
123.45.67.81
123.45.67.82
123.45.67.83
123.45.67.84
123.45.67.85
123.45.67.86
123.45.67.87
123.45.67.88
123.45.67.89
123.45.67.90
123.45.67.91
123.45.67.92
123.45.67.93
123.45.67.94
123.45.67.95
'''
net6 = ipaddress.ip_network('12:3456:78:90ab:cd:ef01:23:30/125')
print(net6)
# 12:3456:78:90ab:cd:ef01:23:30/125
for a in net6:
    print(a)

'''
12:3456:78:90ab:cd:ef01:23:30
12:3456:78:90ab:cd:ef01:23:31
12:3456:78:90ab:cd:ef01:23:32
12:3456:78:90ab:cd:ef01:23:33
12:3456:78:90ab:cd:ef01:23:34
12:3456:78:90ab:cd:ef01:23:35
12:3456:78:90ab:cd:ef01:23:36
12:3456:78:90ab:cd:ef01:23:37
'''

"""Network 也允许像数组一样的索引取值"""

print(net.num_addresses)
# 32
print(net[0])
# 123.45.67.64
print(net[1])
# 123.45.67.65
print(net[-1])
# 123.45.67.95
print(net[-2])
# 123.45.67.94

"""
另外，你还可以执行网络成员检查之类的操作：
"""
a = ipaddress.ip_address('123.45.67.69')
print(a in net)
# True
b = ipaddress.ip_address('123.45.67.123')
print(b in net)
# False

"""
一个 IP 地址和网络地址能通过一个 IP 接口来指定，例如
"""
inet = ipaddress.ip_interface('123.45.67.73/27')
print(inet.network)
# 123.45.67.64/27
print(inet.ip)
# 123.45.67.73

"""
ipaddress 模块有很多类可以表示 IP 地址、网络和接口。当你需要操作网络地址
（比如解析、打印、验证等）的时候会很有用。
要注意的是， ipaddress 模块跟其他一些和网络相关的模块比如 socket 库交集很
少。所以，你不能使用 IPv4Address 的实例来代替一个地址字符串，你首先得显式的
使用 str() 转换它。
"""
a = ipaddress.ip_address('127.0.0.1')
from socket import socket, AF_INET, SOCK_STREAM
s = socket(AF_INET, SOCK_STREAM)
#s.connect((a, 8080))
# TypeError: string or unicode text buffer expected, not IPv4Address
s.connect((str(a), 8080))
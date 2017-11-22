#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' __bytes__()方法 '''
"""
bytes(integer): 返回一个不可变的字节对象，这个对象包含了给定数量的0x00值

bytes(string): 这个会将字符串编码为字节。其他的编码和异常处理的参数会定义编码的具体过程

bytes(something): 这个会调用something.__bytes__()创建字节对象。这里不用编码或者错误
处理参数

基本的object对象没有定义__bytes__()。这意味着所有的类在默认情况下都没有提供__bytes__()
方法。 
"""

a = bytes('a', encoding='utf8')
print(a)
a = a.decode('utf8')
print(a)
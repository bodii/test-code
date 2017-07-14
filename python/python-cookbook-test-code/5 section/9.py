#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第五章：文件与 IO
Desc: 所有程序都要处理输入和输出。这一章将涵盖处理不同类型的文件，包括文本和二
进制文件，文件编码和其他相关的内容。对文件名和目录的操作也会涉及到。

Title;         读取二进制数据到可变缓冲区中
Issue:你想直接读取二进制数据到一个可变缓冲区中，而不需要做任何的中间复制操作。
或者你想原地修改数据并将它写回到一个文件中去
Answer: 为了读取数据到一个可变数组中，使用文件对象的 readinto() 方法。
"""
import os.path

def read_into_buffer(filename):
    buf = bytearray(os.path.getsize(filename))
    with open(filename, 'rb') as f:
        f. readinto(buf)
    return buf

'''
下面是一个演示这个函数使用方法的例子：
'''
with open('../test file/sample.bin', 'wb') as f:
    f.write(b'Hello World')

buf = read_into_buffer('../test file/sample.bin')
print(buf)
# bytearray(b'Hello World')
buf[0:5] = b'Hallo'
print(buf)
# bytearray(b'Hallo World')
with open('../test file/newsample.bin', 'wb') as f:
    f.write(buf)

'''
文件对象的 readinto() 方法能被用来为预先分配内存的数组填充数据，甚至包括
由 array 模块或 numpy 库创建的数组。和普通 read() 方法不同的是， readinto() 填
充已存在的缓冲区而不是为新对象重新分配内存再返回它们。因此，你可以使用它来
避免大量的内存分配操作。比如，如果你读取一个由相同大小的记录组成的二进制文
件时，你可以像下面这样写：
'''
record_size = 32
buf = bytearray(record_size)
with open('../test file/somefile', 'rb') as f:
    while True:
        n = f.readinto(buf)
        if n < record_size:
            break

'''
另外有一个有趣特性就是 memoryview ，它可以通过零复制的方式对已存在的缓冲
区执行切片操作，甚至还能修改它的内容。
'''
print(buf)
# bytearray(b'Hello\r\n\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00')
m1 = memoryview(buf)
m2 = m1[-5:]
print(m2)
# <memory at 0x0000022060130B88>
m2[:] = b'WOELD'
print(buf)


'''
使用 f.readinto() 时需要注意的是，你必须检查它的返回值，也就是实际读取的
'''

'''
字节数。
如果字节数小于缓冲区大小，表明数据被截断或者被破坏了 (比如你期望每次读取
指定数量的字节)。
最后，留心观察其他函数库和模块中和 into 相关的函数 (比如 recv into() ，
pack into() 等)。Python 的很多其他部分已经能支持直接的 I/O 或数据访问操作，这
些操作可被用来填充或修改数组和缓冲区内容。
关于解析二进制结构和 memoryviews 使用方法的更高级例子，请参考 6.12 小节。
'''
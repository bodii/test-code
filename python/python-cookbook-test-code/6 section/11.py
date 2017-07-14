#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第六章：数据编码和处理
Desc: 这一章主要讨论使用 Python 处理各种不同方式编码的数据，比如 CSV 文件，
JSON，XML 和二进制包装记录。和数据结构那一章不同的是，这章不会讨论特殊的
算法问题，而是关注于怎样获取和存储这些格式的数据。

Title;      读写二进制数组数据
Issue: 你想读写一个二进制数组的结构化数据到 Python 元组中。
Answer: 可以使用 struct 模块处理二进制数据。
"""
from struct import Struct

def write_records(records, format, f):
    record_struct = Struct(format)
    for r in records:
        f.write(record_struct.pack(*r))


# if __name__ == '__main__':
#     records = [
#         (1, 2.3, 4.5),
#         (6, 7.8, 90),
#         (12, 13.4, 56.7)
#     ]
#     with open('../test file/data.b', 'wb') as f:
#         write_records(records, '<idd', f)

"""
如果你想将整个文件一次性读取到一个字节字符串中，然后在分片解析。
"""
from struct import Struct

def unpack_recorde(format, data):
    record_struct = Struct(format)
    return (record_struct.unpack_from(data)
            for offset in range(0, len(data), record_struct.size))

if __name__ == '__main__':
    with open('../test file/data.b', 'rb') as f:
        data = f.read()
    for rec in unpack_recorde('<idd', data):
        pass

"""
两种情况下的结果都是一个可返回用来创建该文件的原始元组的可迭代对象。
对于需要编码和解码二进制数据的程序而言，通常会使用 struct 模块。为了声明
一个新的结构体，只需要像这样创建一个 Struct 实例即可
"""
record_struct = Struct('<idd')

"""
结构体通常会使用一些结构码值 i, d, f 等 [参考 Python 文档 ]。这些代码分别代表
某个特定的二进制数据类型如 32 位整数，64 位浮点数，32 位浮点数等。第一个字符
< 指定了字节顺序。在这个例子中，它表示” 低位在前”。更改这个字符为 > 表示高位
在前，或者是 ! 表示网络字节顺序。
产生的 Struct 实例有很多属性和方法用来操作相应类型的结构。 size 属性包含
了结构的字节数，这在 I/O 操作时非常有用。 pack() 和 unpack() 方法被用来打包和
解包数据。
"""
record_struct = Struct('<idd')
print(record_struct.size)
# 20
print(record_struct.pack(1, 2.0, 3.0))
# b'\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@'
#print(record_struct.unpack(_))
"""
有时候你还会看到 pack() 和 unpack() 操作以模块级别函数被调用
"""
import struct
print(struct.pack('<idd', 1, 2.0, 3.0))
# b'\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@'
#print(struct.unpack('<idd', _))

"""
这样可以工作，但是感觉没有实例方法那么优雅，特别是在你代码中同样的结构出
现在多个地方的时候。通过创建一个 Struct 实例，格式代码只会指定一次并且所有的
操作被集中处理。这样一来代码维护就变得更加简单了 (因为你只需要改变一处代码即
可)。
读取二进制结构的代码要用到一些非常有趣而优美的编程技巧。在函数　
read records 中，iter() 被用来创建一个返回固定大小数据块的迭代器，参考
5.8 小节。这个迭代器会不断的调用一个用户提供的可调用对象 (比如 lambda:
f.read(record struct.size) )，直到它返回一个特殊的值 (如 b’‘)，这时候迭代停
止。
"""
f = open('../test file/data.b', 'rb')
chunks = iter(lambda: f.read(20),b'')
print(chunks)
# <callable_iterator object at 0x0000026D06CA2EB8>
for chk in chunks:
    print(chk)

'''
b'\x01\x00\x00\x00ffffff\x02@\x00\x00\x00\x00\x00\x00\x12@'
b'\x06\x00\x00\x00333333\x1f@\x00\x00\x00\x00\x00\x80V@'
b'\x0c\x00\x00\x00\xcd\xcc\xcc\xcc\xcc\xcc*@\x9a\x99\x99\x99\x99YL@'
'''

"""
创建一个可迭代对象的一个原因是它能允许使用一个生成器推导来创建
记录。如果你不使用这种技术，那么代码可能会像下面这样
"""
def read_records(format, f):
    record_struct = Struct(format)
    while True:
        chk = f.read(record_struct.size)
        if chk == b'':
            break
        yield record_struct.unpack(chk)

"""
在 函 数 unpack records() 中 使 用 了 另 外 一 种 方 法 unpack from() 。
unpack from() 对于从一个大型二进制数组中提取二进制数据非常有用，因为它不
会产生任何的临时对象或者进行内存复制操作。你只需要给它一个字节字符串 (或数
组) 和一个字节偏移量，它会从那个位置开始直接解包数据。
如果你使用 unpack() 来代替 unpack from() ，你需要修改代码来构造大量的小的
切片以及进行偏移量的计算。
"""
def unpack_records(format, data):
    record_struct = Struct(format)
    return (record_struct.unpack(data[offset:offset + record_struct.size])
            for offset in range(0, len(data), record_struct.size))

"""
这种方案除了代码看上去很复杂外，还得做很多额外的工作，因为它执行了大量的
偏移量计算，复制数据以及构造小的切片对象。如果你准备从读取到的一个大型字节
字符串中解包大量的结构体的话，unpack from() 会表现的更出色。
在解包的时候，collections 模块中的命名元组对象或许是你想要用到的。它可以
让你给返回元组设置属性名称。
"""

from collections import namedtuple

Record = namedtuple('Record', ['Kind', 'x', 'y'])

with open('../test file/data.p', 'rb') as  f:
    records = (Record(*r) for r in read_records('<idd', f))

for r in records:
    print(r.Kind, r.x, r.y)

"""
如果你的程序需要处理大量的二进制数据，你最好使用 numpy 模块。例如，你可以
将一个二进制数据读取到一个结构化数组中而不是一个元组列表中。
"""
import numpy as np
f = open('../test file/data.b', 'rb')
records = np.fromfile(f, dtype='<i,<d,<d')
print(records)

print(records[0])

print(records[1])

"""
最后提一点，如果你需要从已知的文件格式 (如图片格式，图形文件，HDF5 等)
中读取二进制数据时，先检查看看 Python 是不是已经提供了现存的模块。因为不到万
不得已没有必要去重复造轮子。
"""


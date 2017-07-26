#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第六章：数据编码和处理
Desc: 这一章主要讨论使用 Python 处理各种不同方式编码的数据，比如 CSV 文件，
JSON，XML 和二进制包装记录。和数据结构那一章不同的是，这章不会讨论特殊的
算法问题，而是关注于怎样获取和存储这些格式的数据。

Title;      码解码 Base64 数据
Issue: 你需要使用 Base64 格式解码或编码二进制数据。
Answer: base64 模块中有两个函数 b64encode() and b64decode() 可以帮你解决这个问题。
"""
s = b'hello'
import base64

a = base64.b64encode(s)
print(a)
# b'aGVsbG8='
print(base64.b64decode(a))
# b'hello'

"""
Base64 编码仅仅用于面向字节的数据比如字节字符串和字节数组。此外，编码处
理的输出结果总是一个字节字符串。如果你想混合使用 Base64 编码的数据和 Unicode
文本，你必须添加一个额外的解码步骤。
"""

a = base64.b64encode(s).decode('ascii')
print(a)
# aGVsbG8=
"""
当解码 Base64 的时候，字节字符串和 Unicode 文本都可以作为参数。但是，Uni-
code 字符串只能包含 ASCII 字符。
"""
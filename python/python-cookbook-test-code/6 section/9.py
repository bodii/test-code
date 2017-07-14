#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第六章：数据编码和处理
Desc: 这一章主要讨论使用 Python 处理各种不同方式编码的数据，比如 CSV 文件，
JSON，XML 和二进制包装记录。和数据结构那一章不同的是，这章不会讨论特殊的
算法问题，而是关注于怎样获取和存储这些格式的数据。

Title;      编码和解码十六进制数
Issue: 你想将一个十六进制字符串解码成一个字节字符串或者将一个字节字符串编码成一
个十六进制字符串。
Answer: 如果你只是简单的解码或编码一个十六进制的原始字符串，可以使用　 binascii
模块。
"""
s = b'hello'
import binascii
h = binascii.b2a_hex(s)
print(h)
# b'68656c6c6f'

d = binascii.a2b_hex(h)
print(d)
# b'hello'

"""
类似的功能同样可以在 base64 模块中找到。
"""
import base64
h = base64.b16encode(s)
print(h)
# b'68656C6C6F'
d = base64.b16decode(h)
print(d)
# b'hello'

"""
大部分情况下，通过使用上述的函数来转换十六进制是很简单的。上面两种技术的
主要不同在于大小写的处理。函数 base64.b16decode() 和 base64.b16encode() 只能
操作大写形式的十六进制字母，而 binascii 模块中的函数大小写都能处理。
还有一点需要注意的是编码函数所产生的输出总是一个字节字符串。如果想强制以
Unicode 形式输出，你需要增加一个额外的界面步骤。
"""
h = base64.b16encode(s)
print(h)
#b'68656C6C6F'
print(h.decode('ascii'))
# 68656C6C6F
d = base64.b16decode(h)
print(d)
# b'hello'
print(d.decode('ascii'))
# hello
print(d.decode('utf-8'))
# hello
print(d.decode('gb2312'))
# hello

"""
在解码十六进制数时，函数 b16decode() 和 a2b hex() 可以接受字节或 unicode 字
符串。但是，unicode 字符串必须仅仅只包含 ASCII 编码的十六进制数。
"""
#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第三章：数字日期和时间
Desc: 在 Python 中执行整数和浮点数的数学运算时很简单的。尽管如此，如果你需要执
行分数、数组或者是日期和时间的运算的话，就得做更多的工作了。本章集中讨论的
就是这些主题。

Title;         二八十六进制整数
Issue: 你需要转换或者输出使用二进制，八进制或十六进制表示的整数。
Answer:为了将整数转换为二进制、八进制或十六进制的文本串，可以分别使用 bin() ,
oct() 或 hex() 函数
"""

x = 1234
print(bin(x))
# 0b10011010010
print(oct(x))
# 0o2322
print(hex(x))
# 0x4d2

'''
另外，如果你不想输出 0b , 0o 或者 0x 的前缀的话，可以使用 format() 函数。
'''
print(format(x, 'b'))
# 10011010010
print(format(x, 'o'))
# 2322
print(format(x, 'x'))
# 4d2

'''
整数是有符号的，所以如果你在处理负数的话，输出结果会包含一个负号。
'''
x = -1234
print(format(x, 'b'))
# -10011010010
print(format(x, 'x'))
# -4d2

'''
如果你想产生一个无符号值，你需要增加一个指示最大位长度的值。比如为了显示
32 位的值
'''
x = -1234
print(format(2**32 + x, 'b'))
# 11111111111111111111101100101110
print(format(2**32 + x, 'x'))
# fffffb2e

'''
为了以不同的进制转换整数字符串，简单的使用带有进制的 int() 函数即可
'''
print(int('4d2', 16))
# 1234
print(int('10011010010', 2))
# 1234

'''
import os
>>> os.chmod('script.py', 0755)
File "<stdin>", line 1
os.chmod('script.py', 0755)
^
SyntaxError: invalid token

需确保八进制数的前缀是 0o ，就像下面这样：
os.chmod('script.py', 0o755)
'''

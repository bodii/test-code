#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第七章：函数
Desc: 使用 def 语句定义函数是所有程序的基础。本章的目标是讲解一些更加高级和不
常见的函数定义与使用模式。涉及到的内容包括默认参数、任意数量参数、强制关键
字参数、注解和闭包。另外，一些高级的控制流和利用回调函数传递数据的技术在这
里也会讲解到。

Title;       只接受关键字参数的函数
Issue: 你希望函数的某些参数强制使用关键字参数传递
Answer: 将强制关键字参数放到某个 * 参数或者当个 * 后面就能达到这种效果。
"""

def recv(masize, *, block):
    pass

recv(1024, block=True)

"""
利用这种技术，我们还能在接受任意多个位置参数的函数中指定关键字参数。
"""

def mininum(*values, clip=None):
    m = min(values)
    if clip is not None:
        m = clip if clip > m else m

    return m

"""
很多情况下，使用强制关键字参数会比使用位置参数表意更加清晰，程序也更加具
有可读性
"""
msg = recv(1024, block=False)

"""
另外，使用强制关键字参数也会比使用 **kwargs 参数更好，因为在使用函数 help
的时候输出也会更容易理解
"""
print(help(recv))
"""
强制关键字参数在一些更高级场合同样也很有用。例如，它们可以被用来在使用
*args 和 **kwargs 参数作为输入的函数中插入参数，
"""

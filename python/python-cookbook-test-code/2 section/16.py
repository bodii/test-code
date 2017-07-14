#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第二章：字符串和文本
Desc: 几乎所有有用的程序都会涉及到某些文本处理，不管是解析数据还是产生输出。这
一章将重点关注文本的操作处理，比如提取字符串，搜索，替换以及解析等。大部分
的问题都能简单的调用字符串的内建方法完成。但是，一些更为复杂的操作可能需要
正则表达式或者强大的解析器，所有这些主题我们都会详细讲解。并且在操作 Unicode
时候碰到的一些棘手的问题在这里也会被提及到。

Title;         以指定列宽格式化字符串
Issue:  你有一些长字符串，想以指定的列宽将它们重新格式化
Answer: 使用 textwrap 模块来格式化字符串的输出。
"""

s = "Look into my eyes, look into my eyes, the eyes, the eyes, \
the eyes, not around the eyes, don't look around the eyes, \
look into my eyes, you're under."

'''
下面演示使用 textwrap 格式化字符串的多种方式：
'''
import textwrap
print(textwrap.fill(s, 70))
'''
Look into my eyes, look into my eyes, the eyes, the eyes, the eyes,
not around the eyes, don't look around the eyes, look into my eyes,
you're under.
'''
print(textwrap.fill(s, 40))
'''
Look into my eyes, look into my eyes,
the eyes, the eyes, the eyes, not around
the eyes, don't look around the eyes,
look into my eyes, you're under.
'''
print(textwrap.fill(s, 40, initial_indent='    '))
'''
    Look into my eyes, look into my
eyes, the eyes, the eyes, the eyes, not
around the eyes, don't look around the
eyes, look into my eyes, you're under.
'''
print(textwrap.fill(s, 40, subsequent_indent='   '))
'''
   the eyes, the eyes, the eyes, not
   around the eyes, don't look around
   the eyes, look into my eyes, you're
   under.
'''

'''
textwrap 模块对于字符串打印是非常有用的，特别是当你希望输出自动匹配终端
大小的时候。你可以使用 os.get terminal size() 方法来获取终端的大小尺寸。
'''
import os
print(os.get_terminal_size().columns)
# 未定义(不支持) columns

'''
fill() 方 法 接 受 一 些 其 他 可 选 参 数 来 控 制 tab， 语 句 结 尾 等。
'''
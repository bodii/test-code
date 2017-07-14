#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第二章：字符串和文本
Desc: 几乎所有有用的程序都会涉及到某些文本处理，不管是解析数据还是产生输出。这
一章将重点关注文本的操作处理，比如提取字符串，搜索，替换以及解析等。大部分
的问题都能简单的调用字符串的内建方法完成。但是，一些更为复杂的操作可能需要
正则表达式或者强大的解析器，所有这些主题我们都会详细讲解。并且在操作 Unicode
时候碰到的一些棘手的问题在这里也会被提及到。

Title;         在字符串中处理 html 和 xml
Issue:  你想将 HTML 或者 XML 实体如 &entity; 或 &#code; 替换为对应的文本。再者，
你需要转换文本中特定的字符 (比如<, >, 或 &)。
Answer: 如果你想替换文本字符串中的 ‘<’ 或者 ‘>’ ，使用 html.escape() 函数可以很容易
的完成。
"""

s = 'Elements are written as "<tag>text</tag>".'
import html
print(s)
# Elements are written as "<tag>text</tag>".
print(html.escape(s))
# Elements are written as &quot;&lt;tag&gt;text&lt;/tag&gt;&quot;.
print(html.escape(s, quote=False))
# Elements are written as "&lt;tag&gt;text&lt;/tag&gt;".

'''
如果你正在处理的是 ASCII 文本，并且想将非 ASCII 文本对应的编码实体嵌入进
去，可以给某些 I/O 函数传递参数 errors='xmlcharrefreplace' 来达到这个目。
'''

s = 'Spicy  Jalapeño'
print(s.encode('ascii', errors='xmlcharrefreplace'))
# b'Spicy  Jalape&#241;o'

'''
为了替换文本中的编码实体，你需要使用另外一种方法。如果你正在处理 HTML
或者 XML 文本，试着先使用一个合适的 HTML 或者 XML 解析器。通常情况下，这
些工具会自动替换这些编码值，你无需担心。
有时候，如果你接收到了一些含有编码值的原始文本，需要手动去做替换，通常你
只需要使用 HTML 或者 XML 解析器的一些相关工具函数/方法即可。
'''
s = 'Spicy &quot;Jalape&#241;o&quot.'
from html.parser import HTMLParser
p = HTMLParser()
print(p.unescape(s))
# Spicy "Jalapeño".

t = 'The prompt is &gt;&gt;&gt;'
from xml.sax.saxutils import unescape
print(unescape(t))
# The prompt is >>>

'''
在生成 HTML 或者 XML 文本的时候，如果正确的转换特殊标记字符是一个很容
易被忽视的细节。特别是当你使用 print() 函数或者其他字符串格式化来产生输出的
时候。使用像 html.escape() 的工具函数可以很容易的解决这类问题。
如 果 你 想 以 其 他 方 式 处 理 文 本， 还 有 一 些 其 他 的 工 具 函 数 比 如
xml.sax.saxutils.unescapge() 可以帮助你。然而，你应该先调研清楚怎样使用一个
合适的解析器。比如，如果你在处理 HTML 或 XML 文本，使用某个解析模块比如
html.parse 或 xml.etree.ElementTree 已经帮你自动处理了相关的替换细节。
'''
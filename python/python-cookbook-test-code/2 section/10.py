#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第二章：字符串和文本
Desc: 几乎所有有用的程序都会涉及到某些文本处理，不管是解析数据还是产生输出。这
一章将重点关注文本的操作处理，比如提取字符串，搜索，替换以及解析等。大部分
的问题都能简单的调用字符串的内建方法完成。但是，一些更为复杂的操作可能需要
正则表达式或者强大的解析器，所有这些主题我们都会详细讲解。并且在操作 Unicode
时候碰到的一些棘手的问题在这里也会被提及到。

Title;           在正则式中使用 Unicode
Issue:  你正在使用正则表达式处理文本，但是关注的是 Unicode 字符处理
Answer: 默认情况下 re 模块已经对一些 Unicode 字符类有了基本的支持。比如， \\d 已经
匹配任意的 unicode 数字字符了
"""

import re
num = re.compile('\d+')
print(num.match('123'))
# <_sre.SRE_Match object; span=(0, 3), match='123'>
print(num.match('\u0661\u0662\u0663'))
# <_sre.SRE_Match object; span=(0, 3), match='١٢٣'>

'''
如果你想在模式中包含指定的 Unicode 字符，你可以使用 Unicode 字符对应的转
义序列 (比如 \\uFFF 或者 \\UFFFFFFF )。比如，下面是一个匹配几个不同阿拉伯编码页
面中所有字符的正则表达式：
'''
arabic = re.compile('[\u0600-\u06ff\u0750-\u077f\u08a0-\u08ff]+')
'''
当执行匹配和搜索操作的时候，最好是先标准化并且清理所有文本为标准化格式
(参考 2.9 小节)。但是同样也应该注意一些特殊情况，比如在忽略大小写匹配和大小写
转换时的行为
'''
pat = re.compile('stra\u00dfe', re.IGNORECASE)
s = 'straße'
print(pat.match(s))
# <_sre.SRE_Match object; span=(0, 6), match='straße'>
pat.match(s.upper()) # Doesn't match
print(s.upper()) # STRASSE

'''
混合使用 Unicode 和正则表达式通常会让你抓狂。如果你真的打算这样做的话，最
好考虑下安装第三方正则式库，它们会为 Unicode 的大小写转换和其他大量有趣特性
提供全面的支持，包括模糊匹配。
'''
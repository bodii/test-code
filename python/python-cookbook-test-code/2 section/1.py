#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第二章：字符串和文本
Desc: 几乎所有有用的程序都会涉及到某些文本处理，不管是解析数据还是产生输出。这
一章将重点关注文本的操作处理，比如提取字符串，搜索，替换以及解析等。大部分
的问题都能简单的调用字符串的内建方法完成。但是，一些更为复杂的操作可能需要
正则表达式或者强大的解析器，所有这些主题我们都会详细讲解。并且在操作 Unicode
时候碰到的一些棘手的问题在这里也会被提及到。

Title;        使用多个界定符分割字符串
Issue: 你需要将一个字符串分割为多个字段，但是分隔符 (还有周围的空格) 并不是固定
的。
Answer: string 对象的 split() 方法只适应于非常简单的字符串分割情形，它并不允许有
多个分隔符或者是分隔符周围不确定的空格。当你需要更加灵活的切割字符串的时候，
最好使用 re.split() 方法
"""

line = 'asdf fjdk; afed, fjek,asdf, foo'
import re
print(re.split(r'[;,\s]\s*',line))
# ['asdf', 'fjdk', 'afed', 'fjek', 'asdf', 'foo']

'''
函数 re.split() 是非常实用的，因为它允许你为分隔符指定多个正则模式。比
如，在上面的例子中，分隔符可以是逗号，分号或者是空格，并且后面紧跟着任意个
的空格。只要这个模式被找到，那么匹配的分隔符两边的实体都会被当成是结果中的
元素返回。返回结果为一个字段列表，这个跟 str.split() 返回值类型是一样的。
'''

'''
当你使用 re.split() 函数时候，需要特别注意的是正则表达式中是否包含一个括
号捕获分组。如果使用了捕获分组，那么被匹配的文本也将出现在结果列表中。
'''
fields = re.split(r'(;|,|\s)\s*', line)
print(fields)
# ['asdf', ' ', 'fjdk', ';', 'afed', ',', 'fjek', ',', 'asdf', ',', 'foo']

'''
获取分割字符在某些情况下也是有用的。比如，你可能想保留分割字符串，用来在
后面重新构造一个新的输出字符串
'''
values = fields[::2]
delimiters = fields[1::2] + ['']
print(delimiters) # [' ', ';', ',', ',', ',', '']
print(values) # ['asdf', 'fjdk', 'afed', 'fjek', 'asdf', 'foo']

print(''.join(v + d for v,d in zip(values,delimiters)))
# asdf fjdk;afed,fjek,asdf,foo

'''
如果你不想保留分割字符串到结果列表中去，但仍然需要使用到括号来分组正则表
达式的话，确保你的分组是非捕获分组，形如 (?:...) 。
'''
print(re.split(r'(?:,|;|\s)\s*', line))
# ['asdf', 'fjdk', 'afed', 'fjek', 'asdf', 'foo']
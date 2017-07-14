#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第二章：字符串和文本
Desc: 几乎所有有用的程序都会涉及到某些文本处理，不管是解析数据还是产生输出。这
一章将重点关注文本的操作处理，比如提取字符串，搜索，替换以及解析等。大部分
的问题都能简单的调用字符串的内建方法完成。但是，一些更为复杂的操作可能需要
正则表达式或者强大的解析器，所有这些主题我们都会详细讲解。并且在操作 Unicode
时候碰到的一些棘手的问题在这里也会被提及到。

Title;          最短匹配模式
Issue:  你正在试着用正则表达式匹配某个文本模式，但是它找到的是模式的最长可能匹
配。而你想修改它变成查找最短的可能匹配。
Answer: 这个问题一般出现在需要匹配一对分隔符之间的文本的时候 (比如引号包含的字符
串)。
"""
import re
str_pat = re.compile(r'\"(.*)\"')
text1 = 'Computer says "no."'
print(str_pat.findall(text1))
# ['no.']
text2 = 'Computer says "no." Phone says "yes."'
print(str_pat.findall(text2))
# ['no." Phone says "yes.']

'''
在这个例子中，模式 r'\"(.*)\"' 的意图是匹配被双引号包含的文本。但是在正
则表达式中 * 操作符是贪婪的，因此匹配操作会查找最长的可能匹配。于是在第二个
例子中搜索 text2 的时候返回结果并不是我们想要的。
'''

'''
为了修正这个问题，可以在模式中的 * 操作符后面加上? 修饰符
'''

str_pat = re.compile(r'\"(.*?)\"')
print(str_pat.findall(text2))
# ['no.', 'yes.']
'''
这样就使得匹配变成非贪婪模式，从而得到最短的匹配，也就是我们想要的结果
'''

'''
这一节展示了在写包含点 (.) 字符的正则表达式的时候遇到的一些常见问题。在一
个模式字符串中，点 (.) 匹配除了换行外的任何字符。然而，如果你将点 (.) 号放在开
始与结束符 (比如引号) 之间的时候，那么匹配操作会查找符合模式的最长可能匹配。
这样通常会导致很多中间的被开始与结束符包含的文本被忽略掉，并最终被包含在匹
配结果字符串中返回。通过在 * 或者 + 这样的操作符后面添加一个 ? 可以强制匹配算
法改成寻找最短的可能匹配。
'''
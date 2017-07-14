#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第二章：字符串和文本
Desc: 几乎所有有用的程序都会涉及到某些文本处理，不管是解析数据还是产生输出。这
一章将重点关注文本的操作处理，比如提取字符串，搜索，替换以及解析等。大部分
的问题都能简单的调用字符串的内建方法完成。但是，一些更为复杂的操作可能需要
正则表达式或者强大的解析器，所有这些主题我们都会详细讲解。并且在操作 Unicode
时候碰到的一些棘手的问题在这里也会被提及到。

Title;          多行匹配模式
Issue:  你正在试着使用正则表达式去匹配一大块的文本，而你需要跨越多行去匹配。
Answer: 为了修正这个问题，你可以修改模式字符串，增加对换行的支持
"""
import re
comment = re.compile(r'/\*(.*?)\*/')
text1 = '/* this is a comment */'
text2 = '''/* this is a
     multiline comment */
     '''
print(comment.findall(text1))
# [' this is a comment ']
print(comment.findall(text2))
# []

'''
修正
'''
comment = re.compile(r'/\*((?:.|\n)*?)\*/')
print(comment.findall(text2))
# [' this is a\n     multiline comment ']

'''
在这个模式中， (?:.|\n) 指定了一个非捕获组 (也就是它定义了一个仅仅用来做
匹配，而不能通过单独捕获或者编号的组)

re.compile() 函数接受一个标志参数叫 re.DOTALL ，在这里非常有用。它可以让
正则表达式中的点 (.) 匹配包括换行符在内的任意字符。
'''

comment = re.compile(r'/\*(.*?)\*/',re.DOTALL)
print(comment.findall(text2)) # [' this is a\n     multiline comment ']

'''
对于简单的情况使用 re.DOTALL 标记参数工作的很好，但是如果模式非常复杂或
者是为了构造字符串令牌而将多个模式合并起来 (2.18 节有详细描述)，这时候使用这
个标记参数就可能出现一些问题。如果让你选择的话，最好还是定义自己的正则表达
式模式，这样它可以在不需要额外的标记参数下也能工作的很好
'''


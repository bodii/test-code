#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第二章：字符串和文本
Desc: 几乎所有有用的程序都会涉及到某些文本处理，不管是解析数据还是产生输出。这
一章将重点关注文本的操作处理，比如提取字符串，搜索，替换以及解析等。大部分
的问题都能简单的调用字符串的内建方法完成。但是，一些更为复杂的操作可能需要
正则表达式或者强大的解析器，所有这些主题我们都会详细讲解。并且在操作 Unicode
时候碰到的一些棘手的问题在这里也会被提及到。

Title;          字符串忽略大小写的搜索替换
Issue:  你需要以忽略大小写的方式搜索与替换文本字符串
Answer: 为了在文本操作时忽略大小写，你需要在使用 re 模块的时候给这些操作提供
re.IGNORECASE 标志参数
"""
import re

text = 'UPPER PYTHON, lower python, Mixed Python'
print(re.findall('python', text, flags = re.IGNORECASE))
# ['PYTHON', 'python', 'Python']
print(re.sub('python','snake',text,flags = re.IGNORECASE))
# UPPER snake, lower snake, Mixed snake

'''
最后的那个例子揭示了一个小缺陷，替换字符串并不会自动跟被匹配字符串的大小
写保持一致。为了修复这个，你可能需要一个辅助函数
'''

def matchcase(word):
    def replace(m):
        text = m.group()
        if text.isupper():
            return word.upper()
        elif text.islower():
            return word.lower()
        elif text[0].isupper():
            return word.capitalize()
        else:
            return word
    return replace

print(re.sub('(python|Python|PYTHON)', matchcase('snake'), text))
# UPPER SNAKE, lower snake, Mixed Snake

'''
译者注： matchcase('snake') 返回了一个回调函数 (参数必须是 match 对象)，前
面一节一节提到过， sub() 函数除了接受替换字符串外，还能接受一个回调函数

对于一般的忽略大小写的匹配操作，简单的传递一个 re.IGNORECASE 标志参数就
已经足够了。但是需要注意的是，这个对于某些需要大小写转换的 Unicode 匹配可能
还不够
'''

#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第二章：字符串和文本
Desc: 几乎所有有用的程序都会涉及到某些文本处理，不管是解析数据还是产生输出。这
一章将重点关注文本的操作处理，比如提取字符串，搜索，替换以及解析等。大部分
的问题都能简单的调用字符串的内建方法完成。但是，一些更为复杂的操作可能需要
正则表达式或者强大的解析器，所有这些主题我们都会详细讲解。并且在操作 Unicode
时候碰到的一些棘手的问题在这里也会被提及到。

Title;         字符串搜索和替换
Issue:  你想在字符串中搜索和匹配指定的文本模式
Answer: 对于简单的字面模式，直接使用 str.repalce() 方法即可
"""

text = 'yeah, but no, but yeah, but no, but yeah'
print(text.replace('yeah', 'year'))
# year, but no, but year, but no, but year


'''
对于复杂的模式，请使用 re 模块中的 sub() 函数。为了说明这个，假设你想将形
式为 11/27/201 的日期字符串改成 2012-11-27
'''
import re
text = 'Today is 11/27/2012. PyCon starts 3/13/2013.'
print(re.sub(r'(\d+)/(\d+)/(\d+)', r'\3-\1-\2',text))
# Today is 2012-11-27. PyCon starts 2013-3-13.

'''
sub() 函数中的第一个参数是被匹配的模式，第二个参数是替换模式。反斜杠数字
比如 \3 指向前面模式的捕获组号
如果你打算用相同的模式做多次替换，考虑先编译它来提升性能。
'''
datepat = re.compile(r'(\d+)/(\d+)/(\d+)')
print(datepat.sub(r'\3-\1-\2',text))
# Today is 2012-11-27. PyCon starts 2013-3-13.

'''
对于更加复杂的替换，可以传递一个替换回调函数来代替
'''
from calendar import month_abbr
def change_date(m):
    mon_name = month_abbr[int(m.group(1))]
    return '{} {} {}'.format(m.group(2), mon_name, m.group(3))

print(datepat.sub(change_date, text))
# Today is 27 Nov 2012. PyCon starts 13 Mar 2013

'''
一个替换回调函数的参数是一个 match 对象，也就是 match() 或者 find() 返回的
对象。使用 group() 方法来提取特定的匹配部分。回调函数最后返回替换字符串
'''

'''
如果除了替换后的结果外，你还想知道有多少替换发生了，可以使用 re.subn()
来代替。
'''
newtext, n = datepat.subn(r'\3-\1-\2',text)
print(newtext)
# Today is 2012-11-27. PyCon starts 2013-3-13.
print(n) # 2

'''
关于正则表达式搜索和替换，上面演示的 sub() 方法基本已经涵盖了所有。
'''
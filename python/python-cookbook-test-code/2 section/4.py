#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第二章：字符串和文本
Desc: 几乎所有有用的程序都会涉及到某些文本处理，不管是解析数据还是产生输出。这
一章将重点关注文本的操作处理，比如提取字符串，搜索，替换以及解析等。大部分
的问题都能简单的调用字符串的内建方法完成。但是，一些更为复杂的操作可能需要
正则表达式或者强大的解析器，所有这些主题我们都会详细讲解。并且在操作 Unicode
时候碰到的一些棘手的问题在这里也会被提及到。

Title;         字符串匹配和搜索
Issue:  你想匹配或者搜索特定模式的文本
Answer: 如果你想匹配的是字面字符串，那么你通常只需要调用基本字符串方法就行，比如
str.find() , str.endswith() , str.startswith() 或者类似的方法
"""

text = 'yeah, but no, but yeah, but no, but yeah'

print(text == 'yeah') # False
print(text.startswith('yeah')) # True
print(text.endswith('no')) # False
print(text.find('no')) # 10

'''
对于复杂的匹配需要使用正则表达式和 re 模块。为了解释正则表达式的基本原理，
假设你想匹配数字格式的日期字符串比如 11/27/2012 ，你可以这样做
'''
text1 = '11/27/2012'
text2 = 'Nov 27, 2012'

import re

if re.match(r'\d+/\d+/\d+', text1):
    print('yes')
else:
    print('no')
# yes

if re.match(r'\d+/\d+/\d+' ,text2):
    print('yes')
else:
    print('no')
# no

'''
如果你想使用同一个模式去做多次匹配，你应该先将模式字符串预编译为模式对
象。
'''
datepat = re.compile(r'\d+/\d+/\d+')
if datepat.match(text1):
    print('yes')
else:
    print('no')
# yes
if datepat.match(text2):
    print('yes')
else:
    print('no')
# no

'''
match() 总是从字符串开始去匹配，如果你想查找字符串任意部分的模式出现位
置，使用 findall() 方法去代替。
'''
text = 'Today is 11/27/2012. PyCon starts 3/13/2013.'
print(datepat.findall(text))
# ['11/27/2012', '3/13/2013']

'''
在定义正则式的时候，通常会利用括号去捕获分组。
'''
datepat = re.compile(r'(\d+)/(\d+)/(\d+)')
m = datepat.match('11/27/2013')
print(m)
# <_sre.SRE_Match object; span=(0, 10), match='11/27/2013'>
print(m.group(0)) #  11/27/2013
print(m.group(1)) # 11
print(m.group(2)) # 27
print(m.group(3)) # 2013
print(m.groups()) # ('11', '27', '2013')
month, day, year = m.groups()
print(month) # 11
print(day) # 27
print(year) # 2013

print(datepat.findall(text))
# [('11', '27', '2012'), ('3', '13', '2013')]

for month, day, year in datepat.findall(text):
    print('{}-{}-{}'.format(year,month,day))
'''
2012-11-27
2013-3-13
'''

'''
findall() 方法会搜索文本并以列表形式返回所有的匹配。如果你想以迭代方式返
回匹配，可以使用 finditer() 方法来代替
'''
for m in datepat.finditer(text):
    print(m.groups())
'''
('11', '27', '2012')
('3', '13', '2013')
'''

'''
关于正则表达式理论的教程已经超出了本书的范围。不过，这一节阐述了使用 re
模块进行匹配和搜索文本的最基本方法。核心步骤就是先使用 re.compile() 编译正则
表达式字符串，然后使用 match() , findall() 或者 finditer() 等方法。
当写正则式字符串的时候，相对普遍的做法是使用原始字符串比如 r'(\d+)/(\d
+)/(\d+)' 。这种字符串将不去解析反斜杠，这在正则表达式中是很有用的。如果不
这样做的话，你必须使用两个反斜杠，类似 '(\\d+)/(\\d+)/(\\d+)' 。
'''

'''
需要注意的是 match() 方法仅仅检查字符串的开始部分。它的匹配结果有可能并
不是你期望的那样。
'''
m = datepat.match('11/27/212abcdef')
print(m)
# <_sre.SRE_Match object; span=(0, 9), match='11/27/212'>
print(m.group()) # 11/27/212

'''
如果你想精确匹配，确保你的正则表达式以 $ 结尾
'''
datepat = re.compile(r'(\d+)/(\d+)/(\d+)$')
print(datepat.match('11/27/2012abcdef'))
# None
print(datepat.match('11/27/2012'))
# <_sre.SRE_Match object; span=(0, 10), match='11/27/2012'>

'''
最后，如果你仅仅是做一次简单的文本匹配/搜索操作的话，可以略过编译部分，
直接使用 re 模块级别的函数。比如：
'''
print(re.findall(r'(\d+)/(\d+)/(\d+)', text))

'''
但是需要注意的是，如果你打算做大量的匹配和搜索操作的话，最好先编译正则表
达式，然后再重复使用它。模块级别的函数会将最近编译过的模式缓存起来，因此并
不会消耗太多的性能，但是如果使用预编译模式的话，你将会减少查找和一些额外的
处理损耗。
'''

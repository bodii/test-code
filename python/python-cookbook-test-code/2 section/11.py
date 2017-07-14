#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第二章：字符串和文本
Desc: 几乎所有有用的程序都会涉及到某些文本处理，不管是解析数据还是产生输出。这
一章将重点关注文本的操作处理，比如提取字符串，搜索，替换以及解析等。大部分
的问题都能简单的调用字符串的内建方法完成。但是，一些更为复杂的操作可能需要
正则表达式或者强大的解析器，所有这些主题我们都会详细讲解。并且在操作 Unicode
时候碰到的一些棘手的问题在这里也会被提及到。

Title;           删除字符串中不需要的字符
Issue:  你想去掉文本字符串开头，结尾或者中间不想要的字符，比如空白
Answer: strip() 方法能用于删除开始或结尾的字符。 lstrip() 和 rstrip() 分别从左和
从右执行删除操作。默认情况下，这些方法会去除空白字符，但是你也可以指定其他
字符。
"""

s = ' hello world \n'
print(s.strip())
# hello world
print(s.lstrip())
# hello world

print(s.rstrip())
# #  hello world

t = '-----hello world======='
print(t.lstrip('-'))
# hello world=======
print(t.rstrip('='))
#-----hello world
print(t.strip('-='))
# hello world

'''
这些 strip() 方法在读取和清理数据以备后续处理的时候是经常会被用到的。比
如，你可以用它们来去掉空格，引号和完成其他任务。
但是需要注意的是去除操作不会对字符串的中间的文本产生任何影响。
'''

s = 'hello          world \n'
print(s.strip())
# hello          world

'''
如果你想处理中间的空格，那么你需要求助其他技术。比如使用 replace() 方法
或者是用正则表达式替换。
'''
print(s.replace(' ','',9))
# hello world
import  re
print(re.sub('\s+',' ',s)) # hello world

'''
通常情况下你想将字符串 strip 操作和其他迭代操作相结合，比如从文件中读取
多行数据。如果是这样的话，那么生成器表达式就可以大显身手了。比如：
with open(filename) as f:
    lines = (line.strip() for line in f)
    for line in lines:
        print(line)
在这里，表达式 lines = (line.strip() for line in f) 执行数据转换操作。这
种方式非常高效，因为它不需要预先读取所有数据放到一个临时的列表中去。它仅仅
只是创建一个生成器，并且每次返回行之前会先执行 strip 操作。
对于更高阶的 strip，你可能需要使用 translate() 方法。
'''


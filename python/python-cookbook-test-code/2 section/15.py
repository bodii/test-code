#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第二章：字符串和文本
Desc: 几乎所有有用的程序都会涉及到某些文本处理，不管是解析数据还是产生输出。这
一章将重点关注文本的操作处理，比如提取字符串，搜索，替换以及解析等。大部分
的问题都能简单的调用字符串的内建方法完成。但是，一些更为复杂的操作可能需要
正则表达式或者强大的解析器，所有这些主题我们都会详细讲解。并且在操作 Unicode
时候碰到的一些棘手的问题在这里也会被提及到。

Title;         字符串中插入变量
Issue:  你想创建一个内嵌变量的字符串，变量被它的值所表示的字符串替换掉
Answer: Python 并没有对在字符串中简单替换变量值提供直接的支持。但是通过使用字符
串的 format() 方法来解决这个问题。
"""

s = '{name} has {n} messages.'
print(s.format(name='Guido', n=37))
# Guido has 37 messages.

'''
或者，如果要被替换的变量能在变量域中找到，那么你可以结合使用 format map()
和 vars() 。
'''
name = 'Guido'
n = 37
print(s.format_map(vars()))
# Guido has 37 messages.

'''
vars() 还有一个有意思的特性就是它也适用于对象实例。
'''
class Info:
    def __init__(self,name,n):
        self.name = name
        self.n = n
a = Info('Guido', 37)
print(s.format_map(vars(a)))
# Guido has 37 messages.

'''
format 和 format map() 的一个缺陷就是它们并不能很好的处理变量缺失的情况，
比如：
>>> s.format(name='Guido')
Traceback (most recent call last):
    File "<stdin>", line 1, in <module>
KeyError: 'n'
>>
'''

class safesub(dict):
    '''防止 key 找不到'''
    def __missing__(self,key):
        return '{' + key + '}'

del n
print(s.format_map(safesub(vars())))
# Guido has {n} messages.


'''
如果你发现自己在代码中频繁的执行这些步骤，你可以将变量替换步骤用一个工具
函数封装起来。
'''

import sys
def sub(text):
    return text.format_map(safesub(sys._getframe(1).f_locals))

name = 'Guido'
n = 37
print(sub('Hello {name}'))
# Hello Guido
print(sub('You have {n} messages'))
# You have 37 messages
print(sub('Your favorite color is {color}'))
# Your favorite color is {color}

'''
多年以来由于 Python 缺乏对变量替换的内置支持而导致了各种不同的解决方案。
作为本节中展示的一个可能的解决方案，你可以有时候会看到像下面这样的字符串格
式化代码：
'''
name = 'Guido'
n = 37
# print('%(name) has %(n) messages.' % vars())

'''
字符串模板的使用
'''
import string
s = string.Template('$name has $n messages.')
print(s.substitute(vars()))
# Guido has 37 messages.

'''
然而， format() 和 format map() 相比较上面这些方案而已更加先进，因此应该
被优先选择。使用 format() 方法还有一个好处就是你可以获得对字符串格式化的所有
支持 (对齐，填充，数字格式化等待)，而这些特性是使用像模板字符串之类的方案不
可能获得的。
本机还部分介绍了一些高级特性。映射或者字典类中鲜为人知的 missing () 方
法可以让你定义如何处理缺失的值。在 SafeSub 类中，这个方法被定义为对缺失的值
返回一个占位符。你可以发现缺失的值会出现在结果字符串中 (在调试的时候可能很有
用)，而不是产生一个 KeyError 异常。
sub() 函 数 使 用 sys. getframe(1) 返 回 调 用 者 的 栈 帧。 可 以 从 中 访 问 属 性
f locals 来获得局部变量。毫无疑问绝大部分情况下在代码中去直接操作栈帧应
该是不推荐的。但是，对于像字符串替换工具函数而言它是非常有用的。另外，值得
注意的是 f locals 是一个复制调用函数的本地变量的字典。尽管你可以改变 f locals
的内容，但是这个修改对于后面的变量访问没有任何影响。所以，虽说访问一个栈帧
看上去很邪恶，但是对它的任何操作不会覆盖和改变调用者本地变量的值。
'''
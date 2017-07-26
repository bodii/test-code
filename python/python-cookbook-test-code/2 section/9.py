#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第二章：字符串和文本
Desc: 几乎所有有用的程序都会涉及到某些文本处理，不管是解析数据还是产生输出。这
一章将重点关注文本的操作处理，比如提取字符串，搜索，替换以及解析等。大部分
的问题都能简单的调用字符串的内建方法完成。但是，一些更为复杂的操作可能需要
正则表达式或者强大的解析器，所有这些主题我们都会详细讲解。并且在操作 Unicode
时候碰到的一些棘手的问题在这里也会被提及到。

Title;           将 Unicode 文本标准化
Issue:  你正在处理 Unicode 字符串，需要确保所有字符串在底层有相同的表示。
Answer: 你可以使用 unicodedata 模块先将文本标准化。
"""
s1 = 'Spicy Jalape\u00f1o'
s2 = 'Spicy Jalapen\u0303o'
print(s1) # Spicy Jalapeño
print(s2) # Spicy Jalapeño
print(s1 == s2) # False
print(len(s1)) # 14
print(len(s2)) # 15

'''
这里的文本”Spicy Jalapeño” 使用了两种形式来表示。第一种使用整体字符”ñ”(U
+00F1)，第二种使用拉丁字母”n” 后面跟一个”˜” 的组合字符 (U+0303)
'''
'''
在需要比较字符串的程序中使用字符的多种表示会产生问题。为了修正这个问题，
你可以使用 unicodedata 模块先将文本标准化：
'''
import unicodedata
t1 = unicodedata.normalize('NFC', s1)
t2 = unicodedata.normalize('NFC', s2)
print(t1 == t2)
# True
print(ascii(t1))
# 'Spicy Jalape\xf1o'
t3 = unicodedata.normalize('NFD', s1)
t4 = unicodedata.normalize('NFD', s2)
print(t3 == t4)
# True
print(ascii(t3))
# 'Spicy Jalapen\u0303o'

'''
Python 同样支持扩展的标准化形式 NFKC 和 NFKD，它们在处理某些字符的时候
增加了额外的兼容特性。
'''
s = '\ufb01'
print(s) # ﬁ
print(unicodedata.normalize('NFD', s)) # ﬁ
print(unicodedata.normalize('NFKD', s)) # fi
print(unicodedata.normalize('NFKC', s)) # fi

'''
标准化对于任何需要以一致的方式处理 Unicode 文本的程序都是非常重要的。当处
理来自用户输入的字符串而你很难去控制编码的时候尤其如此。

在清理和过滤文本的时候字符的标准化也是很重要的。比如，假设你想清除掉一些
文本上面的变音符的时候 (可能是为了搜索和匹配)
'''
t1 = unicodedata.normalize('NFD', s1)
print(''.join(c for c in t1 if not unicodedata.combining(c)))
# Spicy Jalapeno

'''
最后一个例子展示了 unicodedata 模块的另一个重要方面，也就是测试字符类的
工具函数。 combining() 函数可以测试一个字符是否为和音字符。在这个模块中还有
其他函数用于查找字符类别，测试是否为数字字符等等。

Unicode 显然是一个很大的主题。如果想更深入的了解关于标准化方面的信息，请
看考 Unicode 官网中关于这部分的说明 Ned Batchelder 在 他的网站 上对 Python 的
Unicode 处理问题也有一个很好的介绍
'''
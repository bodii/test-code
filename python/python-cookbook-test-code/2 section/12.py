#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第二章：字符串和文本
Desc: 几乎所有有用的程序都会涉及到某些文本处理，不管是解析数据还是产生输出。这
一章将重点关注文本的操作处理，比如提取字符串，搜索，替换以及解析等。大部分
的问题都能简单的调用字符串的内建方法完成。但是，一些更为复杂的操作可能需要
正则表达式或者强大的解析器，所有这些主题我们都会详细讲解。并且在操作 Unicode
时候碰到的一些棘手的问题在这里也会被提及到。

Title;           审查清理文本字符串
Issue:  一些无聊的幼稚黑客在你的网站页面表单中输入文本”pýtĥöñ”，然后你想将这些字
符清理掉。
Answer: 文本清理问题会涉及到包括文本解析与数据处理等一系列问题。在非常简单的情形
下，你可能会选择使用字符串函数 (比如 str.upper() 和 str.lower() ) 将文本转为标
准格式。使用 str.replace() 或者 re.sub() 的简单替换操作能删除或者改变指定的
字符序列。你同样还可以使用 2.9 小节的 unicodedata.normalize() 函数将 unicode
文本标准化。
然后，有时候你可能还想在清理操作上更进一步。比如，你可能想消除整个区间上
的字符或者去除变音符。为了这样做，你可以使用经常会被忽视的 str.translate()
方法。
"""

s = 'pýtĥöñ\fis\tawesome\r\n'
print(s)
# pýtĥöñis	awesome

remap = {
    ord('\t') : ' ',
    ord('\f') : ' ',
    ord('\r') : None
}

a = s.translate(remap)
print(a)
# pýtĥöñ is awesome

'''
正如你看的那样，空白字符 \t 和 \f 已经被重新映射到一个空格。回车字符 r 直
接被删除。
你可以以这个表格为基础进一步构建更大的表格。比如，让我们删除所有的和音
符
'''
import unicodedata
import sys

cmb_chrs = dict.fromkeys(c for c in range(sys.maxunicode) if unicodedata.combining(chr(c)))
b = unicodedata.normalize('NFD', a)
print(b)
# pýtĥöñ is awesome
print(b.translate(cmb_chrs))
# python is awesome

'''
上面例子中，通过使用 dict.fromkeys() 方法构造一个字典，每个 Unicode 和音
符作为键，对于的值全部为 None 。
然后使用 unicodedata.normalize() 将原始输入标准化为分解形式字符。然后再
调用 translate 函数删除所有重音符。同样的技术也可以被用来删除其他类型的字符
(比如控制字符等)。
'''

digitmap = {
    c:ord('0') + unicodedata.digit(chr(c))
    for c in range(sys.maxunicode)
    if unicodedata.category(chr(c)) == 'Nd'
}

print(len(digitmap)) # 550

x = '\u0661\u0662\u0663'
print(x.translate(digitmap))
# 123

'''
另一种清理文本的技术涉及到 I/O 解码与编码函数。这里的思路是先对文本做一
些初步的清理，然后再结合 encode() 或者 decode() 操作来清除或修改它。
'''
print(a)
# pýtĥöñ is awesome
b = unicodedata.normalize('NFD', a)
print(b.encode('ascii', 'ignore').decode('ascii'))
# python is awesome


'''
这里的标准化操作将原来的文本分解为单独的和音符。接下来的 ASCII 编码/解码
只是简单的一下子丢弃掉那些字符。当然，这种方法仅仅只在最后的目标就是获取到
文本对应 ACSII 表示的时候生效
'''

'''
文本字符清理一个最主要的问题应该是运行的性能。一般来讲，代码越简单运行越
快。对于简单的替换操作， str.replace() 方法通常是最快的，甚至在你需要多次调
用的时候。
'''

def clean_spaces(s):
    s = s.replace('\r', '')
    s = s.replace('\t', ' ')
    s = s.replace('\f', ' ')
    return s


'''
如果你去测试的话，你就会发现这种方式会比使用 translate() 或者正则表达式
要快很多。
另一方面，如果你需要执行任何复杂字符对字符的重新映射或者删除操作的话，
tanslate() 方法会非常的快。
从大的方面来讲，对于你的应用程序来说性能是你不得不去自己研究的东西。不幸
的是，我们不可能给你建议一个特定的技术，使它能够适应所有的情况。因此实际情
况中需要你自己去尝试不同的方法并评估它。
尽管这一节集中讨论的是文本，但是类似的技术也可以适用于字节，包括简单的替
换，转换和正则表达式。
'''
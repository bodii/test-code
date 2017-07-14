#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第二章：字符串和文本
Desc: 几乎所有有用的程序都会涉及到某些文本处理，不管是解析数据还是产生输出。这
一章将重点关注文本的操作处理，比如提取字符串，搜索，替换以及解析等。大部分
的问题都能简单的调用字符串的内建方法完成。但是，一些更为复杂的操作可能需要
正则表达式或者强大的解析器，所有这些主题我们都会详细讲解。并且在操作 Unicode
时候碰到的一些棘手的问题在这里也会被提及到。

Title;         合并拼接字符串
Issue:  你想将几个小的字符串合并为一个大的字符串
Answer: 如果你想要合并的字符串是在一个序列或者 iterable 中，那么最快的方式就是使
用 join() 方法。
"""
parts = ['Is', 'Chicag', 'Not','Chicago?']
print(' '.join(parts))
# Is Chicag Not Chicago?
print(','.join(parts))
#  Is,Chicag,Not,Chicago?
print(''.join(parts))
# IsChicagNotChicago?

'''
初看起来，这种语法看上去会比较怪，但是 join() 被指定为字符串的一个方法。
这样做的部分原因是你想去连接的对象可能来自各种不同的数据序列 (比如列表，元
组，字典，文件，集合或生成器等)，如果在所有这些对象上都定义一个 join() 方法
明显是冗余的。因此你只需要指定你想要的分割字符串并调用他的 join() 方法去将文
本片段组合起来
'''

'''
如果你仅仅只是合并少数几个字符串，使用加号 (+) 通常已经足够了
'''

a =  'Is Chicago'
b  = 'Not Chicago?'
print(a + ' ' + b)
# Is Chicago Not Chicago?

'''
加号 (+) 操作符在作为一些复杂字符串格式化的替代方案的时候通常也工作的很
好，
'''
print('{} {}'.format(a,b))
# Is Chicago Not Chicago?
print(a + ' '+ b)
# Is Chicago Not Chicago?

a = 'Hello ' 'World'
print(a)
# Hello World


'''
最重要的需要引起注意的是，当我们使用加号 (+) 操作符去连接大量的字符串的
时候是非常低效率的，因为加号连接会引起内存复制以及垃圾回收操作。特别的，你
永远都不应像下面这样写字符串连接代码：
s = ''
for p in parts:
    s += p
这种写法会比使用 join() 方法运行的要慢一些，因为每一次执行 += 操作的时候
会创建一个新的字符串对象。你最好是先收集所有的字符串片段然后再将它们连接起
来
'''

'''
一个相对比较聪明的技巧是利用生成器表达式 (参考 1.19 小节) 转换数据为字符串
的同时合并字符串，
'''
data = ['ACME', 50, 91.1]
print(','.join(str(d) for d in data))
# ACME,50,91.1

'''
同样还得注意不必要的字符串连接操作。有时候程序员在没有必要做连接操作的时
候仍然多此一举。比如在打印的时候：
print(a + ':' + b + ":" + c)
print(':'.join([a, b, c]))
print(a, b, c, sep=':')
'''

'''
当混合使用 I/O 操作和字符串连接操作的时候，有时候需要仔细研究你的程序。
f.write(chunk1 + chunk2)

f.write(chunk1)
f.write(chunk2)
如果两个字符串很小，那么第一个版本性能会更好些，因为 I/O 系统调用天生就
慢。另外一方面，如果两个字符串很大，那么第二个版本可能会更加高效，因为它避
免了创建一个很大的临时结果并且要复制大量的内存块数据。还是那句话，有时候是
需要根据你的应用程序特点来决定应该使用哪种方案
'''

'''
最后谈一下，如果你准备编写构建大量小字符串的输出代码，你最好考虑下使用生
成器函数，利用 yield 语句产生输出片段。
'''

def sample():
    yield 'Is'
    yield 'Chicago'
    yield 'Not'
    yield 'Chicago?'

text = ' '.join(sample())
print(text)
# Is Chicago Not Chicago?

'''
或者你也可以将字符串片段重定向到 I/O：
for part in sample():
    f.write(part)
再或者你还可以写出一些结合 I/O 操作的混合方案：
'''

def combine(source, maxsize):
    parts = []
    size = 0
    for part in source:
        parts.append(part)
        size += len(part)
        if size > maxsize:
            yield ''.join(parts)
            parts = []
            size = 0
        yield ''.join(parts)

#结合文件操作
with open('../test file/filename', 'w') as f:
    for part in combine(sample(), 32768):
        f.write(part)

'''
这里的关键点在于原始的生成器函数并不需要知道使用细节，它只负责生成字符串
片段就行了
'''
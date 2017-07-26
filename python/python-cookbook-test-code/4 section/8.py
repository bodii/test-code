#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第四章：迭代器与生成器
Desc: 迭代是 Python 最强大的功能之一。初看起来，你可能会简单的认为迭代只不过是
处理序列中元素的一种方法。然而，绝非仅仅就是如此，还有很多你可能不知道的，
比如创建你自己的迭代器对象，在 itertools 模块中使用有用的迭代模式，构造生成器
函数等等。这一章目的就是向你展示跟迭代有关的各种常见问题。

Title;      跳过可迭代对象的开始部分
Issue:你想遍历一个可迭代对象，但是它开始的某些元素你并不感兴趣，想跳过它们。
Answer: itertools 模 块 中 有 一 些 函 数 可 以 完 成 这 个 任 务。 首 先 介 绍 的 是
itertools.dropwhile() 函数。使用时，你给它传递一个函数对象和一个可迭代对
象。它会返回一个迭代器对象，丢弃原有序列中直到函数返回 True 之前的所有元素，
然后返回后面所有元素。
"""

'''
为了演示，假定你在读取一个开始部分是几行注释的源文件。
'''
with open('../test file/passwd') as f:
    for line in f:
        print(line, end='')
'''
##
# User Database
#
# Note that this file is consulted directly only when the system is running
# in single-user mode. At other times, this information is provided by
# Open Directory.
...
##
nobody:*:-2:-2:Unprivileged User:/var/empty:/usr/bin/false
root:*:0:0:System Administrator:/var/root:/bin/sh...
'''

'''
如果你想跳过开始部分的注释行的话，可以这样做：
'''
from itertools import dropwhile
with open('../test file/passwd') as f:
    for line in dropwhile(lambda lines: lines.startswith(('#','##','...')), f):
        print(line, end='')
'''
nobody:*:-2:-2:Unprivileged User:/var/empty:/usr/bin/false
root:*:0:0:System Administrator:/var/root:/bin/sh
'''

'''
如果你已经明确知道了要跳过的元素的个数的话，那么可以使用 itertools.islice() 来代替。
'''

from itertools import islice

items = ['a', 'b', 'c', 1, 4, 5, 9]
for x in islice(items,3,None):
    print(x)
'''
4
5
9
'''

'''
在这个例子中， islice() 函数最后那个 None 参数指定了你要获取从第 3 个到最
后的所有元素，如果 None 和 3 的位置对调，意思就是仅仅获取前三个元素恰恰相反，
(这个跟切片的相反操作 [3:] 和 [:3] 原理是一样的)。
'''

'''
函数 dropwhile() 和 islice() 其实就是两个帮助函数，为的就是避免写出下面这
种冗余代码：
'''
with open('../test file/passwd') as f:
    while True:
        line = next(f, '')
        if not line.startswith('#'):
            break

    while line:
        print(line, end='')
        line = next(f, None)

'''
跳过一个可迭代对象的开始部分跟通常的过滤是不同的。比如，上述代码的第一个
部分可能会这样重写：
'''
with open('../test file/passwd') as f:
    lines = (line for line in f if not line.startswith('#'))
    for line in lines:
        print(line, end='')

'''
这样写确实可以跳过开始部分的注释行，但是同样也会跳过文件中其他所有的注
释行。换句话讲，我们的解决方案是仅仅跳过开始部分满足测试条件的行，在那以后，
所有的元素不再进行测试和过滤了。
最后需要着重强调的一点是，本节的方案适用于所有可迭代对象，包括那些事先不
能确定大小的，比如生成器，文件及其类似的对象。
'''

#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第四章：迭代器与生成器
Desc: 迭代是 Python 最强大的功能之一。初看起来，你可能会简单的认为迭代只不过是
处理序列中元素的一种方法。然而，绝非仅仅就是如此，还有很多你可能不知道的，
比如创建你自己的迭代器对象，在 itertools 模块中使用有用的迭代模式，构造生成器
函数等等。这一章目的就是向你展示跟迭代有关的各种常见问题。

Title;      迭代器切片
Issue:你想得到一个由迭代器生成的切片对象，但是标准切片操作并不能做到。
Answer: 函数 itertools.islice() 正好适用于在迭代器和生成器上做切片操作。
"""
def count(n):
    while True:
        yield n
        n += 1

c = count(0)
#print(c[10:20])
'''
Traceback (most recent call last):
  File "python cookbook test code/4 section/7.py", line 21, in <module>
    print(c[10:20])
TypeError: 'generator' object is not subscriptable
'''

import itertools
for x in itertools.islice(c, 10, 20):
    print(x)

'''
10
11
12
13
14
15
16
17
18
19
'''
'''
迭代器和生成器不能使用标准的切片操作，因为它们的长度事先我们并不知道 (并
且也没有实现索引)。函数 islice() 返回一个可以生成指定元素的迭代器，它通过遍
历并丢弃直到切片开始索引位置的所有元素。然后才开始一个个的返回元素，并直到
切片结束索引位置。
这里要着重强调的一点是 islice() 会消耗掉传入的迭代器中的数据。必须考虑到
迭代器是不可逆的这个事实。所以如果你需要之后再次访问这个迭代器的话，那你就
得先将它里面的数据放入一个列表中。
'''
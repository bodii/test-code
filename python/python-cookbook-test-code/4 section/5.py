#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第四章：迭代器与生成器
Desc: 迭代是 Python 最强大的功能之一。初看起来，你可能会简单的认为迭代只不过是
处理序列中元素的一种方法。然而，绝非仅仅就是如此，还有很多你可能不知道的，
比如创建你自己的迭代器对象，在 itertools 模块中使用有用的迭代模式，构造生成器
函数等等。这一章目的就是向你展示跟迭代有关的各种常见问题。

Title;        反向迭代
Issue:你想反方向迭代一个序列
Answer: 使用内置的 reversed() 函数
"""

a = [1, 2, 3, 4]
for x in reversed(a):
    print(x)

'''
4
3
2
1
'''

'''
反向迭代仅仅当对象的大小可预先确定或者对象实现了 reversed () 的特殊方
法时才能生效。如果两者都不符合，那你必须先将对象转换为一个列表才行，比如：
'''

f = open('../test file/somefile.txt')
for line in reversed(list(f)):
    print(line, end='')

'''
要注意的是如果可迭代对象元素很多的话，将其预先转换为一个列表要消耗大量的
内存。
'''

class Countdown:
    def __init__(self, start):
        self.start = start

    def __iter__(self):
        n = self.start
        while n > 0:
            yield n
            n -= 1

    def __reversed__(self):
        n = 1
        while n <= self.start:
            yield n
            n += 1

for rr in reversed(Countdown(30)):
    print(rr)
'''
1
2
3
4
5
6
7
8
9
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
20
21
22
23
24
25
26
27
28
29
30
'''
for rr in Countdown(30):
    print(rr)
'''
30
29
28
27
26
25
24
23
22
21
20
19
18
17
16
15
14
13
12
11
10
9
8
7
6
5
4
3
2
1
'''
#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第四章：迭代器与生成器
Desc: 迭代是 Python 最强大的功能之一。初看起来，你可能会简单的认为迭代只不过是
处理序列中元素的一种方法。然而，绝非仅仅就是如此，还有很多你可能不知道的，
比如创建你自己的迭代器对象，在 itertools 模块中使用有用的迭代模式，构造生成器
函数等等。这一章目的就是向你展示跟迭代有关的各种常见问题。

Title;       排列组合的迭代
Issue:你想迭代遍历一个集合中元素的所有可能的排列或组合
Answer: itertools 模 块 提 供 了 三 个 函 数 来 解 决 这 类 问 题。 其 中 一 个 是
itertools.permutations() ， 它 接 受 一 个 集 合 并 产 生 一 个 元组 序 列， 每 个 元 组
由集合中所有元素的一个可能排列组成。也就是说通过打乱集合中元素排列顺序生成
一个元组，
"""

items = ['a', 'b', 'c']
from itertools import permutations

for p in permutations(items):
    print(p)
'''
('a', 'b', 'c')
('a', 'c', 'b')
('b', 'a', 'c')
('b', 'c', 'a')
('c', 'a', 'b')
('c', 'b', 'a')
'''

'''
如果你想得到指定长度的所有排列，你可以传递一个可选的长度参数
'''

for p in permutations(items, 2):
    print(p)
'''
('a', 'b')
('a', 'c')
('b', 'a')
('b', 'c')
('c', 'a')
('c', 'b')
'''

'''
使用 itertools.combinations() 可得到输入集合中元素的所有的组合。
'''
from itertools import combinations
for c in combinations(items, 3):
    print(c)
# ('a', 'b', 'c')
for c in combinations(items,2):
    print(c)
'''
('a', 'b')
('a', 'c')
('b', 'c')
'''

for c in combinations(items, 1):
    print(c)
'''
('a',)
('b',)
('c',)
'''

'''
对于 combinations() 来讲，元素的顺序已经不重要了。也就是说，组合 ('a',
'b') 跟 ('b', 'a') 其实是一样的 (最终只会输出其中一个)。
在 计 算 组 合 的 时 候， 一 旦 元 素 被 选 取 就 会 从 候 选 中 剔 除 掉 (比 如
如 果 元 素’a’ 已 经 被 选 取 了， 那 么 接 下 来 就 不 会 再 考 虑 它 了)。 而 函 数
itertools.combinations with replacement() 允许同一个元素被选择多次，
'''
"""for c in combinations_with_replacement(items, 3):
    print(c)"""


'''
这一小节我们向你展示的仅仅是 itertools 模块的一部分功能。尽管你也可以自
己手动实现排列组合算法，但是这样做得要花点脑力。当我们碰到看上去有些复杂的
迭代问题时，最好可以先去看看 itertools 模块。如果这个问题很普遍，那么很有可能
会在里面找到解决方案！
'''
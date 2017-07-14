#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第一章：数据结构和算法
Desc: Python 提供了大量的内置数据结构，包括列表，集合以及字典。大多数情况下使
用这些数据结构是很简单的。但是，我们也会经常碰到到诸如查询，排序和过滤等等
这些普遍存在的问题。因此，这一章的目的就是讨论这些比较常见的问题和算法。另
外，我们也会给出在集合模块 collections 当中操作这些数据结构的方法。

Title:         过滤序列元素
Issue:  你有一个数据序列，想利用一些规则从中提取出需要的值或者是缩短序列
代访问
answer:  最简单的过滤序列元素的方法就是使用列表推导
"""

mylist = [1, 4, -5, 10, -7, 2, 3, -1]
print([n for n in mylist if n > 0])  # [1, 4, 10, 2, 3]
print([n for n in mylist if n < 0]) # [-5, -7, -1]

'''
使用列表推导的一个潜在缺陷就是如果输入非常大的时候会产生一个非常大的结果
集，占用大量内存。如果你对内存比较敏感，那么你可以使用生成器表达式迭代产生
过滤的元素。
'''

pos = (n for n in mylist if n > 0)
print(pos) # <generator object <genexpr> at 0x0000020A54D84C50>

for x in pos:
    print(x)
'''
1
4
10
2
3
'''

'''
有时候，过滤规则比较复杂，不能简单的在列表推导或者生成器表达式中表达出
来。比如，假设过滤的时候需要处理一些异常或者其他复杂情况。这时候你可以将过
滤代码放到一个函数中，然后使用内建的 filter() 函数。
'''

values = ['1', '2', '-3', '-', '4', 'N/A', '5']

def is_int(val):
    try:
        x = int(val)
        return True
    except ValueError:
        return False

ivals = list(filter(is_int, values))
print(ivals)
# ['1', '2', '-3', '4', '5']

'''
filter() 函数创建了一个迭代器，因此如果你想得到一个列表的话，就得像示例
那样使用 list() 去转换
'''

'''
列表推导和生成器表达式通常情况下是过滤数据最简单的方式。其实它们还能在过
滤的时候转换数据。
'''

import math
print([math.sqrt(n) for n in mylist if n > 0])
# [1.0, 2.0, 3.1622776601683795, 1.4142135623730951, 1.7320508075688772]

'''
过滤操作的一个变种就是将不符合条件的值用新的值代替，而不是丢弃它们。比
如，在一列数据中你可能不仅想找到正数，而且还想将不是正数的数替换成指定的数。
通过将过滤条件放到条件表达式中去，可以很容易的解决这个问题
'''

clip_neg = [n if n > 0 else 0 for n in mylist]
print(clip_neg) # [1, 4, 0, 10, 0, 2, 3, 0]
clip_pos = [n if n < 0 else 0 for n in mylist]
print(clip_pos) # [0, 0, -5, 0, -7, 0, 0, -1]

'''
另外一个值得关注的过滤工具就是 itertools.compress() ，它以一个 iterable
对象和一个相对应的 Boolean 选择器序列作为输入参数。然后输出 iterable 对象中
对应选择器为 True 的元素。当你需要用另外一个相关联的序列来过滤某个序列的时
候，这个函数是非常有用的。
'''

addresses = [
    '5412 N CLARK',
    '5148 N CLARK',
    '5800 E 58TH',
    '2122 N CLARK'
    '5645 N RAVENSWOOD',
    '1060 W ADDISON',
    '4801 N BROADWAY',
    '1039 W GRANVILLE',
]

counts = [0, 3, 10, 4, 1, 7, 6, 1]
'''
现在你想将那些对应 count 值大于 5 的地址全部输出
'''
from itertools import compress
more5 = [n > 5 for n in counts]
print(more5)
# [False, False, True, False, False, True, True, False]
print(list(compress(addresses, more5)))
# ['5800 E 58TH', '4801 N BROADWAY', '1039 W GRANVILLE']

'''
这里的关键点在于先创建一个 Boolean 序列，指示哪些元素复合条件。然后
compress() 函数根据这个序列去选择输出对应位置为 True 的元素。
和 filter() 函数类似， compress() 也是返回的一个迭代器。因此，如果你需要
得到一个列表，那么你需要使用 list() 来将结果转换为列表类型。
'''
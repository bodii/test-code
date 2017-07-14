#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第一章：数据结构和算法
Desc: Python 提供了大量的内置数据结构，包括列表，集合以及字典。大多数情况下使
用这些数据结构是很简单的。但是，我们也会经常碰到到诸如查询，排序和过滤等等
这些普遍存在的问题。因此，这一章的目的就是讨论这些比较常见的问题和算法。另
外，我们也会给出在集合模块 collections 当中操作这些数据结构的方法。

Title:       序列中出现次数最多的元素
Issue: 怎样找出一个序列中出现次数最多的元素呢？
answer:  collections.Counter 类就是专门为这类问题而设计的，它甚至有一个有用的
most common() 方法直接给了你答案。
"""

words = [
    'look', 'into', 'my', 'eyes', 'look', 'into', 'my', 'eyes',
    'the', 'eyes', 'the', 'eyes', 'the', 'eyes', 'not', 'around', 'the',
    'eyes', "don't", 'look', 'around', 'the', 'eyes', 'look', 'into',
    'my', 'eyes', "you're", 'under'
]

from collections import Counter

word_counts = Counter(words)
print(word_counts)
# Counter({'eyes': 8, 'the': 5, 'look': 4, 'into': 3, 'my': 3, 'around': 2, "you're": 1, "don't": 1, 'not': 1, 'under': 1})

# 出现频率最高的 3 个单词
top_three = word_counts.most_common(3)
print(top_three) # [('eyes', 8), ('the', 5), ('look', 4)]

'''
作为输入， Counter 对象可以接受任意的 hashable 序列对象。在底层实现上，一
个 Counter 对象就是一个字典，将元素映射到它出现的次数上。
'''
print(word_counts['not']) # 1
print(word_counts['the']) # 5

morewords = ['why','are','you','not','looking','in','my','eyes']

w = {}
for word in morewords:
    w[word] = word_counts[word]
print(w)
# {'in': 0, 'you': 0, 'eyes': 8, 'why': 0, 'not': 1, 'looking': 0, 'my': 3, 'are': 0}

'''
或者你可以使用 update() 方法
'''
word_counts.update(morewords)
print(word_counts)
#Counter({'eyes': 9, 'the': 5, 'my': 4, 'look': 4, 'into': 3, 'not': 2, 'around': 2, 'under': 1, 'looking': 1, 'why': 1, 'in': 1, 'are': 1, "don't": 1, 'you': 1, "you're": 1})

'''
Counter 实例一个鲜为人知的特性是它们可以很容易的跟数学运算操作相结合。
'''
a = Counter(words)
b = Counter(morewords)
print(a) # Counter({'eyes': 8, 'the': 5, 'look': 4, 'into': 3, 'my': 3, 'around': 2, 'under': 1, "you're": 1, "don't": 1, 'not': 1})
print(b) # Counter({'my': 1, 'you': 1, 'looking': 1, 'eyes': 1, 'are': 1, 'not': 1, 'in': 1, 'why': 1})

c = a + b
print(c)
# Counter({'eyes': 9, 'the': 5, 'look': 4, 'my': 4, 'into': 3, 'not': 2, 'around': 2, 'looking': 1, "you're": 1, 'you': 1, 'in': 1, 'under': 1, "don't": 1, 'are': 1, 'why': 1})

d = a - b
print(d) #  Counter({'eyes': 7, 'the': 5, 'look': 4, 'into': 3, 'around': 2, 'my': 2, "don't": 1, 'under': 1, "you're": 1})

'''
 Counter 对象在几乎所有需要制表或者计数数据的场合是非常有用的
工具。在解决这类问题的时候你应该优先选择它，而不是手动的利用字典去实现。
'''
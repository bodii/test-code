#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' Counter集合 '''


"""
defaultdict类的一个最常用场景是为事件计数。

from collections import defaultdict

frequency = defaultdict(int)
for k in some_iterator():
    frequency[k] += 1
# 以上代码计算了k在some_iterator()生成的序列中出现的次数。


由于这种需求很常见，因此有一个defaultdict的等价对象提供了和上面的代码一样的功能---Counter。
但是，Conunter集合比一个简单的defaultdict类更加复杂。它可用于确定最常用值的场景，也就是统计
学家说的众数(mode)


from collections import defaultdict

by_value = defaultdict(list)
for k in frequency:
    by_value[ frequency[k] ].append(k)

# 定位并按出现频率排序并显示出的最常用值
for freq in sorted(by_value, reverse=True):
    print( by_value[freq], freq )

# 这会创建一个频率分布图，它向我们展示了一个给定频率的所有值和所有共享这些键值的频率计数。


# 上面的所有属性都是Counter集合的一部分。
from collections import Counter

frequency = Counter(some_iterator())
for k, freq in frequency.most_common():
    print( k, freq )
"""
"""
Counter集合并非仅仅是defaultdict集合的变种。一个Counter对象实现是一个“多重集合”，有时也称为
“包”。
它是一个类似set的集合，但是在包中是允许重复的。它不是一个用下标或者位置来标志元素的序列，顺序在包
中并不重要。它也不是一个键值映射。它像是一个元素就代表它们本身并且顺序无关的集合（set）。但是，它
又不是一个集合，因为，正如这个例子，元素可以重复。
"""

from collections import Counter

# 创建两个包
bag1 = Counter('aardwolves')
bag2 = Counter('zymologies')
print( bag1 )
# Counter({'a': 2, 'v': 1, 'e': 1, 'o': 1, 'r': 1, 'l': 1, 'd': 1, 'w': 1, 's': 1})

print( bag2 )
# Counter({'o': 2, 'z': 1, 'm': 1, 'i': 1, 'l': 1, 's': 1, 'e': 1, 'g': 1, 'y': 1})

print( bag1 + bag2 )
# Counter({'o': 3, 'l': 2, 'e': 2, 's': 2, 'a': 2, 'z': 1, 'w': 1, 'd': 1, 'y': 1, 'r': 1, 'g': 1, 'i': 1, 'm': 1, 'v': 1})

print( bag1 - bag2 )
# Counter({'a': 2, 'v': 1, 'r': 1, 'w': 1, 'd': 1})

print( bag2 - bag1 )
# Counter({'g': 1, 'z': 1, 'i': 1, 'y': 1, 'm': 1, 'o': 1})

bag3 = Counter('中国人民解放军')
print( bag3 )
Counter({'中': 1, '军': 1, '放': 1, '国': 1, '解': 1, '民': 1, '人': 1})

bag4 = Counter('131415926')
print( bag4 )
Counter({'1': 3, '4': 1, '2': 1, '5': 1, '3': 1, '6': 1, '9': 1})

# bag5 = Counter(123456722986)
# print( bag5 )
# TypeError: 'int' object is not iterable

bag6 = Counter(['2', '6'])
print( bag6 )
# Counter({'6': 1, '2': 1})
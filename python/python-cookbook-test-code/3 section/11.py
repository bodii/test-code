#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第三章：数字日期和时间
Desc: 在 Python 中执行整数和浮点数的数学运算时很简单的。尽管如此，如果你需要执
行分数、数组或者是日期和时间的运算的话，就得做更多的工作了。本章集中讨论的
就是这些主题。

Title;          随机选择
Issue: 你想从一个序列中随机抽取若干元素，或者想生成几个随机数。
Answer: random 模块有大量的函数用来产生随机数和随机选择元素。比如，要想从一个序
列中随机的抽取一个元素，可以使用 random.choice() 。
"""

import random
values = [1, 2, 3, 4, 5, 6]
print(random.choice(values))
# 3
print(random.choice(values))
# 4
print(random.choice(values))
# 3
print(random.choice(values))
# 4

'''
为了提取出 N 个不同元素的样本用来做进一步的操作，可以使用 random.sample()
'''
print(random.sample(values, 2))
# [2, 6]
print(random.sample(values, 3))
# [4, 5, 2]
print(random.sample(values, 2))
# [4, 5, 2]
print(random.sample(values, 4))
# [3, 1, 4, 5]

'''
如果你仅仅只是想打乱序列中元素的顺序，可以使用 random.shuffle()
'''

random.shuffle(values)
print(values)
# [1, 6, 2, 3, 5, 4]
random.shuffle(values)
print(values)
# [2, 1, 3, 5, 4, 6]

'''
生成随机整数，请使用 random.randint()
'''
print(random.randint(0, 10))
# 7
print(random.randint(0, 10))
# 4
print(random.randint(0, 10))
# 7
print(random.randint(0, 10))
# 3

'''
为了生成 0 到 1 范围内均匀分布的浮点数，使用 random.random()
'''
print(random.random())
# 0.10368946210757268
print(random.random())
# 0.7925936175007123
print(random.random())
# 0.8861123971242336

'''
如果要获取 N 位随机位 (二进制) 的整数，使用 random.getrandbits()
'''
print(random.getrandbits(200))
# 394856196584730441232903273869689770495036380233817212810827
print(random.getrandbits(100))
# 450464856787754278925780631756


'''
random 模块使用 Mersenne Twister 算法来计算生成随机数。这是一个确定性算法，
但是你可以通过 random.seed() 函数修改初始化种子。
'''
random.seed()
random.seed(12345)
random.seed(b'bytedata')

'''
除了上述介绍的功能，random 模块还包含基于均匀分布、高斯分布和其他分布的
随机数生成函数。比如， random.uniform() 计算均匀分布随机数， random.gauss()
计算正态分布随机数。对于其他的分布情况请参考在线文档。
在 random 模块中的函数不应该用在和密码学相关的程序中。如果你确实需要类似
的功能，可以使用 ssl 模块中相应的函数。比如， ssl.RAND bytes() 可以用来生成一
个安全的随机字节序列。
'''
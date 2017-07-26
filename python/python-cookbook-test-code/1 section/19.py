#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第一章：数据结构和算法
Desc: Python 提供了大量的内置数据结构，包括列表，集合以及字典。大多数情况下使
用这些数据结构是很简单的。但是，我们也会经常碰到到诸如查询，排序和过滤等等
这些普遍存在的问题。因此，这一章的目的就是讨论这些比较常见的问题和算法。另
外，我们也会给出在集合模块 collections 当中操作这些数据结构的方法。

Title:         转换并同时计算数据
Issue:  你需要在数据序列上执行聚集函数 (比如 sum() , min() , max() )，但是首先你需
要先转换或者过滤数据
代访问
answer:  一个非常优雅的方式去结合数据计算与转换就是使用一个生成器表达式参数
"""

nums = [1, 2, 3, 4, 5]
s = sum(x * x for x in nums)
print(s) # 55

import os
# files = os.listdir('dirname')
# if any(name.endswith('.py') for name in files):  # endswith() 方法用于判断 字符串是否以指定后缀结尾，返回boolean
#     print('There  be python!')
# else:
#     print('Sorry, no python.')

s = ('ACME', 50, 123.45)
print(','.join(str(x) for x in s)) # ACME,50,123.45

protfolio = [
    {'name': 'GOOG', 'shares': 50},
    {'name':'YHOO', 'shares': 75},
    {'name':'AOL', 'shares' :20},
    {'name':'SCOX', 'shares': 65}
]
min_shares = min(s['shares'] for s in protfolio)
print(min_shares) # 20

'''
上面的示例向你演示了当生成器表达式作为一个单独参数传递给函数时候的巧妙语
法 (你并不需要多加一个括号)。
s = sum((x * x for x in nums)) # 显示的传递一个生成器表达式对象
s = sum(x * x for x in nums) # 更加优雅的实现方式，省略了括号

使用一个生成器表达式作为参数会比先创建一个临时列表更加高效和优雅。
如果你不使用生成器表达式的话，你可能会考虑使用下面的实现方式：
nums = [1, 2, 3, 4, 5]
s = sum([x * x for x in nums])
这种方式同样可以达到想要的效果，但是它会多一个步骤，先创建一个额外的列
表。对于小型列表可能没什么关系，但是如果元素数量非常大的时候，它会创建一个
巨大的仅仅被使用一次就被丢弃的临时数据结构。而生成器方案会以迭代的方式转换
数据，因此更省内存。
'''

'''
在使用一些聚集函数比如 min() 和 max() 的时候你可能更加倾向于使用生成器版
本，它们接受的一个 key 关键字参数或许对你很有帮助。
'''

min_shares = min(s['shares'] for s in protfolio)
min_shares = min(protfolio, key=lambda s : s['shares'])
print(min_shares) # {'shares': 20, 'name': 'AOL'}
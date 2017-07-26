#!/usr/bin/env python3
# -*- coding:utf8 -*-


# 在列表中做出决策 -- 列表解析
'''
Python 2.0中引入了称为列表解析的特性。它可以用来在列表解除引用操作符（方
括号）中编写小的循环和判定，以此来定义用以限制被访问的元素范围的参数。
'''

# First just print even numbers
everything = range(1, 12)
print([ x for x in everything if x%2 == 0 ])

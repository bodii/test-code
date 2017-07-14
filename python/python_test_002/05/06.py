#!/usr/bin/env python3
# -*- coding:utf-8 -*-

# 为循环生成迭代器
# Python 有一个能够创建迭代器的特殊特性，即range函数

f = range(10, 20)

print(*f)

for i in range(10):
    print("Number is now %d" % i)


for i in range(5, 55, 4):
    print("Number from 5 to 55, by fours: %d" % i)


for i in range(55, 5, -4):
    print("Number from 55 to 5, by fours: %d" % i)

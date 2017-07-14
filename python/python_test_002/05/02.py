#!/usr/bin/env python3
# -*- coding:utf-8 -*-

# 为匿名函数创建名称
# 该名称只在创建名称的作用域内可用

# use lambda with filter, but bind it to a name
filter_me = [1, 2, 3, 4, 6, 7, 8, 11, 12, 14, 15, 19, 22]

# This will only return true for even numbers (because x%2 is 0, or False,
# for odd numbers)

func = lambda x: x%2 == 0
result = filter(func, filter_me)
print(*result)


#!/usr/bin/env python3
# -*- coding:utf-8 -*-

# lambda and filter: 简单匿名函数

filter_me = [1, 2, 3, 4 , 6, 7, 8, 12, 14, 15, 19, 22]
# This will only return true for even numbers(because x%2 is0, or False,
# for odd numbers)

result = filter(lambda x: x%2 == 0, filter_me)
print(result)

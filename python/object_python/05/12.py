#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' list映射和模拟映射 '''

from timeit import timeit

print(
    timeit(
    'i=k.index(10);v[i]=0',
    'k=list(range(20));v=list(range(20))'
    )
)
# 0.544813903999966

print(
    timeit(
        'm[10]=0',
        'm = dict(zip( list(range(20)), list(range(20)) ))'
    )
)
# 0.06433792999996513

"""
很明显，在两个平行的list中执行查找和更新是一个可怕的错误。它比用list.index()
定位一个元素多花的了8.6倍的时间，因为后者通过映射和哈希码来定位一个元素。
"""
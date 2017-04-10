#!/usr/bin/env python3
# -*- coding:utf-8 -*-

'''
Python3.1中，range发生了变化，它不再创建一个列表，而是改为一个迭代器，本质
上它可以以xrange的方式执行。xrange也因此从语言中删除。
range返回一个行为与列表类似的迭代器对象。注意这个对象没有公共接口。而仅有私
有方法，这些私有方法看起来类似于大多数列表和元组所拥有的方法的子集。
'''

xr = range(0, 10)
print(dir(xr))

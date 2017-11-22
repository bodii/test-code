#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 设计要素和折中方案 '''

"""
当创建不可变类时（使用__slots__, 扩展tuple,或重写属性的setter方法), 我们
不能够轻易地创建一个不可变模块。几乎没有一个模块是不可变对象。
"""

"""
from wsgiref.simple_server import make_server, demo_app

with make_server('', 8000, demo_app) as httpd:
    print('Serving HTTP on port 8000...')

    # Respond to requests until process is killed
    httpd.serve_forever()

    # Alternative: serve one request, then exit
    httpd.handle_request()
"""
class A:
    __slots__ = ['a']
    pass

a = A
print(a.a)
# <member 'a' of 'A' objects>
#!/usr/bin/env python3 
# -*- coding:utf-8 -*-

''' __del__() 和 close()方法 '''

"""
    __del__()最常见的用途是确保文件被关闭
    通常，包含文件操作的类都会有类似下面的代码：
    __del__ = close

    这会保证__del__()方法同时也是close()方法。
"""
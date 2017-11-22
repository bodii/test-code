#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 
__bool__()方法 

'''

x = object()
print(bool(x)) # True

x = ''
print(bool(x)) # False

x = {}
print(bool(x)) # False

x = ()
print(bool(x)) # False

x = []
print(bool(x)) # False

x = 0
print(bool(x)) # False



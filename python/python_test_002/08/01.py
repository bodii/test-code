#!/usr/bin/env python3
# -*- coding:utf-8 -*-


import re


s = ('xxx', 'abcxxxxabc')

a = filter((lambda s: re.match(r'xxx', s)), s)

print(*a)

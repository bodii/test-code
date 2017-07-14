#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import re

s = ('xxx', 'abcxxxxabc', 'axxxa')

b = filter(lambda s: re.search(r'xxx', s), s)

print(*b)

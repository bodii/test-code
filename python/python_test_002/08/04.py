#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import re

s = ('xxx', 'abcxxxxabc', 'axxxa', 'xyx', 'x.x', 'axxxxa')

d = filter(lambda s: re.search(r'x\.x', s), s)
print(*d)

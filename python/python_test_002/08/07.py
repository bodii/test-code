#!/usr/bin/env python3


import re

s = ('xxx', 'abcxxxxabc', 'axxxa', 'xyx', 'x.x', 'axxxxa')

g = filter(lambda s: re.search(r'c+', s), s)

print(*g)

#!/usr/bin/env python3


import re

s = ('xxx', 'abcxxxxabc', 'axxxa', 'xyx', 'x.x', 'axxxxa')

f = filter(lambda s: re.search(r'x.+x', s), s)

print(*f)

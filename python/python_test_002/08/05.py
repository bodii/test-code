#!/usr/bin/env python3


import re

s = ('xxx', 'abcxxxxabc', 'axxxa', 'xyx', 'x.x', 'axxxxa')

e = filter(lambda s: re.search(r'x.*x', s), s)
print(*e)

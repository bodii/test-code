#!/usr/bin/env python3
# -*- coding:utf-8 -*-

# use map with a list of lists, to re-order the output.
map_me_again = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
result = map(lambda list: [list[1], list[0], list[2]], map_me_again)
print(*result)

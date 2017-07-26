#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第一章：数据结构和算法
Desc: Python 提供了大量的内置数据结构，包括列表，集合以及字典。大多数情况下使
用这些数据结构是很简单的。但是，我们也会经常碰到到诸如查询，排序和过滤等等
这些普遍存在的问题。因此，这一章的目的就是讨论这些比较常见的问题和算法。另
外，我们也会给出在集合模块 collections 当中操作这些数据结构的方法。

Title:    保留最后 N 个元素
Issue: 在迭代操作或者其他操作的时候，怎样只保留最后有限几个元素的历史记录？
answer: 保留有限历史记录正是 collections.deque 大显身手的时候
"""

from collections import deque
'''
使用 deque(maxlen=N) 构造函数会新建一个固定大小的队列。当新的元素加入并
且这个队列已满的时候，最老的元素会自动被移除掉。
'''
def search(lines, pattern, history=5):
    previous_lines = deque(maxlen = history)
    for li in lines:
        if pattern in li:
            yield li, previous_lines
        previous_lines.append(li)

if __name__ == '__main__':
    with open(r'../test file/somefile.txt') as f:
        for line,prevlines in search(f,'python',5):
            for pline in prevlines:
                print(pline, end='')
            print(line, end='')
            print('-' * 20)


q = deque(maxlen=3)
q.append(1)
q.append(2)
q.append(3)
print(q) # deque([1, 2, 3], maxlen=3)
q.append(4)
print(q) # deque([2, 3, 4], maxlen=3)
q.append(5)
print(q) # deque([3, 4, 5], maxlen=3)
q.append(6)
print(q) # deque([4, 5, 6], maxlen=3)

'''
如果你不设置最大队列大小，那么就会得到一个无限大小队列，你可以在队列的两端执
行添加和弹出元素的操作
'''
q = deque()
q.append(1)
q.append(2)
q.append(3)
print(q) # deque([1, 2, 3])
q.appendleft(4)
print(q) # deque([4, 1, 2, 3])
q.pop()
print(q) # deque([4, 1, 2])
q.popleft()
print(q) # deque([1, 2])

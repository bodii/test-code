#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第四章：迭代器与生成器
Desc: 迭代是 Python 最强大的功能之一。初看起来，你可能会简单的认为迭代只不过是
处理序列中元素的一种方法。然而，绝非仅仅就是如此，还有很多你可能不知道的，
比如创建你自己的迭代器对象，在 itertools 模块中使用有用的迭代模式，构造生成器
函数等等。这一章目的就是向你展示跟迭代有关的各种常见问题。

Title;         手动遍历迭代器
Issue: 你想遍历一个可迭代对象中的所有元素，但是却不想使用 for 循环。
Answer: 通常来讲， StopIteration 用来指示迭代的结尾。然而，如果你手动使用上面演
示的 next() 函数的话，你还可以通过返回一个指定值来标记结尾，比如 None 。
"""


def manual_iter():
    with open('../test file/passwd') as f:
        try:
            while True:
                line = next(f)
                print(line, end='')
        except StopIteration:
            pass

# with open('../test file/passwd') as f:
#     while True:
#         line = next(f)
#         if line is None:
#             break
#         print(line, end='')

'''
大多数情况下，我们会使用 for 循环语句用来遍历一个可迭代对象。但是，偶尔也
需要对迭代做更加精确的控制，这时候了解底层迭代机制就显得尤为重要了。
'''
items = [1, 2, 3]
it = iter(items)
print(next(it))
# 1
print(next(it))
# 2
print(next(it))
# 3
# print(next(it))
'''
Traceback (most recent call last):
  File "python cookbook test code/4 section/1.py", line 46, in <module>
    print(next(it))
StopIteration
'''
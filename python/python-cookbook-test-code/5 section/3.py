#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第五章：文件与 IO
Desc: 所有程序都要处理输入和输出。这一章将涵盖处理不同类型的文件，包括文本和二
进制文件，文件编码和其他相关的内容。对文件名和目录的操作也会涉及到。

Title;       使用其他分隔符或行终止符打印
Issue:你想使用 print() 函数输出数据，但是想改变默认的分隔符或者行尾符
Answer: 可以使用在 print() 函数中使用 sep 和 end 关键字参数，以你想要的方式输出。
"""

print('ACME', 50, 91.5)
# ACME 50 91.5

print('ACME', 50, 91.5, sep=',')
# ACME,50,91.5

print('ACME', 50, 91.5, sep=',', end='!!\n')
'''
ACME,50,91.5!!
|
'''


'''
使用 end 参数也可以在输出中禁止换行。
'''

for i in range(5):
    print(i)
'''
0
1
2
3
4
'''

for i in range(5):
    print(i, end=' ')
#  0 1 2 3 4

'''
当你想使用非空格分隔符来输出数据的时候，给 print() 函数传递一个 seq 参数
是最简单的方案。有时候你会看到一些程序员会使用 str.join() 来完成同样的事情。
'''
row = ('ACME', 50, 91.5)
print(','.join(str(x) for x in row))
# ACME,50,91.5

'''
你当然可以不用那么麻烦，仅仅只需要像下面这样写:
'''
print(*row, sep=',')
# ACME,50,91.5


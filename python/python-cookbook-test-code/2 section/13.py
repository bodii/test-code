#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第二章：字符串和文本
Desc: 几乎所有有用的程序都会涉及到某些文本处理，不管是解析数据还是产生输出。这
一章将重点关注文本的操作处理，比如提取字符串，搜索，替换以及解析等。大部分
的问题都能简单的调用字符串的内建方法完成。但是，一些更为复杂的操作可能需要
正则表达式或者强大的解析器，所有这些主题我们都会详细讲解。并且在操作 Unicode
时候碰到的一些棘手的问题在这里也会被提及到。

Title;           字符串对齐
Issue:  你想通过某种对齐方式来格式化字符串
Answer: 对于基本的字符串对齐操作，可以使用字符串的 ljust() , rjust() 和 center()
方法。
"""

text = 'Hello World'
print(text.ljust(20))
# Hello World         |
print(text.rjust(20))
#          Hello World
print(text.center(20))
#     Hello World     |

'''
所有这些方法都能接受一个可选的填充字符
'''

print(text.rjust(20, '='))
# =========Hello World
print(text.ljust(20, '='))
# Hello World=========
print(text.center(20, '*'))
# ****Hello World*****

'''
函数 format() 同样可以用来很容易的对齐字符串。你要做的就是使用 <,> 或者
ˆ 字符后面紧跟一个指定的宽度。
'''
print(format(text, '>20'))
#          Hello World
print(format(text, '<20'))
#Hello World         |
print(format(text, '^20'))
#     Hello World     |

'''
如果你想指定一个非空格的填充字符，将它写到对齐字符的前面即可
'''

print(format(text, '=>20'))
# =========Hello World
print(format(text, '%<20'))
# Hello World%%%%%%%%%
print(format(text, '*^20'))
# ****Hello World*****


'''
当格式化多个值的时候，这些格式代码也可以被用在 format() 方法中。
'''
print('{:>10s} {:>10s}'.format('Hello', 'World'))
#      Hello      World

'''
format() 函数的一个好处是它不仅适用于字符串。它可以用来格式化任何值，使
得它非常的通用。
'''

x = 1.2345
print(format(x, '>10'))
#    1.2345
print(format(x, '*<10'))
# 1.2345****
print(format(x, '#^10.2f'))
# ###1.23###

'''
在老的代码中，你经常会看到被用来格式化文本的 % 操作符。比如：
'''
print('%-20s' % text)
# Hello World         |
print('%20s' % text)
#          Hello World


'''
但是，在新版本代码中，你应该优先选择 format() 函数或者方法。 format() 要
比 % 操作符的功能更为强大。并且 format() 也比使用 ljust() , rjust() 或 center()
方法更通用，因为它可以用来格式化任意对象，而不仅仅是字符串
'''
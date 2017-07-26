#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第三章：数字日期和时间
Desc: 在 Python 中执行整数和浮点数的数学运算时很简单的。尽管如此，如果你需要执
行分数、数组或者是日期和时间的运算的话，就得做更多的工作了。本章集中讨论的
就是这些主题。

Title;         数字的格式化输出
Issue: 你需要将数字格式化后输出，并控制数字的位数、对齐、千位分隔符和其他的细
节。
Answer:格式化输出单个数字的时候，可以使用内置的 format() 函数
"""

x = 1234.56789
print(format(x, '0.2f'))
# 1234.57
print(format(x, '>10.1f'))
#     1234.6
print(format(x, '<10.1f'))
# 1234.6    |
print(format(x, '^10.1f'))
#   1234.6  |
print(format(x, ','))
# 1,234.56789
print(format(x, '0,.1f'))
# 1,234.6

'''
如果你想使用指数记法，将 f 改成 e 或者 E(取决于指数输出的大小写形式)
'''
print(format(x, 'e'))
# 1.234568e+03
print(format(x, '0.2E'))
# 1.23E+03


'''
同时指定宽度和精度的一般形式是 '[<>ˆ]?width[,]?(.digits)?' ，其中 width
和 digits 为整数，？代表可选部分。同样的格式也被用在字符串的 format() 方法中
'''
print('The value is {:0,.2f}'.format(x))
# The value is 1,234.57

'''
当指定数字的位数后，结果值会根据 round() 函数同样的规则进行四舍五入后返
回。
'''
print(x)
# 1234.56789
print(format(x, '0.1f'))
# 1234.6
print(format(-x, '0.1f'))
# -1234.6

'''
包含千位符的格式化跟本地化没有关系。如果你需要根据地区来显示千位符，你
需要自己去调查下 locale 模块中的函数了。你同样也可以使用字符串的 translate()
方法来交换千位符。
'''
swap_separators = {
    ord('.') : ',',
    ord(',') : '.'
}
print(format(x, ',').translate(swap_separators))
# 1.234,56789

'''
在很多 Python 代码中会看到使用% 来格式化数字的
'''
print('%0.2f' % x)
# 1234.57
print('%10.1f' % x)
#     1234.6
print('%-10.1f' % x)
# 1234.6    |

'''
这种格式化方法也是可行的，不过比更加先进的 format() 要差一点。比如，在使
用% 操作符格式化数字的时候，一些特性 (添加千位符) 并不能被支持
'''
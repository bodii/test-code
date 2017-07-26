#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第三章：数字日期和时间
Desc: 在 Python 中执行整数和浮点数的数学运算时很简单的。尽管如此，如果你需要执
行分数、数组或者是日期和时间的运算的话，就得做更多的工作了。本章集中讨论的
就是这些主题。

Title;         基本的日期与时间转换
Issue: 你需要执行简单的时间转换，比如天到秒，小时到分钟等的转换。
Answer: 为了执行不同时间单位的转换和计算，请使用 datetime 模块 。
"""

from datetime import timedelta

a  = timedelta(days = 2, hours = 6)
b = timedelta(hours = 4.5)
c = a + b
print(c.days)
# 2
print(c.seconds)
# 37800
print(c.seconds / 3600)
# 10.5
print(c.total_seconds() / 3600)
# 58.5

'''
如果你想表示指定的日期和时间，先创建一个 datetime 实例然后使用标准的数学
运算来操作它们。
'''
from datetime import datetime
a = datetime(2016, 4, 10)
print(a)
# 2016-04-10 00:00:00

b = datetime(2012, 12, 21)
d = a - b
print(d.days)
# 1206

now  = datetime.today()
print(now)
# 2016-04-10 10:25:57.152811

print(now + timedelta(minutes = 10))
# 2016-04-10 10:36:34.253785


'''
在计算的时候，需要注意的是 datetime 会自动处理闰年。
'''
a = datetime(2012, 3, 1)
b = datetime(2012, 2, 28)
print(a - b)
# 2 days, 0:00:00
print((a-b).days)
# 2
c = datetime(2013, 3, 1)
d = datetime(2013, 2, 28)
print((c- d).days)
# 1

'''
对大多数基本的日期和时间处理问题， datetime 模块以及足够了。如果你需要执
行更加复杂的日期操作，比如处理时区，模糊时间范围，节假日计算等等，可以考虑
使用 dateutil 模块
许多类似的时间计算可以使用 dateutil.relativedelta() 函数代替。但是，有一
点需要注意的就是，它会在处理月份 (还有它们的天数差距) 的时候填充间隙。
'''

a = datetime(2012, 9, 23)
#print(a + timedelta(months=1))
'''
Traceback (most recent call last):
File "<stdin>", line 1, in <module>
TypeError: 'months' is an invalid keyword argument for this function
>>>


from dateutil.relativedelta import relativedelta

print(a + relativedelta(months=+1))
# datetime.datetime(2012, 10 ,23, 0, 0)
print(a = relativedelta(months=+4))
#datetime.datetime(2013, 1, 23, 0, 0)

# Time between two dates
b = datetime(2012, 12, 21)
d = b - a
print(d)
# datetime.timedelta(89)
d = relativedelta(b, a)
print(d)
# relativedelta(months=+2, days=+28)
print(d.months)
# 2
print(d.days)
# 28
'''
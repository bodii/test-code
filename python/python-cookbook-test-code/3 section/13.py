#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第三章：数字日期和时间
Desc: 在 Python 中执行整数和浮点数的数学运算时很简单的。尽管如此，如果你需要执
行分数、数组或者是日期和时间的运算的话，就得做更多的工作了。本章集中讨论的
就是这些主题。

Title;         计算最后一个周五的日期
Issue: 你需要查找星期中某一天最后出现的日期，比如星期五。
Answer: Python 的 datetime 模块中有工具函数和类可以帮助你执行这样的计算。
"""

from datetime import datetime, timedelta

weekdays = [
    'Monday', 'Tuesday', 'Wednesday', 'Thursday',
    'Friday', 'Saturday', 'Sunday'
]

def get_previous_byday(dayname, start_date=None):
    if start_date is None:
        start_date = datetime.today()
    day_num = start_date.weekday()
    day_num_target = weekdays.index(dayname)
    print(day_num_target)
    days_ago = (7 + day_num - day_num_target) % 7
    if days_ago == 0:
        days_ago = 7
    target_date = start_date - timedelta(days = days_ago)
    return target_date
'''
上面的算法原理是这样的：先将开始日期和目标日期映射到星期数组的位置上 (星
期一索引为 0)，然后通过模运算计算出目标日期要经过多少天才能到达开始日期。然
后用开始日期减去那个时间差即得到结果日期。
'''
print(datetime.today())
# 2016-04-10 10:49:00.732370
print(get_previous_byday('Monday'))
# 2016-04-04 10:49:31.626388
print(get_previous_byday('Tuesday'))
# 2016-04-05 10:50:12.080556
print(get_previous_byday('Friday'))
# 2016-04-08 10:50:58.132358

print(get_previous_byday('Sunday', datetime(2016,4, 10)))
# 2016-04-03 00:00:00

'''
如果你要像这样执行大量的日期计算的话，你最好安装第三方包 python-dateutil
来代替。比如，下面是是使用 dateutil 模块中的 relativedelta() 函数执行同样的
计算：
'''
from datetime import datetime
from dateutil.relativedelta import relativedelta
from dateutil.rrule import *
d = datetime.now()
print(d)
print(d + relativedelta(weekday=FR))

print(d + relativedelta(weekday=FR(-1)))



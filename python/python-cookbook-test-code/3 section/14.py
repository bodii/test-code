#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第三章：数字日期和时间
Desc: 在 Python 中执行整数和浮点数的数学运算时很简单的。尽管如此，如果你需要执
行分数、数组或者是日期和时间的运算的话，就得做更多的工作了。本章集中讨论的
就是这些主题。

Title;        计算当前月份的日期范围
Issue: 你的代码需要在当前月份中循环每一天，想找到一个计算这个日期范围的高效方
法。
Answer: 在这样的日期上循环并需要事先构造一个包含所有日期的列表。你可以先计算出开
始日期和结束日期，然后在你步进的时候使用 datetime.timedelta 对象递增这个日
期变量即可。
"""


'''
下面是一个接受任意 datetime 对象并返回一个由当前月份开始日和下个月开始日
组成的元组对象。
'''
from datetime import datetime , date, timedelta
import calendar

def get_month_range(start_date = None):
    if start_date is None:
        start_date = date.today().replace(day=1)
    _, days_in_month = calendar.monthrange(start_date.year, start_date.month)
    end_date = start_date + timedelta(days = days_in_month)
    return (start_date, end_date)

a_day = timedelta(days=1)
first_day, last_day = get_month_range()
while first_day < last_day:
    print(first_day)
    first_day += a_day

'''
上面的代码先计算出一个对应月份第一天的日期。一个快速的方法就是使用 date
或 datetime 对象的 replace() 方法简单的将 days 属性设置成 1 即可。 replace()
方法一个好处就是它会创建和你开始传入对象类型相同的对象。所以，如果输入参数
是一个 date 实例，那么结果也是一个 date 实例。同样的，如果输入是一个 datetime
实例，那么你得到的就是一个 datetime 实例。
然后，使用 calendar.monthrange() 函数来找出该月的总天数。任何时候只要你
想获得日历信息，那么 calendar 模块就非常有用了。 monthrange() 函数会返回包含
星期和该月天数的元组。

一旦该月的天数已知了，那么结束日期就可以通过在开始日期上面加上这个天数获
得。有个需要注意的是结束日期并不包含在这个日期范围内 (事实上它是下个月的开始
日期)。这个和 Python 的 slice 与 range 操作行为保持一致，同样也不包含结尾。
为了在日期范围上循环，要使用到标准的数学和比较操作。比如，可以利用
timedelta 实例来递增日期，小于号< 用来检查一个日期是否在结束日期之前。
理想情况下，如果能为日期迭代创建一个同内置的 range() 函数一样的函数就好
了。幸运的是，可以使用一个生成器来很容易的实现这个目标：
'''
def date_range(start, stop, step):
    while start < stop:
        yield start
        start += step

for d in date_range(datetime(2012, 9, 1), datetime(2012, 10, 1), timedelta(hours=6)):
    print(d)

'''
这种实现之所以这么简单，还得归功于 Python 中的日期和时间能够使用标准的数
学和比较操作符来进行运算。
'''
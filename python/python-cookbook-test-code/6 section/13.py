#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第六章：数据编码和处理
Desc: 这一章主要讨论使用 Python 处理各种不同方式编码的数据，比如 CSV 文件，
JSON，XML 和二进制包装记录。和数据结构那一章不同的是，这章不会讨论特殊的
算法问题，而是关注于怎样获取和存储这些格式的数据。

Title;       数据的累加与统计操作
Issue: 你需要处理一个很大的数据集并需要计算数据总和或其他统计量。
Answer: 对于任何涉及到统计、时间序列以及其他相关技术的数据分析问题，都可以考虑使
用 Pandas 库
"""

"""
为了让你先体验下，下面是一个使用 Pandas 来分析芝加哥城市的 老鼠和啮齿类动
物数据库 的例子。在我写这篇文章的时候，这个数据库是一个拥有大概 74,000 行数据
的 CSV 文件。
"""

import pandas
rats = pandas.read_csv('../test file/rats.csv', skip_footer=1)
print(rats)
'''
<class 'pandas.core.frame.DataFrame'>
Int64Index: 74055 entries, 0 to 74054
Data columns:
Creation Date 74055 non-null values
Status 74055 non-null values
Completion Date 72154 non-null values
Service Request Number 74055 non-null values
Type of Service Request 74055 non-null values
Number of Premises Baited 65804 non-null values
Number of Premises with Garbage 65600 non-null values
Number of Premises with Rats 65752 non-null values
Current Activity 66041 non-null values
Most Recent Action 66023 non-null values
Street Address 74055 non-null values
ZIP Code 73584 non-null values
X Coordinate 74043 non-null values
Y Coordinate 74043 non-null values
Ward 74044 non-null values
Police District 74044 non-null values
Community Area 74044 non-null values
Latitude 74043 non-null values
Longitude 74043 non-null values
Location 74043 non-null values
dtypes: float64(11), object(9)
'''

print(rats['Current Activity'].unique())

crew_dispatched = rats[rats['Current Activity'] == 'Dispathch Crew']
print(len(crew_dispatched))

print(crew_dispatched['ZIP Code'].value_counts()[:10])
dates = crew_dispatched.groupby('Completion Date')
print(dates)

date_counts = dates.size()
print(date_counts)

date_counts.sort()
print(date_counts[-10:])

"""
嗯，看样子 2011 年 10 月 7 日对老鼠们来说是个很忙碌的日子啊！ˆ ˆ
Pandas 是一个拥有很多特性的大型函数库，我在这里不可能介绍完。但是只要你
需要去分析大型数据集合、对数据分组、计算各种统计量或其他类似任务的话，这个
函数库真的值得你去看一看。

"""
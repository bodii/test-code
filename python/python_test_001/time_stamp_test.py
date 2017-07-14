#!/usr/bin/env python
# -*- coding:utf-8 -*-


"""
Topic:根据约定的日期和时间，返回对应的UTC时间戳
Desc:
"""


import re
from datetime import datetime, timedelta, timezone
def to_timestamp(dt_str=None, tz_str=None):
    t_re = r'[\d]{4}\-(\d|[0-1]\d)\-(\d|[0-2]\d|3[0-1])\s+?(2[0-4]|[0-1]\d|\d)[\:]([0-5]\d|\d)[\:]([0-5]\d|\d)'
    utc_re = r'UTC([\-|\+])((2[0-4]|[0-1]\d|\d):([0-5]\d|\d))'
    if dt_str is not None:
        if re.match(t_re,dt_str):
            if tz_str is not None:
                if re.match(utc_re,tz_str):
                    utc_list = re.match(utc_re,tz_str).groups()
                    if len(utc_list) == 4:
                        hours_num = int(utc_list[0]+utc_list[2])
                        t = datetime.strptime(dt_str,'%Y-%m-%d %H:%M:%S')
                        utc_time = timezone(timedelta(hours=hours_num))
                        ts = t.replace(tzinfo=utc_time)
                        ts = datetime.timestamp(ts)
                        print(ts)

                else:
                    print('UTC格式不正确')
            else:
                t = datetime.strptime(dt_str,'%Y-%m-%d %H:%M:%S')
                utc_time = timezone(timedelta(hours=8))
                ts = t.replace(tzinfo=utc_time)
                ts = datetime.timestamp(ts)
                print(ts)
        else:
            print('时间格式不正确！')
    else:
        print('缺少必要的参数，请确保传入了日期与时间值')
        
        
# 以下为测试内容

t1 = to_timestamp('2015-6-1 08:10:30', 'UTC+7:00')

>>> 1433121030.0

t2 = to_timestamp('2015-5-31 16:10:30', 'UTC-09:00')

>>> 1433121030.0

t3 = to_timestamp('2015-6-01 9:10:30')

>>> 1433121030.0

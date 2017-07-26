#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第六章：数据编码和处理
Desc: 这一章主要讨论使用 Python 处理各种不同方式编码的数据，比如 CSV 文件，
JSON，XML 和二进制包装记录。和数据结构那一章不同的是，这章不会讨论特殊的
算法问题，而是关注于怎样获取和存储这些格式的数据。

Title;      读写 CSV 数据
Issue:你想读写一个 CSV 格式的文件。
Answer: 对于大多数的 CSV 格式的数据读写问题，都可以使用 csv 库
"""

import csv

'''为了写入 CSV 数据，你仍然可以使用 csv 模块，不过这时候先创建一个 writer 对
象。'''

headers = ['Symbol','Price','Date','Time','Change','Volume']
rows = [
    ('AA', 39.48, '6/11/2007', '9:36am', -0.18, 181800),
    ('AIG', 71.38, '6/11/2007', '9:36am', -0.15, 195500),
    ('AXP', 62.58, '6/11/2007', '9:36am', -0.46, 935000),
]

with open('../test file/stocks.csv', 'w', newline='') as f:
    f_csv = csv.writer(f)
    f_csv.writerow(headers)
    f_csv.writerows(rows)

'''
如果你有一个字典序列的数据，可以像这样做
'''
headers = ['Symbol', 'Price', 'Date', 'Time', 'Change', 'Volume']
rows = [
    {'Symbol':'AA', 'Price':39.48, 'Date':'6/11/2007',
    'Time':'9:36am', 'Change':-0.18, 'Volume':181800},
    {'Symbol':'AIG', 'Price': 71.38, 'Date':'6/11/2007',
    'Time':'9:36am', 'Change':-0.15, 'Volume': 195500},
    {'Symbol':'AXP', 'Price': 62.58, 'Date':'6/11/2007',
    'Time':'9:36am', 'Change':-0.46, 'Volume': 935000},
    ]

with open('../test file/stocks2.csv', 'w', newline='') as f:
    f_csv = csv.DictWriter(f, headers)
    f_csv.writeheader()
    f_csv.writerows(rows)

'''
下面向你展示如何将这些数据读取为一个元组的序列
'''
with open('../test file/stocks.csv') as f:
    f_csv = csv.reader(f)
    headers = next(f_csv)
    rows = []
    for row in f_csv:
        if row != []:
            rows.append(tuple(row))
    print(headers)
    print(rows)
'''
['Symbol', 'Price', 'Date', 'Time', 'Change', 'Volume']
[('AA', '39.48', '6/11/2007', '9:36am', '-0.18', '181800'), ('AIG', '71.38', '6/11/2007', '9:36am', '-0.15', '195500'), ('AXP', '62.58', '6/11/2007', '9:36am', '-0.46', '935000')]
'''

'''
在上面的代码中， row 会是一个元组。因此，为了访问某个字段，你需要使用下
标，如 row[0] 访问 Symbol， row[4] 访问 Change。
由于这种下标访问通常会引起混淆，你可以考虑使用命名元组。
'''

from collections import namedtuple
with open('../test file/stocks2.csv') as f:
    f_csv = csv.reader(f)
    headings = next(f_csv)
    Row = namedtuple('Row', headings)
    for r in f_csv:
        if r!=[]:
            row = Row(*r)
    print(row)
'''
Row(Symbol='AXP', Price='62.58', Date='6/11/2007', Time='9:36am', Change='-0.46', Volume='935000')
'''

'''
它允许你使用列名如 row.Symbol 和 row.Change 代替下标访问。需要注意的是这
个只有在列名是合法的 Python 标识符的时候才生效。如果不是的话，你可能需要修改
下原始的列名 (如将非标识符字符替换成下划线之类的)。
另外一个选择就是将数据读取到一个字典序列中去。可以这样做：
'''
with open('../test file/stocks2.csv','r') as f:
    f_csv = csv.DictReader(f)
    for row in f_csv:
        if row!=[]:
            print(row)
'''
{'Price': '39.48', 'Volume': '181800', 'Change': '-0.18', 'Date': '6/11/2007', 'Symbol': 'AA', 'Time': '9:36am'}
{'Price': '71.38', 'Volume': '195500', 'Change': '-0.15', 'Date': '6/11/2007', 'Symbol': 'AIG', 'Time': '9:36am'}
{'Price': '62.58', 'Volume': '935000', 'Change': '-0.46', 'Date': '6/11/2007', 'Symbol': 'AXP', 'Time': '9:36am'}
'''

"""
你应该总是优先选择 csv 模块分割或解析 CSV 数据。例如，你可能会像编写类似
下面这样的代码：
with open('stocks.csv') as f:
for line in f:
row = line.split(',')
# process row
...
使用这种方式的一个缺点就是你仍然需要去处理一些棘手的细节问题。比如，如果
某些字段值被引号包围，你不得不去除这些引号。另外，如果一个被引号包围的字段
碰巧含有一个逗号，那么程序就会因为产生一个错误大小的行而出错。
默认情况下，csv 库可识别 Microsoft Excel 所使用的 CSV 编码规则。这或许也是
最常见的形式，并且也会给你带来最好的兼容性。然而，如果你查看 csv 的文档，就
会发现有很多种方法将它应用到其他编码格式上 (如修改分割字符等)。例如，如果你
想读取以 tab 分割的数据，可以这样做：
# Example of reading tab-separated values
with open('stock.tsv') as f:
f_tsv = csv.reader(f, delimiter='\t')
for row in f_tsv:
# Process row
...
如果你正在读取 CSV 数据并将它们转换为命名元组，需要注意对列名进行合法性
认证。例如，一个 CSV 格式文件有一个包含非法标识符的列头行，类似下面这样：
这样最终会导致在创建一个命名元组时产生一个 ValueError 异常而失败。为了解
决这问题，你可能不得不先去修正列标题。例如，可以像下面这样在非法标识符上使
用一个正则表达式替换：
import re
with open('stock.csv') as f:
f_csv = csv.reader(f)
headers = [ re.sub('[^a-zA-Z_]', '_', h) for h in next(f_csv) ]
Row = namedtuple('Row', headers)
for r in f_csv:
row = Row(*r)
# Process row
还有重要的一点需要强调的是，csv 产生的数据都是字符串类型的，它不会做任何
其他类型的转换。如果你需要做这样的类型转换，你必须自己手动去实现。下面是一
个在 CSV 数据上执行其他类型转换的例子：
col_types = [str, float, str, str, float, int]
with open('stocks.csv') as f:
f_csv = csv.reader(f)
headers = next(f_csv)
for row in f_csv:
# Apply conversions to the row items
row = tuple(convert(value) for convert, value in zip(col_types, row))
...
另外，下面是一个转换字典中特定字段的例子：
print('Reading as dicts with type conversion')
field_types = [ ('Price', float),
('Change', float),
('Volume', int) ]
with open('stocks.csv') as f:
for row in csv.DictReader(f):
row.update((key, conversion(row[key]))
for key, conversion in field_types)
print(row)
通常来讲，你可能并不想过多去考虑这些转换问题。在实际情况中，CSV 文件都
或多或少有些缺失的数据，被破坏的数据以及其它一些让转换失败的问题。因此，除
非你的数据确实有保障是准确无误的，否则你必须考虑这些问题 (你可能需要增加合适
的错误处理机制)。
最后，如果你读取 CSV 数据的目的是做数据分析和统计的话，你可能需要看一看
Pandas 包。Pandas 包含了一个非常方便的函数叫 pandas.read csv() ，它可以加载
CSV 数据到一个 DataFrame 对象中去。然后利用这个对象你就可以生成各种形式的统
计、过滤数据以及执行其他高级操作了。在 6.13 小节中会有这样一个例子。
"""
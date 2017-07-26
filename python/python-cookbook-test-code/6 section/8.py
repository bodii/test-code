#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第六章：数据编码和处理
Desc: 这一章主要讨论使用 Python 处理各种不同方式编码的数据，比如 CSV 文件，
JSON，XML 和二进制包装记录。和数据结构那一章不同的是，这章不会讨论特殊的
算法问题，而是关注于怎样获取和存储这些格式的数据。

Title;      与关系型数据库的交互
Issue: 你想在关系型数据库中查询、增加或删除记录。
Answer: jPython 中表示多行数据的标准方式是一个由元组构成的序列。
"""

stocks = [
    (1,'GOOG', 100, 490.1),
    (2,'AAPL', 50, 545.75),
    (3,'FB', 150, 7.45),
    (4,'HPQ', 75, 33.2),
]

"""
依据 PEP249，通过这种形式提供数据，可以很容易的使用 Python 标准数据库
API 和关系型数据库进行交互。所有数据库上的操作都通过 SQL 查询语句来完成。每
一行输入输出数据用一个元组来表示。
为了演示说明，你可以使用 Python 标准库中的 sqlite3 模块。如果你使用的是一
个不同的数据库 (比如 MySql、Postgresql 或者 ODBC)，还得安装相应的第三方模块
来提供支持。不过相应的编程接口几乎都是一样的，除了一点点细微差别外。
第一步是连接到数据库。通常你要执行 connect() 函数，给它提供一些数据库名、
主机、用户名、密码和其他必要的一些参数。
"""

import sqlite3
db = sqlite3.connect('../test file/database.db')
"""
为了处理数据，下一步你需要创建一个游标。一旦你有了游标，那么你就可以执行
SQL 查询语句了。
"""
c = db.cursor()
#print(c.execute("create table portfolio (id int unsigned  auto_increment PRIMARY KEY unique,symbol text, shares integer, price float)"))
# <sqlite3.Cursor object at 0x000001DABBDAFF80>
#db.commit()
#print(c.executemany("insert into portfolio values (?,?,?,?)", stocks))
#db.commit()

for row in db.execute("select * from portfolio"):
    print(row)

"""
如果你想接受用户输入作为参数来执行查询操作，必须确保你使用下面这样的占位
符 ‘‘?‘‘来进行引用参数
"""
min_price = 100
for row in db.execute('select * from portfolio where price >= ?', min_price):
    print(row)

"""
在比较低的级别上和数据库交互是非常简单的。你只需提供 SQL 语句并调用相应
的模块就可以更新或提取数据了。虽说如此，还是有一些比较棘手的细节问题需要你
逐个列出去解决。
一个难点是数据库中的数据和 Python 类型直接的映射。对于日期类型，通常可以
使用 datetime 模块中的 datetime 实例，或者可能是 time 模块中的系统时间戳。对
于数字类型，特别是使用到小数的金融数据，可以用 decimal 模块中的 Decimal 实例
来表示。不幸的是，对于不同的数据库而言具体映射规则是不一样的，你必须参考相
应的文档。
另外一个更加复杂的问题就是 SQL 语句字符串的构造。你千万不要使用 Python
字符串格式化操作符 (如%) 或者 .format() 方法来创建这样的字符串。如果传递给这
些格式化操作符的值来自于用户的输入，那么你的程序就很有可能遭受 SQL 注入攻击
(参考 http://xkcd.com/327 )。查询语句中的通配符 ? 指示后台数据库使用它自己的字
符串替换机制，这样更加的安全。
不幸的是，不同的数据库后台对于通配符的使用是不一样的。大部分模块使用 ?
或 %s ，还有其他一些使用了不同的符号，比如:0 或:1 来指示参数。同样的，你还是得
去参考你使用的数据库模块相应的文档。一个数据库模块的 paramstyle 属性包含了参
数引用风格的信息。
对于简单的数据库数据的读写问题，使用数据库 API 通常非常简单。如果你要处
理更加复杂的问题，建议你使用更加高级的接口，比如一个对象关系映射 ORM 所提
供的接口。类似 SQLAlchemy 这样的库允许你使用 Python 类来表示一个数据库表，并
且能在隐藏底层 SQL 的情况下实现各种数据库的操作。
"""

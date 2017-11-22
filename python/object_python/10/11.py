#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 从Python对象到SQLite BLOB列的映射 '''

"""
SQLite中包含了一个二进制大对象（Binary Large Object， BLOB）数值类型。我们可以使用
pickle来处理Python对象，然后将它们存入BLOB列中。可以使用字符串来表示Python对象（例如，
使用JSON或YAML格式），也可以使用SQLite中的文本列。

当需要处理金融数据时，在应用程序中需要使用decimal.Decimal数值。
"""

import decimal
import sqlite3
import os.path as opt

current_path = opt.dirname(__file__) + opt.sep

def adapt_currenchy( value ):
    return str(value)

# 转换与适配
sqlite3.register_adapter(decimal.Decimal, adapt_currenchy)

def convert_currency( bytes ):
    return decimal.Decimal(bytes.decode())

sqlite3.register_converter('DECIMAL', convert_currency)

"""
一个adapt_currency()函数，用于完成将decimal.decimal对象适配为数据库中适当的形式。
由于已经注册了适配函数，因此SQLite接口就能够使用所注册的适配函数来完成decimal.Decimal类
对象的转换逻辑了。
convert_currency()函数，用于从SQLite字节对象转换为Python中的decimal.Decimal对象。由
于注册了转换函数，因此DECIMAL类型的列就能被适当地转换为Python中的对象。
一旦定义了适配器和转换器，我们就能将DECIMAL看作一种被完全支持的列类型。除此之外，还要在建立
数据库连接时通过设置detect_types=sqlite3. PARSE_DECLTYPES来通知SQLite。

例如：
"""

decimal_dd1 = """
create table budget(
    year integer,
    month integer,
    category text,
    amount decimal
)
"""

filename = current_path + 'p2_c11_blog.db'
database = sqlite3.connect(filename, detect_types=sqlite3.PARSE_DECLTYPES)
database.execute(decimal_dd1)

insert_budget = """
insert into budget(year, month, category, amount) 
values(:year, :month, :category, :amount)
"""

database.execute( 
    insert_budget, 
    dict(
        year = 2017,
        month = 1,
        category='fuel',
        amount=decimal.Decimal('256.78')
    )
)

database.execute(
    insert_budget,
    dict(
        year = 2017,
        month = 2,
        category = 'fuel',
        amount = decimal.Decimal('287.65')
    )
)

query_budget = """
select * from budget
"""

for row in database.execute(query_budget):
    print( row )

database.close()
#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 用SQLite保存和获取对象 '''

# SQL数据模型 -- 行和表
"""
SQLite支持NULL、INTEGER、REAL、TEXT、BLOB数据。
对应Python中的类型为None、int、float、str和bytes。

在sqlite3模块中加入了datetime.date和datetime.datetime。
"""
"""
    创建在硬盘上面： conn = sqlite3.connect('c:\\test\\test.db')
    创建在内存上面： conn = sqlite3.connect(":memory:")
    conn.create_function("md5", 1, md5sum)

    下面我们一硬盘上面创建数据库文件为例来具体说明：
    conn = sqlite3.connect('c:\\test\\hongten.db')
    其中conn对象是数据库链接对象，而对于数据库链接对象来说，具有以下操作：
        sqlite3.execute('BEGIN')  --开启事务
        commit()            --事务提交
        rollback()          --事务回滚
        conn.close()        --关闭一个数据库链接
        cursor()            --创建一个游标

    cu = conn.cursor()
    这样我们就创建了一个游标对象：cu
    在sqlite3中，所有sql语句的执行都要在游标对象的参与下完成
    对于游标对象cu，具有以下具体操作：

        execute()           --执行一条sql语句
        executemany()       --执行多条sql语句
        executescript()     --执行多条sql语句
        cu.close()          --游标关闭
        fetchone()          --从结果中取出一条记录
        fetchmany()         --从结果中取出多条记录
        fetchall()          --从结果中取出所有记录
        scroll()            --游标滚动

sqlite3.connect('file:path/to/database?mode=ro', uri=True)
"""
"""
SQL语句可以被分为3类：数据定义语言（data definition language, DDL)、数据操纵语言（data
manipulation language, DML) 和数据控制语言(data control language, DCL)。DDL运用于
对数据表、其中的列以及索引进行定义。

DDL的例子：
"""

sql_dd1 = """\
create table Blog (
    id integer primary key autoincrement,
    title text
);
create table Post (
    id integer primary key autoincrement,
    date timestamp,
    title text,
    rst_text text,
    blog_id integer references Blog(id)
);
create table Tag (
    id integer primary key autoincrement,
    phrase text unique on conflict fail
);
create table Assoc_Post_Tag (
    post_id integer references Post(id),
    tag_id integer references Tag(id)
);
"""

import sqlite3
import os.path as opt

current_path = opt.dirname(__file__) + opt.sep
filename = current_path + 'p2_c11_blog.db'
database = sqlite3.connect(filename)

# sql语名创建表集
# database.executescript( sql_dd1 )

# cursor操作数据库
# crsr = database.cursor() # 光标，游标
# for stmt in sql_dd1.split(";"):
#     crsr.execute(stmt)

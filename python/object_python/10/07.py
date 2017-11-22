#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import sqlite3
import os.path as opt

current_path = opt.dirname(__file__) + opt.sep
filename = current_path + 'p2_c11_blog.db'

database = sqlite3.connect(filename, isolation_level='DEFERRED') # deferred
try:
    # 开启事务
    database.execute('BEGIN') 
    database.execute('some statement')
    database.execute('another statement')
    # 提交事务
    database.comment() 
except Exception as e :
    # 事务执行失败，则回滚
    database.rollback()
    # 抛出异常 
    raise e 
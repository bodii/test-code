#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' SQLite update '''

import sqlite3
import os.path as opt

current_path = opt.dirname(__file__) + opt.sep
filename = current_path + 'p2_c11_blog.db'

database = sqlite3.connect(filename)

update_blog = """
update Blog set title=:new_title where title=:old_title
"""
database.execute('BEGIN') # 事务
database.execute(
    update_blog, 
    dict(new_title='2016-2017 Travel', old_title='Travel Blog')
    )
database.commit() # 提交
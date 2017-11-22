#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用SQL的DML语句完成CRUD '''

import sqlite3
import os.path as opt

current_paht = opt.dirname(__file__) + opt.sep
filename = current_paht + 'p2_c11_blog.db'

database = sqlite3.connect(filename)

create_blog = """
insert into Blog(title) values(?)
"""
database.execute(create_blog, ('Travel Blog',))
database.close()




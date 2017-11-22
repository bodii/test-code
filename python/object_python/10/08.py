#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import sqlite3
import os.path as opt

current_path = opt.dirname(__file__) + opt.sep
filename = current_path + 'p2_c11_blog.db'

database = sqlite3.connect(filename, isolation_level='DEFERRED')
with database: # 上下文会完成事务的工作
    database.execute('some statement')
    database.execute('another statement')
    # 在with上下文的最后，database.commit()语句会自动提交。当发生异常时，database.rollback()
    # 会被执行，然后异常会从with语句中的抛出
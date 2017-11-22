#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' SQLite select '''

import sqlite3
import os.path as opt

current_path = opt.dirname(__file__) + opt.sep
filename = current_path + 'p2_c11_blog.db'

database = sqlite3.connect(filename)

query_blog_by_title = """
select * from Blog where title=?
"""

# for blog in database.execute(query_blog_by_title, ('2016-2017 Travel',)):
#     print( blog[0], blog[1] )

crsr = database.cursor()
crsr.execute( query_blog_by_title, ('2016-2017 Travel',))
for blog in crsr.fetchall():
    print( blog[0], blog[1] )

query_blog = """
select * from blog
"""

result = database.execute(query_blog)
for blog in result:
    print( blog[0], blog[1] )
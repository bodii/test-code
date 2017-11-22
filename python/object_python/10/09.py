#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 设计数据库中的主键和外键 '''
"""
设计表关系时，有3种设计方法：
1. 1对多
2. 多对多
3. 1对1

基本表关系可以通过以下一种或两种方式来实现：
1. 显式
2. 隐式
"""


''' 使用SQL处理程序中的数据 '''

import sqlite3
import os.path as opt

current_path = opt.dirname(__file__) + opt.sep
filename = current_path + 'p2_c11_blog.db'

database = sqlite3.connect(filename)

query_blog_by_title = """
select * from blog where title=?
"""

query_post_by_blog_id = """
select * from post where blog_id=?
"""

query_tag_by_post_id = """
select tag.*
from tag
join assoc_post_tag
on tag.id = assoc_post_tag.tag_id
where assoc_post_tag.post_id=?
"""

for blog in database.execute( query_blog_by_title, ('2016-2017 Travel',)) :
    print( 'Blog', blog)
    for post in database.execute( query_post_by_blog_id, (blog[0], )):
        print( 'Post', post)
        for tag in database.execute( query_tag_by_post_id, (post[0], )):
            print( 'Tag', tag )


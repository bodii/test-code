#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' SQLite delete '''

import sqlite3
import os.path as opt

current_path = opt.dirname(__file__) + opt.sep
filename = current_path + 'p2_c11_blog.db'

database = sqlite3.connect(filename)

delete_post_tag_by_blog_title = """
delete from Assoc_Post_Tag 
where post_id in (
    select distinct post_id 
    from Blog
    join Post
    on Blog.id = Post.Blog_id
    where Blog.title = :old_title
)
"""

delete_post_by_blog_title = """
delete from Post where blog_id in (
    select id from Blog where title = :old_title
)
"""

delete_blog_by_title = """
delete from Blog where title = :old_title
"""

try:
    with database:
        title = dict(old_title='2016-2017 Travel')
        database.execute( delete_post_tag_by_blog_title, title )
        database.execute( delete_post_by_blog_title, title )
        database.execute( delete_blog_by_title, title )
    print( 'Delete finished normally.' )
except Exception as e:
    print( 'Rolled Back due to {0}'.fromat(e) )
#!/usr/bin/env python3
# -*- coding:utf-8 -*-


''' 在纯SQL中实现类似于类的处理方式 '''

import sqlite3
import os.path as opt
from collections import defaultdict

current_path = opt.dirname(__file__) + opt.sep
filename = current_path + 'p2_c11_blog.db'

database = sqlite3.connect(filename)

class Blog:
    def __init__( self, title, *posts ):
        self.title = title
        self.entries = list(posts)
    
    def append( self, post ):
        self.entries.append(post)
    
    def by_tag( self ):
        tag_index = defaultdict(list)
        for post in self.entries:
            for tag in post.tags:
                tag_index[tag].append( post )
        return tag_index

    def as_dict( self ):
        return dict(
            title = self.title,
            underline =  "=" * len(self.title),
            entries = [ p.as_dict() for p in self.entries ],
        )

query_by_tag = """
select tag.phrase, post.title, post.id
from tag
join Assoc_Post_Tag on tag.id = Assoc_Post_Tag.tag_id
join post on post.id = Assoc_Post_Tag.post_id
join Blog on post.blog_id = blog.id
where Blog.title=?
"""
result = database.execute(query_by_tag, ('2016-2017 Travel',))
tag_index = defaultdict(list)
for tag,post_title, post_id in result:
    tag_index[tag].append((post_title, post_id))

print( tag_index )
database.close()
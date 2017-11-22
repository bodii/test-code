#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 手动完成从Python对象到数据库中行的映射 '''

# 以下是一个Blog类的定义：
from collections import defaultdict

class Blog:
    def __init__( self, **kw ):
        """Requires title"""
        self.id = kw.pop('id', None)
        self.title = kw.pop('title', None)
        if kw : raise TooManyValues( kw )
        self.entries = list()

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
            underline = '=' * len(self.title),
            entries = [ p.as_dict() for p in self.entries ],
        )

    
# 以下是一个Post类的定义
import datetime

class Post:
    def __init__( self, **kw ):
        """Requires date, title, rst_text."""
        self.id = kw.pop('id', None)
        self.date = kw.pop('date', None)
        self.title = kw.pop('title', None)
        self.rst_text = kw.pop('rst_text', None)
        self.tags = list()
        if kw : raise TooManyValues(kw)

    def append( self, tag ):
        self.tags.append(tag)

    def as_dict( self ):
        return dict(
            date = str(self.date),
            title = self.title,
            underline = '='* len(title),
            rst_text = self.rst_text,
            tag_text = ' '.join(self.tags),
        )

# 定义异常类
class TooManyValues( Exception ):
    pass

import sqlite3

''' 为SQLiter设计一个访问层 '''
class Access:
    
    def open( self, filename ):
        self.database = sqlite3.connect(filename)
        # 设置row_factory为sqlite3.Row类，而不是一个简单的元组。
        # Row类允许通过数字索引和列名来访问。
        self.database.row_factory = sqlite3.Row
    
    def get_blog( self, id ):
        query_blog = """
        select * from blog where id=?
        """
        row = self.database.execute(query_blog,(id,)).fetchone()
        blog = Blog( id = row['id'], title = row['title'] )
        return blog

    def add_blog( self, blog ):
        insert_blog = """
        insert into Blog(title) values(:title)
        """
        self.database.execute(insert_blog, dict(title=blog,title))
        blog.id = self.get_last_id()
        return blog

    def get_last_id( self ):
        get_last_id = """
        select last_insert_rowid()
        """
        row = self.database.execute(get_last_id).fetchone()
        return row[0]

    # 从数据库中获取一个单独的Post对象
    def get_post( self, id ):
        query_post = """
        select * from post where id=?
        """
        row = self.database.execute(query_post, (id,)).fetchone()
        post = Post(id=row['id'], title=row['title'],
        date=row['date'], rst_text=row['rst_text'])

        query_tags = """
        select tag.*
        from tag join assoc_post_tag on tag.id=assoc_post_tag.tag_id
        where assoc_post_tag.post_id=?
        """
        results = self.database.execute(query_tags, (id,))
        for id, tag in results:
            post.append(tag)
        return post

    def add_post( self, blog, post ):
        insert_post = """
        insert into post(title, date, rst_test, blog_id)
        values (:title, :date, :rst_test, :blog_id)
        """

        query_tag = """
        select * from tag where phrase=?
        """

        insert_tag = """
        insert into tag(phrase) values (?)
        """

        insert_association = """
        insert into assco_post_tag(post_id, tag_id)
        values (:post_id, :tag_id)
        """

        with self.database:
            self.database.execute(insert_post,
                dict(
                    title = post.title,
                    date = post.date,
                    rst_text = post.rst_text,
                    blog_id = blog.id
                )
            )
            post.id = self.get_last_id()
            for tag in post.tags:
                tag_row = self.database.execute(query_tag, (tag,)).fetchone()
                if tag_row is not None:
                    tag_id = tag_row['id']
                else:
                    self.database.execute(insert_tag, (tag,))
                    tag_id = self.get_last_id()
                    self.database.execute(insert_association, 
                        dict(tag_id=tag_id, post_id=post_id)
                    )
        
        return post

    def blog_iter( self ):
        query = """
        select * from blog
        """

        results = self.database.execute(query)
        for row in results:
            blog = Blog( id=row['id'], title=row['title'])
            yield blog

    def post_iter( self, blog ):
        query = """
        select id from post where blog_id=?
        """

        results = self.database.execute(query, (blog.id,))
        for row in results:
            yield self.get_post(row['id'])

    
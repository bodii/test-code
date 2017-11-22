#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 为shelve设计数据访问层 '''

"""
应用程序中修改和保存博客文章的部分，将分成两层：应用层和数据层。其中，又将应用层分成两层。
1. 应用处理(Application processing):这些对象不是持久的。这些类会包括程序中所有的行
为。这些类会响应用户选择的命令、菜单项、按钮和其他处理元素。
2. 问题域数据模型（Problem domain data model）: 这些对象会被保存到shelf上。这些对象
包含了程序的所有状态。

在数据层中，依据数据存储的复杂度，可能会有若干特性。我们只关注其中的两个：
1. 访问（Access）：这些组件为问题域中的对象提供了统一的访问方法。我们会定义一个提供访问Blog
和Post实例的的方法的Access类，它也会管理用于定位shelf上的Blog和Post对象的键。
2. 持久化（Persistence）：这些组件将问题域对象序列化之后保存到持久化存储模块中。这是shelve
模块。
"""

import shelve

class Access:
    def new( self, filename ):
        self.database = shelve.open( filename, 'n')
        self.max = {'Post': 0, 'Blog': 0}
        self.sync()

    def open( self, filename ):
        self.database = shelve.opne(filename, 'w')
        self.max = self.database['_DB:max']

    def close( self ):
        if self.database:
            self.database['_DB:max'] = self.max
            self.database.close()

    def sync( self ):
        self.database['_DB:max'] = self.max
        self.database.sync()

    def quit( self ):
        self.close()

    # 用来更新shelf上的Blog和Post对象的一些不同的方法。
    def add_blog( self, blog):
        self.max['Blog'] += 1
        key = 'Blog:{id}'.format(id = self.max['Blog'])
        blog._id = key
        self.database[blog._id] = blog
        return blog

    def get_blog( self, id ):
        return self.database[id]

    def add_post( self, blog, post ):
        self.max['post'] += 1
        try:
            key = '{blog}:Post:{id}'.format(blog = blog._id, id = self.max['post'])
        except AttributeError:
            raise OperationError( 'Blog not added' )
        post._id = key
        self.database[post._id] = post
        return post

    def get_post( self, id ):
        return self.database[id]
    
    def replace_post( self, post ):
        self.database[post._id] = post
        return post

    def delete_post( self, post ):
        del self.database[post._id]

    # 迭代器使用的搜索方法，用来查询Blog和Post实例
    def __iter__( self ):
        for k in self.database:
            if k[0]  == '_' : continue
            yield self.database[k]

    def blog_iter( self ):
        for k in self.database:
            if not k.startswith('Blog:') : continue
            if ':Post:' in k : continue # Skip children
            yield self.database[k]

    def post_iter( self, blog ):
        key = '{blog}:Post:'.format(blog = blog._id)
        for k in self.database:
            if not k.startswith(key) : continue
            yield self.database[k]

    def title_iter( self, blog, title ):
        return ( p for p in self.post_iter(blog) if p.title == title )


# 编写演示脚本

from contextlib import closing

with closing( Access() ) as access:
    access.new( 'blog' )
    access.add_blog( b1 )
    # b1._id is set
    for post in p2, p3:
        access.add_post( b1, post )
    b = access.get_blog( b1._id )
    print( b._id, b )
    for p in access.post_iter( b ):
        print( p._id, p )
    access.quit()
    
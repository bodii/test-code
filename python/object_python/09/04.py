#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 用索引提高性能 '''

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


"""
每个Blog中增加了一个用于定位所有相关Post的索引。
"""
class Access2( Access ):
    def add_blog( self, blog ):
        self.max['Blog'] += 1
        key = 'Blog:{id}'.format(id=self.max['Blog'])
        blog._id = key
        blog._post_list = []
        self.database[blog._id] = blog
        return blog

    def add_post( self, blog, post ):
        self.max['Post'] += 1
        try:
            key = '{blog}:Post:{id}'.format(blog = blog._id, id = self.max['Post'])
        except AttributeError:
            raise OperationError( 'Blog not added' )
        post._id = key
        post._blog = blog._id
        self.database[post._id] = post
        blog._post_list.append( post._id )
        self.database[blog._id] = blog
        return post
    
    def delete_post( self, blog, post ):
        del self.database[post._id]
        blog = self.database[blog._id]
        blog._post_list.remove( post._id )
        self.database[blog._id] = blog

    def __iter__( self ):
        for k in self.database:
            if k[0] == '_' : continue
            yield self.database[k]

    def blog_iter( self ):
        for k in self.database:
            if not k.startswith('Blog') : continue
            if ':Post:' in k : continue # skip children
            yield self.database[k]

    def post_iter( self, blog ):
        for k in blog._post_list:
            yield self.database[k]

    def title_iter( self, blog, title ):
        return ( p for p in self.post_iter(blog) if p.title == title )



''' 创建顶层索引 '''

"""
也可以在shelf上添加一个用于定位所有Blog实例的顶层索引
"""

class Access3( Access2 ):
    def new( self, *args, **kw ):
        super().new( *args, **kw )
        self.database['_DB:Blog'] = list()

    def add_blog( self, blog ):
        self.max['Blog'] += 1
        key = 'Blog:{id}'.format(id = self.max['Blog'])
        blog._id = key
        blog._post_list = []
        self.database[blog._id] = blog
        self.database['_DB:Blog'].append( blog._id )
        return blog

    def blog_iter( self ):
        return ( self.database[k] for k in self.database['_DB:Blog'] )


''' 相关更多的索引维护工作 '''
"""
以下是CRUD处理过程中的Create部分
"""

from collections import defaultdict

class Access4( Access2 ):
    def new( self, *args, **kw):
        super().new(*args, **kw)
        self.database['_DB:Blog'] = list()
        self.database['_DB:Blog_Title'] = defaultdict(list)

    def add_blog( self, blog ):
        self.max['Blog'] += 1
        key = 'Blog:{id}'.format(id=self.max['Blog'])
        blog._id = key
        blog._post_list = []
        self.database[blog._id] = blog
        self.database['_DB:Blog'].append( blog._id )
        blog_title = self.database['_DB:Blog_Title']
        blog_title[blog.title].append( blog._id )
        self.database['_DB:Blog_Title'] = blog_title
        return blog
        # 添加了两个索引：一个简单的Blog键列和另一个保存了与某个给定标题相关的键的defaultdict.

    def update_blog( self, blog ):
        """Replace this Blog; update index."""
        
        self.database[blog._id] = blog
        blog_title = self.database['_DB:Blog_Title']
        # Remove key from index in old spot.
        empties = []
        for k in blog_title:
            if blog._id in blog_title[k]:
                blog_title[k].remove(blog._id)
                if len(blog_title[k]) == 0 : empties.append(k)
        # Cleanup zero-length lists from defaultdict
        for k in empties:
            del blog_title[k]
        # put key into index in new spot
        blog_title[blog.title].append(blog._id) 
        self.database['_DB:Blog_Title'] = blog_title

    def blog_iter( self ):
        return ( self.database[k] for k in self.database['_DB:Blog'] )

    def blog_title_iter( self, title ):
        blog_title = self.database['_DB:Blog_Title']
        return ( self.database[k] for k in blog_title[title] )


''' 用writeback代替更新索引 '''

"""
我们可以用writeback=True模式打开shelf，这种模式通过保存每个对象的缓存版本追踪所有可变对象
的更改。相比于跟踪所有访问过的对象来检测和保存更改这种会给shelf模块带来沉重负担的方法，这里
展示的方式会更新一个可变对象然后强制shelf更新这个特定对象的持久化版本。
"""

"""
通过使用工厂函数，我们可以保证应用程序的所有部分都可以一致地协作。
"""
def make_blog( *args, **kw ):
    version = kw.pop('_version', 1)
    if version == 1: return Blog(*args, **kw)
    elif version == 2: return Blog2(*args, **kw)
    else: raise Exception( 'Unknown Version {0}'.fromat(version))

class Blog:
    @staticmethod
    def version( self, version ):
        self.version = version

    @staticmethod
    def blog( self, *args, **kw ):
        if self.version == 1 : return Blog1(*args, **kw)
        elif self.version == 2 : return Blog2(*args, **kw)
        else : raise Exception('Unknown Version {0}'.format(self.version))

# 使用这个工厂
blog = Blog.version(2).blog(title=this, other_attribut=that)


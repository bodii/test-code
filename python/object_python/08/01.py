#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 序列化和保存--JSON、YAML、Pickle、CSV和XML '''

"""
Python 常用的术语

Python的常用术语主要是与dump和load有关：

dump(object, file):用于将对象转储到一个指定文件中。
dump(object): 返回一个对象的字符串表示。
load(file): 从指定文件中加载一个对象，返回新构造的对象。
loads(string): 从一个字符串加载一个对象，返回构造的对象。


用于转储或加载的文件可以是任何类似于文件的对象。可以用read()和readline()这样的一些
方法来实现加载，但我们需要的比这些还要多。因此，可以使用io.StringIO对象和urllib.request
对象作为加载的数据源
"""


''' 定义用于持久化的类 '''


import datetime

class Post:
    def __init__( self, date, title, rst_text, tags ):
        self.date = date
        self.title = title
        self.rst_text = rst_text
        self.tags = tags

    def as_dict( self ):
        return dict(
            date = str(self.date),
            title = self.title,
            underline = '-'*len(self.title),
            rst_text = self.rst_text,
            tag_text = ' '.join(self.tags),
        )

# 如下是一个微博文章的集合，封装了一个集合，因为对集合扩展的方式有些不够稳定。
from collections import defaultdict

class Blog:
    def __init__( self, title, posts=None ):
        self.title = title
        self.entries = posts if posts is not None else []

    def append( self, post ):
        self.entries.append(post)

    def by_tag( self ):
        tag_index = defaultdict(list)
        for post in self.entries:
            for tag in post.tags:
                tag_index[tag].append(post.as_dict())
        return tag_index

    def as_dict( self ):
        return dict(
            title = self.title,
            underline = '='*len(self.title),
            entries = [p.as_dict() for p in self.entries],
        )

travel = Blog( 'Travel' )
travel.append(
    Post(
        date = datetime.datetime(2013, 11, 14, 17, 25),
        title = 'Hard Aground',
        rst_text = """Some embarrassing revalation.
        Including """,
        tags = ('#RedRanger', '#whitby42', '#ICW'),
    )
)
travel.append(
    Post(
        date = datetime.datetime(2013, 11, 18, 15, 30),
        title = 'Anchor Follies',
        rst_text = """Some witty epigram. Including < & > 
        characters. """,
        tags = ('#RedRanger', '#Whitby42', '#Mistakes'),    
    )
)

# 渲染博客与文章列表
from jinja2 import Template

blog_template = Template("""
{{title}}
{{underline}}

{% for e in entries %}
{{e.title}}
{{e.underline}}

{{e.rst_text}}

:date: {{e.date}}

:tags: {{e.tag_text}}
{% endfor %}
Tag Index
=========
{% for t in tags %}
*     {{t}}
      {% for post in tags[t] %}

      _ '{{post.title}}' _
      {% endfor %}
{% endfor %}
""")

print( blog_template.render( tags=travel.by_tag(), **travel.as_dict()) )

# result
"""
Travel
======


Hard Aground
------------

Some embarrassing revalation.
                Including

:date: 2013-11-14 17:25:00

:tags: #RedRanger #whitby42 #ICW

Anchor Follies
--------------

Some witty epigram. Including < & >
        characters.

:date: 2013-11-18 15:30:00

:tags: #RedRanger #Whitby42 #Mistakes

Tag Index
=========

*     #ICW


      _ 'Hard Aground' _


*     #Whitby42


      _ 'Anchor Follies' _


*     #Mistakes


      _ 'Anchor Follies' _


*     #RedRanger


      _ 'Hard Aground' _


      _ 'Anchor Follies' _


*     #whitby42


      _ 'Hard Aground' _
"""
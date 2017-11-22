#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用YAML进行转储和加载 '''

"""
关于YAML， yaml.org网页中是这样描述的：
YAML是一种人性化的，跨语言的，基于Unicode编码进行数据序列化，围绕敏捷编程语言中数值
类型而设计的语言。

在python标准库文档中关于json模块的部分是这样描述的：
JSON是YAML1.2的一个子集。基于这个模块的默认设置（尤其是默认的分隔符），将生成JSON，
也同样是YAML1.0和1.1的子集。这个模块也可用于YAML的序列化器。

从技术的角度上来看，可使用josn模块来生成YAML数据。然而json模块不能被用来完成很复杂的
YAML数据的反序列操作。使用YAML有两点优势：首先，它的格式本身很复杂，我们可以为对象的附
加信息进行编码;其次，pyYAML的实现在底层很好地集成了Python，可以很简单地为Pyhton对象
进行YAML编码。YAML的缺点是它的使用范围不像JSON这样广泛。
"""

import yaml
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
text = yaml.dump(travel)
print( text )
"""
!!python/object:__main__.Blog
entries:
- !!python/object:__main__.Post
  date: 2013-11-14 17:25:00
  rst_text: "Some embarrassing revalation.\n        Including "
  tags: !!python/tuple ['#RedRanger', '#whitby42', '#ICW']
  title: Hard Aground
- !!python/object:__main__.Post
  date: 2013-11-18 15:30:00
  rst_text: "Some witty epigram. Including < & > \n        characters. "
  tags: !!python/tuple ['#RedRanger', '#Whitby42', '#Mistakes']
  title: Anchor Follies
title: Travel
"""


"""
YAML中包含了11个标准的标签。yaml模块包含了很多为Python定制的标签以及5个complexPython
标签。
一个列表中的项将以一个块序列的形式呈现。每项以一个-序列为起始，其余部分以两个空格缩进。
当集合或元组足够小，可缩进为一行。当长度增加时，就显示为多行。
"""

copy = yaml.load(text)
print( copy )
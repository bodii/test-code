#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用JSON进行转储和加载 '''

"""
JSON(javaScript Object Notation)是一种轻量级的数据交换格式，易于人阅读和编写。同
时也易于机器解析和生成。

如下几中Python类型对应了JSON中所使用的JavaScript类型：
Python                  JSON
--------------------------------
dict                    object
list, tuple             array
str                     string
int, float              number
True                    true
False                   false
None                    null

除了以上定义的几种类型，其他类型都不支持。
"""

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

import json
print( json.dumps( travel.as_dict(), indent=4 ) )

"""
JSON这种数据表现形式也有几点不方便的地方：
1. 如果必须将Python对象重写为字典，应该提供一种更好的转换方式，不必额外地创建字典对象。
2. 当我们加载这个JSON数据时，并不能够轻易地恢复之前的Blog和Post对象。当我们使用json.
load()时，并不会得到Blog和Post对象，仅仅能够得到dict和集合对象。在创建Blog和Post对象
时，我们需要额外的信息。
3. 对于对象内__dict__中的一些值我们并不希望保存，例如Post中下划线的文字。
"""

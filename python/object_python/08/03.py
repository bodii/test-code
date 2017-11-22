#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 自定义JSON编码 '''

"""
类的提示可以提供3部分信息：一个__class__键，作为目标类的命名;__args__键，将提供一个
位置参数值的序列;一个__kw__值，将提供一个关键字参数值的字典，包含了__init__()的所有
选项。
"""

import datetime
import json
from collections import defaultdict


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

def blog_encode( object ):
    if isinstance(object, datetime.datetime):
        return dict(
            __class__ = 'datetime.datetime',
            __args__ = [],
            __kw__ = dict(
                year = object.year,
                month = object.month,
                day = object.day,
                hour = object.hour,
                minute = object.minute,
                second = object.second,
            )
        )
    elif isinstance(object, Post):
        return dict(
            __class__ = 'Post',
            __args__ = [],
            __kw__ = dict(
                date = object.date,
                title = object.title,
                rst_text = object.rst_text,
                tags = object.tags,
            )
        )
    elif isinstance(object, Blog):
        return dict(
            __class__ = 'Blog',
            __args__ = [
                object.title,
                object.entries,
            ],
            __kw__ = {}

        )   
    else:
        return json.JSONEncoder.default(o)

"""
这个函数演示了对3个类对象编码操作的两种不同风格。
1. 将datetime.datetime对象编码为一个字典，其中包含了独立的字段值。
2. 将Post对象编码为一个字典，也包含了独立的字段值。
3. 将Blog实例编码为由标题和文章组成的一个序列。
"""

text = json.dumps(travel, indent=4, default=blog_encode)
print( text )
"""
{
    "__kw__": {},
    "__class__": "Blog",
    "__args__": [
        "Travel",
        [
            {
                "__kw__": {
                    "date": {
                        "__kw__": {
                            "hour": 17,
                            "second": 0,
                            "month": 11,
                            "minute": 25,
                            "day": 14,
                            "year": 2013
                        },
                        "__class__": "datetime.datetime",
                        "__args__": []
                    },
                    "rst_text": "Some embarrassing revalation.\n        Including ",
                    "tags": [
                        "#RedRanger",
                        "#whitby42",
                        "#ICW"
                    ],
                    "title": "Hard Aground"
                },
                "__class__": "Post",
                "__args__": []
            },
            {
                "__kw__": {
                    "date": {
                        "__kw__": {
                            "hour": 15,
                            "second": 0,
                            "month": 11,
                            "minute": 30,
                            "day": 18,
                            "year": 2013
                        },
                        "__class__": "datetime.datetime",
                        "__args__": []
                    },
                    "rst_text": "Some witty epigram. Including < & > \n        characters. ",
                    "tags": [
                        "#RedRanger",
                        "#Whitby42",
                        "#Mistakes"
                    ],
                    "title": "Anchor Follies"
                },
                "__class__": "Post",
                "__args__": []
            }
        ]
    ]
}
"""
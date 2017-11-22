#!/usr/bin/env python3
# -*- coding:utf-8 -*-

# 如下是一个微博文章的集合，封装了一个集合，因为对集合扩展的方式有些不够稳定。
from collections import defaultdict

class Blog:
    """ Blog博客文章类 """
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
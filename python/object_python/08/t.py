#!/usr/bin/env python3
# -*- coding:utf-8 -*-

class Post:
    def __init__(self, date, title, rst_text, tags):
        self.date = date
        self.title = title
        self.rst_text = rst_text
        self.tags = tags

    def as_dict( self ):
        return dict(
            date = str(self.date),
            title = self.title,
            underline = '-' * len(self.title),
            rst_text = self.rst_text,
            tag_text = ' '.join(self.tags),
        )

# 博客
from collections import defaultdict
from datetime import datetime

class Blog:
    def __init__( self, title, posts=None):
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
            underline = '=' * len(self.title),
            entries = [ p.as_dict() for p in self.entries ],
        )

# 添加博客内容

travel = Blog('Travel')

travel.append(
    Post(
        date = datetime(2017, 9, 3, 10, 12),
        title = 'test1',
        rst_text = 'test1 0000000',
        tags = ('#tag1', '#tag2', '#tag3')
    )
)

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
*       {{t}}
        {% for post in tags[t] %}
        _ '{{post.title}}' _
        {% endfor %}
{% endfor%}
""")

print(blog_template.render( tags=travel.by_tag(), **travel.as_dict() ))
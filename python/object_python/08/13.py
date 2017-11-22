#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 加载XML文档 '''

"""
从一个XML文档中加载Python对象分为两步。首先，我们需要对XML文本解析，用于创建文档对象。然后
，需要对用于生成Python对象的文档对象进行检查。XML格式具有极大的灵活性，从XML到Python的序列
化方法并不是唯一的。
"""

import ast
import xml
import io
from Post import *
import datetime

doc = xml.parsers( io.StringIO(text.decode('utf-8')) )
xml_blog = doc.getroot()

blog = Blog( xml_blog.findtext('title') )
for xml_post in xml_blog.findall('entries/entry'):
    tags = [t.text for t in xml_post.findall( 'tags/tag') ]
    post = Post(
        date = datetime.datetime.strptime(
            xml_post.findtext('date'), '%Y-%m-%d %H:%M:%S'
        ),
        title = xml_post.findtext('title'),
        tags = tags,
        rst_text = xml_post.findtext('rst_text')
    )
    blog.append(post)
render( blog )
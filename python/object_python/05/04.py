#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' OrderedDict集合 '''

"""
OrderedCollecgtion集合很聪明地使用了两种存储结构。使用一个基本的dict对象类型用来
匹配争键和值。另外，使用存储了键的双向列表用来维护插入顺序。
OrderedDict的一常用场景是处理HTML或者XML文件，这些文件中对象的顺序必须保留，但是
对象之间有可能通过ID或者IDREF属性互相引用。通过将ID用作字典的键，我们可以优化对象间
的这种关系。我们可以使用OrderedDict结构保持源文档中对象的顺序。
"""

# 两个XML
"""
1.
<blog>
    <topics>...</topics> <indices>...</indices>
</blog>

2. 
<topics>
    <entry ID='UUID98766'>
        <title>first</title>
        <body>more words</body>
    </entry>
    <entry ID='UUID86543'>
        <title>second</title>
        <body>words</body>
    </entry>
    <entry ID='UUID64319'>
        <title>third</title>
        <body>text</body>
    </entry>
</topics>

每一个topics都包含一系列的entry。每一个entry都有一个唯一的ID。这里在暗示它们应该
属于唯一标识符（UUID）
"""

from collections import OrderedDict
import xml.etree.ElementTree as ET


parser = ET.XMLParser(encoding="utf-8")
doc = ET.parse( 'blog.xml', parser=parser )

topics = OrderedDict()

for topic in doc.findall('topics/entry'):
    topics[topic.attrib['ID']] = topic

for topic in topics:
    print( topic, topics[topic].find('title').text)

"""
UUID98766 first
UUID86543 second
UUID64319 third
"""
#！/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic:第六章：数据编码和处理
Description: 这一章主要讨论使用Python 处理各种不同方式编码的数据，比如CSV 文件，
JSON，XML 和二进制包装记录。和数据结构那一章不同的是，这章不会讨论特殊的
算法问题，而是关注于怎样获取和存储这些格式的数据。

Title:      利用命名空间解析 XML 文档
Issue:你想解析某个 XML 文档，文档中使用了 XML 命名空间。
Answers:使用 xml.etree.ElementTree 模块可以很容易的处理这些任务。
"""
"""
。
>>> # Some queries that work
>>> doc.findtext('author')
'David Beazley'
>>> doc.find('content')
<Element 'content' at 0x100776ec0>
>>> # A query involving a namespace (doesn't work)
>>> doc.find('content/html')
>>> # Works if fully qualified
>>> doc.find('content/{http://www.w3.org/1999/xhtml}html')
<Element '{http://www.w3.org/1999/xhtml}html' at 0x1007767e0>
>>> # Doesn't work
>>> doc.findtext('content/{http://www.w3.org/1999/xhtml}html/head/title')
>>> # Fully qualified
>>> doc.findtext('content/{http://www.w3.org/1999/xhtml}html/'
... '{http://www.w3.org/1999/xhtml}head/{http://www.w3.org/1999/xhtml}title')
'Hello World'
"""
"""
你可以通过将命名空间处理逻辑包装为一个工具类来简化这个过程：
"""
class XMLNamespaces:
    def __init__(self, **kwargs):
        self.namespaces = {}
        for name, uri in kwargs.items():
            self.register(name, uri)

    def register(self, name, uri):
        self.namespaces[name] = '{'+uri+'}'

    def __call__(self, path):
        return path.format_map(self.namespaces)

"""
通过下面的方式使用这个类：
"""
ns = XMLNamespaces(html='https://www.w3.org/1999/xhtml/')
print(doc.find(ns('content/{html}html')))
print(doc.findtext(ns('content/{html}html/{html}head/{html}title')))
"""
解析含有命名空间的 XML 文档会比较繁琐。上面的 XMLNamespaces 仅仅是允许你
使用缩略名代替完整的 URI 将其变得稍微简洁一点。
很不幸的是，在基本的 ElementTree 解析中没有任何途径获取命名空间的信息。
但是，如果你使用 iterparse() 函数的话就可以获取更多关于命名空间处理范围的信
息。例如：
"""
from xml.etree.ElementTree import iterparse
for evt, elem in iterparse('ns2.xml', ('end', 'start-ns', 'end-ns')):
    print(evt, elem)
'''
end <Element 'author' at 0x10110de10>
start-ns ('', 'http://www.w3.org/1999/xhtml')
end <Element '{http://www.w3.org/1999/xhtml}title' at 0x1011131b0>
end <Element '{http://www.w3.org/1999/xhtml}head' at 0x1011130a8>
end <Element '{http://www.w3.org/1999/xhtml}h1' at 0x101113310>
end <Element '{http://www.w3.org/1999/xhtml}body' at 0x101113260>
end <Element '{http://www.w3.org/1999/xhtml}html' at 0x10110df70>
end-ns None
end <Element 'content' at 0x10110de68>
end <Element 'top' at 0x10110dd60>
>>> elem # This is the topmost element
<Element 'top' at 0x10110dd60>
'''
"""
最后一点，如果你要处理的 XML 文本除了要使用到其他高级 XML 特性外，还要
使用到命名空间，建议你最好是使用 lxml 函数库来代替 ElementTree 。例如， lxml
对利用 DTD 验证文档、更好的 XPath 支持和一些其他高级 XML 特性等都提供了更
好的支持。这一小节其实只是教你如何让 XML 解析稍微简单一点。
"""
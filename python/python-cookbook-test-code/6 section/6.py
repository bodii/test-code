#！/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic:第六章：数据编码和处理
Description: 这一章主要讨论使用Python 处理各种不同方式编码的数据，比如CSV 文件，
JSON，XML 和二进制包装记录。和数据结构那一章不同的是，这章不会讨论特殊的
算法问题，而是关注于怎样获取和存储这些格式的数据。

Title:      解析和修改 XML
Issue:你想读取一个 XML 文档，对它最一些修改，然后将结果写回 XML 文档。
Answers:使用 xml.etree.ElementTree 模块可以很容易的处理这些任务。
"""
from xml.etree.ElementTree import parse, Element

doc = parse('../test file/pared.xml')
root = doc.getroot()
print(root)
# <Element 'urlset' at 0x0000000000B351D8>

print(root.remove(root.find('sri')))
print(root.remove(root.find('cr')))
print(root.getchildren().index(root.find('nm')))
"""
>>> e = Element('spam')
>>> e.text = 'This is a test'
>>> root.insert(2, e)
>>> # Write back to a file
>>> doc.write('newpred.xml', xml_declaration=True)

修改一个 XML 文档结构是很容易的，但是你必须牢记的是所有的修改都是针对父
节点元素，将它作为一个列表来处理。例如，如果你删除某个元素，通过调用父节点
的 remove() 方法从它的直接父节点中删除。如果你插入或增加新的元素，你同样使
用父节点元素的 insert() 和 append() 方法。还能对元素使用索引和切片操作，比如
element[i] 或 element[i:j]
如果你需要创建新的元素，可以使用本节方案中演示的 Element 类。我们在 6.5 小
节已经详细讨论过了。
"""
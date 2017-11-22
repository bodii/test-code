#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 为容器和集合设计类 '''

"""
当处理更复杂的容器或者集合时，类的设计会变得更复杂。第1个问题是容器的范围，我们必须确定
shelf上对象的粒度。
当使用容器时，我们可以将整个容器作为一个单一的复杂对象保存到shelf上。在某种程度上，这种
做法可能违背了shelf上保存多个对象的初衷。保存一个巨大的容器为我们带来了粗粒度的存储结构。
如果改变容器中的一个对象，那么整个容器都必须重新序列化并保存。如果可以高效地将全部对象都
保存在一个单一容器中，那么为什么还要使用shelve？因此，我们必须找到一个符合程序需求的平衡
点。
另外一个选择是将集合分解为独立的元素。用这种方法的话，最高层的Blog对象不再是一个合适的Python
容器。父类可能会通过键的集合来获取每个子类，每个子对象可以用键获取父对象，这种键的使用方法
在面向对象设计中并不常用。通常，对象只是简单地包含了指向其他对象的引用。当使用shelve（或者
其他数据库）时，我们必须通过键使用间接引用。
现在每个子类都有两个键：它自己的主键和一个指向父对象主键的外键。这就带来了第2个问题，如何用
字符串表示父类和它们的子类的键？
"""


''' 用外键引用对象 '''

"""
我们用来唯一标识一个对象的键是[主键]。当子对象引用父对象时，需要添加一些额外的设计。要
格式化子对象的主键？有两种常用的设计子类主键的方法，这两种方法都基于类的对象间是何种依
赖关系。

"Child:cid": 当子类可以独立于父类存在时，会考虑使用这种格式。例如，发票中的一个条目
代表一个产品，即使没有这个代表产品的发票条目，这个产品依然可以存在。

"Parent:pid:Child:cid":当子类不能脱离父类而独立存在时，我们会考虑使用这种格式。一
个用户地址不能离开用户而单独存在。当子类完全依赖于父类时，子类的键可以包含父类的键来反
映这种依赖关系。


与父类的设计一样，如果将主键和所有与子类有关的外键都保存起来是最简单的方法。建议不要在
__init__()方法中初始化它们，因为它们只是持久化的一部分。
"""

import time

class Post:
    def __init__( self, date, title, rst_text, tags ):
        self.date = date
        self.title = title
        self.rst_text = rst_text
        self.tags = tags

    def as_dict(self):
        return dict(
            date = str(self.date),
            title = self.title,
            underline = '-' * len(self.title),
            rst_text = self.rst_text,
            tag_text = ' '.join(self.tags),
        )

date = time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(time.time()))

p2 = Post(
            date = date,
            title = 'Hard Aground',
            rst_text = """Some embarrassing revelation.
                    Including ☺ and ⚓""",
            tags = ("#RedRanger", "#Whitby42", "#ICW"),
        )

p3 = Post(
    date = date,
    title = 'Anchor Follies',
    rst_text = """Some witty  epigram.Including < & > characters.""",
    tags = ('#RedRanger', '#Whitby42', '#Mistakes'),
)


class Blog:
    def __init__( self, title, *posts ):
        self.title = title
    
    def as_dict( self ):
        return dict(
            title = self.title,
            underline = '=' * len(self.title),
        )

b1 = Blog( title='Travel Blog' )

"""
现在我们可以通过设置属性和分配定义关系的键来将这些Post和它们所属的Blog联系起来。
"""

import shelve

shelf = shelve.open('blog')
b1._id = 'Blog:1'
shelf[b1._id] = b1

owner = shelf['Blog:1']
"""
我们用了主键来定位owner对象。一个真实的应用程序可能会根据标题来搜索这个对象，可能也
会创建索引来优化搜索的过程。
"""

# 将owner的键分配给每个Post对象，然后保存这些对象。
p2._parent = owner._id
p2._id = p2._parent + ':Post:2'
shelf[p2._id] = p2

p3._parent = owner._id
p3._id = p3._parent + ':Post:3'
shelf[p3._id] = p3

print( list(shelf.keys()) )
# ['Blog:1:Post:3', 'Blog:1', 'Blog:1:Post:2']

print( p2._parent )
# Blog:1
print( p2._id )
# Blog:1:Post:2
#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 用Shelve保存和获取对象 '''

import shelve
from contextlib import closing
import os.path

pwd = os.path.dirname(__file__) + os.path.sep

with closing( shelve.open(pwd + 'some_file') ) as shelf:
    # process( shelf )
    pass

"""
打开一个shelf，然后将它提供给应用程序中那些真正完成需求的函数使用。当这个过程完成后，这个
上下文可以确保shelf被关闭。如果process()函数抛出一个异常，shelf仍然会被正确地关闭。

由于键用于定位值，因此键必须是唯一的。这就为设计类带来了一些需要考虑的因素，因为必须提供合适
且唯一的键。一些情况下，问题域中会包含一个属性，这个属性就是明显的唯一的键。在那种情况下，
可以简单地用这个属性创建键：shelf[object.key_attribute] = object。这是最简单的情况，
但是并不通用。

对于没有明显主键的情形，我们可以尝试用一些值的组合创建唯一的组合键（composite key）。
这并不总是一个非常好的主意，因为现在键不是原子的，对于键中任何一个部分的改变都会带来数据
更新的问题。

最简单的方法通常是使用代理键(surrogate key)设计模式。这个键不依赖于对象中的数据，它是对象
的一个替代品。这意味着对象的任何属性都可以被改变，并且不会带来什么副作用或者限制。Python
内部的对象ID是代理键的一种示例。一个shelf键的字符串表示可以遵循这种模式：class:oid。

字符串键包含了对象所属的类和当前类示例的唯一标识符。我们可以用这种形式的键简单地将不同的类的对
象保存在一个单一的shelf上。即使只准备在shelf上保存一种类型的对象。这种格式对于保存索引的命名
空间、用于管理的元数据和以后的扩展都是非常有帮助的。

当我们有了一个适合的业务主键后，我们可能想做一些后面这样的事情持久化shelf上的对象：
self[object.__class__.__name__ + ':' + object.key_attribute] = object。
"""

class Blog:
    def __init__( self, title, *posts ):
        self.title = title
    
    def as_dict( self ):
        return dict(
            title = self.title,
            underline = '=' * len(self.title),
        )

b1 = Blog( title='Travel Blog' )

# 可以用下面的方式将这个简单的对象保存在shelf上

import shelve

# 创建一个新的shelf，文件名是blog。
shelf = shelve.open('blog')
# 创建一个键名为'Blog:1',值为Blog的实例b1,赋值给_id属性的键
b1._id = 'Blog:1'
shelf[b1._id] = b1

# 当调用shelf['Blog:1']时，它会从shelf上取回原始的Blog实例。
print( shelf['Blog:1'] )
# <__main__.Blog object at 0x7fd057e9b898>
print( shelf['Blog:1'].title )
# Travel Blog
print( shelf['Blog:1']._id )
# Blog:1
print( list(shelf.keys()) )
# ['Blog:1']
shelf.close()


shelf = shelve.open('blog')
results = [ shelf[k] for k in shelf.keys() if k.startswith('Blog:') and 
            shelf[k].title == 'Travel Blog' ]

r0 = results[0]
print( r0.title )
# Travel Blog
print(r0._id)
# Blog:1

"""
打开shelf访问对象，用restuls生成器表达式遍历shelf上的每个元素，查询出所有以’Blog:‘为键的开始并且对象的
title属性是’Travel Blog‘的元素。
"""
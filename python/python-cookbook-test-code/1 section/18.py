#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第一章：数据结构和算法
Desc: Python 提供了大量的内置数据结构，包括列表，集合以及字典。大多数情况下使
用这些数据结构是很简单的。但是，我们也会经常碰到到诸如查询，排序和过滤等等
这些普遍存在的问题。因此，这一章的目的就是讨论这些比较常见的问题和算法。另
外，我们也会给出在集合模块 collections 当中操作这些数据结构的方法。

Title:          映射名称到序列元素
Issue:  你有一段通过下标访问列表或者元组中元素的代码，但是这样有时候会使得你的代
码难以阅读，于是你想通过名称来访问元素。
代访问
answer:  collections.namedtuple() 函数通过使用一个普通的元组对象来帮你解决这个问
题。这个函数实际上是一个返回 Python 中标准元组类型子类的一个工厂方法。你需要
传递一个类型名和你需要的字段给它，然后它就会返回一个类，你可以初始化这个类，
为你定义的字段传递值等。
"""

from collections import namedtuple

Subscriber = namedtuple('Subscriber', ['addr', 'joined'])
sub = Subscriber('jonesy@exalmple.com', '2012-10-19')
print(sub) # Subscriber(addr='jonesy@exalmple.com', joined='2012-10-19')
print(sub.addr) # jonesy@exalmple.com
print(sub.joined) # 2012-10-19

'''
尽管 namedtuple 的实例看起来像一个普通的类实例，但是它跟元组类型是可交换
的，支持所有的普通元组操作，比如索引和解压。
'''
print(len(sub)) # 2
addrs, joineds = sub
print(addrs) # jonesy@exalmple.com
print(joineds) # 2012-10-19

'''
命名元组的一个主要用途是将你的代码从下标操作中解脱出来。因此，如果你从数
据库调用中返回了一个很大的元组列表，通过下标去操作其中的元素，当你在表中添
加了新的列的时候你的代码可能就会出错了。但是如果你使用了命名元组，那么就不
会有这样的顾虑。
'''

def compute_cost(records):
    total = 0.0
    for rec in records:
        total += rec[1] * rec[2]
    return total

'''
下标操作通常会让代码表意不清晰，并且非常依赖记录的结构。下面是使用命名元
组的版本：
'''
Stock = namedtuple('Stock', ['name', 'shares', 'price'])
def compute_cost(records):
    total =0.0
    for rec in records:
        s = Stock(*rec)
    total += s.shares * s.price
    return total

'''
命名元组另一个用途就是作为字典的替代，因为字典存储需要更多的内存空间。如
果你需要构建一个非常大的包含字典的数据结构，那么使用命名元组会更加高效。但
是需要注意的是，不像字典那样，一个命名元组是不可更改的。
'''
s = Stock('ACME', 100, 123.45)
print(s)
# Stock(name='ACME', shares=100, price=123.45)

print(s.shares == 75) # False

'''
如果你真的需要改变然后的属性，那么可以使用命名元组实例的 replace() 方法，
它会创建一个全新的命名元组并将对应的字段用新的值取代。
'''
s = s._replace(shares=75)
print(s)
# Stock(name='ACME', shares=75, price=123.45)

'''
replace() 方法还有一个很有用的特性就是当你的命名元组拥有可选或者缺失字
段时候，它是一个非常方便的填充数据的方法。你可以先创建一个包含缺省值的原型
元组，然后使用 replace() 方法创建新的值被更新过的实例
'''

Stock = namedtuple('Stock', ['name', 'shares', 'price', 'date', 'time'])
stock_prototype = Stock('', 0, 0.0, None, None)

def dict_to_stock(s):
    return stock_prototype._replace(**s)

a = {'name': 'ACME', 'shares': 100, 'price': 123.45}
print(dict_to_stock(a))
# Stock(name='ACME', shares=100, price=123.45, date=None, time=None)

b = {'name': 'ACME', 'shares': 100, 'price': 123.45, 'date': '12/17/2012'}
print(dict_to_stock(b))
# Stock(name='ACME', shares=100, price=123.45, date='12/17/2012', time=None)

'''
最后要说的是，如果你的目标是定义一个需要更新很多实例属性的高效数据结构，
那么命名元组并不是你的最佳选择。这时候你应该考虑定义一个包含 slots 方法的
类
'''

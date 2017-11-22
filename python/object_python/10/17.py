#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用索引提高性能 '''

"""
如：
create index ix_blog_title on blog(title);

大多数情况下，通过在数据库中创建一个unique索引来实现unique和primary key的约束。
"""

''' 添加ORM层 '''

''' 设计ORM友好类 '''
"""
SQLAlchemy需要创建一个定义性的基类（declarative base class)。这个基类为应用的类定义
提供了元类。对于元数据而言它是我们为数据库所定义的库。默认地，很容易将这个类称为Base。
"""

from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, Table
from sqlalchemy import BigInteger, Boolean, Date, DateTime, Enum, \
    Float, Integer, Interval, LargeBinary, Numeric, PickleType, \
    SmallInteger, String, Text, Time, Unicode, UnicodeText, ForeignKey
from sqlalchemy.orm import relationship, backref

# SQLAlchemy的元类由declarative_base()函数来创建
Base = declarative_base()

# 所创建的Base对象必须是将要定义的持久化类的元类。

# 以下是Blog类
class Blog( Base ):
    __tablename__ = 'blog'
    id = Column(Integer, primary_key=True)
    title = Column(String)
    
    def as_dict( self ):
        return dict(
            title = self.title,
            underline = '='*len(self.title),
            entries = [ e.as_dict() for e in self.entries ]
        )

# post类
class Post( Base ):
    __tablename__ = 'post'
    id = Column(Integer, primary_key=True)
    title = Column(String)
    date = Column(DateTime)
    rst_text = Column(UnicodeText) # 任何字符都是Unicode
    blog_id = Column(Integer, ForeignKey('blog.id')) # 外键
    blog = relationship('blog', backref='entries')
    tags = relationship('tag', secondary=assoc_post_tag, backref='posts')

    def as_dict( self ):
        return dict(
            title = self.title,
            underline = '-'*len(self.title),
            date = self.date,
            rst_text = self.rst_text,
            tags = [ t.phrase for t in self.tags ],
        )

class Tag( Base ):
    __tablename__ = 'tag'
    id = Column(Integer, primary_key)
    phrase = Column(String, unique=True)


assoc_post_tag = Table('Assoc_post_tag',
    Column('post_id', Integer, ForeignKey('post.id')),
    Column('tag_id', Integer, ForeignKey('tag.id'))
)
#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用ORM层创建模型 '''

from sqlalchemy import create_engine
import os.path as opt

cur_path = opt.dirname(__file__) + opt.sep

engine = create_engine('sqlite:///'+ cur_path + 'p2_c11_blog2.db', echo=True)
Base.metadata.create_all(engine) 
"""
echo = True意味着，生成的SQL语句会写入日志记录。

create_all，创建所有包括的表。
drop_all, 用于删除所有数据。
"""


''' 使用ORM层操作对象 '''

from sqlalchemy.orm import sessionmaker
Session = sessionmaker(bind=engine)
session = Session()

blog = Blog( title='Travel 2017' )
session.add(blog)



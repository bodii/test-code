#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 安全性和eval() '''

"""
有时，一个文件在网络传输过程中可能被中间人篡改成为不可靠的文件，对于这样的场景需要制定
一些规则。
如果有必要，可以将eval()替换为一个字典，完成从名字到类的映射。我们可以将eval(som_dict
['__class__']) 改为 {'Post': Post, 'Blog': Blog, 'datetime.datetime':datetime.
datetime}[some_dict['__class__']]
"""


''' 重构编码函数 '''

"""
如果要修改类库中类的编码行为，例如datetime,就需要在应用程序中对datetime.datetime
进行扩展。这样一来，就需要确定在我们的应用中使用了扩展版本而不是类库版本，而完全避免使用
内置的datetime类并不是非常容易。
"""
@property
def _json( self ):
    return dict(
        __class__ = self.__class__.__name__,
        __kw__ = {},
        __args__ = [ self.title, self.entries ],
    )

# 这个特性被用来提供解码函数所使用的初始化参数。可以向Post添加这个两个特性

def _json( slef ):
    return dict(
        __class__ = self.__class__.__name__,
        __kw__ = dict(
            date = self.date,
            title = self.title,
            rst_text = self.rst_text,
            tags = self.tags,
        ),
        __args__ = [],
    )

# 对于Blog，这个特性将用来提供初始化参数。它们会被解码函数使用，可以修改编码器
import datetime
import json

def blog_encode_2( object ):
    if isinstance(object, datetime.datetime):
        return dict(
            __class__ = 'datetime.datetime',
            __args__ = [],
            __kw__ = dict(
                year = object.year,
                month = object.month,
                day = object.day,
                hour = object.hour,
                minute = object.minute,
                second = object.second,
            )
        )
    else:
        try:
            encoding = object._json()
        except AttributeError:
            encoding = json.JSONEncoder.default(o)
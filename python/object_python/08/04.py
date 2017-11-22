#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 自定义JSON解码 '''

"""
JSON解码器中“对象钩子”是一个被dict调用的函数，用于检验是否被正确地表达成了自定义对象。
如果dict没有被hook函数识别，那么它仅是一个字典并且应当被直接返回。
"""
import json

def blog_decode( some_dict ):
    if set(some_dict.keys()) == set(['__class__', '__args__', '__kw__']):
        class_ = eval(some_dict['__class__'])
        return class_(*some_dict['__args__'], **some_dict['__kw__'])
    else:
        return some_dict


blog_data = json.loads(text, object_hook = blog_decode)

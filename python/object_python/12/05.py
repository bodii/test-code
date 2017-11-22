#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 通过SimpleNamespace进行配置 '''

"""
使用types.SimpleNamespace对象来根据需要进行属性的添加，这和类定义是类似的。
在定义一个类时，所有的赋值语句都在类中完成。当创建一个SimpleNamespace对象时，需要
显式指定每一个需要获取的Namespace的名称。
"""
import types
config = types.SimpleNamespace(
    param1 = 'some value',
    param2 = 3.14,
)

print( config )

config5 = types.SimpleNamespace()
config5.dealer_rule = Hit17()
config5.split_rule = NoReSplitAces()
config5.table = Table( decks=6, limit=50, dealer=config5.dealer_rule,
                split=config5.split_rule, payout=(3, 2))


def make_config():
    config = types.SimpleNamespace()
    # set the default values
    config.some_option = default_value
    return config

config = make_config()
config.some_option = another_value


def make_config( **kw ):
    config = types.SimpleNamespace()
    # set the default value
    config.some_option = kw.get('some_option', default_value) 
    return config

config = make_config( some_option=auther_value)



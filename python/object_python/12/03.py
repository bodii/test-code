#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 实现具有层次结构的配置 '''

"""
"""
import this

# print( this.__file__ )
import os.path
import os

config_name = 'someapp.config'
config_locations = (
    os.path.expanduser('~thisapp/'),
    '/etc',
    os.path.expanduser('~/'),
    os.path.curdir,
)   

candidates = (os.path.join(dir, config_name)
    for dir in config_locations)
config_names = [ name for name in candidates if os.path.exists(name) ]


''' 使用INI文件保存配置 '''

"""
一个INI文件看起来如下(blackjack.ini)：
; Default casino rules
[table]
    dealer = Hit17
    split = NoResplitAces
    decks = 6
    limit = 50
    payout = (3, 2)

[player]
    play = SomeStrategy

[simulator]
    samples = 100

我们将参数分成了3个部分。

"""

import configparser
config = configparser.ConfigParser()
config.read('blackjack.ini')
"""
当为ConfigParser.read()提供了一个文件名列表时，它将按顺序读取每一个文件。
"""
def main_ini( config ):
    dealer_nm = config.get('table', 'dealer', fallback='Hit17')
    dealer_rule = {
                    'Hit17' : Hit17(),
                    'Stand17' : Stand17()
                }.get(dealer_nm, Hit17())
    split_nm = config.get('table', 'split', fallback='ReSplit')
    split_rule = {'ReSplit': ReSplit(),
                    'NoReSplit': NoReSplit(),
                    'NoReSplitAces': NoReSplitAces()
                }.get(split_nm, ReSplit())
    decks = config.getint('table', 'decks', fallback=6)
    limit = config.getint('table', 'limit', fallback=100)
    """
    获取一个值，可使用ConfigParser来解析：这个类可以直接处理str、int、float和
    bool。这个类中包括了一个复杂的映射，从字符串到布尔类型，使用了大量的公共代码
    以及Ture和False的同义词。
    非内置类型的处理：对于payout的情况，我们使用了一个字符串值，‘（3, 2）’，
    Config.parser中并没有直接支持这个数值类型。有两种方法可以解决：
    1. 可以自己试着解析，或者把它的值当作合法的Python表达式然后使用Python来完成。
    这样的话，就需要使用eval()函数。
    """
    payout = eval( config.get('table', 'payout', fallback='(3, 2)'))
    table = Table(decks = decks, limit = limit, dealer=dealer_rule,
                split=split_rule, payout=payout)


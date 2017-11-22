#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用JSON和YAML文件存储配置 '''

import json
import os.path
current_path = os.path.dirname(__file__) + os.path.sep
file = open(current_path + 'config.json')
config = json.load(file)

# print( config.get('table', {}).get('decks', 7) )
print( config.get('table').get('decks') ) # get第二个参数是默认值


def main_nested_dict( config ):
    dealer_nm = config.get('table', {}).get('dealer', 'Hit17')
    dealer_rule = {
        'Hit17': Hit17(),
        'Stand17': Stand17()
        }.get(dealer_nm, Hit17())
    split_rule = {
        'ReSplit': ReSplit(),
        'NoReSplit' : NoReSplit(),
        'NoReSplitAces' : NoReSplitAce()
        }.get(split_nm, ReSplit())
    decks = config.get('table', {}).get('decks', 6)
    limit = config.get('table', {}).get('limit', 100)
    payout = config.get('table', {}).get('payout', (3, 2))
    table = Table( decks=decks, limit=limit, dealer=dealer_rule,
            split=split_rule, payout=payout )
    
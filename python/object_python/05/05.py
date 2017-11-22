#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' defaultdic子类 '''

"""
当一个键不存在时，默认的dict类型会抛出一个异常。defaultdic集合类型会执行一个给定的函数，并将执行结果
作为这个不存在的键值存入字典中。

defaultdict类的一个常见应用是为对象创建索引。当有一些对象共同包含一个键时，我们可以创建一个对象列表共
享这个键。

例：
outcomes = defualtdict(list)
self.play_hand( table, hand )
outcome = self.get_payout()
outcomes[hand.dealer_card.rank].append(outcome)

outcomes[rank]中的值是一个模拟的收益列表，我们可以求这些值的平均数或总数来统计收益情况。 
"""
"""
一些情况下，我们可能会想用defaultdict集合提供一个常量。比起container.get(key, 'N/A'),
更倾向于使用container[key]。当key不存在时，总是返回一个字符串常量。实现这个行为的难点在于
defaultdict类会使用一个无参的函数创建默认值，没有办法为其指定一个常量。
"""

from collections import defaultdict

messages = defaultdict( lambda: 'N/A')
messages['error1'] = 'Full Error Text'
print( messages['other'] )
# N/A

print( [k for k in messages if messages[k] == 'N/A'] )
# ['other']

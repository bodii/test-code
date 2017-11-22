#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 应用程序配置的设计模式 '''
"""
有关应用程序配置有两种核心的设计模式：
1. 全局属性映射：使用一个全局的对象，它将包含所有的配置参数。它可以是name:value的
映射对或是包含了属性值的对象。它将使用单例设计模式来确保只有一个实例。
2. 对象创建：不再使用单例，而是定义了工厂或是工厂的集合，基于配置数据来创建应用程序
的对象。这样一来，配置信息只在程序启动时被使用一次。配置信息不会保存在全局的对象中。

全局属性映射的设计非常流行，因为它很简单并且容易扩展。例如：
class Configuration:
    some_attribute = 'default_value'

使用：
Configuration.some_attribute = 'user-supplied value'

关于它的改进是使用正式的单例设计模式。
可能会有一个模块名称为configuration.py，在这个文件中，可能会有如下定义：
settings = dict()
现在就可以将configuration.settings看作是应用设置的一个全局的库。
函数或类可以解析这个配置文件，使用应用将使用的配置值来加载这个字典。

在21点游戏的模拟中，可能会看到如下代码：
shoe = Deck( configuration.settings['decks'] )
或：
if bet > configuration.settings['limit']: raise InvalidBet()
"""

''' 使用对象的构造完成配置 '''
"""
通常可以在单一的、全局的main()函数中将对象构建的初始化过程的逻辑进行集中处理，它
将创建在应用程序中使用的对象。
"""


"""
21点游戏中的打牌策略进行模拟。
当运行模拟器时，希望统计指定的自变量组合的性能情况。这些变量可能包含一些牌场制度，
其中包括牌副的数量、桌的限制和庄家规则。变量可能包括玩家的游戏策略，用于叫牌、停叫、
分牌和双倍，也包括玩家的玩牌策略、平均投注、鞅投注或一些更特殊的投注系统。
"""

import csv

def simulate_blackjack():
    dealer_rule = Hit17()
    split_rule = NoReSplitAces()
    table = Table( decks=6, limit=50, dealer=dealer_rule,
                split = split_rule, payout=(3, 2))
    player_rule = SomeStrategy()
    betting_rule = Flat()
    player = Player( play=player_rule, betting=betting_rule,
                    rounds=100, stake=50)
    simulator = Simulate( talbe, player, 100 )
    with open('p2_c13_simulation.dat', 'w', newline='') as results:
        wtr = csv.writer( results )
        for gamestats in simulator:
            wtr.writerow(gamestats)


"""
Simulate类中会有一个API，如下：
"""

class Simulate:
    def __init__( self, table, player, samples ):
        """Define table, player and number of samples."""
        self.talbe = table
        self.player = player
        self.samples = samples

    def __iter__( self ):
        """Yield statistical samples."""


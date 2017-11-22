#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 设计命令类 '''

# 下面是抽象的命令基类：
class Command:
    def set_config( self, config ):
        self.__dict__.update(config.__dict__)
    
    config = property(fset=set_config)

    def run( self ):
        pass

# 一旦对象配置完成，我们就可以通过调用run()方法来设置它并开始执行命令工作
"""
如:
main = SomeCommand()
main.config = config
main.run()
"""

import ast
import csv

# 下面是一个实现21点模拟操作的具体子类。
class Simulate_Command( Command ):
    dealer_rule_map = {'Hit17': Hit17, 'Stand17': Stand17}
    split_rule_map = {'ReSplit': ReSplit, 'NoReSplit': NoReSplit, 
    'NoReSplitAces': NoReSplitAces}
    player_rule_map = {'SomeStrategy': SomeStrategy,
    'AnotherStrategy': AnotherStrategy}
    betting_rule_map = {'Flat': Flat, 'Martingale': Martingale, 
    'OneThreeTwoSix': OneThreeTwoSix}

    def run( self ):
        dealer_rule = self.dealer_rule_map[self.dealer_rule]()
        split_rule = self.split_rule_map[self.split_rule]()
        try:
            payout = ast.literal_eval( slef.payout )
            assert len(payout) == 2
        except Exception as e:
            raise Exception('Invalid payout {0}'.format(self.payout))
        
        table = Table( decks=self.decks, limit=self.limit, 
        dealer=dealer_rule, split=split_rule, payout=payout)
        player_rule = self.player_rule_map[self.player_rule]()
        betting_rule = self.betting_rule_map[self.betting_rule]()
        player = Player( play=player_rule, betting=betting_rule, 
        rounds=self.rounds, stake=self.stake)
        simulate = Simulate(table, player, self.samples)
        with open( self.outputfile, 'w', newline='' ) as target:
            wtr = csv.write(target)
            wtr.writerows(simulate)


if __name__ == '__main__':
    with Logging_Config():
        with Application_Config() as config:
            main = Simulate_Command()
            main.config = config
            main.run()

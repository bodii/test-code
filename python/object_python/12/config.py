#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' config.py '''

# SomeStrategy setup

# Table
dealer_rule = Hit17()
split_rule = NoReSplitAces()
table = Table( decks=6, limit=50, dealer=dealer_rule,
        split=split_rule, payout=(3,2))

# Player
player_rule = SomeStrategy()
betting_rule = Flat()
player = Player( play=player_rule, betting=betting_rule,
            rounds=100, stake=50)

# Simulation
outputfile = 'p2_c13_simulation6.dat'
sample = 100
#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用压平的JSON配置 '''

import json
from collections import ChainMap

js = {
    "player.betting": "Flat",
    "player.play": "SomeStrategy",
    'player.rounds': 100,
    "player.stake": 50,
    "table.decks": 6,
    "table.dealer": "Hit17",
    "table.limit": 50,
    "table.payout": [3, 2],
    "table.split": "NoResplitAces",
    "simulator.outputfile": "p2_c13_simulation.dat",
    "simulator.samples": 100
}

# config = ChainMap( *[json.load(js) for file in reversed(config_names)] )
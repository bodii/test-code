#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 带参数选项 '''

import argparse

parser = argparse.ArgumentParser('')
parser.add_argument('-b', '--bet', action='store', default='Flat',
choices=['Flat', 'Martingale', 'OneThreeTwoSix'], dest='betting_rule')

parser.add_argument('-s', '--stake', action='store', default=50, type=int)

args = parser.parse_args()

print( args )
print( vars(args) )
print( args.stake )


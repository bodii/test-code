#!/usr/bin/env python3
# -*- coding:utf-8 -*-


''' 所有其他参数 '''


import argparse

__version__ = '3.3.2'

parser = argparse.ArgumentParser()

parser.add_argument('-v', '--version', action='version', version=__version__)

args = parser.parse_args()

print( args )

# run 08.py -v
# result: 3.3.2


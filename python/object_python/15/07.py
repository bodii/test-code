#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 位置参数 '''


import argparse

parser = argparse.ArgumentParser(prog='file')

parser.add_argument('input_filename', action='store')
parser.add_argument('output_filename', action='store')

args = parser.parse_args()

print( args )

# run: 15/07.py 15/05.py 15/06.py
# result:Namespace(input_filename='15/05.py', output_filename='15/06.py')
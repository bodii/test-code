#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用日志记录器，作为参数 '''

import argparse
import logging

parser = argparse.ArgumentParser(description='process some logging')
parser.add_argument('--debug', action='store_const', const=logging.DEBUG,
default=logging.INFO, dest='logging_level')

args = parser.parse_args()

print( args )
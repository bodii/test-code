#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import os
import sys
import argparse


''' 用argparse解析命令行 '''

"""
使用argparse包含以下4个步骤：
1. 创建ArgumentParser。
2. 定义命令行选项和参数。可以通过用ArgumentParse.add_argument()方法函数添加参数
来定义。
3. 通过解析sys.argv命令行创建用于描述选项、选项参数和总体命令行参数的命名空间对象。
4. 用创建好的命令空间对象配置应用程序和处理参数。
"""

# 创建ArgumentParser解析对象
parser = argparse.ArgumentParser(description='Process some integers.')
# 添加参数
parser.add_argument('integers', metavar='N', type=int, nargs='+',
                    help='an integer for the accumulator'
                    )

parser.add_argument('--sum', dest='accumulate', action='store_const',
const=sum, default=max, help='sum the integers(default: find the max)')

# ./02.py 1 2 3 4 --sum  [ return 10 ]
args = parser.parse_args()
print(args.accumulate(args.integers))
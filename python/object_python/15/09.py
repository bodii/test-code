#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 集成命令行选项和环境变量 '''

"""
在定义命令行参数时，显式设置值：这样做的优点是可以在帮助消息中显示默认值。它只针对环境变量
和命令行选项重合的情况。我们可能希望用SIM_SAMPLES环境变量提供一个可以被覆盖的默认值。

parser.add_argument('--samples', action='store', default=
int(os.environ.get('SIM_SAMPLES', 100)), type=int, help='Samples to generate')


在解析过程中隐式设置值：
config4 = argparse.Namespace()
config4.samples = int(os.environ.get('SIM_SAMPLES', 100))
config4a = parser.parse_args(namespace=config4)
"""

''' 让配置文件理解None '''
import os

def nint( x ):
    if x is None: return x
    return int(x)

env_values = [
    ('samples', nint(os.environ.get('SIM_SAMPLES', None))),
    ('stake', nint(os.environ.get('SIM_STAKE', None))),
    ('rounds', nint(os.environ.get('SIM_ROUNDS', None))),
]
print( env_values )
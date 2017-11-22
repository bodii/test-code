#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用ChainMap '''

"""
将映射连接在一起的场景非常符合Python中关于本地与全局的概念。当我们在Python中使用
一个变量时，会按照先本地命名空间、后全局命名空间的顺序搜索这个变量。除了搜索这两个
命名空间外，Python还会在本地命名空间中设置一个变量，但是这个变量不会影响到全局命名
空间。这种默认的行为（没有使用global或者nonlocal语句）也是ChainMap的工作方式。
当我们的程序开始运行时，我们常常会从命令行参数、配置文件、操作系统环境变量甚至有可
能是基于安装范围的配置中获取属性。我们希望可以将这些属性整合到一个类似于字典的结构中，、
这样我们就可以轻易地对配置进行定位。
"""

import argparse
import json
import os

parser = argparse.ArgumentParser(description='''Process some
intgeers.''')
parser.add_argument( '-c', '--configuration', type=open, nargs='?')
parser.add_argument( '-p', '--playerclass', type=str, nargs='?',
 default='Simple' )

if cmdline.configuration:
    config_file = json.load( options.configuration )
    options.configuration.close()
else:
    config_file = {}

with open('defaults.json') as installation:
    defaults = json.load( installation )


from collections import ChainMap
options = ChainMap( vars(cmdline), config_file, os.environ, defaults )

"""
上面的代码我们展示不同来源的配置：
1. 命令行参数。 如令牌参数( playerclass\configuration ) 
2. defautls.json文件以及另外一个可以寻找配置值的地方

基于上面的代码，我们可以建立一个单一的ChainMap对象使用案例，在这个案例中，我们可以
在指定的位置中查找参数。ChainMap实例会按顺序在每一个映射中搜索指定值。这为我们提供
了一种简洁、易于使用的处理运行时选项和参数的方法。
"""

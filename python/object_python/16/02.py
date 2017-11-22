#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 全局模块和模块项 '''


# from cards import *

import cards

# 创建Card类
print(cards.Card())

''' 包的设计 '''

"""
一个包由一个目录和一个__init__.py文件组成。目录名称必须为适当的Python名称，OS
的名称中包含了很多在Python的命名中不允许使用的字符。

通常使用以下3种方式来进行包的设计：
1. 对于简单的包，使用一个目录和一个空的__init__.py文件。这个包的名称将被作为内部
模块名称的限定词，如： import package.module

2. 一个模块中的包可以包含一个__init__.py文件作为模块的定义，可以从包目录中导入其他
模块。或者，它可以作为包含了最上层模块和被限定的子模块的大规模设计中的一部分。可以使用
以下代码进行导入：import package

3. 目录是__init__.py文件中可替代的一种实现方法，如使用：import package
"""
# try:
#     from some_algorithm.long_version import *
# except ImportError as e:
#     from some_algorithm.short_version import *



''' 主脚本和__main__模块的设计 '''
from wsgiref.util import setup_testing_defaults
from wsgiref.simple_server import make_server

def simple_app(environ, start_response):
    setup_testing_defaults(environ)

    status = '200 OK'
    headers = [('Contest-teyp', 'text/plain; charset=utf-8')]

    start_response(status, headers)

    ret = [("%s: %s\n" % (key, value)).encode('utf-8') 
            for key, value in environ.items()]
    return ret

with make_server('', 8000, simple_app) as httpd:
    print('Serving on port 8000...')
    httpd.serve_forever()

import multiprocessing

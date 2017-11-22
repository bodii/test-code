#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 在配置中使用Python的exec() '''

import os.path

current_path = os.path.dirname(__file__) + os.path.sep

with open(current_path + 'config.py') as py_file:
    code = compile(py_file.read(), 'config.py', 'exec')
config = {}
exec( code, globals(), config )
simulate( config['table'], config['player'],
        config['outputfile'], config['samples'])

# compile()函数创建对象
# exec()函数直接执行代码
"""
exec()函数提供了3个参数：代码、存放全局变量的字典以及用于存放将要被创建的本地变量
的字典。代码块执行后，随着赋值语句的执行，创建了本地变量字典。
"""

# 由于字典格式不方便，因此可使用一种设计模式
class AttrDict( dict ):
    def __getattr__( self, name ):
        return self.get( name, None )
    def __setattr__( self, name, value ):
        self[name] = value
    def __dir__( self ):
        return list(self.keys())

# 在使用这个类时，仅当键和Python变量名匹配时才可以。
config = AttrDict()

# 然后，可以使用一种简单的属性格式（config.table、config.player)来实现对象的创建
# 初始化。
class Configuration:
    def __init__( self, **kw ):
        self.__dict__.update(kw)

# 然后就可以使用命名的属性来将简单的dict转换为对象：
config = Configuration( **config )


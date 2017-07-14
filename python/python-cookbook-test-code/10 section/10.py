#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十章：模块与包
Desc: 模块与包是任何大型程序的核心，就连 Python 安装程序本身也是一个包。本章重
点涉及有关模块和包的常用编程技术，例如如何组织包、把大型模块分割成多个文件、
创建命名空间包。同时，也给出了让你自定义导入语句的秘籍。

Title;               通过字符串名导入模块
Issue:你想导入一个模块，但是模块的名字在字符串里。你想对字符串调用导入命令。
Answer: 使用 importlib.import module() 函数来手动导入名字为字符串给出的一个模块或者
包的一部分。
"""

import importlib

math = importlib.import_module('math')
print(math.sin(2))
# 0.9092974268256817
mod = importlib.import_module('urllib.request')
#u = mod.urlopen('http://www.python.org')

"""
import module 只是简单地执行和 import 相同的步骤，但是返回生成的模块对象。
你只需要将其存储在一个变量，然后像正常的模块一样使用。
如果你正在使用的包， import module() 也可用于相对导入。但是，你需要给它一
个额外的参数。例如
"""
import importlib
# Same as 'from . import b'
b = importlib.import_module('.b',__package__)

"""
使用 import module() 手动导入模块的问题通常出现在以某种方式编写修改或覆盖
模块的代码时候。例如，也许你正在执行某种自定义导入机制，需要通过名称来加载
一个模块，通过补丁加载代码。
在旧的代码，有时你会看到用于导入的内建函数 import ()。尽管它能工作，但是
importlib.import module() 通常更容易使用。
自定义导入过程的高级实例见 10.11 小节
"""
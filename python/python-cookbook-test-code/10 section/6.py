#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十章：模块与包
Desc: 模块与包是任何大型程序的核心，就连 Python 安装程序本身也是一个包。本章重
点涉及有关模块和包的常用编程技术，例如如何组织包、把大型模块分割成多个文件、
创建命名空间包。同时，也给出了让你自定义导入语句的秘籍。

Title;               重新加载模块
Issue:你想重新加载已经加载的模块，因为你对其源码进行了修改。
Answer: 使用 imp.reload() 来重新加载先前加载的模块。
"""
"""
import spam
import imp

imp.reload(spam)
# <module 'spam' from './spam.py'>

重新加载模块在开发和调试过程中常常很有用。但在生产环境中的代码使用会不安
全，因为它并不总是像您期望的那样工作。
reload() 擦除了模块底层字典的内容，并通过重新执行模块的源代码来刷新它。模
块对象本身的身份保持不变。因此，该操作在程序中所有已经被导入了的地方更新了
模块。
尽管如此， reload() 没有更新像”from module import name” 这样使用 import 语句
导入的定义。举个例子：
# spam.py
def bar():
    print('bar')

def grok():
    print('grok')

现在启动交互式会话：

import spam
from spam import grok
spam.bar()
# bar

grok()
# grok

不退出 Python 修改 spam.py 的源码，将 grok() 函数改成这样：
def grok():
    print('New grok')

现在回到交互式会话，重新加载模块，尝试下这个实验：
import imp
imp.reload(spam)
# <module 'spam' from './spam.py'>

spam.bar()
bar

grok()
#grok

spam.grok()
# New grok

在这个例子中，你看到有 2 个版本的 grok() 函数被加载。通常来说，这不是你想
要的，而是令人头疼的事。
因此，在生产环境中可能需要避免重新加载模块。在交互环境下调试，解释程序并
试图弄懂它。

"""
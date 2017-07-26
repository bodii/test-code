#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十章：模块与包
Desc: 模块与包是任何大型程序的核心，就连 Python 安装程序本身也是一个包。本章重
点涉及有关模块和包的常用编程技术，例如如何组织包、把大型模块分割成多个文件、
创建命名空间包。同时，也给出了让你自定义导入语句的秘籍。

Title;                构建一个模块的层级包
Issue:你想将你的代码组织成由很多分层模块构成的包。
Answer: 封装成包是很简单的。在文件系统上组织你的代码，并确保每个目录都定义了一
个 __init__.py 文件。
"""

""" 例如:
graphics
    __init__.py
    primitive
        __init__.py
        line.py
        fill.py
        text.py
    formats
        __init__.py
        png.py
        jpg.py

一旦你做到了这一点，你应该能够执行各种 import 语句，如下：
import graphices.primitive.line
from graphices.primitive import line'
import graphices.formats.jpg as jpg

定义模块的层次结构就像在文件系统上建立目录结构一样容易。文件 __init__.py 的
目的是要包含不同运行级别的包的可选的初始化代码。举个例子，如果你执行了语句
import graphics，文件 graphics/__init__.py 将被导入, 建立 graphics 命名空间的内容。
像 import graphics.format.jpg 这样导入，文件 graphics/__init__ .py 和文件 graphics/
graphics/formats/__init__.py 将在文件 graphics/formats/jpg.py 导入之前导入。
绝大部分时候让__init__.py 空着就好。但是有些情况下可能包含代码。举个例子，
__init__.py 能够用来自动加载子模块:
# graphics/formats/__init__.py
from . import jpg
from . import png

像这样一个文件, 用户可以仅仅通过 import grahpics.formats 来代替 import graphics.formats.jpg 以及 import graphics.formats.png。

像这样一个文件, 用户可以仅仅通过 import grahpics.formats 来代替 import graphics.formats.jpg 以及 import graphics.formats.png。
__init__.py 的其他常用用法包括将多个文件合并到一个逻辑命名空间，这将在 10.4
小节讨论。
敏锐的程序员会发现，即使没有 __init__.py 文件存在， python 仍然会导入包。如果
你没有定义 __init__.py 时，实际上创建了一个所谓的“命名空间包”，这将在 10.5 小节
讨论。万物平等，如果你着手创建一个新的包的话，包含一个 __init__.py 文件吧。
 """

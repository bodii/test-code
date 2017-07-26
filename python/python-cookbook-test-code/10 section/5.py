#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十章：模块与包
Desc: 模块与包是任何大型程序的核心，就连 Python 安装程序本身也是一个包。本章重
点涉及有关模块和包的常用编程技术，例如如何组织包、把大型模块分割成多个文件、
创建命名空间包。同时，也给出了让你自定义导入语句的秘籍。

Title;               利用命名空间导入目录分散的代码
Issue:你可能有大量的代码，由不同的人来分散地维护。每个部分被组织为文件目录，如
一个包。然而，你希望能用共同的包前缀将所有组件连接起来，不是将每一个部分作
为独立的包来安装。
Answer: 在统一不同的目录里统一相同的命名空间，但是要删去用来将组件联合起来
的 init .py 文件。
"""

"""
从本质上讲，你要定义一个顶级 Python 包，作为一个大集合分开维护子包的命名
空间。这个问题经常出现在大的应用框架中，框架开发者希望鼓励用户发布插件或附
加包。
在统一不同的目录里统一相同的命名空间，但是要删去用来将组件联合起来
的 init .py 文件。假设你有 Python 代码的两个不同的目录如下：
foo package
    spam
        blah.py

bar package
    spam
        grok.py

在这 2 个 目 录 里， 都 有 着 共 同 的 命 名 空 间 spam。 在 任 何 一 个 目 录 里 都 没
有 init .py 文件。
让我们看看，如果将 foo-package 和 bar-package 都加到 python 模块路径并尝试导
入会发生什么
import sys
sys.path.extend({'foo-package','bar-package'})
import spam.blah
import spam.grok

两个不同的包目录被合并到一起，你可以导入 spam.blah 和 spam.grok，并且它们
能够工作。

在这里工作的机制被称为“包命名空间”的一个特征。从本质上讲，包命名空间是
一种特殊的封装设计，为合并不同的目录的代码到一个共同的命名空间。对于大的框
架，这可能是有用的，因为它允许一个框架的部分被单独地安装下载。它也使人们能
够轻松地为这样的框架编写第三方附加组件和其他扩展。
包命名空间的关键是确保顶级目录中没有 init .py 文件来作为共同的命名空间。
缺失 init .py 文件使得在导入包的时候会发生有趣的事情：这并没有产生错误，解释
器创建了一个由所有包含匹配包名的目录组成的列表。特殊的包命名空间模块被创建，
只读的目录列表副本被存储在其 path 变量中。举个例子：
import spam
print(spam.__path__)
# _NamespacePath(['foo-package/spam', 'bar-package/spam'])

在定位包的子组件时，目录 path 将被用到 (例如, 当导入 spam.grok 或者
spam.blah 的时候).
包命名空间的一个重要特点是任何人都可以用自己的代码来扩展命名空间。举个例
子，假设你自己的代码目录像这样：
my-package
    spam
        custom.py
如果你将你的代码目录和其他包一起添加到 sys.path，这将无缝地合并到别的
spam 包目录中：
import spam.custom
import spam.grok
import spam.blah

一个包是否被作为一个包命名空间的主要方法是检查其 file 属性。如果没有，那
包是个命名空间。这也可以由其字符表现形式中的“ namespace”这个词体现出来。
spam.__file__
# Traceback (most recent call last):
# File "<stdin>", line 1, in <module>
# AttributeError: 'module' object has no attribute '__file__'
spam
# <module 'spam' (namespace)>


更多的包命名空间信息可以查看 PEP 420

"""

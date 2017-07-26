#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十章：模块与包
Desc: 模块与包是任何大型程序的核心，就连 Python 安装程序本身也是一个包。本章重
点涉及有关模块和包的常用编程技术，例如如何组织包、把大型模块分割成多个文件、
创建命名空间包。同时，也给出了让你自定义导入语句的秘籍。

Title;               使用相对路径名导入包中子模块
Issue:将代码组织成包, 想用 import 语句从另一个包名没有硬编码过的包的中导入子模
块。
Answer: import 语句的 . 和‘‘..‘‘看起来很滑稽, 但它指定目录名. 为当前目录， ..B 为目录../
B。这种语法只适用于 import。
"""

"""
mypackage
    __init__.py
    A
        __init__.py
        spam.py
        grok.py
    B
        __init__.py
        bar.py

如果模块 mypackage.A.spam 要导入同目录下的模块 grok，它应该包括的 import
语句如下：
from . import grok

如果模块 mypackage.A.spam 要导入不同目录下的模块 B.bar，它应该使用的 import 语句如下：
from ..B import bar

两个 import 语句都没包含顶层包名，而是使用了 spam.py 的相对路径。

在包内，既可以使用相对路径也可以使用绝对路径来导入。举个例子：
from mypackage.A import grok
from . import grok
import grok

像 mypackage.A 这样使用绝对路径名的不利之处是这将顶层包名硬编码到你的源
码中。如果你想重新组织它，你的代码将更脆，很难工作。举个例子，如果你改变了
包名，你就必须检查所有文件来修正源码。同样，硬编码的名称会使移动代码变得困
难。举个例子，也许有人想安装两个不同版本的软件包，只通过名称区分它们。如果
使用相对导入，那一切都 ok，然而使用绝对路径名很可能会出问题。
import 语句的 . 和‘‘..‘‘看起来很滑稽, 但它指定目录名. 为当前目录， ..B 为目录../
B。这种语法只适用于 import。举个例子：
from . import grok
import .grok

尽管使用相对导入看起来像是浏览文件系统，但是不能到定义包的目录之外。也就
是说，使用点的这种模式从不是包的目录中导入将会引发错误。
最后，相对导入只适用于在合适的包中的模块。尤其是在顶层的脚本的简单模块
中，它们将不起作用。如果包的部分被作为脚本直接执行，那它们将不起作用例如：
% python3 mypackage/A/spam.py

另一方面，如果你使用 Python 的 -m 选项来执行先前的脚本，相对导入将会正确
运行。
% python3 -m mypackage.A.spam
更多的包的相对导入的背景知识, 请看 PEP 328 .
"""





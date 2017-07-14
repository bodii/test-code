#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十章：模块与包
Desc: 模块与包是任何大型程序的核心，就连 Python 安装程序本身也是一个包。本章重
点涉及有关模块和包的常用编程技术，例如如何组织包、把大型模块分割成多个文件、
创建命名空间包。同时，也给出了让你自定义导入语句的秘籍。

Title;               安装私有的包
Issue:你想要安装一个第三方包，但是没有权限将它安装到系统 Python 库中去。或者，
你可能想要安装一个供自己使用的包，而不是系统上面所有用户。
Answer: Python 有一个用户安装目录，通常类似”˜/.local/lib/python3.3/site-packages”。要
强制在这个目录中安装包，可使用安装选项“ –user”。
"""
"""
python3 setup.py install --user

或是
pip install --user packagename

在 sys.path 中用户的“ site-packages”目录位于系统的“ site-packages”目录之前。
因此，你安装在里面的包就比系统已安装的包优先级高（尽管并不总是这样，要取决
于第三方包管理器，比如 distribute 或 pip）。

通常包会被安装到系统的 site-packages 目录中去，路径类似“/usr/local/lib/
python3.3/site-packages”。不过，这样做需要有管理员权限并且使用 sudo 命令。就
算你有这样的权限去执行命令，使用 sudo 去安装一个新的，可能没有被验证过的包有
时候也不安全。
安装包到用户目录中通常是一个有效的方案，它允许你创建一个自定义安装。
另外，你还可以创建一个虚拟环境，这个我们在下一节会讲到。
"""
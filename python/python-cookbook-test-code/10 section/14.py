#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十章：模块与包
Desc: 模块与包是任何大型程序的核心，就连 Python 安装程序本身也是一个包。本章重
点涉及有关模块和包的常用编程技术，例如如何组织包、把大型模块分割成多个文件、
创建命名空间包。同时，也给出了让你自定义导入语句的秘籍。

Title;               创建新的 Python 环境
Issue:你想创建一个新的 Python 环境，用来安装模块和包。不过，你不想安装一个新的
Python 克隆，也不想对系统 Python 环境产生影响。
Answer: 你可以使用 pyvenv 命令创建一个新的“虚拟”环境。这个命令被安装在 Python
解释器同一目录，或 Windows 上面的 Scripts 目录中。
"""
"""
bash % pyvenv Spam
bash %
传给 pyvenv 命令的名字是将要被创建的目录名。当被创建后， Span 目录像下面这
样：
bash % cd Spam
bash % ls
bin include lib pyvenv.cfg
bash %

在 bin 目录中，你会找到一个可以使用的 Python 解释器：
bash % Spam/bin/python3
3
Python 3.3.0 (default, Oct 6 2012, 15:45:22)
[GCC 4.2.1 (Apple Inc. build 5666) (dot 3)] on darwin
Type "help", "copyright", "credits" or "license" for more information.
>>> from pprint import pprint
>>> import sys
>>> pprint(sys.path)
['',
'/usr/local/lib/python33.zip',
'/usr/local/lib/python3.3',
'/usr/local/lib/python3.3/plat-darwin',
'/usr/local/lib/python3.3/lib-dynload',
'/Users/beazley/Spam/lib/python3.3/site-packages']
这个解释器的特点就是他的 site-packages 目录被设置为新创建的环境。如果你要
安装第三方包，它们会被安装在那里，而不是通常系统的 site-packages 目录。
创建虚拟环境通常是为了安装和管理第三方包。正如你在例子中看到的那样，
sys.path 变量包含来自于系统 Python 的目录，而 site-packages 目录已经被重定位到
一个新的目录。
有了一个新的虚拟环境，下一步就是安装一个包管理器，比如 distribute 或 pip。
但安装这样的工具和包的时候，你需要确保你使用的是虚拟环境的解释器。它会将包
安装到新创建的 site-packages 目录中去。
尽管一个虚拟环境看上去是 Python 安装的一个复制，不过它实际上只包含了少量
几个文件和一些符号链接。素有标准库函文件和可执行解释器都来自原来的 Python 安
装。因此，创建这样的环境是很容易的，并且几乎不会消耗机器资源。
默认情况下，虚拟环境是空的，不包含任何额外的第三方库。如果你想将一个已经
安装的包作为虚拟环境的一部分，可以使用“ –system-site-packages”选项来创建虚拟
环境，例如：
bash % pyvenv --system-site-packages Spam
bash %
跟多关于 pyvenv 和虚拟环境的信息可以参考 PEP 405
"""

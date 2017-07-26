#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十章：模块与包
Desc: 模块与包是任何大型程序的核心，就连 Python 安装程序本身也是一个包。本章重
点涉及有关模块和包的常用编程技术，例如如何组织包、把大型模块分割成多个文件、
创建命名空间包。同时，也给出了让你自定义导入语句的秘籍。

Title;               将文件夹加入到 sys.path
Issue:你无法导入你的 Python 代码因为它所在的目录不在 sys.path 里。你想将添加新目
录到 Python 路径，但是不想硬链接到你的代码。
Answer: 有两种常用的方式将新目录添加到 sys.path。第一种，你可以使用 PYTHONPATH
环境变量来添加。第二种方法是创建一个.pth 文件，将目录列举出来
"""
"""
有两种常用的方式将新目录添加到 sys.path。第一种，你可以使用 PYTHONPATH
环境变量来添加。例如：
bash % env PYTHONPATH = /some/dir:/other/dir python3
Python 3.3.0 (default, Oct 4 2012, 10:17:33)
[GCC 4.2.1 (Apple Inc. build 5666) (dot 3)] on darwin
Type "help", "copyright", "credits" or "license" for more information.

"""
import sys
print(sys.path)
# ['', '/some/dir', '/other/dir', ...]
"""
在自定义应用程序中，这样的环境变量可在程序启动时设置或通过 shell 脚本。
第二种方法是创建一个.pth 文件，将目录列举出来，像这样：
# myapplication.pth
/some/dir
/other/dir

这个.pth 文件需要放在某个 Python 的 site-packages 目录，通常位于/usr/local/
lib/python3.3/site-packages 或者 ˜/.local/lib/python3.3/sitepackages。当解释器启动
时， .pth 文件里列举出来的存在于文件系统的目录将被添加到 sys.path。安装一个.pth
文件可能需要管理员权限，如果它被添加到系统级的 Python 解释器。

比起费力地找文件，你可能会倾向于写一个代码手动调节 sys.path 的值。例如:
import sys
sys.path.insert(0, '/some/dir')
sys.path.insert(0, '/other/dir')

虽然这能“工作”，它是在实践中极为脆弱，应尽量避免使用。这种方法的问题是，
它将目录名硬编码到了你的源。如果你的代码被移到一个新的位置，这会导致维护问
题。更好的做法是在不修改源代码的情况下，将 path 配置到其他地方。如果您使用模
块级的变量来精心构造一个适当的绝对路径，有时你可以解决硬编码目录的问题，比
如 file 。举个例子
import sys
from os.path import abspath, join, dirname
sys.path.insert(0, abspath(dirname('__file__'), 'src'))

这将 src 目录添加到 path 里，和执行插入步骤的代码在同一个目录里。
site-packages 目录是第三方包和模块安装的目录。如果你手动安装你的代码，它将
被安装到 site-packages 目录。虽然.pth 文件配置的 path 必须出现在 site-packages 里，
但代码可以在系统上任何你想要的目录。因此，你可以把你的代码放在一系列不同的
目录，只要那些目录包含在.pth 文件里。

"""
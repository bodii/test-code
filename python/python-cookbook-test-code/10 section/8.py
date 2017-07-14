#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十章：模块与包
Desc: 模块与包是任何大型程序的核心，就连 Python 安装程序本身也是一个包。本章重
点涉及有关模块和包的常用编程技术，例如如何组织包、把大型模块分割成多个文件、
创建命名空间包。同时，也给出了让你自定义导入语句的秘籍。

Title;               读取位于包中的数据文件
Issue:你的包中包含代码需要去读取的数据文件。你需要尽可能地用最便捷的方式来做这
件事。
Answer: import pkgutil
data = pkgutil.get_data(__package__, 'somedata.dat')
"""
"""
假设你的包中的文件组织成如下
mypackage
    __init__.py
    somedata.dat
    spam.py
现在假设 spam.py 文件需要读取 somedata.dat 文件中的内容。你可以用以下代码
来完成：
# spam.py
import pkgutil
data = pkgutil.get_data(__package__, 'somedata.dat')
由此产生的变量是包含该文件的原始内容的字节字符串。
要读取数据文件，你可能会倾向于编写使用内置的 I/ O 功能的代码，如 open()。
但是这种方法也有一些问题。

0
首先，一个包对解释器的当前工作目录几乎没有控制权。因此，编程时任何 I/O 操
作都必须使用绝对文件名。由于每个模块包含有完整路径的 file 变量，这弄清楚它的
路径不是不可能，但它很凌乱。
第二，包通常安装作为.zip 或.egg 文件，这些文件像文件系统上的一个普通目录一
样不会被保留。因此，你试图用 open() 对一个包含数据文件的归档文件进行操作，它
根本不会工作。
pkgutil.get data() 函数是一个读取数据文件的高级工具，不用管包是如何安装以及
安装在哪。它只是工作并将文件内容以字节字符串返回给你
get data() 的第一个参数是包含包名的字符串。你可以直接使用包名，也可以使用
特殊的变量，比如 package 。第二个参数是包内文件的相对名称。如果有必要，可以
使用标准的 Unix 命名规范到不同的目录，只有最后的目录仍然位于包中。
"""
#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十三章：脚本编程与系统管理
Description: 许多人使用 Python 作为一个 shell 脚本的替代，用来实现常用系统任务的自动化，
如文件的操作，系统的配置等。本章的主要目标是描述光宇编写脚本时候经常遇到的
一些功能。例如，解析命令行选项、获取有用的系统配置数据等等。第 5 章也包含了
与文件和目录相关的一般信息。

Title:            复制或者移动文件和目录
Issue:你想哟啊复制或移动文件和目录，但是又不想调用 shell 命令。
Answer: shutil 模块有很多便捷的函数可以复制文件和目录。使用起来非常简单，
"""

import shutil
'''
# copy src to dst.(cp src dst)
shutil.copy(src, dst)
# copy files, but preserve metadata(cp -p src dst)
shutil.copy2(src.dst)
# copy directory tree( cp -R src dst)
shutil.copytree(src, dst)
# move src to dst(mv src dst)
shutil.move(src, dst)

"""
这些函数的参数都是字符串形式的文件或目录名。底层语义模拟了类似的 Unix 命
令，如上面的注释部分。
默认情况下，对于符号链接而已这些命令处理的是它指向的东西。例如，如果源文
件是一个符号链接，那么目标文件将会是符号链接指向的文件。如果你只想复制符号
链接本身，那么需要指定关键字参数 follow symlinks , 如下：
如果你想保留被复制目录中的符号链接，像这样做：
"""
shutil.copytree(src, dst, symlinks=True)

"""
copytree() 可以让你在复制过程中选择性的忽略某些文件或目录。你可以提供一
个忽略函数，接受一个目录名和文件名列表作为输入，返回一个忽略的名称列表。例
如：
"""
def ignore_pyc_files(dirname, filenames):
    return [name in filenames if name.endswith('.pyc')]
shutil.copytree(src, dst, ignorg=ignorg_pyc_files)

"""
copytree() 可以让你在复制过程中选择性的忽略某些文件或目录。你可以提供一
个忽略函数，接受一个目录名和文件名列表作为输入，返回一个忽略的名称列表。例
如：
"""
'''
filename = '/Users/guido/programs/spam.py'
import os.path
os.path.basename(filename)
# 'spam.py'
os.path.dirname(filename)
# /Users/guido/programs
os.path.split(filename)
# ('/Users/guido/programs', 'spam.py')
os.path.join('/new/dir', os.path.basename(filename))
# /new/dir/spam.py
os.psth.expanduser('~/guido/programs/spam.py')
# /Users/guido/programs/spam.py
"""
使用 copytree() 复制文件夹的一个棘手的问题是对于错误的处理。例如，在复制
过程中，函数可能会碰到损坏的符号链接，因为权限无法访问文件的问题等等。为了
解决这个问题，所有碰到的问题会被收集到一个列表中并打包为一个单独的异常，到
了最后再抛出。
"""
try:
    shutil.copytree(src, dst)
except shutil.Error as e:
    for src, dst, msg in e.args[0]:
        print(dst, src, msg)

"""
如果你提供关键字参数 ignore dangling symlinks=True ，这时候 copytree() 会
忽略掉无效符号链接。
本节演示的这些函数都是最常见的。不过， shutil 还有更多的和复制数据相关的
操作。它的文档很值得一看，参考 Python documentation
"""
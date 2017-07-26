#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十三章：脚本编程与系统管理
Description: 许多人使用 Python 作为一个 shell 脚本的替代，用来实现常用系统任务的自动化，
如文件的操作，系统的配置等。本章的主要目标是描述光宇编写脚本时候经常遇到的
一些功能。例如，解析命令行选项、获取有用的系统配置数据等等。第 5 章也包含了
与文件和目录相关的一般信息。

Title:            通过重定向/管道/文件接受输入
Issue:你希望你的脚本接受任何用户认为最简单的输入方式。包括将命令行的输出通过管
道传递给该脚本、重定向文件到该脚本，或在命令行中传递一个文件名或文件名列表
给该脚本。
Answer: Python 内置的 fileinput 模块让这个变得简单。
"""
import fileinput

with fileinput.input() as f_input:
    for line in f_input:
        print(line, end='')

"""
那么你就能以前面提到的所有方式来为此脚本提供输入。假设你将此脚本保存为
filein.py 并将其变为可执行文件，那么你可以像下面这样调用它，得到期望的输出：
$ ls | ./filein.py
$ ./filein.py /etc/passwd
$ ./filein.py < /etc/passwd

fileinput.input() 创建并返回一个 FileInput 类的实例。该实例除了拥有一些
有用的帮助方法外，它还可被当做一个上下文管理器使用。因此，整合起来，如果我
们要写一个打印多个文件输出的脚本，那么我们需要在输出中包含文件名和行号，如
下所示

import fileinput
with fileinput.input('/etc/passwd') as f:
    for line in f:
        print(f.filename(), f.lineno(), line, end='')
/etc/passwd 1 ##
/etc/passwd 2 # User Database
/etc/passwd 3 #
通过将它作为一个上下文管理器使用，可以确保它不再使用时文件能自动关闭，而
且我们在之类还演示了 FileInput 的一些有用的帮助方法来获取输出中的一些其他信
息。
"""



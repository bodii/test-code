#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十三章：脚本编程与系统管理
Description: 许多人使用 Python 作为一个 shell 脚本的替代，用来实现常用系统任务的自动化，
如文件的操作，系统的配置等。本章的主要目标是描述光宇编写脚本时候经常遇到的
一些功能。例如，解析命令行选项、获取有用的系统配置数据等等。第 5 章也包含了
与文件和目录相关的一般信息。

Title:            获取终端的大小
Issue:你需要知道当前终端的大小以便正确的格式化输出。
Answer: 使用 os.get terminal size() 函数来做到这一点。
"""
import os
sz = os.get_terminal_size()
print(sz)
print(sz.columns)
print(sz.lines)

"""
有太多方式来得知终端大小了，从读取环境变量到执行底层的 ioctl() 函数等等。
不过，为什么要去研究这些复杂的办法而不是仅仅调用一个简单的函数呢？
"""

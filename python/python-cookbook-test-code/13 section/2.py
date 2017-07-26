#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十三章：脚本编程与系统管理
Description: 许多人使用 Python 作为一个 shell 脚本的替代，用来实现常用系统任务的自动化，
如文件的操作，系统的配置等。本章的主要目标是描述光宇编写脚本时候经常遇到的
一些功能。例如，解析命令行选项、获取有用的系统配置数据等等。第 5 章也包含了
与文件和目录相关的一般信息。

Title:            终止程序并给出错误信息
Issue:你想向标准错误打印一条消息并返回某个非零状态码来终止程序运行
Answer: 抛出一个 SystemExit 异常
"""

"""
你有一个程序像下面这样终止，抛出一个 SystemExit 异常，使用错误消息作为参
数。
raise SystemExit('It failed!')

它会将消息在 sys.stderr 中打印，然后程序以状态码 1 退出
本节虽然很短小，但是它能解决在写脚本时的一个常见问题。也就是说，当你想要
终止某个程序时，你可能会像下面这样写：
import sys
sys.stderr.write('It failed!\n')
raise SystemExit(1)
如果你直接将消息作为参数传给 SystemExit() ，那么你可以省略其他步骤，比如
import 语句或将错误消息写入 sys.stderr
"""
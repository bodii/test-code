#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十三章：脚本编程与系统管理
Description: 许多人使用 Python 作为一个 shell 脚本的替代，用来实现常用系统任务的自动化，
如文件的操作，系统的配置等。本章的主要目标是描述光宇编写脚本时候经常遇到的
一些功能。例如，解析命令行选项、获取有用的系统配置数据等等。第 5 章也包含了
与文件和目录相关的一般信息。

Title:            运行时弹出密码输入提示
Issue:你写了个脚本，运行时需要一个密码。此脚本是交互式的，因此不能将密码在脚本
中硬编码，而是需要弹出一个密码输入提示，让用户自己输入。
Answer: 这时候 Python 的 getpass 模块正是你所需要的。
"""
import getpass

user = getpass.getuser()
passwd = getpass.getpass()

if svc_login(user, passwd):
    print('Yay!')
else:
    print('Boo!')

"""
在此代码中， svc login() 是你要实现的处理密码的函数，具体的处理过程你自己
决定。

注意在前面代码中 getpass.getuser() 不会弹出用户名的输入提示。它会根据该
用户的 shell 环境或者会依据本地系统的密码库（支持 pwd 模块的平台）来使用当前用
户的登录名，
如果你想显示的弹出用户名输入提示，使用内置的 input 函数：
user = input('Enter your username: ')

还有一点很重要，有些系统可能不支持 getpass() 方法隐藏输入密码。这种情况
下， Python 会提前警告你这些问题（例如它会警告你说密码会以明文形式显示）
"""
#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用命令行 '''

"""
通常， shell会用一些构成OS API的信息来启动应用程序
1. shell会为每个应用程序提供环境变量的集合。在Python中，这些集合可以通过os.EnvironmentError
访问。
2. shell会准备3种标准文件。在Python中，这3种文件对应的是sys.stdin, sys.stdout和
sys.stderr。还有一些其他的模块，例如fileinput可以访问sys.stdin。
3. shell会将命令行解析为一些单词，命令行一部分可以通过sys.argv来访问。Python会提供
原始命令。对于POSIX操作系统，shell可能会替换shell的环境变量并展开通配符文件名。在Windows
中，简单的cmd.exe shell不会为我们展开文件名。
3. 操作系统也维护一些上下文设置，如当前工作目录、用户ID和用户组。这些可以通过os模块访问。


在应用程序中使用sys.exit().如果正常结束，Python会返回0。
"""

import os
import sys

# print( os.environ )

print( sys.stdin )
print( sys.stdout )
print( sys.stderr )

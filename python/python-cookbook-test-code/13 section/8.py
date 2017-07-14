#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十三章：脚本编程与系统管理
Description: 许多人使用 Python 作为一个 shell 脚本的替代，用来实现常用系统任务的自动化，
如文件的操作，系统的配置等。本章的主要目标是描述光宇编写脚本时候经常遇到的
一些功能。例如，解析命令行选项、获取有用的系统配置数据等等。第 5 章也包含了
与文件和目录相关的一般信息。

Title:           创建和解压归档文件
Issue:你需要创建或解压常见格式的归档文件（比如.tar, .tgz 或.zip）
Answer: shutil 模块拥有两个函数—— make archive() 和 unpack archive() 可派上用场。
"""
import shutil

shutil.unpack_archive('Python-3.3.0.tgz')
shutil.make_archive('py33','zip','Python-3.3.0')
# /Users/beazley/Downloads/py33.zip

"""
make archive() 的 第 二 个 参 数 是 期 望 的 输 出 格 式。 可 以 使 用
get archive formats() 获取所有支持的归档格式列表。例如：
shutill.get_archive_formats()
[('bztar', "bzip2'ed tar-file"), ('gztar', "gzip'ed tar-file"),
('tar', 'uncompressed tar file'), ('zip', 'ZIP file')]


Python 还有其他的模块可用来处理多种归档格式（比如 tarfile, zipfile, gzip, bz2）
的底层细节。不过，如果你仅仅只是要创建或提取某个归档，就没有必要使用底层库
了。可以直接使用 shutil 中的这些高层函数。
这些函数还有很多其他选项，用于日志打印、预检、文件权限等等。参考 shutil 文
档
"""
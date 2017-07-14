#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十三章：脚本编程与系统管理
Description: 许多人使用 Python 作为一个 shell 脚本的替代，用来实现常用系统任务的自动化，
如文件的操作，系统的配置等。本章的主要目标是描述光宇编写脚本时候经常遇到的
一些功能。例如，解析命令行选项、获取有用的系统配置数据等等。第 5 章也包含了
与文件和目录相关的一般信息。

Title:          启动一个 WEB 浏览器
Issue:你想通过脚本启动浏览器并打开指定的 URL 网页
Answer: webbrowser 模块能被用来启动一个浏览器，并且与平台无关。
"""
import webbrowser
webbrowser.open('http://www.python.org')

"""
它会使用默认浏览器打开指定网页。如果你还想对网页打开方式做更多控制，还可
以使用下面这些函数：
"""
webbrowser.open_new('http://www.python.org')

webbrowser.open_new_tab('http://www.python.org')

"""
这样就可以打开一个新的浏览器窗口或者标签，只要浏览器支持就行。
如果你想指定浏览器类型，可以使用 webbrowser.get() 函数来指定某个特定浏览
器。例如：
"""
c = webbrowser.get('IE')
c.open('http://www.python.org')
c.open_new_tab('http://docs.python.org')

"""
对于支持的浏览器名称列表可查阅 ‘Python 文档 <http://docs.python.org/3/library/webbrowser.html>
在脚本中打开浏览器有时候会很有用。例如，某个脚本执行某个服务器发布任务，
你想快速打开一个浏览器来确保它已经正常运行了。或者是某个程序以 HTML 网页格
式输出数据，你想打开浏览器查看结果。不管是上面哪种情况，使用 webbrowser 模块
都是一个简单实用的解决方案。
"""
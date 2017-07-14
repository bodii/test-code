#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十四章：测试、调试和异常
Description: 试验还是很棒的，但是调试？就没那么有趣了。事实是，在 Python 测试代码之前
没有编译器来分析你的代码，因此使的测试成为开发的一个重要部分。本章的目标是
讨论一些关于测试、调试和异常处理的常见问题。但是并不是为测试驱动开发或者单
元测试模块做一个简要的介绍。因此，笔者假定读者熟悉测试概念。

Title:                 测试 stdout 输出
Issue: 你的程序中有个方法会输出到标准输出中（sys.stdout）。也就是说它会将文本打印
到屏幕上面。你想写个测试来证明它，给定一个输入，相应的输出能正常显示出来。
Answer: 使用 unittest.mock 模块中的 patch() 函数，使用起来非常简单，可以为单个测
试模拟 sys.stdout 然后回滚，并且不产生大量的临时变量或在测试用例直接暴露状态
变量。
"""

"""
作为一个例子，我们在 mymodule 模块中定义如下一个函数
"""
# mymodule.py
def urlprint(protocol, host, domain):
    url = '{}://{}.{}'.format(protocol, host, domain)
    print(url)

"""
默认情况下内置的 print 函数会将输出发送到 sys.stdout 。为了测试输出真
的在那里，你可以使用一个替身对象来模拟它，然后使用断言来确认结果。使用
unittest.mock 模块的 patch() 方法可以很方便的在测试运行的上下文中替换对象，
并且当测试完成时候自动返回它们的原有状态。下面是对 mymodule 模块的测试代码：
"""
from io import StringIO
from unittest import TestCase
from unittest.mock import patch


class TestURLPrint(TestCase):
    def test_url_gets_to_stdout(self):
        protocol = 'http'
        host = 'www'
        domain = 'example.com'
        expected_url = '{}://{}.{}\n'.format(protocol, host, domain)

        with patch('sys.stdout', new=StringIO()) as fake_out:
            urlprint(protocol, host, domain)
            self.assertEqual(fake_out.getvalue(), expected_url)

"""
urlprint() 函 数 接 受 三 个 参 数， 测 试 方 法 开 始 会 先 设 置 每 一 个 参 数 的 值。
expected url 变量被设置成包含期望的输出的字符串。
unittest.mock.patch() 函数被用作一个上下文管理器，使用 StringIO 对象来代
替 sys.stdout . fake out 变量是在该进程中被创建的模拟对象。在 with 语句中使用
它可以执行各种检查。当 with 语句结束时， patch 会将所有东西恢复到测试开始前的
状态。有一点需要注意的是某些对 Python 的 C 扩展可能会忽略掉 sys.stdout 的配
置二直接写入到标准输出中。限于篇幅，本节不会涉及到这方面的讲解，它适用于纯
Python 代码。如果你真的需要在 C 扩展中捕获 I/O，你可以先打开一个临时文件，然
后将标准输出重定向到该文件中。更多关于捕获以字符串形式捕获 I/O 和 StringIO
对象请参阅 5.6 小节
"""
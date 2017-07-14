#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十四章：测试、调试和异常
Description: 试验还是很棒的，但是调试？就没那么有趣了。事实是，在 Python 测试代码之前
没有编译器来分析你的代码，因此使的测试成为开发的一个重要部分。本章的目标是
讨论一些关于测试、调试和异常处理的常见问题。但是并不是为测试驱动开发或者单
元测试模块做一个简要的介绍。因此，笔者假定读者熟悉测试概念。

Title:               忽略或期望测试失败
Issue: 你想在单元测试中忽略或标记某些测试会按照预期运行失败。
Answer: unittest 模块有装饰器可用来控制对指定测试方法的处理
"""
import unittest
import os
import platform

class Tests(unittest.TestCase):
    def test_0(self):
        self.assertTrue(True)

    @unittest.skip('skipped test')
    def test_1(self):
        self.fail('should have failed!')

    @unittest.skipIf(os.name=='posix', 'Not supported on Unix')
    def test_2(self):
        import winreg

    @unittest.skipUnless(platform.system()=='Darwin', 'Mac specific test')
    def test_3(self):
        self.asserTrue(True)

    @unittest.expectedFailure
    def test_4(self):
        self.assertEqual(2+2, 5)

if __name__== '__main__':
    unittest.main()

"""
如果你在 Mac 上运行这段代码，你会得到如下输出
bash % python3 testsample.py -v
test_0 (__main__.Tests) ... ok
test_1 (__main__.Tests) ... skipped 'skipped test'
test_2 (__main__.Tests) ... skipped 'Not supported on Unix'
test_3 (__main__.Tests) ... ok
test_4 (__main__.Tests) ... expected failure
----------------------------------------------------------------------
Ran 5 tests in 0.002s
OK (skipped=2, expected failures=1)
"""

"""
skip() 装饰器能被用来忽略某个你不想运行的测试。 skipIf() 和 skipUnless()
对于你只想在某个特定平台或 Python 版本或其他依赖成立时才运行测试的时候非常有
用。 使用 @expected 的失败装饰器来标记那些确定会失败的测试，并且对这些测试你
不想让测试框架打印更多信息。
忽略方法的装饰器还可以被用来装饰整个测试类，比如：
"""
@unittest.skipUnless(platform.system() == 'Darwin', 'Mac specific tests')
class DarwinTests(unittest.TestCase):
    pass

#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十四章：测试、调试和异常
Description: 试验还是很棒的，但是调试？就没那么有趣了。事实是，在 Python 测试代码之前
没有编译器来分析你的代码，因此使的测试成为开发的一个重要部分。本章的目标是
讨论一些关于测试、调试和异常处理的常见问题。但是并不是为测试驱动开发或者单
元测试模块做一个简要的介绍。因此，笔者假定读者熟悉测试概念。

Title:                在单元测试中测试异常情况
Issue: 你想写个测试用例来准确的判断某个异常是否被抛出。
Answer: 对于异常的测试可使用 assertRaises() 方法。
"""

"""
例如，如果你想测试某个函数抛出了 ValueError 异常，像下面这样写：
"""
import unittest

def parse_int(s):
    return int(s)

class TestConversion(unittest.TestCase):
    def test_bad_int(self):
        self.assertRaises(ValueError, parse_int, 'N/A')

"""
如果你想测试异常的具体值，需要用到另外一种方法
"""
import errno

class TestIO(unittest.TestCase):
    def test_file_not_found(self):
        try:
            f = open('/file/not/found')
        except IOError as e:
            self.assertEqual(e.errno, errno.ENOENT)
        else:
            self.fail('IOError not raised')

"""
assertRaises() 方法为测试异常存在性提供了一个简便方法。一个常见的陷阱是
手动去进行异常检测。
"""
class TestConversion(unittest.TestCase):
    def test_bad_int(self):
        try:
            r = parse_int('N/A')
        except ValueError as e:
            self.assertEqual(type(e), ValueError)

"""
这种方法的问题在于它很容易遗漏其他情况，比如没有任何异常抛出的时候。那么
你还得需要增加另外的检测过程，
"""

class TestConversion(unittest.TestCase):
    def test_bad_int(self):
        try:
            r = parse_int('N/A')
        except ValueError as e:
            self.assertEqual(type(e), ValueError)
        else:
            self.fail('ValueError not raised')

"""
assertRaises() 方法会处理所有细节，因此你应该使用它。
assertRaises() 的一个缺点是它测不了异常具体的值是多少。为了测试异常值，
可以使用 assertRaisesRegex() 方法，它可同时测试异常的存在以及通过正则式匹配
异常的字符串表示。
"""

class TestConversion(unittest.TestCase):
    def test_bad_int(self):
        self.assertRaisesRegex(ValueError, 'invalid literal .*',
                               parse_int, 'N/A')

"""
assertRaises() 和 assertRaisesRegex() 还有一个容易忽略的地方就是它们还能
被当做上下文管理器使用：
"""
class TestConversion(unittest.TestCase):
    def test_bad_int(self):
        with self.assertRaisesRegex(ValueError, 'invalid literal .*'):
            r = parse_int('N/A')

"""
但你的测试涉及到多个执行步骤的时候这种方法就很有用了。
"""



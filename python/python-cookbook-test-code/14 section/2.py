#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十四章：测试、调试和异常
Description: 试验还是很棒的，但是调试？就没那么有趣了。事实是，在 Python 测试代码之前
没有编译器来分析你的代码，因此使的测试成为开发的一个重要部分。本章的目标是
讨论一些关于测试、调试和异常处理的常见问题。但是并不是为测试驱动开发或者单
元测试模块做一个简要的介绍。因此，笔者假定读者熟悉测试概念。

Title:                在单元测试中给对象打补丁
Issue: 你写的单元测试中需要给指定的对象打补丁，用来断言它们在测试中的期望行为
（比如，断言被调用时的参数个数，访问指定的属性等）。
Answer: unittest.mock.patch() 函数可被用来解决这个问题。 patch() 还可被用作一个
装饰器、上下文管理器或单独使用，尽管并不常见。
"""
"""
例如，下面是一个将它当做装饰器使用的例子:
"""
from unittest.mock import patch

@patch('example.func')
def test1(x, mock_func):
    example.func(x)
    mock_func.assert_called_with(x)

"""
最后，你还可以手动的使用它打补丁
"""
p = patch('example.func')
mock_func = p.start()
example.func(x)
mock_func.assert_called_with(x)
p.stop()

"""
如果可能的话，你能够叠加装饰器和上下文管理器来给多个对象打补丁。
"""
@patch('example.func1')
@patch('example.func2')
@patch('example.func3')
def test1(mock1, mock2, mock3):
    pass

def test2():
    with patch('example.patch1') as mock1, \
        patch('example.patch2') as mock2, \
        patch('example.patch3') as mock3:
        pass

"""
patch() 接受一个已存在对象的全路径名，将其替换为一个新的值。原来的值会在
装饰器函数或上下文管理器完成后自动恢复回来。默认情况下，所有值会被 MagicMock
实例替代。例如
"""
x = 42
with patch('__main__.x'):
    print(x)

"""
不过，你可以通过给 patch() 提供第二个参数来将值替换成任何你想要的
"""
with patch('__main__.x', 'patched_value'):
    print(x)

"""
被用来作为替换值的 MagicMock 实例能够模拟可调用对象和实例。他们记录对象
的使用信息并允许你执行断言检查，
"""
from unittest.mock import MagicMock
m = MagicMock(return_value = 10)
m(1, 2, debug=True)
# 10
m.assert_called_with(1, 2, debug=True)
m.assert_called_with(1, 2)
# AssertionError: Expected call: mock(1, 2)
# ctual call: mock(1, 2, debug=True)

m.upper.return_value = 'HELLO'
m.upper('hello')
'HELLO'
assert m.upper.called

m.split.return_value = ['hello', 'world']
m.split('hello world')
#  ['hello', 'world']
m.split.assert_called_with('hello world')

m['blah']
# <MagicMock name='mock.__getitem__()' id='4314412048'>
m.__getitem__.called
# True
m.__getitem__.assert_called_with('blah')
"""
一般来讲，这些操作会在一个单元测试中完成。例如，假设你已经有了像下面这样
的函数：
"""
# example.py
from urllib.request import urlopen
import csv

def dowprices():
    u = urlopen('http://finance.yahoo.com/d/quotes.csv?s=@~DJI&f=sl1')
    lines = (line.decode('utf-8') for line in u)
    rows = (row for row in csv.reader(lines) if len(row) ==2)
    prices = {name:float(price) for name, price in rows}
    return prices

"""
正常来讲，这个函数会使用 urlopen() 从 Web 上面获取数据并解析它。在单元测
试中，你可以给它一个预先定义好的数据集。下面是使用补丁操作的例子:
"""
import unittest
from unittest.mock import patch
import io

sample_data = io.BytesIO(b'''\
"IBM",91.1\r
"AA",13.25\r
"MSFT",27.72\r
\r''')

class Tests(unittest.TestCase):
    @patch('example.urlopen', return_value=sample_data)
    def test_dowprices(self, mock_urlopen):
        p = example.dowprices()
        self.assertTrue(mock_urlopen.called)
        self.assertEqual(p,
                         {'IBM':91.1,
                          'AA':13.25,
                          'MSFT':27.72})

if __name__ == '__main__':
    unittest.main()

"""
本例中，位于 example 模块中的 urlopen() 函数被一个模拟对象替代，该对象会
返回一个包含测试数据的 ByteIO().
还 有 一 点， 在 打 补 丁 时 我 们 使 用 了 example.urlopen 来 代 替
urllib.request.urlopen 。当你创建补丁的时候，你必须使用它们在测试代码
中的名称。由于测试代码使用了 from urllib.request import urlopen , 那么
dowprices() 函数中使用的 urlopen() 函数实际上就位于 example 模块了。
本节实际上只是对 unittest.mock 模块的一次浅尝辄止。更多更高级的特性，请
参考 官方文档
"""
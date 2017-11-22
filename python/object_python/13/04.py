#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用命名的日志记录器 '''

"""
有4种常见的需要用logging.getLogger()为Loggers命名的情况，我们通常会选择和
应用程序结构一致的名称。
1. 模块名： 当模块包含大量的小函数或都创建大量对象的类时，我们可能会声明一个模
块级的全局Logger实例。例如，当扩展tuple时，我们不希望每个实例中都包含一个Logger
的引用。通常会定义一个全局的Logger，而且通常这个创建过程在模块的头部完成。
    如：
    import logging
    logger = logging.getLogger(__class__.__name__)

2. 对象实例： 在前面的代码中，当我们在__init__()方法中创建Logger时，展示过在对
象实例中使用Logger的例子。这里，每个对象中的Logger都是唯一的，想单纯通过全名业区
分可能会导致一些令人误解的情况，因为一个类可以有多个实例。一个更好的设计是在日志记
录器的名称中包含一个唯一的实例标识符。
    如：
    def __init__( self, player_name ):
        self.name = player_name
        self.logger = logging.getLogger(
            "{0}.{1}".format(self.__class__.__qualname__, player_name)
            )

3. 类名： 之前我们定义简单的装饰器时，演示过如何在类中使用Logger。我们可以用
__class__.__qualname__作为Logger的名字并且将Logger作为一个整体赋值给类。该类的
所有实例会共享这个Logger。

4. 函数名： 对于经常使用的小函数，经常会使用前面展示的模块级日志。对于很少使用的大型
函数，可以会在函数中创建日志。
    如：
    def main():
        log = logging.getLogger('main')

"""

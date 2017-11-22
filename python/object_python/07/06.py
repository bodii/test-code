#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 禁用类的一个功能 '''

"""
通过重新定义一个方法为只包含pass的方法，我们关闭了初始化时的洗牌功能。对于从一个子类中、
删除一个功能来说，这个过程显得有些冗长。还有一个方法可以从子类中删除功能：将方法名设置为
None。

如：
_init_shuffle=None

上面的代码需要在基类中增加一些代码用于兼容方法缺失的情形：
try:
    self._init_shuffle()
except AttributeError, TypeError:
    pass

这是从子类中删除一个功能的方法中比较显式的方式。这段代码说明了方法可能不存在或者是
被有意设为了None。然而另一种设计是将对_init_shuffle()的调用从__init__()中移到
__enter__()方法中。这种方法需要使用上下文管理器，它会让对象按我们预期的行为工作。
"""

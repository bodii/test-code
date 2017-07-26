#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第八章：类与对象
Desc: 本章主要关注点的是和类定义有关的常见编程模型。包括让对象支持常见的
Python 特性、特殊方法的使用、类封装技术、继承、内存管理以及有用的设计模
式。

Title;        子类中扩展 property
Issue:在子类中，你想要扩展定义在父类中的 property 的功能。
Answer: ，它定义了一个 property
"""

class Person:
    def __init__(self, name):
        self.name = name

    @property
    def name(self):
        return self._name
    @name.setter
    def name(self, value):
        if not isinstance(value, str):
            raise TypeError('Expected a string')
        self._name = value

    @name.deleter
    def name(self):
        raise AttributeError("Can't delete attribute")

"""
下面是一个示例类，它继承自 Person 并扩展了 name 属性的功能
"""

class SubPerson(Person):

    @property
    def name(self):
        print('Getting name')
        return super().name

    @name.setter
    def name(self, value):
        print('String name to', value)
        super(SubPerson, SubPerson).name.__set__(self, value)

    @name.deleter
    def name(self):
        print('Deleting name')
        super(SubPerson, SubPerson).name.__delete__(self)

s = SubPerson('Guido')
# String name to Guido

s.name
# Getting name

s.name = 'Larry'
# String name to Larry
#s.name = 42
#  raise TypeError('Expected a string')
"""
如果你仅仅只想扩展 property 的某一个方法，那么可以像下面这样写
"""
class SubPerson(Person):

    @Person.name.getter
    def name(self):
        print("Getting name")
        return super().name

    @Person.name.setter
    def name(self, value):
        print('Setting name to', value)
        super(SubPerson, SubPerson).name.__set__(self, value)

"""
在子类中扩展一个 property 可能会引起很多不易察觉的问题，因为一个 property
其实是 getter、setter 和 deleter 方法的集合，而不是单个方法。因此，但你扩展
一个 property 的时候，你需要先确定你是否要重新定义所有的方法还是说只修改其中
某一个。
在 第 一 个 例 子 中， 所 有 的 property 方 法 都 被 重 新 定 义。 在 每 一 个 方 法 中，
使 用 了 super() 来 调 用 父 类 的 实 现。 在 setter 函 数 中 使 用 super(SubPerson,
SubPerson).name. set (self, value) 的语句是没有错的。为了委托给之前定义
的 setter 方法，需要将控制权传递给之前定义的 name 属性的 set () 方法。不过，
获取这个方法的唯一途径是使用类变量而不是实例变量来访问它。这也是为什么我们
要使用 super(SubPerson, SubPerson) 的原因
"""

"""
如果你只想重定义其中一个方法，那只使用 @property 本身是不够的。
"""
class SubPerson(Person):

    @property
    def name(self):
        print("Getting name")
        return super().name

# s = SubPerson('Guido')
# AttributeError: can't set attribute


"""
你应该像之前说过的那样修改代码
"""
class SubPerson(Person):

    @Person.name.getter
    def name(self):
        print("Getting name")
        return super().name
"""
这么写后，property 之前已经定义过的方法会被复制过来，而 getter 函数被替换。
然后它就能按照期望的工作了
"""
s = SubPerson('Guido')
print(s.name)

s.name = 'Larry'
print(s.name)

# s.name = 42
# raise TypeError('Expected a string')

"""
在这个特别的解决方案中，我们没办法使用更加通用的方式去替换硬编码的
Person 类名。如果你不知道到底是哪个基类定义了 property，那你只能通过重新
定义所有 property 并使用 super() 来将控制权传递给前面的实现。
值的注意的是上面演示的第一种技术还可以被用来扩展一个描述器 (在 8.9 小节我
们有专门的介绍)。
"""
class String:
    def __int__(self, name):
        self.name = name

    def __get__(self, instance, cls):
        if instance is None:
            return self
        return instance.__dict__[self.name]

    def __set__(self, instance, value):
        if not isinstance(value, str):
            raise TypeError("Expected a string")
        instance.__dict__[self.name] = value

class Person:
    name = String('name')

    def __init__(self, name):
        self.name = name

class SubPerson(Person):

    @property
    def name(self):
        print('Getting name')
        return super().name

    @name.setter
    def name(self, value):
        print('Setting name to', value)
        super(SubPerson, SubPerson).name.__set__(self, value)

    @name.deleter
    def name(self):
        print('Deleting name')
        super(SubPerson, SubPerson).name.__delete__(self)

"""
最后值的注意的是，读到这里时，你应该会发现子类化 setter 和 deleter 方法其
实是很简单的。这里演示的解决方案同样适用，但是在 Python 的 issue 页面 报告的一
个 bug，或许会使得将来的 Python 版本中出现一个更加简洁的方法。
"""
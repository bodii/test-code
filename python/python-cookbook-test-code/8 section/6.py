#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第八章：类与对象
Desc: 本章主要关注点的是和类定义有关的常见编程模型。包括让对象支持常见的
Python 特性、特殊方法的使用、类封装技术、继承、内存管理以及有用的设计模
式。

Title;         创建可管理的属性
Issue:你想给某个实例 attribute 增加除访问与修改之外的其他处理逻辑，比如类型检查
或合法性验证。
Answer: 名自定义某个属性的一种简单方法是将它定义为一个 property。
"""

"""
例如，下面的代码定
义了一个 property，增加对一个属性简单的类型检查：
"""

class Person:
    def __init__(self, first_name):
        self.first_name = first_name

    @property
    def first_name(self):
        return self._first_name

    @first_name.setter
    def first_name(self, value):
        if not isinstance(value, str):
            raise TypeError('Expected a string')
        self._first_name = value

    @first_name.deleter
    def first_name(self):
        raise AttributeError("Can't delete attribute")

"""
上述代码中有三个相关联的方法，这三个方法的名字都必须一样。第一个方法是一
个 getter 函数，它使得 first name 成为一个属性。其他两个方法给 first name 属
性添加了 setter 和 deleter 函数。需要强调的是只有在 first name 属性被创建后，
后面的两个装饰器 @first name.setter 和 @first name.deleter 才能被定义。
property 的一个关键特征是它看上去跟普通的 attribute 没什么两样，但是访问它
的时候会自动触发 getter 、setter 和 deleter 方法。
"""

a = Person('Guido')
print(a.first_name)
# Guido
#a.first_name = 42
print(a.first_name)

#del a.first_name

"""
在实现一个 property 的时候，底层数据 (如果有的话) 仍然需要存储在某个地方。
因此，在 get 和 set 方法中，你会看到对 first name 属性的操作，这也是实际数据保
存的地方。另外，你可能还会问为什么 init () 方法中设置了 self.first name 而
不是 self. first name 。在这个例子中，我们创建一个 property 的目的就是在设置
attribute 的时候进行检查。因此，你可能想在初始化的时候也进行这种类型检查。通
过设置 self.first name ，自动调用 setter 方法，这个方法里面会进行参数的检查，
否则就是直接访问 self. first name 了。

还能在已存在的 get 和 set 方法基础上定义 property。
"""

class Person:
    def __init__(self, first_name):
        self.set_first_name(first_name)

    def get_first_name(self):
        return self._first_name

    def set_first_name(self, value):
        if not isinstance(value, str):
            raise TypeError('Expected a string')
        self._first_name = value

    def del_first_name(self):
        raise AttributeError("Can't delete attribute" )

    name = property(get_first_name, set_first_name, del_first_name)

"""
一个 property 属性其实就是一系列相关绑定方法的集合。如果你去查看拥有
property 的类，就会发现 property 本身的 fget、fset 和 fdel 属性就是类里面的普通方
法。
"""
# print(Person.first_name.fget)
# # <function Person.first_name at 0x0000020262A7D598>
# print(Person.first_name.fset)
# # <function Person.first_name at 0x000002E32C9AD950>
# print(Person.first_name.fdel)
# # <function Person.first_name at 0x00000201BBC3D950>

"""
不要写这种没有做任何其他额外操作的 property。首先，它会让你的代码变得很
臃肿，并且还会迷惑阅读者。其次，它还会让你的程序运行起来变慢很多。最后，这
样的设计并没有带来任何的好处。特别是当你以后想给普通 attribute 访问添加额外
的处理逻辑的时候，你可以将它变成一个 property 而无需改变原来的代码。因为访问
attribute 的代码还是保持原样。
Properties 还是一种定义动态计算 attribute 的方法。这种类型的 attributes 并不会
被实际的存储，而是在需要的时候计算出来。
"""

import math

class Circle:
    def __init__(self, radius):
        self.radius = radius

    @property
    def area(self):
        return math.pi * self.radius

    @property
    def diameter(self):
        return self.radius ** 2

    @property
    def perimeter(self):
        return 2 * math.pi * self.radius

"""
在这里，我们通过使用 properties，将所有的访问接口形式统一起来，对半径、直
径、周长和面积的访问都是通过属性访问，就跟访问简单的 attribute 是一样的。如果
不这样做的话，那么就要在代码中混合使用简单属性访问和方法调用。
"""
c = Circle(4.0)
print(c.radius)
# 4.0
print(c.area)
# 12.566370614359172
print(c.perimeter)
# 25.132741228718345

"""
尽管 properties 可以实现优雅的编程接口，但有些时候你还是会想直接使用 getter
和 setter 函数。
"""
p = Person('Guido')
print(p.get_first_name())
# Guido

p.set_first_name('Larry')
"""
这种情况的出现通常是因为 Python 代码被集成到一个大型基础平台架构或程序
中。例如，有可能是一个 Python 类准备加入到一个基于远程过程调用的大型分布式系
统中。这种情况下，直接使用 get/set 方法 (普通方法调用) 而不是 property 或许会更
容易兼容。
"""

"""
最后一点，不要像下面这样写有大量重复代码的 property 定义

class Person:
    def __init__(self, first_name, last_name):
        self.first_name = first_name
        self.last_name = last_name

    @property
    def first_name(self):
        return self._first_name

    @first_name.setter
    def first_name(self, value):
        if not isinstance(value, str):
            raise TypeError('Expected a string')
        self._first_name = value

    # Repeated property code, but for a different name (bad!)
    @property
    def last_name(self):
        return self._last_name

    @last_name.setter
    def last_name(self, value):
        if not isinstance(value, str):
            raise TypeError('Expected a string')
        self._last_name = value

重复代码会导致臃肿、易出错和丑陋的程序。好消息是，通过使用装饰器或闭包，
有很多种更好的方法来完成同样的事情。可以参考 8.9 和 9.21 小节的内容
"""
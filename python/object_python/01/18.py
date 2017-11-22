#!/usr/bin/env python3 
# -*- coding:utf-8 -*-

''' __new__()方法和元类型 '''
"""
__new__()方法的另一种用途，作为元类型的一部分，主要是为了控制如何创建一个类。这和之前的
如何用__new__()控制一个不可变对象是完全不同的。
一个元类型创建一个类。一旦类对象被创建，我们就可以用这个类对象创建不同的实例。所有类的元
类型都是type,type()函数被用来创建类对象
另外，type()函数还可以被用作显示当前对象类型。
"""

# 直接使用type()作为构造器创建一个新的但是几乎完全没有任何用处的类：
Useless = type( 'Useless', (), {} )

print(Useless())
# <__main__.Useless object at 0x7fc7ed5b6be0>

u = Useless
u.attr = 1
print(dir(u))
"""
['__class__', '__delattr__', '__dict__', '__dir__', '__doc__', '__eq__',
 '__format__', '__ge__', '__getattribute__', '__gt__', '__hash__', '__init__',
  '__le__', '__lt__', '__module__', '__ne__', '__new__', '__reduce__', 
  '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__',
   '__subclasshook__', '__weakref__', 'attr']
"""

# 下面这个和type(类名， (), {}) 一样
class Useless:
    pass


''' 元类型示例1 --- 有序的属性 '''

"""
实现的3个具体步骤：

1. 创建一个元类型。元类型的__prepare__()和__new__()方法会改变目标类创建的方法，会将原本
的dict类替换为OrderedDict类。

2. 创建一个基于此元类型的抽象基类。这个抽象类简化了其他类继承这个元类型的过程。

3. 创建一个继承于这个抽象基类的子类，这样它就可以获得元类型的默认行为。
"""
import collections

class Ordered_Attributes( type ):
    
    @classmethod
    def __prepare__( metacls, name, bases, **kwds ):
        return collections.OrderedDict()

    def __new__( cls, name, bases, namespace, **kwds ):
        result = super().__new__( cls, name, bases, namespace )
        result._order = tuple(n for n in namespace if not n.startswith('__'))
        return result
    """
    这个类用自定义的__prepare__()和__new__()方法扩展了内置的默认元类型type.

    __prepare__()方法会在类创建之前执行，它的工作是创建初始的命令空间对象，类定义最后被添加
    到这个对象中。这个方法可以用来处理任何在类的主体开始执行前需要的准备工作。

    __new__()静态方法在类的主本被加入命名空间后开始执行。它的参数是要创建的类对象、类
    名、基类的元组和创建好的命名空间匹配对象。这个例子很经典：它将__new__()的真正工作
    委托给了基类；一个元类型基类是内置的type;然后我们使用type.__new__()创建一个稍后
    可以修改的默认类。
    这个例子中的__new__()方法向类中增加了一个_order属性，用于存储原始的属性创建顺序。
    当我们定义新的抽象基类时，我们可以用这个元类型而非type。
    """

class Order_preserved( metaclass = Ordered_Attributes):
        pass

class Something( Order_preserved ):
    
    this = 'text'

    def z( self ):
        return False

    b = 'order is preserved'
    a = 'more text'

print( Something._order )
# ('this', 'z', 'b', 'a')
print( dir(Something) )


''' 元类型示例2---自引用 '''


class Unit:
    """ Full name for the unit. """
    factor = 1.0
    standard = None # Reference to the apperopriate StandardUnit
    name = '' # Abbreviation of the unit's name

    @classmethod
    def value( class_, value ):
        if value is None: return None
        return value / class_.factor

    @classmethod
    def convert( class_, value ):
        if value is None: return None
        return value * class_.factor

    """
    这个类的目的是Unit.value()可以将一个值从给定的单位转换为标准单位，而Unit.convert()
    方法可以将一个值从标准单位转换为给定的单位
    """

# @standard
class INCH:
    pass

INCH.standard = INCH

# 定义一个可以向类定义中插入一个循环引用的元类型
class UnitMeta( type ):
    def __new__( cls, name, bases, dict ):
        new_class = super().__new__( cls, name, bases, dict )
        new_class.standard = new_class
        return new_class
    """ 这段代码强制地将变量stanard作为类定义的一部分 """

class Standard_Unit( Unit, metaclass=UnitMeta ):
    pass

class INCH( Standard_Unit ):
    """ Inches """
    name = "in"

class FOOT( Unit ):
    """ Feet """
    name = 'ft'
    standard = INCH
    factor = 1/12

class CENTIMETER( Unit ):
    """ Centimeters """
    name = 'cm'
    standard = INCH
    factor = 2.54


class METER( Unit ):
    """ Meters """
    name = 'm'
    standard = INCH
    factor = .0254

""" 我们将INCH定为标准单位，其他单位需要转换英寸或者从英寸转换而来。
    在每一个单位类中，我们都提供了一些文档信息：全名写在docstring中并且用name属性记录缩写。
    从Unit继承而来的convert()和value()方法会自动应用转换因子。
"""

x_std = INCH.value( 159.625 )
print( FOOT.convert( x_std ) )
# 13.302083333333332
print( METER.convert( x_std ) )
# 4.054475
print( METER.factor )
# 0.0254

print( INCH.standard.__name__ )
# INCH
print( FOOT.standard.__name__ )
# INCH
#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 创建修饰符 '''

"""
修饰符类必须实现以下3个方法的一个或多个：

1. Descriptor.__get__( self, instance, owner ) -> object:在这个方法中，instance
参数来自访问对象的self变量。owner变量是拥有者类的对象。如果这个修饰符在类中被调用，
instance参数默认值将为None。此方法负责返回修饰符的值。

2. Descriptor.__set__( self, instance, value ): 在这个方法中，instance参数是被
访问对象的self变量，而value参数为即将赋的新值。

3. Descriptor.__delete__( self, instance ): 在这个方法中，instance参数是被访问对象
的self变量，并在这个方法中实现属性值的删除。

有时，修饰符类也需要在__init__()函数中初始化修饰符内部的一些状态。


基于方法的定义，如下是两种不同的修饰符类型：
1. 非数据修饰符： 这类修饰符需要定义__set__()或__delete__()或两者皆有，但不能定义
__get__()。非数据修饰符对象经常用于构建一些复杂表达式的逻辑。它可能是一个可调用对象，
可能包含自己的属性或方法。一个不可变的非数据修饰符必须实现__set__()函数，而逻辑只是
单纯地抛出AtrributeError异常。这类修饰符的设计相对简单一些，因为接口更灵活。

2. 数据修饰符： 这类修饰符至少要定义__get__()函数。通常，可通过定义__get__()和__set
__()函数来创建一个可变对象。这类修饰符不能定义自己内部的属性或方法，因为它通常是不可见的。
对修饰符属性的访问，也相应地转换为对修饰符中的__get__()、__set__()或__delete__()
方法的调用。这样对设计是一个挑战，因此不会作为首要选择。


第一种情况。使用__get__()和__set__()函数创建数据修饰符，以及不使用__get__()方法的情
况下创建非数据修饰符。

第2种情况（数据包含在拥有者实例中），正是@property装饰器的用途。比起传统的特性，修饰符
带来的好处是，它把计算逻辑从拥有者类搬到了修饰符类中。而完全采用这样的设计思路来设计类是
片面的，有些场景不能获得最大的收益。如果计算逻辑相当复杂，使用策略模式则更好。

第3种情况，@staticmethod 和 @classmethod 装饰器的实现就很好。

"""


''' 使用非数据修饰符 '''

# 如下，一个简单的非数据修饰符类的实现，但未包含__get__()函数
class UnitVlaue_1:
    ''' Measure and unit combined. '''
    def __init__( self, unit ):
        self.value = None
        self.unit = unit
        self.default_format = '5.2f'

    def __set__( self, instance, value ):
        self.value = value

    def __str__( self ):
        return '{value:{spec}} {unit}'.format(
        spec = self.default_format,
        **self.__dict__
    )

    def __format__( self, spec= '5.2f'):
        # print( 'formatting', spec)
        if spec == '': spec = self.default_format
        return '{value:{spec}} {unit}'.format( spec=spec, **self.__dict__ )


# 以下这个类用来完成速率一时间一距离的计算
class RTD_1:
    rate = UnitVlaue_1( 'kt' )
    time = UnitVlaue_1( 'hr' )
    distance = UnitVlaue_1( 'nm' )

    def __init__( self, rate=None, time=None, distance=None ):
        if rate is None:
            self.time = time
            self.distance = distance
            self.rate = distance / time

        if time is None:
            self.rate = rate
            self.distance = distance
            self.time = distance / rate

        if distance is None:
            self.rate = rate
            self.time = time
            self.distance = rate * time

    def __str__( self ):
        return 'rate: {0.rate} time: {0.time} distance:{0.distance}'.format( self )


m1 = RTD_1( rate=5.8, distance=12 )
print( str(m1) )
# rate:  5.80 kt time:  2.07 hr distance:12.00 nm

print( 'Time:', m1.time.value, m1.time.unit )
# Time: 2.0689655172413794 hr

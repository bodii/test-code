#!/usr/bin/env python3 
# -*- coding:utf-8 -*-

''' __new__()方法和不可变对象 '''

"""
__new__方法的一个用途是初始化不可变对象。__new__()方法中允许创建未初始化的对象。
这允许我们在__init__()方法被调用之前先设置对象的属性。

由于不可变类的__init__()方法很难重载，因此__new__方法提供了一种扩展这种类的方法。
"""

class Float_Fail( float ):
    def __init__( self, value, unit ):
        super().__init__( value )
        self.unit = unit

# 我们试图（不合理地）初始化一个不可变对象

# s2 = Float_Fail( 6.5, 'knots' )
# TypeError: float() takes at most 1 argument (2 given)

"""
可以看到，对于内置的float类，我们不能简单地重载__init__方法。对于其他的内置不可变
类型，也有类似的问题。我们不能在不可变对象self上设置新的属性值，因为这是不可变性的
定义。我们只能在对象创建的过程中设置属性值，对象创建之后__new__()方法就会被调用。
__new__()方法天生就是一个静态方法。即使没有使用@staticmethod修饰符，它也是静态
的。它没有使用self变量，因为它的工作是创建最终会被赋值给self变量的对象。
这种情况下，我们会使用的方法签名是__new__( cls, *args, **kw ).cls变量是准备创建
的类的实例。

__new__()方法的默认实现：
return super().__new__( cls )将调用基类的__new__()方法创建对象。这个工作最终委托
给了object.__new__()，这个方法创建了一个简单的空对象。除了cls以外，其他的参数和关键字
最终都会传递给__init__()方法，这是Python定义的标准行为。
"""
"""
当创建一个内置的不可变类型的子类时，不能重载__init__()方法。取而代之的是，我们必须通过
重载__new__()方法在对象创建的过程中扩展基类的行为。
"""

class Float_units( float ):
    def __new__( cls, value, unit ):
        obj = super().__new__( cls, value )
        obj.unit = unit 
        return obj

speed = Float_units( 6.5, 'knots' )
print( speed ) # 6.5
print( speed * 10 ) # 65.0
print( speed.unit ) # knots

"""
注意，像speed * 10这种表达式不会创建一个Float_units对象。这个类的定义继承了float中所
有的运算：float的所有算术特殊方法也都只会创建float对象。
"""
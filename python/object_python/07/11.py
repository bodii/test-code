#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 创建类装饰器 '''

import logging

class UglyClass1:
    def __init__( self ):
        self.logger = logging.getLogger(self.__class__.__qualname__)
        self.logger.info( 'New thing' )
    
    def method( self, *args ):
        self.logger.info( 'method %r', args )

"""
这个类的缺点是它创建了一个对象，它不属于类操作的一部分却是类的一个独立方面的logger实例。
我们不希望用这种额外的方面污染我们的类。但是这并不是问题的全部。虽然logging.getLogger()
非常高效，但是并非完全没有代价。我们希望每次创建UglyClass1实例时，可以避免这种额外的消耗。
"""

# 我们将国。logger从类中每个独立的对象中分离出来，把它变成一个类级的实例变量
class UglyClass2:
    logger = logging.getLogger('UglyClass2')

    def __init__( self ):
        self.logger.info( 'New thing' )
    
    def method( self, *args ):
        self.logger.info( 'method %r', args )

"""
这个类的优点是它只调用了logging.getLogger()一次。但是，它有严重的DRY问题。在定义中
，没有办法自动设置类名。由于类还没有被创建，因此必须重复类名。
"""

# 可以用下面的这个小装饰器来解决DRY问题。
def logged( class_ ):
    class_.logger = logging.getLogger( class_.__qualname__ )
    return class_
"""
这个装饰器修改了类的定义，它添加了logger引用作为类级属性。
"""

@logged
class SomeClass:
    def __init__( self ):
        self.logger.info( 'New thing' )

    def method( self, *args ):
        self.logger.info( 'method %r', args )

"""
现在，我们类中包含了一个每个方法都可以使用的logger属性。logger的值不是对象的一个功能，
这样就保证了这个方面和类的其他方面是分离的。这个属性还有一个额外的好处就是，它在导入模块
时创建了logger实例，减少了一些日志记录的开销。
"""

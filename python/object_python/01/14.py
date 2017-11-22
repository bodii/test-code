#!/usr/bin/env python3 
# -*- coding:utf-8 -*-

''' 循环引用和垃圾回收 '''

class Parent:
    def __init__( self, *children ):
        self.children = list(children)
        for child in self.children:
            child.parent = self

    def __del__( self ):
        print( 'Removing {__class__.__name__} {id:d}'.format(__class__= self.__class__,
        id= id(self)))


class Child:
    def __del__( self ):
        print( 'Removing {__class__.__name__} {id:d}'.format(__class__= self.__class__,
        id= id(self)))


p = Parent( Child(), Child() )
print( id(p) )

del p
print( '_' * 30 )
p = Parent()
print( id(p) )
del p

""" 由于互相之间有引存在，因此我们不能从内存中删除Parent实例和它包含的Child实例的集合。
如果我们导入垃圾回收器的接口--gc，我们就可以回收和显示这些不能被删除的对象 """

import gc
print(gc.collect()) # 7

print(gc.garbage)

""" 为了让引用计数器归零，我们需要删除所有Parent对象中的children列表，或者删除所有
Child实例中对Parent的引用
注意，即使把清理资源的代码放在__del__()方法中，我们也没办法解决循环引用的问题。因为
__del__()方法是在循环引用被解除并且引用计数器已经归零之后被调用的。当有循环引用时，我们
不能只是简单地依赖于Python中计算引用数量的机制来清理内存中的无用对象。我们必须显式地解除
循环引用或者使用可以保证垃圾回收的weakref引用。
 """
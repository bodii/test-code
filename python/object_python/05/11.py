#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 一些设计原则 '''
"""
将集合分成两个类：TreeNode 和 Tree。
TreeNode类会包含所有的元素和more、less和parent引用。
例如，为了能够使用__contains__()或者discard()搜索一个特定的元素会用一个简单的递归将这个搜索工作
委托给节点自身来完成，算法的描述如下：
1. 如果目前元素和当前元素相等，那么返回self。
2. 如果目标元素比self.item小，那么递归地使用less.find(target item)继续搜索目标元素。
3. 如果目标元素比self.item大，那么递归地使用more.find(target.item)继续搜索目标元素。
用类似的方式将更多维护树结构的工作委托给TreeNode类完成。。
第2个类会使用外观模式(Facade)定义Tree。外观模式也被成为包装模式（Wrapper），主要目的是为一个特定的接口增加属性。
我们会为MutableSet抽象基类提供必要的外部接口。
"""

''' 定义Tree类 '''

from collections import abc

class Tree( abc.MutableSet ):
    def __init__( self, iterable=None ):
        self.root = TreeNode(None)
        self.size = 0
        if iterable:
            for item in iterable:
                self.add( item )
    
    def add( self, item ):
        self.root.add( item )
        self.size += 1

    def discard( self, item ):
        try:
            self.root.more.remove( item )
            self.size -= 1
        except KeyError:
            pass

    def __contains__( self, item ):
        try:
            self.root.more.find( item )
            return True
        except KeyError:
            return False

    def __iter__( self ):
        for item in iter( self.root.more ):
            yield item
    
    def __len__( self ):
        return self.size

"""
初始化方法和Counter对象类似，这个类会接受一个可迭代对象作为参数，并加载对象中的所有元素。
add()和discard()方法会持续更新节点总数。这样当我们需要知道当前节点总数时，就不用遍历整棵树。这些方法也将
工作委托给了位于根部的TreeNode对象。
__contains__()特殊方法会执行递归查找。当发生KeyError异常时，会返回False。
__iter__()特殊方法是一个生成器函数。它也将实际的工作委托给了TreeNode类的递归迭代器。
我们定义了discard(),当试图忽略不存在的键时，可变集合需要这个方法可以忽略异常。抽象基类中提供了一个remove()
的默认实现，当一个键不存在时会抛出一个异常。两个方法函数都必须定义，我们基于remove（）定义了discard(),
但是会忽略不存在时remove()抛出的异常。
在一些情况下，基于discard()定义remove()可能会更容易，如果发现问题就抛出一个异常。
"""


''' 定义TreeNode类 '''

import weakref

class TreeNode:
    def __init__( self, item, less=None, more=None, parent=None ):
        self.item = item
        self.less = less
        self.more = more
        if parent is not None:
            self.parent = parent

    @property
    def parent( self ):
        return self.parent_ref()
    
    @parent.setter 
    def parent( self, value ):
        self.parent_ref = weakref.ref(value)

    def __repr__( self ):
        return 'TreeNode({item!r},{less!r},{more!r}'.format( **self.__dict__ )

    def find( self, item ):
        if self.item is None: # Root
            if self.more: return self.more.find(item)
        elif self.item == item: return self
        elif self.item > item and self.less : return self.less.find(item)
        elif self.item < item and self.more : return self.more.find(item)
        raise KeyError

    def __iter__( self ):
        if self.less:
            for item in iter(self.less):
                yield item
        yield self.item
        
        if self.more:
            for item in iter(self.more):
                yield item

    """
    定义了初始化两种不同节点的基本方法。唯一必要的参数为元素本身;两个子树和父节点引用作为可选参数。
    这些属性用来确保parent属性以强引用的方式出现，虽然实际上它是weakref属性。
    在一个TreeNode父节点对象和它的孩子节点对象之间存在互相引用，这种循环引用让删除TreeNode对象
    变得很困难。可以用一个weakref打破这种循环引用。
    接下来是find()方法，它会递归地在树中遍历子树，搜索目标元素。
    __iter__()方法会按顺序遍历当前节点和它的所有子树。和往常一样，这是一个生成器函数，它从每一个
    子树集合的迭代器中生成要返回的值。尽管可以创建一个基于Tree类的独立迭代器，但是这样做没有任何
    好处，因为生成器函数可以完成我们需要的所有功能。
    """
    # 下面是这个类的第2部分，实现了向树中添加新节点。
    def add( self, item ):
        if self.item is None:
            if self.more:
                self.more.add( item )
            else:
                self.more = TreeNode( item, parent=self )
        elif self.item >= item:
            if self.less:
                self.less.add( itme )
            else:
                self.less = TreeNode( item, parent=self )
        elif self.item < item:
            if self.more:
                self.more.add( item )
            else:
                self.more = TreeNode( item, parent=self )
    # 这个方法递归地搜索要插入节点的正确位置。这个方法的结构和find()方法类似


    # 从树中删除节点
    def remove( self, item ):
        if self.item is None or item > self.item:
            if self.more:
                self.more.remove(item)
            else:
                raise KeyError
        elif item < self.item:
            if self.less:
                self.less.remove(item)
            else:
                raise KeyError
        else: # self.item == item
            if self.less and self.more:
                successor = self.more._least()
                self.item = successor.item
                successor.remove(successor.item)
            elif self.less:
                self._replace(self.less)
            elif self.more:
                self._replace(self.more)
            else:
                self._replace(None)

    def _least( self ):
        if self.less is None:
            return self.less._least()

    def _replace( self, new=None ):
        if self.parent:
            if self == self.parent.less:
                self.parent.less = new
            else:
                self.parent.more = new
        if new is not None:
            new.parent = self.parent

"""
当删除一个没有孩子的节点时，我们可以简单地删除它然后将与父节点的引用改为None;
当删除一个有一个孩子的节点时，我们可以用这个孩子代替当前节点在父节点中的引用
当有两个孩子时，我们需要调整树的结构。我们首先找到后继节点（在more子树中的最小
节点）。可以用这个后继节点的值替换准备删除的节点。然后，可以删除之前那个重复的
后继节点。
两个私有方法:_least方法会在一棵给定的树上查询出最小节点。_replace()方法检查
父节点，以确定是否需要更新less或者more属性。
"""


''' 演示二叉树集合 '''

s1 = Tree( ['Item 1', 'Another', 'Middle'] )
print( list(s1) )
# ['Another', 'Item 1', 'Middle']
print( len(s1) )
# 3

s2 = Tree( ['Anothter', 'More', 'Yet More'] )
union = s1 | s2
print( list(union) )
# ['Another', 'Anothter', 'Item 1', 'Middle', 'More', 'Yet More']

print( len(union) )
# 6
union.remove( 'Another' )
print( list(union) )
# ['Anothter', 'Item 1', 'Middle', 'More', 'Yet More']
#!/usr/bin/env python3
# -*- coding:utf-8 -*-
# PW @ 2019-11-09 19:35:06

"""
链表:
链表(linked list)是一种常见的基础数据结构，是一种线性表，但是不像
顺序表一样连续存储数据，而是在每一个节点(数据存储单元)里存放下一个节点的位置信息。

链表结构可以充分利用计算机内存空间，实现灵活的内存动态管理。
"""

"""
单向链表
单向链表也叫单链表，是链表中最简单的一种形式，它的每个节点包含两个域，一个信息域
（元素域)和一个链接域。这个链接指向链表中的下一个节点，而最后一个节点的链接域则指
向一个空值。

a. 表元素域elem用来存放具体的数据。
b. 链接域next用来存放下一个节点的位置(ptyhon中的标识)
c. 变量p指向链表的头节点(首节点)的位置，从p出发能找到表中的任意节点。
"""
class Node:
    """
    单向链表类
    """
    def __init__(self, elem):
        self.elem = elem # 当前元素
        self.next = None # 下一个节点位置
        

# 构造单向链表类
class SingleLinkList:
    def __init__(self, node=None):
        if node != None:
            self.__head = Node(node)
        else:
            self.__head = None

    # 添加头元素
    def add(self, item):
        headNode = Node(item)
        headNode.next = self.__head
        self.__head = headNode

    # 添加尾部元素
    def append(self, item):
        node = Node(item)
        currentNode = self.__head
        if  self.is_empty() == False:
            while currentNode.next != None:
                currentNode = currentNode.next
            
            currentNode.next = node
        else:
            currentNode = node

    # 添加指定位置的元素
    def insert(self, index, item) -> bool:
        prevNode =  self.__head
        currentNode = prevNode
        node = Node(item)
        length = self.length()

        if length+1 < index:
            return False

        if length == 0 and index == 1:
            currentNode = node
            return True
    
        for i in range(1, length + 1):
            if i == index:
                node.next = currentNode
                prevNode.next = node
                return True
            else:
                prevNode = currentNode
                currentNode = currentNode.next

    # 删除元素
    def remove(self, item) -> bool:
        currentNode = self.__head
        prevNode = currentNode
        status = False

        while currentNode != None:
            if currentNode.elem == item:
                prevNode = currentNode.next
                status = True
            else:
                currentNode = currentNode.next
            prevNode.next = currentNode.next

        return status



    # 返回所有元素的总长度
    def length(self) -> int:
        count = 0
        currentNode = self.__head
        while currentNode != None:
            count += 1
            currentNode = currentNode.next

        return count

    # 判断是否为空
    def is_empty(self) -> bool:
        return self.__head == None

    # 遍历当前链表
    def travel(self) -> str:
        currentNode = self.__head
        while currentNode != None:
            print(currentNode.elem, end='\t')
            currentNode = currentNode.next
        print('')
         


    # 查询一个元素
    def search(self, item) -> bool:
        currentNode = self.__head
        while currentNode != None:
            if currentNode.elem == item:
                return True

        return False


if __name__ == '__main__':
    # 初始化元素值为20的单向链表
    singleLinkList = SingleLinkList(20)
    singleLinkList.add(40)
    singleLinkList.add(17)
    singleLinkList.add(36)
    singleLinkList.add(51)
    singleLinkList.add(17)
    singleLinkList.append(66)
    singleLinkList.remove(17)
    # singleLinkList.insert(7, 14)
    singleLinkList.travel()

    # 初始化一个空的单向链表
#    singleLinkList = SingleLinkList()

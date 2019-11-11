#!/usr/bin/env python3
# -*- coding:utf-8 -*-
# PW @ 2019-11-11 20:54:54


"""
双向链表
一种更复杂的链表是“双向链表”或“双面链表"。
每个节点有两个链接：一个指向前一个节点，当此节点为第一个节点时，指向空值；
而另一个指向下一个节点，当此节点为最后一个节点时，指向空值。
"""

class Node:
    """
    双向链表节点
    """
    def __init__(self, elem):
        self.elem = elem
        self.prev = None
        self.next = None


class DoubleLinketList:
    """
    双向链表
    """
    def __init__(self, item = None):
        """
        初始化双向链表
        """
        self.__head = None

        if item is not None:
            self.__head = Node(item)


    def is_empty(self) -> bool:
        """
        检测链表是否为空
        """
        return self.__head is None


    def length(self) -> int:
        """
        返回链表的长度
        """
        count = 0
        currentNode = self.__head
        while currentNode is not None:
            count += 1
            currentNode = currentNode.next

        return count


    def append(self, item) -> bool:
        """
        在链表尾部添加元素
        """
        node = Node(item)
        currentNode = self.__head
        length = self.length()
        for i in range(length - 1):
            currentNode = currentNode.next

        node.prev = currentNode
        currentNode.next = node

        return True


    def add(self, item) -> bool:
        """
        在链表头部添加元素
        """
        node = Node(item)
        currentNode = self.__head
        currentNode.prev = node
        node.next = currentNode
        self.__head = node

        return True


    def insert(self, index, item) -> bool:
        """
        在指定节点插入元素
        """
        node = Node(item)
        currentNode = self.__head
        prevNode = currentNode.prev
        length = self.length()
        if index > length + 1:
            return False
        
        # 如果在第一个节点插入元素
        if index == 1:
            node.next = currentNode
            currentNode.prve = node
            self.__head = node
            return True

        # 如果在中间节点插入元素
        for i in range(1, length + 1):
            if index == i:
                currentNode.prev = node
                prevNode.next = node
                node.next = currentNode
                node.prev = prevNode
                return True
            else:
                currentNode = currentNode.next
                prevNode = currentNode

        # 如果在最后一个节点插入元素
        if index == length + 1:
            node.prev = currentNode
            currentNode.next = node
            return True

        return False


    def remove(self, item) -> bool:
        """
        删除指节点的元素
        """
        currentNode = self.__head
        prevNode = currentNode.prev
        length = self.length()
    
        status = False
        for i in range(length):
            if currentNode.elem == item:
                if prevNode is not None:
                    prevNode.next = currentNode.next
                else:
                    self.__head = currentNode.next
                status = True
            else:
                prevNode = currentNode
            currentNode = currentNode.next
                
        status = False
        return status
        

    def search(self, item) -> bool:
        """
        从双向链表中查询某个元素
        """
        currentNode = self.__head
        length = self.length()
        for i in range(length):
            if currentNode.elem == item:
                return True
            else:
                currentNode = currentNode.next

        return False


    def travel(self) -> str:
        """
        打印双向链表的所有节点元素
        """
        currentNode = self.__head
        length = self.length()
        for i in range(length):
            print(currentNode.elem, end="\t")
            currentNode = currentNode.next

        print('')


if __name__ == '__main__':
    doubleLinkList = DoubleLinketList(3)
    doubleLinkList.add(16)
    doubleLinkList.append(9)
    doubleLinkList.insert(1, 10)
    doubleLinkList.travel()
    print("search 4 is ", doubleLinkList.search(4))
    print("search 9 is ", doubleLinkList.search(9))
    doubleLinkList.remove(10)
    doubleLinkList.travel()


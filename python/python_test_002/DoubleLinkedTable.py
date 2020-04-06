#!/usr/bin/env python3
# -*- coding:utf-8 -*-

from typing import *

# 节点类
class Node:
    def __init__(self, k = None, v = None):
        """
        节点类
        :param: key 键
        :param: value 值
        """
        self.prev = None
        self.next = None
        self.key = k
        self.value = v

    def __str__(self):
        return '{%s: %s}' % (self.key, self.value)

    def __repr__(self):
        return '{%s: %s}' % (self.key, self.value)

# 双向链表类
class DoubleLinkedTable:
    def __init__(self):
        self.size = 0
        self.__head = Node()
        

    def is_empty(self):
        """
        判断是否为空链表
        :return: bool
        """
        return self.size == 0

    def length(self):
        """
        返回链表元素的长度
        :return: int 长度
        """
        return self.size

    def contains(self, key):
        """
        检测链表中是否包含元素的key
        :param: any key 元素索引
        :return: bool 返回是否包含
        """
        root = self.__head
        while (root.next is not None):
            if root.next.key == key:
                return True
            root = root.next
        
        return False


    def add(self, node: Node):
        """
        向链表的头部添加节点
        :param: Node node 要添加的节点元素
        """
        root = self.__head
        node.next = root.next
        node.prev = root
        root.next = node
        
        self.size += 1

    def append(self, node: Node):
        """
        向链表的尾部添加节点
        :param: Node node 要添加的节点元素
        """
        root = self.__head
        while (root.next is not None):
            root = root.next

        node.prev = root
        root.next = node
        self.size += 1

    def insert(self, key, node: Node):
        """
        向链表指定的key索引后添加节点
        :param: key 指定索引
        :param: node 要添加的节点
        :return: bool 是否插入成功
        """
        root = self.__head
        result = False

        while(root.next is not None):
            if (root.next.key == key):
                node.prev = root.next
                node.next = root.next.next
                root.next.next = node
                self.size += 1
                result = True
            root = root.next

        return result


    def remove(self, key):
        """
        删除链表中指定key对应的节点
        :param: key 指定key
        :return: 返回被删除节点key对应的value
        """
        root = self.__head
        result = None

        while (root.next is not None):
            if (root.next.key == key):
                newNode = root.next
                result = newNode.value
                root.next = newNode.next
                self.size -= 1
                
                break
            root = root.next

        return result

    def pop(self):
        """
        从链表的尾部弹出一个节点元素
        :return: 返回被删除的节点
        """
        root = self.__head.next
        while(root.next is not None):
            root = root.next

        result = root
        root.prev.next = root.next
        self.size -= 1
        
        return result

    def shift(self):
        """
        从链表的头部弹出一个节点元素
        :return: 返回被删除的节点
        """
        root = self.__head
        result = root.next
        root.next = root.next.next
        self.size -= 1

        return result

    def get(self, key):
        """
        获取链表中指定key对应的值
        :param: key 指定key
        :return: 返回查到的值或None
        """
        root = self.__head
        result = None

        while (root.next is not None):
            if (root.next.key == key):
                result = root.next.value
                break
            root = root.next

        return result

    def set(self, key, newVal):
        """
        设置链表中指定key的新值
        :param: key 指定的key 
        :param: newVal 新值
        :return: 是否设置成功
        """
        result = False
        root = self.__head
        while(root.next is not None):
            if (root.next.key == key):
                root.next.value = newVal
                result = True
            root = root.next
        
        return result

    def __str__(self):
        """
        将当前链表格式化成字符串
        :return: str 格式化后的字符串
        """
        linkedStr = 'head->'
        root = self.__head

        while (root.next is not None):
            linkedStr += (" %s ->" % root.next)
            root = root.next
        linkedStr += 'tail'
        return linkedStr


if __name__ == '__main__':
    linked = DoubleLinkedTable()
    for i in range(5):
        linked.append(Node(i, i+1))

    print(linked)

    print("del key is 1 value is:", linked.remove(1))
    print(linked)
    linked.add(Node(11,12))
    print(linked)
    linked.append(Node(7, 11))
    print(linked)
    
    print("链表的长度是:", linked.length())
    print("链表中是否包含9：" , linked.contains(9))
    print("链表中是否包含7：" , linked.contains(7))
    print("链表中是否包含11：" , linked.contains(11))
    print("链表中获取索引为7的值：" , linked.get(7))

    print("将链表中索引为3的值设置为20：" )
    linked.set(3, 20)
    print(linked)

    print("从链表中弹出末尾节点：", linked.pop())
    print(linked)

    print("从链表中移出头部节点：", linked.shift())
    print(linked)

    insert = Node(33, 19)
    print("从链表中索引为3之后插入：%s 是否成功:" %  insert,  linked.insert(3, insert))
    print(linked)




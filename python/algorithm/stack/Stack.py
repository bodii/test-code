#!/usr/bin/env python3
# -*- coding:utf-8 -*-
# PW @ 2019-11-12 21:54:32

# 栈
class Stack:
    def __init__(self):
        self.__container = []

    def push(self, item):
        """
        向尾部压入元素
        """
        self.__container.append(item)

    def pop(self):
        """
        向头部弹出一个元素
        """
        return self.__container.pop()

    def peek(self):
        """
        向头部获取一个元素
        """
        return self.__container[0]

    def is_empty(self) -> bool:
        """
        判断是否为空
        """
        return self.size() <= 0 

    def size(self) -> int:
        """
        获取栈的长度
        """
        return len(self.__container) 

    def getAll(self) -> list:
        """
        获取全部剩余内容
        """
        return self.__container



if __name__ == '__main__':
    stack = Stack()
    stack.push(4)
    stack.push(7)
    stack.push('a')
    stack.pop()
    print(stack.peek())
    print("stack length: ", stack.size())
    print("is empty: ", stack.is_empty())
    print("all: ", stack.getAll())


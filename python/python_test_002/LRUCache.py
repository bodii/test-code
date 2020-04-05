#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from DoubleLinkedTable import DoubleLinkedTable, Node
from typing import *

class LRUCache:
    """
    Least recently used cache
    最近最小使用缓存算法
    """
    def __init__(self, capacity: int = 10 ):
        self.capacity = capacity # 设置缓存的容量，没有时默认为10
        self.map = {}
        self.list = DoubleLinkedTable()

    def get(self, key):
        """
        获取缓存中key键对应的值
        :param: key 键
        :return: 返回键对应的值，如果不存在则返回-1
        """
        if key in self.map:
            node = self.map[key]
            self.list.remove(key)
            self.list.add(node)
            return node.value
        else:
            return -1

    def put(self, key, value):
        """
        向缓存中添加key和value对应关系
        :param: key 键
        :param: value 值
        """
        if key in self.map:
            node = self.map.get(key)
            self.list.remove(key)
            node.value = value
        else:
            node = Node(key, value)
            if self.list.size >= self.capacity:
                oldNode = self.list.pop()
                self.map.pop(oldNode.key)
            self.map[key] = node
        self.list.add(node)


    def __str__(self):
        return str(self.list)

    def __repr__(self):
        return str(self.list)

if __name__ == '__main__':
    cache = LRUCache(2)
    cache.put(1, 1)
    print(cache)
    cache.put(2, 2)
    print(cache)
    print(cache.get(1))
    print(cache)
    cache.put(3, 3)
    print(cache)


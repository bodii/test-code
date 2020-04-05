#!/usr/bin/env python3
# -*- coding:utf-8 -*-

from typing import *

from DoubleLinkedTable import DoubleLinkedTable, Node

class FIFOCache:
    """
    先进先出缓存类
    尾进头出
    """
    def __init__(self, capacity: int):
        self.size = 0  # 缓存当前的大小
        self.map = {} # 映射关系
        self.list = DoubleLinkedTable() # 双
        self.capacity = capacity or 10 # 缓存当前的最大容量,如果没有默认值为10

    def get(self, key):
        """
        获取指定缓存key对应的值
        :param: key 缓存时的键
        :return: 返回键对应的值
        """
        if key not in self.map:
            return -1
        else:
            node = self.map.get(key)
            return node.value

    def put(self, key, value) :
        """
        向缓存的末尾添加元素{key : value}
        :param: key 键
        :param: value 值
        """
        if self.capacity == 0:
            return 

        if key in self.map:
            node = self.map.get(key)
            node.value = value
            self.list.set(key, node)
        else:
            if self.size == self.capacity:
                node = self.list.shift()
                del self.map[node.key]
                self.size -= 1
            node = Node(key, value)
            self.list.append(node)
            self.map[key] = node
            self.size += 1

    def __str__(self) -> str:
        return str(self.list)
    
    def __repr__(self) -> str:
        return str(self.list)


if __name__ == '__main__':
    cache = FIFOCache(2)
    cache.put(1, 1)
    print(cache)
    cache.put(2, 2)
    print(cache)
    print(cache.get(1))
    cache.put(3, 3)
    print(cache)
    print(cache.get(2))
    print(cache)
    cache.put(4, 4)
    print(cache)

            
        

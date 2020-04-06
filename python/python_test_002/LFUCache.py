#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from DoubleLinkedTable import DoubleLinkedTable, Node
from typing import *

class LFUNode(Node):
    """
    least frequently used node
    最不经常使用节点类
    """
    def __init_(self, key, value):
        """
        最不经常使用节点类
        :param: key 键
        :param: value 值
        """
        self.freq = 0
        super().__init__(key, value)


class LFUCache:
    """
    least frequently used caching algorithm
    最不经常使用缓存算法
    """
    def __init__(self, capacity : int = 10):
        self.capacity = capacity
        self.map = {}
        self.freq_map = {}

    def __update_freq(self, node: LFUNode = None):
        """
        更新频率
        :param: node 频率节点
        """
        freq = node.freq
        self.freq_map[freq].remove(node.key)
        if self.freq_map[freq].size == 0:
            del self.freq_map[freq]
        
        freq += 1
        node.freq = freq
        if freq not in self.freq_map:
            self.freq_map[freq] = DoubleLinkedTable()
        self.freq_map[freq].append(node)

    def get(self, key):
        """
        获取指定key对应的值
        :param: key 键
        :param: 返回键对应的值，如果不存在则返回-1
        """
        if key not in self.map:
            return -1

        node = self.map.get(key)
        self.__update_freq(node)
        return node.value

    def put(self, key, value):
        """
        向缓存中添加键和值
        :param: key 键
        :param: value 值
        """
        if self.capacity <= 0:
            return 

        if key in self.map:
            node = self.map.get(key)
            node.value = value
            self.__update_freq(node)
        else:
            if self.capacity <= len(self.map):
                min_freq = min(self.freq_map)
                node = self.freq_map[min_freq].pop()
                del self.map[node.key]
                
            node = LFUNode(key, value)
            node.freq = 1
            self.map[key] = node
            if node.freq not in self.freq_map:
                self.freq_map[node.freq] = DoubleLinkedTable()
            self.freq_map[node.freq].append(node)

    def __str__(self):
        result = '*********************\n'
        for f, node in self.freq_map.items():
            result += f"freq:  %d" % f
            result += "\n\t\t{node}\n".format(node=str(node))
        result +=  '*********************\n'
        return result
        

if __name__ == '__main__':
    cache = LFUCache(2)
    cache.put(1, 1)
    print(cache)
    cache.put(2, 2)
    print(cache)
    cache.put(3, 3)
    print(cache.get(2))
    print(cache.get(2))
    print(cache)
    print(cache.get(3))
    print(cache.get(3))
    print(cache)
    cache.put(4, 4)
    print(cache)

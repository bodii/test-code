#!/usr/bin/env python3
# -*- coding:utf-8 -*-
# PW @ 2019-11-14 21:12:53

# 定义二叉树的节点
class Node:
    def __init__(self, elem):
        self.elem = elem
        self.lchild = None
        self.rchild = None


# 树类
class Tree:
    def __init__(self):
        self.root = None

    def add(self, elem):
        """
        添加树节点元素
        """
        node = Node(elem)

        if self.root is None:
            self.root = node
        else:
            queue = []

            while queue:
                curNode = queue.pop(0)
                if curNode.lchild is None:
                    curNode.lchild = node
                    return
                else:
                    queue.append(curNode.lchild)

                if curNode.rchild is None:
                    curNode.rchild = node
                    return
                else:
                    queue.append(curNode.rchild)



if __name__ == '__main__':
    tree = Tree()
    tree.add(3)
    tree.add(5)


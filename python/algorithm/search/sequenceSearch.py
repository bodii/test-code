#!/usr/bin/env python3
# -*- coding:utf-8 -*-
# PW @ 2019-11-08 22:08:10

def sequenceSearch(alist : list, v : int):
    for i in range(len(alist)):
        if alist[i] == v:
            return i

    return -1


if __name__ == '__main__':

    alist = [ 4, 3, 5, 1, 233, 66, 99 ]
    index = sequenceSearch(alist, 99)
    if index != -1:
        print("找到了， 索引值是: " , index)
    else:
        print("没有打到")

#!/usr/bin/env python3
# -*- coding:utf-8 -*-
# PW @ 2019-11-04 22:09:00


def bubbleSort(alist: list):
    """
    冒泡排序法
    """
    llen = len(alist)
    for i in range(0, llen -1):
        for j in range(i+1, llen):
            if alist[i] > alist[j]:
                alist[i], alist[j] = alist[j], alist[i]


alist = [54, 26, 93, 17, 77, 31, 44, 55, 20]

bubbleSort(alist)
print(alist)

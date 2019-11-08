#!/usr/bin/env python3
# -*- coding:utf-8 -*-
# PW @ 2019-11-05 21:03:26


def insertSort(alist : list):
    """
    插入排序
    """
    n = len(alist)
    for j in range(1, n):
        i = j 
        while i > 0:
            if alist[i] < alist[i-1]:
                alist[i], alist[i-1] = alist[i-1], alist[i]
            else: 
                break

            i -= 1


if __name__ == "__main__":
    alist = [ 54, 226, 93, 17, 31, 44, 55, 20 ]

    print("原列表：")
    print(alist)
    print("排序后的列表: ")
    insertSort(alist)
    print(alist)

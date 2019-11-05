#!/usr/bin/env python3
# -*- coding:utf-8 -*-
# PW @ 2019-11-04 21:27:50


def select_sort(alist: list):
    """
    选择排序
    """
    llen = len(alist)
    for i in range(llen - 1):
        min_index = i
        for j in range(i+1, llen):
            if alist[j] < alist[min_index]:
                min_index = j

        if min_index != i:
            alist[i], alist[min_index] = alist[min_index], alist[i]

        

if __name__ == '__main__':
    alist = [7, 5, 3, 6, 44, 22, 99, 11]
    print('原列表: ')
    print(alist)

    print('排序后的列表: ')
    select_sort(alist)

    print(alist)

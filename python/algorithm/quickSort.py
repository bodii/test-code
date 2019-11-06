#!/usr/bin/env python3
# -*- coding:utf-8 -*-
# PW @ 2019-11-05 22:14:25

def quickSort(alist : list, start : int, end : int):
    """
    快速排序
    """
    # 如果递归完成，则退出
    if start >= end: 
        return 

    datum = alist[start]  # 基准数
    low = start
    high = len(alist) - 1
    while low < high:
        # 从右向左
        while low < high and alist[high] >= datum: 
            high -= 1
        alist[low] = alist[high]

        # 从左向右
        while low < high and alist[low] > datum:
            low += 1
        alist[high] = alist[low]

    alist[low] = datum

    # 比基准数小的即左边
    quickSort(alist, start, low - 1)
    # 比基准数大的即右边
    quickSort(alist, low + 1, end)



if __name__ == '__main__':

    alist = [ 54, 26, 93, 17, 77, 31, 44, 55, 20 ]
    
    print("原列表: ")
    print(alist)

    print("排序后的列表: ")
    quickSort(alist, 0, len(alist) - 1)
    print(alist)

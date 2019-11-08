#!/usr/bin/env python3
# -*- coding:utf-8 -*-
# PW @ 2019-11-08 23:00:48

def binarySearch(alist : list, v : int):
    """
    二分查找
    循环实现
    """
    first = 0
    last = len(alist) - 1
    while first <= last:
        mid = (first + last) // 2
        if alist[mid] == v:
            return True
        elif alist[mid] > v:
            last = mid - 1
        else:
            first = mid + 1

    return False


def binarySearch2(alist: list, v: int):
    """
    二分查找
    递归实现
    """
    n = len(alist)
    if n == 0:
        return False
    mid = n // 2
    if alist[mid] == v:
        return True
    elif alist[mid] < v:
        return binarySearch2(alist[mid+1:], v)
    else:
        return binarySearch2(alist[0:mid-1], v)
        


if __name__ == '__main__':
    alist = [ 1, 2, 4, 3, 7, 5, 6, 9, 8 ]
   
    if binarySearch2(alist, 9):
        print("找到了")
    else:
        print("没有找到")

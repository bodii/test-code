#!/usr/bin/env python3
# -*- coding:utf-8 -*-
# PW @ 2019-11-08 11:34:02

def mergeSort(alist : list): 
    """
    归并排序法
    该算法采用经典的分治(divide-and-conquer)策略
    分治法将问题分(divide)成一些小的问题然后递归求解，而治(conquer)的阶段则将分的阶段得到的
    各答案“修补”在一起，即分而治之。
    分阶段可以理解为就是递归拆分子序列的过程，递归深度为log2n。
    """
    if len(alist) <= 1:
        return alist

    num = int(len(alist) / 2)
    left = mergeSort(alist[:num])
    right = mergeSort(alist[num:])

    r_pointer, l_pointer = 0, 0
    result = []
    while l_pointer < len(left) and r_pointer < len(right):
        if left[l_pointer] <= right[r_pointer]:
            result.append(left[l_pointer])
            l_pointer += 1
        else:
            result.append(right[r_pointer])
            r_pointer += 1

    result.extend(list(left[l_pointer:]))
    result.extend(list(right[r_pointer:]))

    return result


if __name__ == "__main__":
    alist = [ 50, 10, 20, 30, 70, 40, 80, 60 ]

    print("原列表: ")
    print(alist)

    nlist = mergeSort(alist)
    print("排序后的列表: ")
    print(nlist)




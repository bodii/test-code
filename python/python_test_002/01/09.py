#!/usr/bin/env python3
# -*- coding:utf-8 -*-


def do_plus(oneParam, twoParam):
    try:
        if type(oneParam) != int or type(twoParam) != int:
            raise TypeError('传入的两个参数都必须为int类型！')
    except TypeError as e:
        print(e)
    return oneParam + twoParam

if __name__ == '__main__':
    print(do_plus(1, 2))

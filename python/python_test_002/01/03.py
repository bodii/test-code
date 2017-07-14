#!/usr/bin/env python3
# -*- coding:utf8 -*-

def in_fridge():

    try:
        count = fridge[wanted_food]
    except Exception as e:
        count = 0
    return count

if __name__ == '__main__':
   result = in_fridge()
   print(result)

#!/usr/bin/env python3
# -*- coding:utf8 -*-

def in_fridge ():
    """ This is a function to see if the fridge has a food.
    fridge has to be a dictionary defined outside of the function.
    the food to be searched for is in the string wanted_food"""
    try:
        count = fridge[wanted_food]
    except Exception as e:
        count = 0
    return count


if __name__ == '__main__':
    fridge = in_fridge()
    print(fridge)
    print(in_fridge.__doc__)


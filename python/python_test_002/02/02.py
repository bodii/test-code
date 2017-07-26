#!/usr/bin/env python3
# -*- coding:utf-8 -*-

class Fridge:
    """This class implements a frideg where ingredients can be
    added and rmoeved individually, or in groups."""

    def __init__(self, items={}):
        """optionally pass in an initial dictionary of items"""

        if type(items) != type({}):
            raise TypeError("Fridege requires a dictionary but was given %s" % type(ites))
        self.items = items
        return

    def __add_multi(self, food_name, quantity):
        
        if (not food_name in self.items):
            self.items[food_name] = 0

        self.items[food_name] = self.items[food_name] + quantity


    def add_one(self, food_name):
        """
        add_one(food_name) - adds a single food_name to the fridge
        returns True
        Raises a TypeError if food_name is not a string.
        """

        if type(food_name) != type(""):
            raise TypeError("add_one requires a string, given a  %s" % type(food_name))
        else:
            self.__add_multi(food_name, 1)

        return True

    def add_many(self, food_dict):
        
        if type(food_dict) != type({}):
            raise TypeError("add_many requires a dictionary, got a %s" % food_dict)

        for item in food_dict.keys():
            print(item, food_dict[item])
            self.__add_multi(item, food_dict[item])
            return



if __name__ == "__main__":
    
    f = Fridge({"eggs":6, "milk":4, "cheese":3})
    print(f.items)

    print(f.add_one('grape'))

    print(f.items)

    f.add_many({"mushroom":5, "cheese":0, 'kk':55})

    print(f.items)


#!/usr/bin/env python3
# -*- coding:utf-8 -*-


from fridge import Fridge

class Omelet:
    
    def __init__(self, kind="cheese"):

        self.set_kind(kind)
    
    def __ingreditents__(self):

        return self.needed_ingredients

    def get_kind(self):

        return self.kind

    def set_kind(self, kind):
        possible_ingredients = self.__known_kinds(kind)

        if possible_ingredients == False:
            return False

        else:
            self.kind = kind
            self.needed_ingredients = possible_ingredients


    def set_new_kind(self, name, ingredients):
        self.kind = name
        self.needed_ingredients = ingredients
        return

    def __known_kinds(self, kind):
        
        if kind == "cheese":
            return {"eggs":2, "milk":1, "cheese":1}
        elif kind == "mushroom":
            return {"eggs":2, "milk":1, "cheese":1, "mushroom":2}
        elif kind == "onion":
            return {"eggs":2, "milk":1, "cheese":1, "onion":1}
        else:
            return False
    
    def get_ingredients(self, fridge):
        self.from_fridge = fridge.get_ingredients(self)

    def mix(self):
        for ingredient in self.from_fridge.keys():
            print("Mixing %d %s for the %s omelet" %
                    (self.from_fridge[ingredient], ingredient, self.kind))
        self.mixed = True

    def make(self):
        if self.mixed:
            print("Cooking the %s omelet!" % self.kind)
            self.cooked = True
    


        

if __name__ == '__main__':
    o = Omelet('cheese')
    f = Fridge({"cheese":5, "milk":4, "eggs": 12})
    o.get_ingredients(f)
    print(o.mix())

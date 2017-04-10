#!/usr/bin/env python3

def make_omelet(cmelet_type):
    """This will make an omelet. you can either pass in a dictionary
    that contains all of the ingredients for your omelet, or provide 
    a string to select a type of omelet this function already knows
    about"""

    if type(omelet_type) == type({}):
        print("omelet_type is a dictionary with ingredients")
        return make_food(omelet_type, "omelet")
    elif type(omelet_type) == type(""):
        omelet_ingredients = get_omelet_ingredients(omelet_type)
        return make_food(omelet_ingredients, omelet_type)
    else:
        print("I don't think I can make this kind of omelet: %s" % 
                omelet_type)

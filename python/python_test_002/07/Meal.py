#!/usr/bin/env python3
# -*- coding:utf-8 -*-


class Meal:
    '''Holds the food and drink used in a meal.
    In true object-oriented tradition, this clas
    includes setter methods for the food and drink.

    Call printIt to pretty-print the values.
    '''

    def __init__(self, food='omelet', drink='coffee'):
        '''Initialize to default values.'''
        self.name = 'generic meal'
        self.food = food
        self.drink = drink

    def printIt(self, prefix=''):
        '''Print the data nicely.'''
        print(prefix, 'A fine', self.name, 'with', self.food, 'and', self.drink)

    # Setter for the food.
    def setFood(self, food='omelet'):
        self.food = food

    # Setter for the drink.
    def setDrink(self, drink='coffee'):
        self.drink = drink

    # Setter for the name.
    def setName(self, name=''):
        self.name = name



#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import Meal

class Dinner(Meal):
    '''Holds the food and drink for dinner.'''

    def __init__(self):
        '''Initialize with steak and merlot.'''
        Meal.__init__(self, 'steak', 'merlot')
        self.setName('dinner')

    def printIt(self, prefix=''):
        '''Print even more nicely.'''
        print(prefix, 'A gourmet', self.name, 'with', self.food, 'and', 
                self.drink)

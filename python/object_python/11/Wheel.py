#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用可回调类创建WGI应用程序 '''

import random


__all__ = ['Wheel', 'Zero', 'DoubleZero']

__name__ = "Wheel"

class Wheel:
    def  __init__( self ):
        self.ng = random.Random()
        self.bins = [
            {
                str(n) : (35, 1),
                self.redblack(n) : (1, 1),
                self.hilo(n) : (1, 1),
                self.evenodd(n) : (1, 1),
            } for n in range(1, 37)
        ]

    @staticmethod
    def redblack( n ):
        return 'Red' if n in (1, 3, 5) else 'Black'
    
    @staticmethod
    def hilo( n ):
        return 'Hi' if n >= 19 else 'Lo'

    @staticmethod
    def evenodd( n ):
        return 'Even' if n % 2 == 0 else 'Odd'

    def spin( self ):
        return self.ng.choice( self.bins ) 


class Zero:
    def __init__( self ):
        super().__init__()
        self.bins += [{ '0' : {35, 1} }]

class DoubleZero:
    def __inist__( self ):
        super().__init__()
        self.bins += [{ '00': {35, 1} }]

    
        

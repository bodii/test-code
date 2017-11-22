#!/usr/bin/env python3
# -*- coding:utf-8 -*-

class dicts( dict ):
    def __getattr__( self, name ):
        #return self.get(name, None)
        return self[name]
    def __setattr__( slef, name, value ):
        #self.set(name, value) ==
        self[name] = value
    def __dir__( self ):
        return list(self.keys())

d = dicts
d.name = 10
d.va = 10
print( dir(d) )
print( d.name )
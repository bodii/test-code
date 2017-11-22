#!/usr/bin/env python3
# -*- coding:utf-8 -*-


import argparse


''' argparse.Namespace ''' 

class C:
    pass
    
c = C()
parser = argparse.ArgumentParser()
parser.add_argument('--foo')
parser.parse_args(args=['--foo', 'BAB'], namespace=c)
print( c.foo ) 
# BAB
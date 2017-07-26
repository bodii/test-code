#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import os

def make_another_file():
    if os.path.isfile('test.txt'):
        print('You are trying to create a file that already exists!')
    else:
        f = open('test.txt', 'w')
        f.write('This is how you create a new text file.\n')
        f.close()


if __name__ == '__main__':
    make_another_file()

#!/usr/bin/env python3
# -*- coding:utf-8 -*-

def make_text_file():
    a = open('test.txt', 'w')
    a.write('This is how you create a new text file.\n')
    a.close()

if __name__ == '__main__':
    make_text_file()

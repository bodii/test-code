#!/usr/bin/env python3
# -*- coding:utf-8 -*-

def read_file():
    a = open('test.txt', 'r')
    print(a.readline())
    a.close()

if __name__ == '__main__':
    read_file()

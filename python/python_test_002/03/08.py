#!/usr/bin/env python3
# -*- coding:utf-8 -*-

def read_file():
    a = open('test.txt', 'r')
    lineText = a.readlines()
    for line in lineText:
        print(line)
        print(len(line))

    a.close()


if __name__ == '__main__':
    read_file()

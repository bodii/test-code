#!/usr/bin/env python3
# -*- coding:utf-8 -*-

def add_some_text():
    a = open('test.txt', 'a')
    a.write('Here is some additional text!\n')
    a.close()


if __name__ == '__main__':
    add_some_text()

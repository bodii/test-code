#!/usr/bin/env python3
# -*- coding:utf-8 -*-

def even_more_text():
    a = open('test.txt', 'a')
    a.write(
    '''here is
    more
    text
    ''')
    a.close()


if __name__ == '__main__':
    even_more_text()

#!/usr/bin/env python3
# -*- coding:utf-8 -*-

def print_line_lengths():
    a = open('test.txt', 'r')
    text = a.readlines()
    for line in text:
        print(len(line))


if __name__ == '__main__':
    print_line_lengths()

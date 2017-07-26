#!/usr/bin/env python3
# -*- coding:utf-8 -*-


def use_logging(func):

    def wrapper(*args, **kwargs):
        print("%s is running" % func.__name__)
        return func(*args)
    return wrapper

@use_logging
def foo():
    print("i am foo")


if __name__ == '__main__':
    foo()

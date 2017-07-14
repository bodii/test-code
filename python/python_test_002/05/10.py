#!/usr/bin/env python3
# -*- coding:utf-8 -*-

# 基于字典的字符串格式化

person = {"name" : "James", "camera": "nikon", "handedness": "lefty",
        "baseball_team": "angels", "instrument": "guitar"}

print("%(name)s, %(camera)s, %(baseball_team)s" % person)

person["height"] = 16
person["weight"] = 80
print("%(name)s, %(camera)s, %(baseball_team)s, %(height)2.2f, %(weight)2.2f"
        % person)



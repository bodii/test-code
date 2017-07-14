#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import string
person = {"name": "James", "camera": "nikon", "handedness": "lefty",
        "baseball_team": "angels", "instrument": "guitar"}
person["height"] = 1.6
person['weight'] = 80

t = string.Template("$name is $height m high and $weight kilos")
print(t.substitute(person))

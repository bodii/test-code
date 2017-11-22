#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 将JSON写入文件 '''

import json


with open('temp.json', 'w',encoding='UTF-8') as target:
    json.dump( trave13, target, separators=(',',':'), 
    default=blog_j2_encode 
    )

with open( 'some_source.json','r', encoding='UTF-8') as source:
    objects = json.load(source, ojbect_hook=blog_decode)


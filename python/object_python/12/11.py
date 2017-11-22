#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用XML文件  -- PLIST以及其他格式保存配置 '''
"""
一种常见的方式是使用XML将配置数据表示在.plist文件中
"""

"""
与Python类型兼容的XML标签
Python类型      Plist标签
str             <string>
float           <real>
int             <integer>
datetime        <date>
boolean         <true/> 或 <false/>
bytes           <data>
list            <array>
dict            <dict>
"""

import plistlib
import os.path

current_path = os.path.dirname(__file__) + os.path.sep
file = current_path + '.plist'


with open(file, 'rb') as f:
    p1 = plistlib.readPlist(f)
    print(p1)
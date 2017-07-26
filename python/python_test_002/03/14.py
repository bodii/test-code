#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import os.path

print(os.path.splitext("image.jpg"))

(name, extension) = os.path.splitext('image.jpg')
print(name)
print(extension)

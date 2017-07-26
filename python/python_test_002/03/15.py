#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import os.path

parts = os.path.splitext("image.jpg")
extension = parts[1]
print(extension)

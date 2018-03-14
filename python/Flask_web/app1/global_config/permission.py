#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 权限配置 '''

class Permission:
    FOLLOW = 0x01
    COMMENT = 0x02
    WRITE_ARTICLES = 0x04
    MODERATE_COMMENTS = 0x08
    ADMINISTER = 0x80
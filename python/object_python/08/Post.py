#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import datetime

class Post:
    """ 定义用于持久化的类 """ 
    def __init__( self, date, title, rst_text, tags ):
        self.date = date
        self.title = title
        self.rst_text = rst_text
        self.tags = tags

    def as_dict( self ):
        return dict(
            date = str(self.date),
            title = self.title,
            underline = '-'*len(self.title),
            rst_text = self.rst_text,
            tag_text = ' '.join(self.tags),
        )
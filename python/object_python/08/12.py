#!/usr/bin/env python3
# -*- coding:utf-8 -*-


''' 使用XML转储和加载 '''

from Post import *
from Blog import *
import datetime

travel = Blog( 'Travel' )
travel.append(
    Post(
        date = datetime.datetime(2013, 11, 14, 17, 25),
        title = 'Hard Aground',
        rst_text = """Some embarrassing revalation.
        Including """,
        tags = ('#RedRanger', '#whitby42', '#ICW'),
    )
)
travel.append(
    Post(
        date = datetime.datetime(2013, 11, 18, 15, 30),
        title = 'Anchor Follies',
        rst_text = """Some witty epigram. Including < & > 
        characters. """,
        tags = ('#RedRanger', '#Whitby42', '#Mistakes'),    
    )
)

class Blog_X( Blog ):
    def xml( self ):
        children = '\n'.join( c.xml() for c in self.entries )
        return """\
    <blog>
        <title>{0.title}</title>
        <entries>
        {1}
        </entries
    </blog>""".format( self, children )

class Post_X( Post ):
    def xml( self ):
        tags = ''.join( "<tag>{0}</tag>".format(t) for t in self.tags )
        return """\
    <entry>
        <title>{0.title}</title>
        <date>{0.date}</date>
        <tags>{1}</tags>
        <text>{0.rst_text}</text>
    </entry>""".format( self, tags )


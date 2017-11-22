#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用特性文件存储配置 '''

import re
import os.path
import io
from collections import ChainMap

class PropertyParser:
    
    def read_string( self, data ):
        return self._parse(data)

    def read_file( self, file ):
        data = file.read()
        return self.read_string(data)

    def read( self, filename ):
        with open(filename) as file:
            return self.read_file(file)

    def load( self, file_or_name ):
        if isinstance( file_or_name, io.TextIOBase):
            self.loads(file_or_name.read())
        else:
            with open(filename) as file:
                self.loads(file.read())

    def loads( self, string ):
        return self._parse(data)
    
    # 每个配制元素的正则匹配规则
    key_element_pat = re.compile(r'(.*?)\s*(?<!\\)[:=\s]\s*(.*)')
    def _parse( self, data ):
        logical_lines = (line.strip() 
        for line in re.sub(r'\\\n\s*', '', data).splitlines())
        non_empty = (line for line in logical_lines if len(line) !=0)
        non_comment = (line for line in non_empty if not( line.startswith('#') or line.startswith('!')))
        for line in non_comment:
            ke_match = self.key_element_pat.match(line)
            if ke_match:
                key, element = ke_match.group(1), ke_match.group(2)
            else:
                key, element = line, ''
            key = self._escape(key)
            element = self._escape(element)
            yield key, element
            
    def _escape( self, data ):
        #d1 = re.sub(r'\\([:#!=\s])', lambda x: x.group(1), data)
        #d2 = re.sub(r'\\u([0-9A-Fa-f]+)', lambda x: chr(int(x.group(1), 16)), d1)
        d2 = re.sub(
            r'\\([:#!=\s])|\\u(0-9A-Fa-f)+', 
            lambda x: x.group(1) if x.group(1) else chr(int(x.group(2), 16)), 
            data
            )
        return d2

'''
config = ChainMap(
    *[dict(pp.read(file)) for file in reversed(candidate_list)]
)
'''
current_path = os.path.dirname(__file__) + os.path.sep
filename = current_path + 'properties';
pp = PropertyParser()
config = pp.read(filename)

for co in config:
    print( co[0], co[1] )

"""
player.betting Flat
player.play SomeStrategy
player.ronds 100
player.stake 50
table.dealer Hit17
table.decks 6
table.limit 50
table.payout (3, 2)
table.split NoResplitAces
simulator.outputfile p2_c13_simulation8.dat
simulator.samples 100
"""
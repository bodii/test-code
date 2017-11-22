#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import xml.etree.ElementTree as xml

class Configuration:
    def read_file( self, file ):
        self.config = XML.parse(file)

    def read( self, filename ):
        self.config = XML.parse(filename)

    def read_string( self, text ):
        self.config = XML.fromstring(text)
        
    # stake = int(config.get('player.stake', 50))
    def get( self, qual_name, default ):
        section, _, item = qual_name.partitin('.')
        query = './{0}/{1}'.format( section, item )
        node = self.config.find(query)
        if node is None: return default
        return node.text

    # stake = config['player']['stake']
    def __getitem__( self, section ):
        query = './{0}'.format(section)
        parent = self.config.find(query)
        return dict((item.tag, item.text) for item in parent )

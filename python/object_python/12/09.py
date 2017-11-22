#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import yaml
import os.path

if __name__ == '__main__':
    current_path = os.path.dirname(__file__) + os.path.sep
    file = open(current_path +'config.yaml')
    config = yaml.load(file)
    print( config['table'], config['player'] )
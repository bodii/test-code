#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import logging
import yaml
import os.path
import logging.config

BASE_PATH = os.path.dirname(__file__)
filename = 'tailHandler.yaml'
filename = os.path.join(BASE_PATH, filename)

with open(filename) as file:
    config = yaml.load(file)
    logging.config.dictConfig(config)
    log = logging.getLogger('test.demo8')
    print( 'Last 5 before error' )
    for i in range(20):
        log.debug( 'Message {:d}'.format(i) )
    log.error( 'Error causes dump of last 5' )
    print( 'Last 5 before shutdown' )
    for i in range(20, 40):
        log.debug( 'Message {:d}'.format(i) )
    logging.shutdown()
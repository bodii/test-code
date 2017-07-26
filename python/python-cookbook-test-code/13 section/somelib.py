#!/usr/bin/env python
# -*- coding:utf-8 -*-

import logging

log = logging.getLogger(__name__)
log.addHandler(logging.NullHandler)

def func():
    log.critical('A Critical Error!')
    log.debug('A debug message')

#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用warnings模块 '''

import warnings
import urllib
import json
import socket
import logging

try:
    import simulation_model_l as model 
except ImportError as e:
    warnings.warn(e)
if 'model' not in globals():
    try:
        import simulation_model_2 as model 
    except ImportError as e:
        warnings.warn(e)
if 'model' not in globals():
    raise ImportError('Missing simulation_model_1 and simulation')


log = logging.getLogger('test')
try:
    with urllib.request.urlopen('http://host/resource/', timeout=30) as resource:
        content = json.load(resource)
except socket.timeout as e:
    log.warn('Missing information from http://host/resource')
    content = []


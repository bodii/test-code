#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import shelve
import os.path as oph

c_path = oph.dirname(__file__)
sep = oph.sep
filename = c_path + sep + 'test.db'

# db = shelve.open(filename, 'n')
# db['db_name:test:table_name'] = 'test1'
# db['db_name:test:data'] = {'a': 1}
# db.sync()
# db.close()

# db = shelve.open(filename, 'r')
# print(db.get('db_name:test:data'))
# db.close() 


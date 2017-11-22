#!/usr/bin/env python3
# -*- coding:utf-8 -*-


import sqlite3
import hashlib

def md5sum(t):
    return hashlib.md5(t).hexdigest()

con = sqlite3.connect(':memory:')
con.create_function('md5', 1, md5sum)
cur = con.cursor()
cur.execute('select md5(?)', (b'foo',))
print(cur.fetchone()[0])
# acbd18db4cc2f85cedef654fccc4a4d8  
cur.close()
con.close()      
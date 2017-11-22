#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 字节或ASCII编码的Unicode用户名和密码 '''\

from Authentication import Authentication

class Users( dict ):
    def __init__( self, *args, **kw ):
        # Can never match -- keys are the same.
        self[''] = Authentication(b'__dummy', b"Doesn't Matter")
    
    def add( self, authentication ):
        if authentication.username == '':
            raise KeyError('Invalid Authentication')
        self[authentication.username] = authentication

    def match( self, username, password ):
        if username in self and username != '':
            return self[username].match(password)
        else:
            return self[''].match(b"Something which doesn't match")\


# 创建用户

users = Users()
users.add( Authentication(b'Aladdin', b'open sesame') )

al = Authentication(b'Aladdin', b'open sesame')
print( al )
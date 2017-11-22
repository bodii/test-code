#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' WSGI验证程序 '''

import base64
from collections.abc import Callable

class WSGI( Callable ):
    def __call__( self, environ, start_response ):
        raise NotImplementedError

class Authenticate( WSGI ):
    def __init__( self, users, target_app ):
        self.users = users
        self.target_app = target_app

    def __call__( self, environ, start_response ):
        if 'HTTP_AUTHORIZATION' in environ:
            scheme, credentials = environ['HTTP_AUTHORIZATION'].split()
            if scheme == 'Basic':
                username, password = base64.b64decode(credentials).split(b':')
            if self.users.match(username, password):
                environ['Authenticate.username'] = username
                return self.target_app(envirn, start_response)
            status = '401 UNAUTHORIZED'
            header = [
                        ('Content-type', 'text/plain; charset=utf-8'),
                        ('WWW-Authenticate', 'Basic realm="roulette@localhost"')
                    ]
            start_response(status, header)
            return ['Not authorized'.encode('utf-8')]


#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 使用可回调类创建WGI应用程序 '''

from Wheel import Wheel, Zero, DoubleZero

from collections.abc import Callable
import json

class Wheel2( Wheel, Callable ):
    def __call__( self, environ, start_response ):
        winner = self.spin()
        status = '200 OK'
        headers = [( 'Content-type', 'application/json; charset=utf-8' )]
        start_response(status, headers)
        return [ json.dumps(winner).encode('utf-8') ]

class American2( Zero, DoubleZero, Wheel2 ):
    pass

class European2( Zero, Wheel2 ):
    pass



import wsgiref.utill
import sys

class Wheel3( Callable ):
    def __init__( self ):
        self.am = American2()
        self.eu = European2()

    def __call__( self, environ, start_response ):
        request = wsgiref.util.shift_path_info(environ)
        print( 'Wheel3', request, file=sys.stderr)

        if request.lower().startswith('eu'):
            response = self.en(environ, start_response)
        else:
            response = self.am(environ, start_response)
        return response

    

    

    
        

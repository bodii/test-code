#!/usr/bin/env python3
# -*- coding:utf-8 -*-


''' 多层REST服务 '''

from collections import defaultdict
from collections.abc import Callable
from Wheel import Wheel, Zero, DoubleZero
import sys
import wsgiref
import json
import http.client
import time

class Table:
    def __init__( self, stake=100 ):
        self.bets = defaultdict(int)
        self.stake = stake
    
    def place_bet( self, name, amount ):
        self.bets[name] += amount
    
    def clear_bets( self, name ):
        self.bets = defaultdict(int)

    def resolve( self, spin ):
        """spin is a dict with bet:(x:y)."""
        details = []
        while self.bets:
            bet, amount = self.bets.popitem()
            if bet in spin:
                x, y = spin[bet]
                self.stake += amount * x / y
                details.append( (bet, amount, 'win') )
            else:
                self.stake -= amount
                details.append( (bet, amount, 'lose') )
        return details



# 定义WSGI的两个帮助类

class WSGI( Callable ):
    def __call__( self, environ, start_response ):
        raise NotImplementedError

class RESTException( Exception ):
    pass

class American( Zero, DoubleZero, Wheel ):
    pass


class Roulette( WSGI ):
    def __init__( self, wheel ):
        self.table = Table(100)
        self.rounds = 0
        self.wheel = wheel

    def __call__( self, environ, start_response ):
        # print( environ, file=sys.stderr)
        # wsgiref.util.shift_path_info函数会解析路径并且根据'/'取出第1个单词。
        app = wsgiref.util.shift_path_info(environ)
        try:
            if app.lower() == 'player':
                return self.player_app( environ, start_response )
            elif app.lower() == 'bet':
                return self.bet_app( environ, start_response )
            elif app.lower() == 'wheel':
                return self.wheel_app( environ, start_response )
            else:
                raise RESTException(
                    '404 NOT_FOUND', 
                    'Unknown app in {SCRIPT_NAME}/{PATH_INFO}'.format_map(environ)
                )
        except RESTException as e:
            status = e.args[0]
            headers = [('Content-type', 'text/plain;charset=utf-8')]
            start_response( status, headers, sys.exc_info() )
            return [ repr(e.args).encode('utf-8') ]
    
    # 定义一个全局异常处理器，它会将任何RESTEXception实例转化成一个合适的RESTful
    # 响应。没有处理的异常会被转化为wsgiref提供的通用状态码500。
    def player_app( self, environ, start_response ):
        if environ['REQUEST_METHOD'] == 'GET':
            details = dict( stake=self.table.stake, rounds=self.rounds)
            status = '200 OK'
            headers = [('Content-type', 'application/json;charset=utf-8')]
            start_response(status, headers)
            return [ json.dumps(details).encode('utf-8') ]
        else:
            raise RESTException('405 METHOD_NOT_ALLOWED',
            "Method '{REQUEST_METHOD}' not allowed".format_map(environ)
            )
    
    def bet_app( self, environ, start_response ):
        if environ['REQUEST_METHOD'] == 'GET':
            details = dict( self.table.bets )
        elif environ['REQUEST_METHOD'] == 'POST':
            size = int(environ['CONTENT_LENGTH'])
            raw = environ['wsgi.input'].read(size).decode('utf-8')
            try:
                data = json.loads( raw )
                if isinstance(data, dict): data = [data]
                for detail in data:
                    self.table.place_bet( detail['bet'], int(detail['amount']))
            except Exception as e:
                raise RESTException('403 FORBIDDEN',
                "Bet {raw!r}".format(raw=raw)
                )
                details = dict( self.table.bets )
        else:
            raise RESTException('405 METHOD_NOT_ALLOWED',
            "Method '{REQUEST_METHOD}' not allowed".format_map(environ)
            )
        
        status = '200 OK'
        headers = [('Content-type', 'application/json; charset=utf-8')]
        start_response(status, headers)
        return [ json.dumps(details).encode('utf-8') ]

    def wheel_app( self, environ, start_response ):
        if environ['REQUEST_METHOD'] == 'POST':
            size = environ['CONTENT_LENGTH']
            if size != '':
                raw = environ['wsgi.input'].read(int(size))
                raise RESTException("403 FORBIDDEN", 
                    "Data '{raw!r}' not allowed".format(raw=raw)
                )
            spin = self.wheel.spin()
            payout = self.table.resolve( spin )
            self.rounds += 1
            details = dict( spin=spin, payout=payout,
                stake = self.table.stake, rounds = self.rounds
            )
            status = '200 OK'
            headers = [('Content-type', 'application/json;charset=utf-8')]
            start_response(status, headers)
            return [ json.dumps( details ).encode('utf-8') ]
        else:
            raise RESTException('405 METHOD_NOT_ALLOWED',
                "Method '{REQUEST_METHOD}' not allowed".format_map(environ)
            )

''' 创建roulette服务器 '''
def roulette_server_3(count=1):
    from wsgiref.simple_server import make_server
    from wsgiref.validate import validator
    wheel = American()
    roulette = Roulette(wheel)
    debug = Validator(roulette)
    httpd = make_server('', 8080, debug)
    if count is None:
        httpd.server_forever()
    else:
        for c in range(count):
            httpd.handle_request()


''' 创建roulette客户端 '''
def roulette_client(method='GET', path='/', data=None):
    rest = http.client.HTTPConnection('localhost', '8080')
    if data:
        header = {'Content-type': 'application/json; charset=utf-8'}
        params = json.dumps(data).encode('utf-8')
        rest.request(method, path, params, header)
    else:
        rest.request(method, path)
    response = rest.getresponse()
    raw = response.read().decode('utf-8')
    if 200 <= response.status < 300:
        document = json.loads(raw)
        return document
    else:
        print(response.stauts, response.reason)
        print(response.getheaders())
        print(raw)


''' 验证客户端与服务器 '''
import concurrent.futures
with concurrent.futures.ProcessPoolExecutor() as executor:
    executor.submit(roulette_server_3, 4)
    time.sleep(3)
    print(roulette_client('GET', '/player/'))
    print(roulette_client('POST', '/bet/', {'bet': 'Black', 'amout': 2}))
    print(roulette_client('GET', '/bet/'))
    print(roulette_client('POST', '/wheel/'))
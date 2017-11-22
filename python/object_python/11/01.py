#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' [ 传输与共享对象 ] '''


''' 用REST实现CRUD操作 '''

''' 创建简单的REST应用程序和服务器 '''


import random

# 定义一个简化的轮盘
class Wheel:
    """Abstract, zero bins omiited."""
    def __init__( self ):
        self.rng = random.Random() # Random对象 == random对象
        self.bins = [
            {
                str(n): (35, 1),
                self.redblack(n): (1, 1),
                self.hilo(n): (1, 1),
                self.evenodd(n): (1, 1),
            } for n in range(1, 37)
        ]
    
    @staticmethod
    def redblack( n ):
        return 'Red' if n in (1, 3, 5, 7, 9, 12, 14, 16, 18, 
            19, 21, 23, 25, 27, 30, 32, 34, 36) else 'Black'

    @staticmethod
    def hilo( n ):
        return 'Hi' if n >= 19 else 'Lo'

    @staticmethod
    def evenodd( n ):
        return 'Even' if n % 2 == 0 else 'Odd'

    def spin( self ):
        return self.rng.choice( self.bins )

"""
Wheel类包含了一个箱子的列表。每个箱子都是一个dict,键代表的是当球落到箱子中时会成为赢家
的注，值是赔率。这里我们只定义了一个很短的下注列表，完整的轮盘下注列表非常庞大。

"""

class Zero:
    def __init__( self ):
        super().__init__()
        self.bins += [ {'0' : {35, 1}} ]

class DoubleZero:
    def __init__( self ):
        super().__init__()
        self.bins += [ {'00': (35, 1)} ]

"""
Zero mixin类将bins初始化为单零。DoubleZero mixin类将bins初始化为双零。这些是相对简单
的箱子，只有定了这个号码本身才能羸得这些箱子的投注。
"""

class American( Zero, DoubleZero, Wheel ):
    pass

class European( Zero, Wheel ):
    pass

american = American()
european = European()

print( 'SPIN', american.spin() )
# SPIN {'12': (35, 1), 'Hi': (1, 1), 'Red': (1, 1), 'Even': (1, 1)}
"""
dict中的键是投注的名称，值是包含两个元素、用于表示赔率的元组。
上面例子中Red 12是赢家，如果我们在12上下注，因为赔率是35比1,所以会赢得35倍于投注的回
报。其他投注的赔率是1比1， 我们会赢得和投注一样的回报。
"""

"""
我们定义一个简单路径来决定使用哪种轮盘的WSGI应用程序。
一个类似于http://localhost:8080/european/的URI会使用欧式轮盘。另一种路径使用美式轮盘。
"""

import sys
import wsgiref.util
import json

def wheel(environ, start_response):
    # wsgiref.util.shift_path_info()函数扫描environ['PATH_INFO']的值
    # 这会对请求中的路径信息的一个层次进行解析，它会返回找到的值，当没有提供路径时，
    # 它会返回None。
    request = wsgiref.util.shift_path_info(environ) # 1. Parse.
    # 记录日志则必须将信息写到sys.stderr中。任何写到sys.stderr的信息都会作为WSGI
    # 应用程序响应的一部分
    print( 'wheel', request, file=sys.stderr ) # 2. Loggin.
    if request.lower().startswith('eu'): # 3. Evaluate.
        winner = european.spin()
    else:
        winner = american.spin()

    status = '200 OK' # 4. Request
    headers = [('Content-type', 'application/json; charset=utf-8')]
    start_response(status, headers)
    return [ json.dumps(winner).encode('UTF-8') ]


# 启动这个服务器的演示版本
from wsgiref.simple_server import make_server
# wsgiref.simple_server.make_server()函数创建了服务器对象。

def roulette_server(count=1):
    httpd = make_server('', 8080, wheel)
    if count is None:
        httpd.serve_forever()
    else:
        for c in range(count):
            httpd.handle_request()


''' 实现REST客户端 '''
import http.client
import json

def json_get( path='/' ):
    rest = http.client.HTTPConnection('localhost', 8080)
    rest.request('GET', path)
    response = rest.getresponse()
    raw = response.read().decode('utf-8')
    if response.status == 200:
        document = json.loads(raw)
        print( document )
    else:
        print( raw )

"""
RESTful API的基本使用方法。
http.client模块包含了4个步骤：
1. 用HTTPConnection()建立连接。
2. 用命令和路径发送请求。
3. 获取响应。
4. 读取响应中的数据。
"""


''' 演示RESTful服务并创建单元测试 '''
"""
用concurrent.futures模块创建一个独立的子进程来运行服务器。
"""
import concurrent.futures
import time

with concurrent.futures.ProcessPoolExecutor() as executor:
    executor.submit( roulette_server, 4)
    time.sleep(2)
    json_get()
    json_get()
    json_get('/european/')
    json_get('/european/')


"""
通过创建concurrent.futures.ProcessPoolExecutor实例创建了一个独立的进程。
运行了json_get()客户端函数再次，用于读取默认路径/。然后，用GET请求两次'/european/'
executor.submit()函数让进程池执行roulette_server(4)函数。现在服务器会处理4个请求，
然后终止。由于ProcessPoolExecutor是一个上下文管理器，因此可以保证所有的资源都会被适当
地释放。
"""


if __name__ == '__main__':
    roulette_server()
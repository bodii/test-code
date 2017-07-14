#!/usr/bin/env python
# -*- coding:utf-8 -*-

#先运行server.py
#再在浏览地址栏输入服务器地址和端口号，以及参数，就可以访问了

# 1.server.py
"""
Topic:
Desc:此文件为WSGI(Web Server Gateway Interface) 服务器网关接口的模拟服务器文件
     保存为 server.py
     先运行
"""

# 导入wsgiref模块
from wsgiref.simple_server import make_server
#导入自己编写的application函数
from hello import application

# 创建一个服务器，IP地址为空，端口号是：8000 ,处理函数是application
httpd = make_server('', 8000, application)

print('Serving HTTP on port 8000...')

# 开始监听HTTP请求
httpd.server_forever()


# 2.hello.py (第一种内容，静态）
"""
Topic:
Desc:此文件为WSGI(Web Server Gateway Interface) 服务器网关接口的 http请求文件
     保存为 hello.py
     后运行
"""
def application(environ, start_response):
        start_response('200 OK',[('Content-Type','text/html')])
        return [b'<h1>Hello,Web!</h1>']
        
# 3.hello.py （第二种内容，动态）
def application(environ,start_response):
        start_response('200 OK',[('Content-Type','text/html')])
        body = '<h1>hello,%s!</h1>' % (environ['PATH_INFO'][1:2] or 'web')
        return [body.encode('utf-8')]

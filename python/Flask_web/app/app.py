#!/usr/bin/env python3
# -*- coding:utf-8 -*-

from flask import Flask # Flask 对象
from flask import render_template # 模板
from flask import current_app # 当前app对象
from flask import request # 获取或设置请求信息
from flask import make_response # 生成response对象 
from flask import redirect # 重定向
from flask import abort  # 特殊生成响应
from flask.ext.script import Manager
from flask.ext.bootstrap import Bootstrap


# 实例化网站应用
app = Flask(__name__)

# 实例化bootstrap
bootstrap = Bootstrap(app)

# 网站应用内容
@app.route('/')
def hello():
    # 添加一个response对象
    # response = make_response('<h1>This document carries a cookie!</h1>')
    # response.set_cookie('answer', value='42', max_age=6000)
    # return response
    return render_template('index.html')

@app.route('/user/<name>')
def user(name):
    return render_template('user.html', name=name)

@app.route('/test')
def test():
    return render_template('test.html')


if __name__ == '__main__':
    app.run(port=8000, debug=True)
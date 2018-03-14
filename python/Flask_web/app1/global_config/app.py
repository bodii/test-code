#!/usr/bin/env python3
#-*- coding:utf-8 -*-

__all__ = [
    'BASE_DIR',
    'BASE_NAME',
    'OPT',
    'SEP',
    'Config',
    'DevelopmentConfig',
    'TestingConfig',
    'ProductionConfig',
    'site_config',
    'default_config',
    ]

'''
全局配制项 
'''
import os
import os.path as opt

# 应用根目录
BASE_DIR = opt.abspath(opt.join(opt.dirname(__file__), '../'))
# 应用根目录-目录级名称
BASE_NAME = opt.basename(BASE_DIR)
# 目录对象
OPT = opt
# 目录对象分隔符
SEP = opt.sep

class Config:
    # 网站host
    HOST = '127.0.0.1'
    # 网站端口号
    PORT = '8000'
    SERVER_NAME = HOST + ':' + PORT
    # 网站名称
    SITE_NAME = 'app'
    # admin 管理员句子
    ADMIN = 'bodii'
     

    # 模板目录
    TEMPLATE_PATH = OPT.join(BASE_DIR, 'template')
    # 静态文件目录（css、js、image等）
    STATIC_PATH = OPT.join(BASE_DIR, 'static')
    # 自定义扩展类目录
    library_PATH = OPT.join(BASE_DIR, 'library')
    # 数据库文件目录
    DATABASE_PATH = OPT.join(BASE_DIR, 'database')
    # 模型目录
    MODEL_PATH = OPT.join(BASE_DIR, 'model')
    # 全站全局配制目录
    GLOBAL_CONFIG_PATH = OPT.join(BASE_DIR, 'global_config')
    # 路由目录
    ROUTE_PATH = OPT.join(BASE_DIR, 'route')
    # 数据迁移目录    
    MIGRATIONS_PATH = OPT.join(BASE_DIR, 'migrations')
    # 测试目录
    TEST_PATH = OPT.join(BASE_DIR, 'test')
    # 虚拟环境目录 
    VENV_PATH = OPT.join(BASE_DIR, 'venv')

    # 必须配置项
    SECRET_KEY = os.getenv('SECRET_KEY') or 'hard to guess string'
    SQLALCHEMY_DATABASE_URI = 'sqlite:///' + OPT.join(DATABASE_PATH, 'data.sqlite')
    SQLALCHEMY_COMMIT_ON_TEARDOWN = True
    SQLALCHEMY_TRACK_MODIFICATIONS = True

    # 指定每页显示的数量
    FLASKY_POSTS_PER_PAGE = 20


    @staticmethod
    def init_app(app):
        """初始化应用
        
        :param app objcet:
        """
        pass


# 开发时的配制
class DevelopmentConfig(Config):
    DEBUG = True
    # 网站host
    HOST = '127.0.0.1'
    # 网站端口号
    PORT = '8000'
    SERVER_NAME = HOST + ':' + PORT
    # 网站名称
    SITE_NAME = 'app_dev'


# 测试时的配制
class TestingConfig(Config):
    TESTING = True
    # 网站名称
    SITE_NAME = 'app_test'
    # 网站host
    HOST = '127.0.0.1'
    # 网站端口号
    PORT = '8000'
    SERVER_NAME = HOST + ':' + PORT


# 正式环境时的配制
class ProductionConfig(Config):
    # 网站host
    HOST = '127.0.0.1'
    # 网站端口号
    PORT = '8000'
    SERVER_NAME = HOST + ':' + PORT


site_config = {
    'development' : DevelopmentConfig,
    'testing' : TestingConfig,
    'production' : ProductionConfig,

    'default' : DevelopmentConfig,
}

default_config = site_config['default']

#!/usr/bin/env python3
# -*- coding:utf-8 -*-


__all__  = [
    'db',
    'bootstrap',
    'mail',
    'moment',
    'principal',
    'login_manager',
    ]

import os
from flask import Flask, render_template, current_app
from flask_bootstrap import Bootstrap 
from flask_mail import Mail
from flask_moment import Moment
from flask_login import LoginManager
from flask_sqlalchemy import SQLAlchemy
from flask_principal import Principal
from global_config.app import site_config, OPT
from flask_pagedown import PageDown

bootstrap = Bootstrap()
mail = Mail()
moment = Moment()
db = SQLAlchemy()
principal = Principal()
login_manager = LoginManager()
pagedown = PageDown()

login_manager.session_protection = 'strong'
login_manager.login_view = 'auth.login'


def create_app(config_name):
    app = Flask(__name__, static_folder='../static', template_folder='../templates')
    app.config.from_object(site_config[config_name])
    site_config[config_name].init_app(app)

    bootstrap.init_app(app)
    mail.init_app(app)
    moment.init_app(app)
    login_manager.init_app(app)
    db.init_app(app)
    pagedown.init_app(app)

    # 注册蓝本
    from route.index import index as index_blueprint
    from route.auth import auth as auth_blueprint
    from route.error import error as error_blueprint
    app.register_blueprint(index_blueprint)
    app.register_blueprint(auth_blueprint)
    app.register_blueprint(error_blueprint)

    return app

# app = create_app(os.getenv('config') or 'default')

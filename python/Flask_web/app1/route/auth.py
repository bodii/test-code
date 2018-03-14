#!/usr/bin/env python3
# -*- coding:utf-8 -*-

from flask_login import (
    login_required, 
    login_user,
    logout_user,
    current_user,
)
from flask import (
    render_template, 
    redirect, 
    request, 
    url_for,
    flash,
    session,
)
from model.user import User
from library.userForm import LoginForm, RegistrationForm
from library.mail import *
from flask import Blueprint
from global_config.app import default_config
from main import db

# 注册蓝图
auth = Blueprint('auth', __name__)

''' 保护路由，只让认证用户访问 '''
@auth.route('/secret')
@login_required
def secret():
    return 'Omly authenticated users are allowed!'


''' 用户登录 '''
@auth.route('/login', methods=['GET', 'POST'])
def login():
    form = LoginForm()
    if form.validate_on_submit():
        user = User.query.filter_by(email=form.email.data).first()
        if user is not None and user.verify_password(form.password.data):
            login_user(user, form.remember_me.data)
            if default_config.ADMIN:
                send_email(
                    default_config.ADMIN,
                    'New User',
                    'mail/new_user',
                    user = user
                    )
            return redirect(request.args.get('next') or url_for('index.home'))
        flash('Invalid username or password.')
    return render_template(
                'auth/login.html', 
                form = form,
                title = 'login',
                )


''' 用户登出 '''
@auth.route('/logout')
@login_required
def logout():
    logout_user()
    flash('You have been logged out.')
    return redirect(url_for('index.home'))


''' 注册 '''
@auth.route('/register', methods=['GET', 'POST'])
def register():
    form = RegistrationForm()
    if form.validate_on_submit():
        user = User(
            email=form.email.data,
            username = form.username.data,
            password = form.password.data,
            )
        db.session.add(user)
        db.session.commit()
        token = User.token()
        send_email( # 注册提交后发送邮件
            User.email,
            'Confirm Your Account',
            'auth/email/confirm',
            user = User,
            token = token,
            )
        flash('You can now login.')
        flash('A confirmation email has been sent to you by email.')
        return redirect(url_for('auth.login'))
    return render_template(
                    'auth/register.html', 
                    form = form,
                    title = 'register',
                    )


''' 账号确认 '''
@auth.route('/confirm/<token>')
@login_required
def confirm(token):
    if current_user.confirmed:
        return redirect(url_for('index.home'))
    if current_user.confirm(token):
        flash('You have confirmed your account. Thanks!')
    else:
        flash('The confirmation link is invalid or has expired.')
    return redirect(url_for('index.home'))


'''
同时满足以下3个条件时，before_app_request处理程序会拦截请求
1. 用户已登录(current_user.is_authenticated()必须返回True)
2. 用户的账户还未确认
3. 请求的端点(使用request.endpoin获取)不在认证蓝本中。访问认证路由要获取权限，
因为这些路由的作用是让用户确认账户或执行其他账户管理操作。
如果请深圳市满足以上3个条件，则会被重定向到 unconfirmed路由，显示一个确认账户
的相关信息 
'''
# ''' 处理程序中过滤未确认的账户 '''
# @auth.before_app_request
# def before_request():
#     if current_user.is_authenticated():
#         current_user.ping()
#         if not current_user.confirmed \
#             and request.endpoint[:5] != 'auth.':
#             return redirect(url_for('auth.unconfirmed'))
#
#
# @auth.route('/unconfirmed')
# def unconfirmed():
#     if current_user.is_anonymous() or current_user.confirm:
#         return redirect(url_for('index.home'))
#     return render_template(
#                     'auth/uncnfirmed.html',
#                     title = 'uncnfirmed',
#                     )

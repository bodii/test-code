#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 用户模型 '''

from werkzeug.security import generate_password_hash, check_password_hash
from flask_login import UserMixin, AnonymousUserMixin
from main import login_manager, db
from library.token import TokenCase
from model.role import Role
from flask import current_app, request
from global_config.permission import Permission
from global_config.app import default_config
from datetime import datetime
from hashlib import md5


class User(UserMixin, db.Model, TokenCase):
    __tablename__ = 'user'
    id = db.Column(db.Integer, primary_key=True)
    email = db.Column(db.String(64), unique=True, index=True)
    username = db.Column(db.String(64), unique=True, index=True)
    password_hash = db.Column(db.String(128))
    role_id = db.Column(db.Integer, db.ForeignKey('roles.id'))
    confirmed = db.Column(db.Boolean, default=False)
    name = db.Column(db.String(64)) # 真实姓名
    location = db.Column(db.String(64)) # 所在地
    about_me = db.Column(db.Text()) # 自我介绍
    member_since = db.Column(db.DateTime(), default=datetime.now) # default参数可以接受函数作为默认值
    last_seen = db.Column(db.DateTime(), default=datetime.now)
    posts = db.relationship('Post', backref='author', lazy='dynamic')
    followed = db.relationship(
        'Follow',
        foreign_keys = [Follow.follower_id],
        backref=db.backref('follower', lazy='joined'),
        lazy='dynamic',
        cascade='all, delete-orphan',
        )
    followers = db.relationship(
        'Follow',
        foreign_keys=[Follow.followed_id],
        backref=db.backref('followed', lazy='joined'),
        lazy='dynamic',
        cascade='all, delete-orphan', 
        )

    def __init__(self, **kw):
        super(User, self).__init__(**kw)
        if self.role is None:
            if self.email == current_app.config['ADMIN']:
                self.role = Role.query.filter_by(permissions=Permission.ADMINISTER).first()
            if self.role is None:
                self.role = Role.query.filter_by(default=True).first()

    def __repr__(self):
        return '<User %r>' % self.username

    @property
    def password(self):
        raise AttributeError('password is not a readable attribute')

    @password.setter
    def password(self, password):
        """设置加密密码"""
        self.password_hash = generate_password_hash(password)

    def verify_password(self, password):
        """检验密码是否正确"""
        return check_password_hash(self.password_hash, password)
    
    @login_manager.user_loader
    def load_user(self, user_id):
        """通过user_id获取用户信息"""
        return self.query.get(int(user_id))

    def token(self, expires=3600):
        """生成一个token令牌，有效期为1小时"""
        return self.get_token({'confirm': self.id}, expires)
    
    def confirm(self, token=''):
        """检验token是否有效，如果有效则把confirmed设为True"""
        if self.checkToken(token, self.id):
            self.confirmed = True
            db.session.add(self)
            return True
        return False
    
    def can(self, permissions):
        ''' 检测用户权限是否存在 '''
        return self.role is not None and \
            (self.role.permissions & permissions) == permissions

    def is_administrator(self):
        ''' 是否是管理员 '''
        return self.can(Permission.ADMINISTER)

    def ping(self):
        ''' 更新最后访问时间 '''
        self.last_seen = datetime.now()
        db.session.add(self)

    def change_email(self, token):
        self.email = new_email
        self.avatar_hash = md5(self.email.encode('utf-8')).hexdigest()
        db.session.add(self)
        return True

    def gravatar(self, size=100, default='identicon', rating='g'):
        ''' 获取用户头像 '''
        if request.is_secure:
            url = 'https://secure.gravatar.com/avatar'
        else:
            url = 'http://www.gravatar.com/avatar'
        hash = self.avatar_hash or md5(self.email.encode('utf-8')).hexdigest()
        ''' Gravatar查询字符串参数：
            s 图片大小，单位为像素
            r 图片级别。可选值有‘g’、‘pg’、‘r’、‘x’
            d 没有注册Gravatar服务的用户使用的默认图片生成方式。可选值有：
        “404”，返回404错误；默认图片的URL;图片生成器"mm"、"identicon"、
        "monsterid"、"wavatar"、"retro"或"blank"之一
            fd 强制使用默认头像
        '''
        return '{url}/{hash}?s={size}&d={default}&r={rating}'.format(
                url=url, hash = hash, size = size,
                default = default, rating = rating
            )
    
    @staticmethod
    def generate_fake(count=100):
        """生成虚拟用户"""
        from sqlalchemy.exc import IntegrityError
        from random import seed
        import forgery_py

        seed()
        for i in range(count):
            u = User(
                email = forgery_py.internet.email_address(),
                username = forgery_py.internet.user_name(True),
                password = forgery_py.lorem_ipsum.word(),
                confirmed = True,
                name = forgery_py.name.full_name(),
                location = forgery_py.address.city(),
                about_me = forgery_py.lorem_ipsum.sentence(),
                member_since = forgery_py.date.date(True)
            )
            db.session.add(u)
            try:
                db.session.commit()
            except IntegrityError:
                db.session.rollback()


    ''' 关注关系的辅助方法 '''
    def follow(self, user):
        if not self.is_following(user):
            f = Follow(follower=self, followed=user)
            db.session.add(f)

    def unfollow(self, user):
        f = self.followed.filter_by(followed_id=user.id).first()
        if f:
            db.session.delete(f)

    def is_following(self, user):
        return self.followed.filter_by(
                        followed_id=user.id
                        ).first() is not None

    def is_followed_by(self, user):
        return self.followers.filter_by(
                        follower_id=user.id
                        ) is not None

    


class AnonymousUser(AnonymousUserMixin):
    def can(self, permissions):
        return False

    def is_administrator(self):
        return False

login_manager.anonymous_user = AnonymousUser


''' 关注关联表模型 '''
class Follow(db.Model):
    __tablename__ = 'follows'
    follower_id = db.Column(
        db.Integer, 
        db.ForeignKey('users.id'),
        primary_key=True,
        )
    followed_id = db.Column(
        db.Integer,
        db.ForeignKey('users.id'),
        primary_key=True,
        )
    timestamp = db.Column(db.DateTime, default=datetime.now)
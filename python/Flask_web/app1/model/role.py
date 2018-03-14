#!/usr/bin/env python3
#-*- coding:utf-8 -*-

''' 角色模型 '''
from werkzeug.security import generate_password_hash, check_password_hash
from flask_login import UserMixin
from main import db
from library.token import TokenCase
from global_config.permission import Permission

''' 
程序权限配置示例 

--------------------------------------------------------------
操作                 位值                说明
关注用户              0b00000001(0x01)   关注其他用户
在他人的文章中发表评论   0b00000010(0x02)   在他人撰写的文章中发布评论
写文章                0b00000100(0x04)   写原创文章
管理他人发表的评论      0b00001000(0x08)   查处他人发表的不当评论
管理员权限             0b10000000(0x80)   管理网站
--------------------------------------------------------------
'''

class Role(db.Model):
    __tablename__ = 'roles'
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(64), unique=True)
    default = db.Column(db.Boolean, default=False, index=True)
    permissions = db.Column(db.Integer) # 权限
    users = db.relationship('User', backref='role', lazy='dynamic')

    def __repr__(self):
        return '<Role %r>' % self.name

    @staticmethod
    def insert_roles():
        ''' 创建角色 '''
        roles = {
            'User': (Permission.FOLLOW | 
                     Permission.COMMENT |
                     Permission.WRITE_ARTICLES, 
                     True
                     ),
            'Moderator': ( Permission.FOLLOW |
                           Permission.COMMENT |
                           Permission.WRITE_ARTICLES |
                           Permission.MODERATE_COMMENTS,
                           False
                           ),
            'Administrator':(0xff, False),
        }

        for r in roles:
            role = Role.query.filer_by(name=r).first()
            if role is None:
                role = Role(name=r)
            role.permissions = roles[r][0]
            role.default = roles[r][1]
            db.session.add(role)
        db.session.commit()
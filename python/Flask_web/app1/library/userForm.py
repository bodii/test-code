#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 权限表单 '''

from flask_wtf import FlaskForm
from wtforms import (
    StringField,
    PasswordField,
    BooleanField,
    SubmitField,
    TextAreaField,
    SelectField,
)
from wtforms.validators import (
    Required,
    Length,
    Email,
    Regexp,
    EqualTo,
)
from wtforms import ValidationError
from model.user import User
from model.role import Role


class LoginForm(FlaskForm):
    """登录表单"""
    email = StringField('Email', validators = [
            Required(),
            Length(1, 64),
            Email(),
        ]
    )
    password = PasswordField('Password', validators = [
            Required(),  
        ]
    )
    remember_me = BooleanField('Keep me logged in')
    submit = SubmitField('Log In')


class RegistrationForm(FlaskForm):
    """注册表单"""
    email = StringField('Email', validators = [
            Required(),
            Length(1, 64),
            Email(),
        ]
    )
    username = StringField('Username', validators = [
            Required(),
            Length(1, 64),
            Regexp(
                '^[A-Za-z][A-Za-z0-9_.]*$', 
                0,
                'Usernames must have only letters,\
                 numbers, dots or underscores'
                ),
        ]
    )
    password = PasswordField('Password', validators = [
            Required(),
            EqualTo(
                'password2',
                message = 'Passwords must match.' 
                ),
        ]
    )
    password2 = PasswordField(
        'Confirm password', 
        validators = [Required()]
        )
    submit = SubmitField('Register')

    def validate_email(self, field):
        if User.query.filter_by(email=field.data).first():
            raise ValidationError('Email already registered.')

    def validate_username(self, field):
        if User.query.filter_by(username=field.data).first():
            raise ValidationError('Username already in use.')


class EditProfileForm(FlaskForm):
    ''' 资料编辑表单 '''
    name = StringField('Real name', validators=[Length(0, 64)])
    location = StringField('Location', validators=[Length(0, 64)])
    about_me = TextAreaField('About me')
    submit = SubmitField('Submit')


class EditProfileAdminForm(FlaskForm):
    ''' 管理员资料编辑表单 '''
    email = StringField('Email', validatores=[
                        Required(), 
                        Length(1, 64), 
                        Email()
                        ]
                    )
    username = StringField('Username', validators=[
                        Required(),
                        Length(1, 64),
                        Regexp(
                            '^[A-Za-z][A-Za-z0-9_.]*$', 0,
                            'Usernames must have only letters \
                            numbers, dots or underscores'
                            )    
                        ]
                    )
    confirmed = BooleanField('Confirmed')
    role = SelectField('Role', coerce=int) # 把字段的值转换为整数
    name = StringField('Real name', validators=[Length(0, 64)])
    location = StringField('Location', validators=[Length(0, 64)])
    about_me = TextAreaField('About me')
    submit = SubmitField('Submit')

    def __init__(self, user, *args, **kw):
        super(EditProfileAdminForm, self).__init__(*args, **kw)
        self.role.choices = [  # SelectField实例必须在其choices属性中设置各选项
            (role.id, role.name) 
            for role in Role.query.order_by(Role.name).all()
            ]
        self.user = user

    def validate_email(self, field):
        if field.data != self.user.name and \
            User.query.filter_by(email=field.data).first():
            raise ValidationError('Email already registered.')

    def validate_username(self, field):
        if field.data != self.user.username and \
            User.query.filter_by(username=field.data).first():
            raise ValidationError('Username already in use.')

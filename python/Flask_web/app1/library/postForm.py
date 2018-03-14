#!/usr/bin/env python3
#-*- coding:utf-8 -*-


''' 博客文章表单 '''

from flask_wtf import FlaskForm
from wtforms import (
    StringField,
    PasswordField,
    BooleanField,
    TextAreaField,
    SelectField,
    SubmitField,
)
from wtforms.validators import (
    Required,
    Length,
    Email,
    Regexp,
    EqualTo,
)
# 转换成Markdown富文本编辑器
from flask_pagedown.fields import PageDownField

class PostForm(FlaskForm):
    body = PageDownField("What's on your mind?", validators=[Required()])
    submit = SubmitField('Submit')
    
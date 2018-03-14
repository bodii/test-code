#!/usr/bin/env python3
# -*- coding:utf-8 -*-

__all__ = [
    'send_email',
    'send_async_email',
    ]

from flask_mail import Mail, Message
from main import mail
from flask import render_template
from threading import Thread
from global_config.mail import Config as mail_config

# 发送邮件
def send_email(to, subject, template, **kw):
    msg = Message(
        mail_config.FlASKY_MAIL_SUBJECT_PREFIX + subject,
        sender = mail_config.FLASKY_MAIL_SENDER,
        recipients=[to]
        )
    msg.body = render_template(template + '.txt', **kw)
    msg.html = render_template(template + '.html', **kw)
    mail.send(msg)

# 异步发送邮件
def send_async_email(app, msg):
    with app.app_context():
        mail.send(msg)

def thread_send_email(to, subject, template, **kw):
    msg = Message(
        mail_config.FlASKY_MAIL_SUBJECT_PREFIX + subject,
        sender = mail_config.FLASKY_MAIL_SENDER,
        recipients=[to]
        )
    msg.body = render_template(template + '.txt', **kw)
    msg.html = render_template(template + '.html', **kw)
    thread = Thread(target=send_async_email, args=[app, msg])
    thread.start()
    return thread
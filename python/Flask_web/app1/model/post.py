#!/usr/bin/env python3
#-*- coding:utf-8 -*-


''' 博客文章模型 '''

from main import db
from datetime import datetime

# 将文章的Markdown源文本保存的数据库中
from markdown import markdown
import bleach


class Post(db.Model):
    __tablename__ = 'posts'
    id = db.Column(db.Integer, primary_key=True)
    body = db.Column(db.Text)
    timestamp = db.Column(db.DateTime, index=True, default=datetime.now)
    author_id = db.Column(db.Integer, db.ForeignKey('users.id'))
    body_html = db.Column(db.Text)

    @staticmethod
    def generate_fake(count=100):
        """生成虚拟博客文章"""
        from random import seed, randint
        from sqlalchemy.exc import IntegrityError
        import forgery_py

        seed()
        user_count = User.query.count()
        for i in range(count):
            u = User.query.offset(randint(0, user_count - 1)).first()
            self.body = forgery_py.lorem_ipsum.sentences(randint(1,
            self.timestamp = forgery_py.date.date(True)
            self.author = u
            db.session.add(self)
            db.session.commit()

    @staticmethod
    def on_changed_body(target, value, oldvalue, initiator):
        allowed_tags = [
                        'a', 'abbr', 'acronym', 'b', 'blockquote', 'code',
                        'em', 'i', 'li', 'ol', 'pre', 'strong', 'ul',
                        'h1', 'h2', 'h3', 'p',
                        ]
        target.body_html = bleach.linkify(bleach.clean(
                            markdown(value, output_format='html'),
                            tags=allowed_tags, strip=True
                        ))

db.event.listen(Post.body, 'set', 'Post.on_changed_body')

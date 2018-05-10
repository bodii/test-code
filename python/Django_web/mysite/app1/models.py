#!/usr/bin/env python3
#-*- coding:utf-8 -*-


from django.db import models
from django.utils.encoding import python_2_unicode_compatible
from django.utils import timezone
import datetime

# Create your models here.

@python_2_unicode_compatible
class Question( models.Model ):
    """问题和发布日期"""
    question_text = models.CharField(max_length=200)
    pub_date = models.DateTimeField('date published')

    def __str__( self ):
        return self.question_text

    def was_published_recently( self ):
        return self.pub_date >= timezone.now() - datetime.timedelta(days=1)


@python_2_unicode_compatible
class Choice( models.Model ):
    """选择的文本和投票计数"""
    question = models.ForeignKey(Question, on_delete=models.CASCADE)
    choice_text = models.CharField(max_length=200)
    votes = models.IntegerField(default=0)

    def __str__( self ):
        return self.choice_text

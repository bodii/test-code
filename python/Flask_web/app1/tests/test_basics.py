#!/usr/bin/env python3
# -*- coding:utf-8 -*-


''' 单元测试 '''


import unittest
from flask import current_app
from main import create_app, db


class BasiceTestCase(unittest.TestCase):
    def setUp(self):
        """尝试创建一个测试环境"""
        self.app = create_app('testing')
        self.app_context = self.app.app_context()
        self.app_context.push()
        db.create_all()

    def tearDown(self):
        """测试完成的清除工作"""
        db.session.remove()
        db.drop_all()
        self.app_context.pop()

    def test_app_exists(self):
        """要执行的测试方法->是否创建成功"""
        self.assertFalse(current_app is None)

    def test_app_is_testing(self):
        """要执行的测试方法->获取测试环境的网站配置"""
        self.assertTrue(current_app.site_config['TESTING'])
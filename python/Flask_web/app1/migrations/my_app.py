#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 数据库迁移 '''

from flask_migrate import Migrate, MigrateCommand
from main import app, db
from flask_script import Manager 

manager = Manager(app)

migrate = Migrate(app, db)
manager.add_command('db', MigrateCommand)
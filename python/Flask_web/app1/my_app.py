#!/usr/bin/env python3
# -*- coding:utf-8 -*-

import os, sys
sys.path.append(os.path.abspath(os.path.dirname(__file__)))

from flask_script import Manager, Shell
from flask_migrate import Migrate, MigrateCommand
from main import create_app, db
from model.role import Role
from model.user import User

app = create_app(os.getenv('config') or 'default')
manager = Manager(app)
migrate = Migrate(app, db)

def make_shell_context():
    return dict(app=app, db=db, User=User, Role=Role)
manager.add_command('shell', Shell(make_context=make_shell_context))
manager.add_command('db', MigrateCommand)

# 启动单元测试的命令
@manager.command
def test():
    """Run the unit tests."""
    import unittest
    tests = unittest.TestLoader.discover('tests')
    unittest.TextTestRunner(verbosity=2).run(tests)

if __name__ == '__main__':
    app.run()

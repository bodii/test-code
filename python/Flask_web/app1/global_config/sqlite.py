#!/usr/bin/env python3
# -*- coding:utf-8 -*-


__all__ = [
    'config', 
    'default',
    ]


"""
linux = "sqlite:////absolute/path/to/database"
windows = "sqlite:///c:absoute/path/to/database"
"""

from app import default_config, OPT

SQLITE_PATH = default_config.SQLITE_PATH

class Config:
    db_filename = 'data.sql'
    PATH = OPT.join(SQLITE_PATH, db_filename)

        
class DevelopmentConifg(Config):
    db_filename = 'data-dev.sql'
    PATH = OPT.join(SQLITE_PATH, db_filename)


class TestingConfig(Config):
    db_filename = 'data-test.sql'
    PATH = OPT.join(SQLITE_PATH, db_filename)


class ProductionConfig(Config):
    db_filename = 'data.sql'
    PATH = OPT.join(SQLITE_PATH, db_filename)

config = {
    'development' : DevelopmentConifg,
    'testing' : TestingConfig,
    'production' : ProductionConfig,

    'default' : DevelopmentConifg,
}

default_config = config['default']
#!/usr/bin/env python3
# -*- coding:utf-8 -*-


''' 管理嵌套的配置上下文 '''

import yaml
import logging

class Logging_Config:
    def __enter__( self, fielname="logging.config" ):
        logging.config.dictConfig(yaml.load(fielname))

    def __exit__( self, *exc ):
        logging.shutdown()


class Application_Config:
    def __enter__( self ):
        # Build os.environ defaults.
        # Load files.
        # Build ChainMap from environs and files.
        # Parse command-line arguments.
        return namespace

    def __exit__( self, *exc ):
        pass

if __name__ == '__main__':
    with Logging_Config():
        with Application_Config() as config:
            simulate_blackjack(config)
#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 另一种方法关闭logging '''
"""
使用atexit handler来关闭logging
"""

import atexit
import sys
import logging

if __name__ == "__main__":
    logging.config.dictConfig( yaml.load('log_config.yaml') )
    atexit.register(logging.shutdown)
    try:
        application = Main()
        status = application.run()
    except Exception as e:
        logging.exception( e )
        status = 2
    sys.exit(status)

"""
当应用程序退出时，会调用给定的函数。如果main()函数中已经正确地处理了异常，那么
可以用简单行多的 status = main(); sys.exit(status)替代try:块。
"""
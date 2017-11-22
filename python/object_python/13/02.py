#!/usr/bin/env python3
# -*- coding:utf-8 -*-


''' 开始和关闭日志记录系统 '''
"""
logging.shutdown()确保所有的缓冲区都被刷新
当处理顶层的错误和异常时，我们有两种明确的技术用于确保所有的缓冲区都被写入，其中
一种技术是在try:代码块中使用finally语句。
"""
import logging
import sys

if __name__ == "__main__":
    logging.config.dictConfig( yaml.load('log_config.yaml') )
    try:
        application = Main()
        status = application.run()
    except Exception as e:
        logging.exception( e )
        status = 2
    finally:
        logging.shutdown()
    sys.exit( status )


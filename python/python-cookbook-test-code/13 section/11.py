#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十三章：脚本编程与系统管理
Description: 许多人使用 Python 作为一个 shell 脚本的替代，用来实现常用系统任务的自动化，
如文件的操作，系统的配置等。本章的主要目标是描述光宇编写脚本时候经常遇到的
一些功能。例如，解析命令行选项、获取有用的系统配置数据等等。第 5 章也包含了
与文件和目录相关的一般信息。

Title:           给简单脚本增加日志功能
Issue:你希望在脚本和程序中将诊断信息写入日志文件。
Answer: 打印日志最简单方式是使用 logging 模块。
"""
import logging

def main():
    logging.basicConfig(
        filename ='app.log',
        level = logging.error
    )

    hostname = 'www.python.org'
    item = 'spam'
    filename = 'data.csv'
    mode = 'r'

    logging.critical('Host %s unknown', hostname)
    logging.error("Couldn't find %s", item)
    logging.warning('Feature is deprecated')
    logging.info('Opening file %s, mode=%s', filename, mode)
    logging.debug('Got here')

if __name__ == '__main__':
    main()

"""
上面五个日志调用（critical(), error(), warning(), info(), debug()）以降序方式表示
不同的严重级别。 basicConfig() 的 level 参数是一个过滤器。所有级别低于此级别
的日志消息都会被忽略掉。每个 logging 操作的参数是一个消息字符串，后面再跟一个
或多个参数。构造最终的日志消息的时候我们使用了% 操作符来格式化消息字符串。
运行这个程序后，在文件 app.log 中的内容应该是下面这样：
CRITICAL:root:Host www.python.org unknown
ERROR:root:Could not find 'spam'
If you want to change the output or level of output, you can change the parameters to the basicConfig() call. For example: 如果你想改变输出等级，你可以修改
basicConfig() 调用中的参数。例如：
logging.basicConfig(
    filename='app.log',
    level=logging.WARNING,
    format='%(levelname)s:%(asctime)s:%(message)s')

CRITICAL:2012-11-20 12:27:13,595:Host www.python.org unknown
ERROR:2012-11-20 12:27:13,595:Could not find 'spam'
WARNING:2012-11-20 12:27:13,595:Feature is deprecated

上面的日志配置都是硬编码到程序中的。如果你想使用配置文件，可以像下面这样
修改 basicConfig() 调用：
"""
import logging
import logging.config

def main():
    logging.config.fileConfig('logconfig.ini')

"""
如果你想修改配置，可以直接编辑文件 logconfig.ini 即可

尽管对于 logging 模块而已有很多更高级的配置选项，不过这里的方案对于简单
的程序和脚本已经足够了。只想在调用日志操作前先执行下 basicConfig() 函数方法，
你的程序就能产生日志输出了。
如 果 你 想 要 你 的 日 志 消 息 写 到 标 准 错 误 中， 而 不 是 日 志 文 件 中， 调 用
basicConfig() 时不传文件名参数即可。例如：
logging.baseicConfig(level=logging.INFO)

basicConfig() 在程序中只能被执行一次。如果你稍后想改变日志配置，就需要先
获取 root logger ，然后直接修改它。例如：
logging.getLogger().level = loging.DEBUG

需要强调的是本节只是演示了 logging 模块的一些基本用法。它可以做更多更高
级的定制。关于日志定制化一个很好的资源是 Logging Cookbook
"""


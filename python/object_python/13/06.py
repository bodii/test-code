#!/usr/bin/env python3
# -*- coding:utf-8 -*-

''' 定义指向多个目标输出的handler '''

"""
有些情况下，我们需要将日志发送到多个不同的目标，如下所示：
    可能希望有两份日志，这样可以提高运营的可靠性。
    可能希望使用复杂的Filter对象为不同的消息创建子集。
    可能希望为不同的目标定义不同的级别，通过这种方式分离调试信息和指示性信息。
    可能希望基于Loggers的名称定义不同的handlers来代表不同的日志源
"""
"""
尽管我们可以通过logging模块的API配置目标输出，但是，更清晰一些的通常做法是将日志
细节定义在配置文件中。一个优雅的解决方案是用YAML标记作为配置字典。可以用一种相对直
接的方法加载字典--使用logging.config.dictConfig(yaml.load(somefile))

!!在YAML配置文件中，必须以URI风格语法中的ext://sys.stderr来命名外部的Python资
源。
audit_file是一个会写入指定文件中的FileHandler默认情况下，会用a模式打开文件，向
文件末尾添加内容。
"""

import logging.config
import yaml
import os.path

current_paht = os.path.dirname(__file__) + os.path.sep

with open(current_paht + 'logging_config.yaml') as config:
    config_dict = yaml.load(config)
    # print(config_dict['version'])
    logging.config.dictConfig(config_dict)
    """
    将YAML文本转化为dict，然后用dictConfig()函数使用指定的字典配置日志。
    """
    # 创建两个Logger对象，一个属于verbose家族树，另一个属于audit家族树
    verbose = logging.getLogger('verbose.example.SomeClass')
    audit = logging.getLogger('audit.example.SomeClass')
    verbose.info('Verbose inforamtion')
    audit.info('Audit record with before and after')
    """
    logging模块会使用老式的%风格的格式化说明，这些格式化说明与str.format()
    方法的不同。当这们定义格式化参数时，使用{风格的格式说明，这和str.format()
    是一致的。
    """


''' 管理传播规则 '''

"""
如果需要添加一个根级别的日志记录器来捕获其他名称，那么要注意传播法则。
对配置文件的修改：
loggers:
    verbose:
        handlers: [console]
        level: INFO
        propagate: False # Added
    audit:
        handlers: [audit_file]
        level: INFO
        propagate: False # Added
    root: # Added
        handlers: [console]
        level: INFO

关闭了两个低层的日志记录器的传播功能：verbose和audit，添加了一个新的根级别的日志
记录器，由于这个日志记录器没有名称，所有我们声明了一个与loggers：平行的顶层字典root
如果我们不关心这两个低层日志记录器的传播功能，那么verbose或者audit的每条记录都会被
处理两次。
"""

''' 理解配置 '''
"""
basicConfig()方法能够保留在完成配置之前创建的所有日志处理器。但是，logging.config.
dictConfig()方法会默认禁用所有在配置完成前创建的日志处理器。
当组建一个庞大复杂的应用程序时，在import过程中，我们可能会创建模块级的日志处理器。主脚本
引入的模块可能在logging.config创建完成之前创建日志记录器。同样地，任何全局的对象或者类
定义都可能在配置完成前创建日志记录器。
我们通常需要在配置文件中添加下这行代码：
disable_existing_loggers: False
这行代码可以确保配置完成前创建的所有日志处理器仍然会将信息传播到配置文件创建的根日志记录器
中。
"""


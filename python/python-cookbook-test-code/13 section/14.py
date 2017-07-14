#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十三章：脚本编程与系统管理
Description: 许多人使用 Python 作为一个 shell 脚本的替代，用来实现常用系统任务的自动化，
如文件的操作，系统的配置等。本章的主要目标是描述光宇编写脚本时候经常遇到的
一些功能。例如，解析命令行选项、获取有用的系统配置数据等等。第 5 章也包含了
与文件和目录相关的一般信息。

Title:           限制内存和 CPU 的使用量
Issue:你想对在 Unix 系统上面运行的程序设置内存或 CPU 的使用限制。
Answer: resource 模块能同时执行这两个任务。
"""
import signal
import resource
import os

def time_exceeded(signo, tframe):
    print("Time's up!")
    raise SystemExit(1)

def set_max_runtime(seconds):
    soft, hard = resource.getrlimi(resource.RLIMIT_CPU)
    resource.setrlimit(resource.RLIMIT_CPG, (seconds, hard))
    signal.signal(signal.SIGXCPU, time_exceeded)

if __name__ == '__main__':
    set_max_runtime(15)
    while True:
        pass

"""
程序运行时， SIGXCPU 信号在时间过期时被生成，然后执行清理并退出。
要限制内存使用，设置可使用的总内存值即可，如下
"""
import resource
def limit_momory(maxsize):
    soft, hard = resource.getrlimi(resource.RLIMIT_AS)
    resource.setrlimit(resource.RLIMIT_AS, (maxsize, hard))

"""
像这样设置了内存限制后，程序运行到没有多余内存时会抛出 MemoryError 异常。

在本节例子中， setrlimit() 函数被用来设置特定资源上面的软限制和硬限制。软
限制是一个值，当超过这个值的时候操作系统通常会发送一个信号来限制或通知该进
程。硬限制是用来指定软限制能设定的最大值。通常来讲，这个由系统管理员通过设
置系统级参数来决定。尽管硬限制可以改小一点，但是最好不要使用用户进程去修改。
setrlimit() 函数还能被用来设置子进程数量、打开文件数以及类似系统资源的限
制。更多详情请参考 resource 模块的文档。
需要注意的是本节内容只能适用于 Unix 系统，并且不保证所有系统都能如期工作。
比如我们在测试的时候，它能在 Linux 上面正常运行，但是在 OS X 上却不能。
"""
#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十三章：脚本编程与系统管理
Description: 许多人使用 Python 作为一个 shell 脚本的替代，用来实现常用系统任务的自动化，
如文件的操作，系统的配置等。本章的主要目标是描述光宇编写脚本时候经常遇到的
一些功能。例如，解析命令行选项、获取有用的系统配置数据等等。第 5 章也包含了
与文件和目录相关的一般信息。

Title:            执行外部命令并获取它的输出
Issue:你想执行一个外部命令并以 Python 字符串的形式获取执行结果
Answer: 使用 subprocess.check output() 函数。
"""
import subprocess
out_bytes = subprocess.check_output(['netstat', '-a'])
"""
这段代码执行一个指定的命令并将执行结果以一个字节字符串的形式返回。如果你
需要文本形式返回，加一个解码步骤即可。例如：
"""
out_text = out_bytes.decode('utf-8')
"""
如果被执行的命令以非零码返回，就会抛出异常。下面的例子捕获到错误并获取返
回码：
"""
try:
    out_bytes = subprocess.check_output(['cmd', 'arg1', 'arg2'])
except subprocess.CalledProcessError as e:
    out_bytes = e.output
    code = e.returncode

"""
默认情况下， check output() 仅仅返回输入到标准输出的值。如果你需要同时收
集标准输出和错误输出，使用 stderr 参数
"""
out_bytes = subprocess.check_output(['cmd', 'arg1', 'arg2'],
                                   stderr=subprocess.STDOUT)
"""
如果你需要用一个超时机制来执行命令，使用 timeout 参数
"""
try:
    out_bytes = subprocess.check_output(['cmd', 'arg1', 'arg2'],timeout=5)
except subprocess.TimeoutExpired as e:
    pass

"""
通常来讲，命令的执行不需要使用到底层 shell 环境（比如 sh、 bash）。一个字符
串列表会被传递给一个低级系统命令，比如 os.execve() 。如果你想让命令被一个
shell 执行，传递一个字符串参数，并设置参数 shell=True . 有时候你想要 Python 去
执行一个复杂的 shell 命令的时候这个就很有用了，比如管道流、 I/O 重定向和其他特
性。例如：
"""
out_bytes = subprocess.check_output('grep python | wc > out', shell=True)
"""
需要注意的是在 shell 中执行命令会存在一定的安全风险，特别是当参数来自于用
户输入时。这时候可以使用 shlex.quote() 函数来讲参数正确的用双引用引起来。

使用 check output() 函数是执行外部命令并获取其返回值的最简单方式。但是，
如果你需要对子进程做更复杂的交互，比如给它发送输入，你得采用另外一种方法。
这时候可直接使用 subprocess.Popen 类。
"""
import subprocess

text = b'''
hello world
this is a test
goodbye
'''
p = subprocess.Popen(['wc'],
                     stdout = subprocess.PIPE,
                     stdin = subprocess.PIPE)

stdout, stderr = p.communicate(text)
out = stdout.decode('utf-8')
err = stderr.decode('utf-8')
"""
ubprocess 模块对于依赖 TTY 的外部命令不合适用。例如，你不能使用它来自
动化一个用户输入密码的任务（比如一个 ssh 会话）。这时候，你需要使用到第三方模
块了，比如基于著名的 expect 家族的工具（pexpect 或类似的）
"""
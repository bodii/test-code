#!/usr/bin/env python
# -*- coding:utf-8 -*-

"""
Topic: 第十三章：脚本编程与系统管理
Description: 许多人使用 Python 作为一个 shell 脚本的替代，用来实现常用系统任务的自动化，
如文件的操作，系统的配置等。本章的主要目标是描述光宇编写脚本时候经常遇到的
一些功能。例如，解析命令行选项、获取有用的系统配置数据等等。第 5 章也包含了
与文件和目录相关的一般信息。

Title:           通过文件名查找文件
Issue:你需要写一个涉及到文件查找操作的脚本，比如对日志归档文件的重命名工具，你
不想在 Python 脚本中调用 shell，或者你要实现一些 shell 不能做的功能。
Answer: 查找文件，可使用 os.walk() 函数，传一个顶级目录名给它。
"""
import os
import sys
def findfile(start, name):
    for relpath, dirs, files in os.walk(start):
        if name in files:
            full_path = os.path.join(start, relpath, name)
            print(os.path.normpath(os.path.abspath(full_path)))

# if __name__ == '__main__':
#     findfile(sys.argv[1], sys.argv[2])

"""
保存脚本为文件 findfile.py，然后在命令行中执行它。指定初始查找目录以及名字
作为位置参数，

os.walk() 方法为我们遍历目录树，每次进入一个目录，它会返回一个三元组，包
含相对于查找目录的相对路径，一个该目录下的目录名列表，以及那个目录下面的文
件名列表。
对于每个元组，只需检测一下目标文件名是否在文件列表中。如果是就使用
os.path.join() 合并路径。为了避免奇怪的路径名比如 ././foo//bar ，使用了另
外两个函数来修正结果。第一个是 os.path.abspath() , 它接受一个路径，可能是相
对路径，最后返回绝对路径。第二个是 os.path.normpath() ，用来返回正常路径，可
以解决双斜杆、对目录的多重引用的问题等。
尽管这个脚本相对于 UNIX 平台上面的很多查找公交来讲要简单很多，它还有跨
平台的优势。并且，还能很轻松的加入其他的功能。我们再演示一个例子，下面的函
数打印所有最近被修改过的文件：
"""
import os
import time
def modified_within(top, seconds):
    now = time.time()
    for path, dirs, files in os.walk(top):
        for name in files:
            fullpath = os.path.join(path, name)
            if os.path.exists(fullpath):
                mtime = os.path.getmtime(fullpath)
                if mtime > (now - seconds):
                    print(fullpath)

if __name__ == '__main__':
    import sys
    if len(sys.argv) != 3 :
        print('USage: {} dir seconds'.format(sys.argv[0]))
        raise SystemExit(1)
    modified_within(sys.argv[1], float(sys.argv[2]))

"""
在此函数的基础之上，使用 os,os.path,glob 等类似模块，你就能实现更加复杂的操
作了。可参考 5.11 小节和 5.13 小节等相关章节。
"""
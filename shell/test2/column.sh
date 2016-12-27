#!/bin/bash

# 通过printf命令打印出表格的标题，而命令ls -l | sed 1d输出当前目录下的
# 文件列表，其中ls命令的输出经过管道传递给sed 1d命令进行处理，目的是删
# 除目录列表中的第一行，也就是 total ** 字符串，因为我们不希望它出现在
# 表格中。可以看到，两条命令之间通过分号分隔，然后通过括号（）把这两条
# 命令组合在一起，这样两条命令就会在运行在当前shell的子shell中，并且这
# 两部分输出的数据会被一起传送到column命令的标准输入中进行处理。指定选
# 项-t以后， column命令会自动对输入数据进行格式化，并判断列的个数，从
# 而创建一个表格并输出到标准输出。
(printf "PERMISSIONS LINKS OWNER GROUP SIZE NN DD HH:MM FILE_NAME\n";
ls -l | sed 1d) | column -t

exit 0

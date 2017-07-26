#!/bin/bash

# 首先对文件/etc/passwd进行排序，然后添加行号
# 再对标准输入进行复制，一个输出到标准输出
# 另一个输出到文件sort.out中

sort /etc/passwd | cat -n | tee sort.out

exit 0

#!/bin/bash

echo 
echo "Block output is redirected to file 'block1.out;"

# 重定向若干个命令序列的标准输出
# 显示子shell的当前工作目录
( date; cd /etc; echo -n "Current Working dir: ";pwd; ) > block_subshell.out

# 打印父shell的当前工作目录
echo "Current Working dir:$PWD"
echo 

exit 0

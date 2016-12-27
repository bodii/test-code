#!/bin/bash

echo 
echo "Block output is redirected to file 'block_current.out'."

# 把多个命令的输出一起重定向到文件中
{ date; cd /etc; echo -n "Current Working dir: "; pwd; } > block_current.out

# 打印当前工作目录
echo "Current Working dir: $PWD"
echo 


exit 0


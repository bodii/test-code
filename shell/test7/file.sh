#!/bin/bash

echo 
echo " Check file type using 'file' command:"


# 显示内核文件的信息
echo

echo " Kernel file:"
echo -n "   "
file /boot/vmlinuz-`uname -r`

# 二进制程序/bin/ls
echo
echo " Binary program 'ls':"
echo -n "   "
file /bin/ls

# 主目录
echo
echo "  Home directory:"
echo -n "   "
file ~

# 符号链接文件
echo
echo "  Sybolic link file:"
echo -n "   "
file /usr/bin/locate

# gzip压缩的数据
echo 
echo "  gzip compressed data:"
echo -n "    "
file /usr/share/man/man1/ls.1.gz

# shell脚本文件
echo
echo "  Shell script:"
echo -n "    "
file ./file.sh
echo 

exit 0

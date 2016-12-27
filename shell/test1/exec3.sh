#!/bin/bash

echo "文件重定向--“文件描述符”练习"
echo " “<>” 文件可读写 重定向"

echo "创始文件描述符 是一个大于2的整数（0, 1, 2 已经被使用了)"
exec 3<>exec3.out

echo "这是第一句测试的话" >&3
echo "关闭文件描述符3"
exec 3>&-

exit 0



#!/bin/bash

# 检测脚本被执行时，命令行的参数个数
if [ $# -lt 2 ]
then
	#参数个数小于2, 脚本退出执行
	echo "Error!Arguments are not enough."
	echo "Usage:$0 operator file."
	exit 1 

elif [ $# -gt 2 ]
then
	# 参数个数大于2，脚本退出执行
	echo "Error!You sepcify Too many arguments."
	echo "Usage:$0 operator file."
	exit 1

fi

# 执行到这里，说明正确地的指定了2 个命令行参数
echo "My shell script begins running."

exit 0

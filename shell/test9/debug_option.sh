#!/bin/bash

# 定义debug 函数用来打开debug模式
debug () {
	# 当脚本内定义的变量DEBUG=true时，debug模式将被打开
	if [ "$DEBUG" = "true" ]
	then
		if [ "$1" = "on" -o "$1" = "ON" ] 
		then
			set -x
		elif [ "$1" = "off" -o "$1" = "OFF" ]
		then
			set +x
		fi
	fi
}

# 检查命令行选项"-debug"
if [ "$1" = "-debug" -o "$1" = "-d" ]
then
	# 为变量DEBUG赋值打开debug
	DEBUG="true"
fi

# 打开debug trace模式
debug on
echo "do stuff..."

# 关闭debug trace模式
debug off


exit 0

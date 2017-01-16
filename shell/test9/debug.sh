#!/bin/bash

# debug 调试函数，用于调试shell程序代码或代码块
# export DEBUG=true 打开调试模式
# unset DEBUG 关闭调试模式
# 使用方式 DEBUG on 或 DEBUG off 
debug() {
	if [ "$DEBUG" = "true" ]
	then
		if [ "$1" = "on" -o "$1" = "ON" ]
		then
			# 打开 debug trace
			set -x
		elif [ "$1" = "off" -o "$1" = "OFF" ]
		then
			# 关闭 debug trace
			set +x
		fi
	fi
}

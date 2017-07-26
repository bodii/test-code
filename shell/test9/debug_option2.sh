#!/bin/bash

# 定义debug函数，它可以执行任传递给它的命令
DEBUG () {
	if [ "$DEBUG" = "true" ] 
	then
		# 把接收到的参数当做命令执行
		$@
	fi
}

TODAY="MON"

# 调用debug 函数并传递给它要执行的命令
DEBUG echo "TOGAY=$TODAY"

if [ "$TODAY" = "MON" ]
then
	TOMORROW="TUE"
else
	TOMORROW="FRI"
fi

# 调用debug函数并传递给它要执行的命令
DEBUG echo "TOMORROW=$TOMORROW"
YESTERDAY="SUN"

# 调用debug函数并传递给它要执行的命令
DEBUG echo "YESTERDAY=$YESTERDAY

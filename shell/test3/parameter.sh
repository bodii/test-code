#!/bin/bash

if [ -n "$1" ]
then
	echo 
	# 显示脚本文件名及脚本如何被调用
	echo -e "Script Name:\t\t\t`basename $0`"
	echo -e "How Scritp is invoked:\t\t$0 $1"

	# 显示指定的参数
	echo -e "The first parameter:\t\t$1"

	if [ -f "$1" ]
	then
		echo
		echo "The content of file '$1':"
		echo "---------------------------------"
		echo "`cat $1`"
		echo "---------------------------------"
	else

		echo "File $1 don't exist."
		exit 1
	fi

else

	# 如果没有指定参数，显示帮助信息
	echo 
	echo "Show the content of a file."
	echo 
	echo "Usage:"
	echo "`basename $0` [filename]"
fi

echo 
exit 0
	


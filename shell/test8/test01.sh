#!/bin/bash

if [ "$#" -eq 0 ]
then
	echo "请使用 `basename $0` [FILE] 的格式载入一个文件"
	exit 1
fi

if [ -f "$1" ]
then
	echo "$1:这是一个文件"

	if [ -s "$1" ] 
	then
		echo "$1:这不是一个空文件"
	else
		echo "$1:这是一个空文件"
	fi
else
	echo "$1:这不是一个文件"
fi

echo 
exit 0

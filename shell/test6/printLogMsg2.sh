#!/bin/bash

# 定义pringMsg函数
printMsg ()
{
	if [ $# -lt 2 ]
	then
		echo "printMsg() needs 2 or more arguments..."

		return 1
	fi

	PREFIX=$1

	shift

	echo "$PREFIX: $@"
}


logMsg ()
{
	if [ -n "$1" ]
	then
		echo "`date` $@" >> arguments2.log
	else
		echo "logMsg() needs some arguments"
		return 1
	fi
}

printMsg "from printMsg()" "it works..." && logMsg "From logMsg()" "it works..."

#!/bin/bash

# 定义函数printMsg用来输出信息到终端屏幕，该函数至少需要两个参数
printMsg ()
{
	# 如果参数个数小于两个，打印错误信息并且退出状态为1
	if [ $# -lt 2 ]
	then
		echo "printMsg() needs 2 or more arguments..."

		# 使用return语句退出函数的执行，并指定退出状态
		return 1
	fi

	# 保存第一个参数
	PREFIX="$1"

	# 在参数列表中去除第一个参数
	shift

	# 以参数1、参数2、参数3、参数4的形式打印出参数列表
	echo "$PREFIX: $@"

	# 函数执行到这里会以最后一个命令的退出状态退出
}

# 定义函数logMsg用来记录信息到log文件,该函数至少可以接收一个参数
logMsg()
{
	# 判断是否指定了参数的另一种方式，查看第一个参数是否为空
	if [ -n "$1" ]
	then
		# 如果参数不为空，把日期和参数列表输出到log文件中
		echo "`date` $@ " >> arguments.log
	else
		# 如果参数为空，输出错误信息并让函数以退出状态1返回
		echo "logMsg() needs some arguments"
		return 1
	fi
}

# 通过两个参数调用函数printMsg,如果函数以非零值退出，才会调用logMsg函数
# 由于指定了足够的参数，printMsg函数会返回成功值0,所认logMsg不会被调用
printMsg "From printMsg()" "it works..." && logMsg "From logMsg()" "it works..."


#!/bin/bash

# 定义函数pidof_process 用来返回一个程序名对应的PID
pidof_process () 
{
	# 该函数在调用时需要指定一个参数，即所查询的程序名
	if [ $# -eq 0 ]
	then
		echo "Error, function pidof_process() need on argument!"
		exit 1
	fi


	# 通过命令pidof得到程序名对应的PID
	# 并把多个PID值之间的空格符替换为换行符\n,从而获得一个PID列表
	# 最终获得的PID列表会被保存到全局变量PID中
	PID=`pidof "$1" | tr ' ' '\n'`

	if [ -z "$PID" ]
	then
		# 如果pidof命令的输出为空，则说明参数S1中保存的程序名没有被执行
		PID="$1 is not running!"

		# 退出状态1表示程序没有被执行
		return 1
	else
		# 以状态0退出，表示成功获得了进程的PID
		return 0
	fi
}

echo

# 以程序名apache2作为参数调用函数pidof_process,获得apache服务器的所有进程
# PID但是并没有把结果保存在一个变量中

pidof_process 'apache2'

# 检查函数的退出状态
if [ $? -eq 0 ]
then
	echo -e "Pid(s) of process \"apache2\" : \n $PID"
	echo
else
	echo -e "Pid(s) of process \"apache2\" :Not found\n"
	echo
fi

pidof_process 'firefox'

if [ $? -eq 0 ]
then
	echo -e "Pid(s) of process \"firefox\" : \n $PID"
	echo
else
	echo -e "Pid(s) of process \"firefox\" :Not found\n"
	echo
fi

pidof_process 'fcitx'

if [ $? -eq 0 ]
then
	echo -e "Pid(s) of process \"fictx\" : \n $PID"
	echo
else
	echo -e "Pid(s) of process \"fictx\" : \n $PID"
	echo
fi

exit 0

#!/bin/bash

# 定义函数pidof_process 用来返回一个程序名对应的PID
pidof_process () 
{
	# 该函数在调用时需要指定一个参数，即所查询的程序名
	if [ $# -ne 1 ]
	then
		# 如果在调用时参数个数不为1,则显示错误信息并以状态1退出
		echo "Error,Function pidof_process() need one argument!"
		exit 1
	fi


	# 通过命令pidof得到程序名对应的PID
	# 并把多个PID值之间的空格符替换为换行符\n,从而获得一个PID列表
	PID=`pidof "$1" | tr ' ' '\n'`

	if [ -z "$PID" ]
	then
		# 如果pidof命令的输出为空，说明参数$1中保存的程序名没有被执行
		echo " $1 is not running!"
		
		# 退出状态1表示程序没有被执行
		return 1
	else
		# 如果程序正在运行，则打印出PID列表
		echo "$PID"

		# 并以状态0退出，表示成功获得了进程的PID
		return 0
	fi
}

echo

# 以程序名apache2 作为参数调用函数[pidof_process, 获得apache服务器的所
# 有进程PID

APACHE_PID=`pidof_process "apache2"`

# 检查函数退出状态
if [ $? -eq 0 ]
then
	# 如果函数pidof_process 以状态0退出，说明成功获得了PID
	# 打印apache服务器的PID列表
	echo -e "Pid(s) of process \"apache\": \n $APACHE_IPD"
	echo
else
	# 如果函数pidof_process以状态0退出，说明程序没有运行
	echo -e "Pid(s) of process \"apache\": Not found\n"
	echo

fi

# 查询浏览器firefox的进程PID
FIREFO_PID=`pidof_process "firefox"`

# 检查函数的退出状态
if [ $? -eq 0 ]
then
	echo -e "pid(s) of process \"firefox\": \n $FIREFOX_PID"
	echo
else
	echo -e "pid(s) of process \"firefox\": Not fount!\n"
	echo
fi

# 查询后台进程fcitx的PID
FCITX_PID=`pidof_process "fcitx"`

if [ $? -eq 0 ]
then
	echo -e "Pid(s) of process \"fcitx\": \n $FCITX_PID"
	echo
else
	echo -e "Pid(s) of process \"fcitx\": Not found!\n"
	echo
fi


exit 0

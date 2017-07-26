#!/bin/bash

# 定义函数EXpire_Handler用来捕捉SIGALARM 信号
Expire_Handler()
{
	echo 
	echo "Got SIGALARM signal,Waiting for your info too long!"
	echo "Bye."

	# 从脚本中以退出代码14退出，用来表示接收到了SIGALRM信号
	exit 14

}

# 定义函数，用来设置一个定时器
Start_Timer()
{
	# 如果没有指定参数，默认为10秒
	local INTERVAL=${1:-10}

	# 检查参数大于0
	if [ $INTERVAL -gt 0 ]
	then
		# 15秒后，发送信号SIGALRM到脚本进程本身
		sleep $INTERVAL && kill -s 14 $$ &

		# 记住后台进程的PID，用来杀死定时器
		# 如果用户在最后期限以前输入信息
		TIMERPID=$!
	else
		echo "Error:Interval must be positive integer!"
		exit 1
	fi
}

# 定义函数，用来杀死后台进程，从而去除定时器
Unset_Timer()
{
	# 首先杀死子进程
	kill `pgrep -P $TIMERPID`

	# 然后再杀死父进程
	kill $TIMERPID
}

# 设置定时器回调函数Expire_Handler
trap Expire_Handler 14

echo 
echo "You have only 15 seconds to enter your info!"
echo

# 把定时器设置为15秒的时间 
Start_Timer 15

read -p "Please Enter your ID:" id

# 如查用户在最后期限以前输入了信息，则移除定时器
Unset_Timer

echo
echo "Your ID is :$id"
echo "All Done."

exit 0

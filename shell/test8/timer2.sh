#!/bin/bash

# 定义退出函数Expire_Handler
Expire_Handler ()
{
	echo 
	echo "Got SIGALARM signal, Waiting for Your info too long!"
	echo "Bye."

	# 从脚本中以退出代码14退出，用来表示接收到了SIGALRM信号
	exit 14
}

# 定时器
Start_Time ()
{
	# 定时器的时间
	local INTERVAL=${1:-10}

	# 判断定时是否大于0
	if [ $INTERVAL -gt 0 ]
	then
		sleep $INTERVAL && kill -s 14 $$ &

		#记住后台进程的PID，用来杀死定时器
		TIMERPID=$!
	else
		echo "Error:Interval must be positive integer!"
		exit 1
	fi
}

# 定义杀死后台进程的函数
Unset_Timer ()
{
	# 首先杀死子进程
	kill `pgerp -P $TIMERPID`

	# 然后再杀死父进程
	kill $TIMEPID
}

# 定时器捕捉SIGALARM
trap Expire_Handler 14

echo 
echo "You have only 15 seconds to enter your info!"
echo

# 设置定时器
Start_Time 15

read -p "Please Enter your ID:" id

# 如果用户在最后期限以前输入了信息，则移除定时器
Unset_Timer

echo
echo "Your ID is: $id"
echo "All Done."

exit 0

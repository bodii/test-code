#!/bin/bash

# 定义要测试的网络
NETWORK=192.168.1

# 定义检测的起始地址
IP=1

# 只要IP地址小于130,循环就会被执行
while [ "$IP" -lt "130" ]
do
	echo -en "Pinging ${NETWORK}.${IP}..."

	# 变量NETWORK和IP组合在一起构成IP地址
	# 执行命令Ping检测IP地址的连通性，并且不显示任何输出
	ping -c1 -w1 ${NETWORK}.${IP} &> /dev/null

	# 根据命令ping的退出状态值是否等于0,判断网络是否连通
	if [ $? -eq 0 ]
	then
		# 如果ping 的退出状态等于0,打印OK
		echo "OK"
		exit 0
	else
		echo "Failed"
	fi

	# IP的数量加1
	let IP=$IP+1
done

exit 0


#!/bin/bash

# 定义检查IP地址的函数checkip
checkip ()
{
	# 对第一个参数进行匹配
	case $1 in
		"" | *[!0-9.]* | *[!0-9])
			return 1
			;;
	esac

	# 把内部分隔符（之前提到的分隔符和这里的内部分隔符是两个概念，
	# 内部分隔符Internal field Seperator 是指bash如何分隔Shell脚本
	# 中的各个字段，默认会把空格和TAB字符当做内部分隔符，可以通过
	# 修改内部变量IFS来改变内部分隔符)Internal Field Seperator 由
	# 空格修改为点
	# 目的是把位置变量$1中保存的地址拆散为若干个数字
	# 使用修饰符local标明这是一个本地变量，不会影响函数外的行为
	local IFS=.

	# 打散位置变量$1中的IP地址，并把结果保存在位置变量$1,$2,$3,...中
	set -- $1

	# 用逻辑与&&串联多个测试表达式，从而判断地址是否合法
	# 测试条件包括：
	# 1.参数被打散以后应该得到4 个数字
	# 2.被打散的4个数字的值都小于等于255
	# 如果任一位置变量的值为空，使用999作为它的默认值
	if [ $# -eq 4 ] && [ ${1:-999} -le 255 ] && [ ${2:-999} -le 255 ] &&
		[ ${3:-999} -le 255 ] && [ ${4:-999} -le 255 ]
	then
		# 在所有条件都满足的情况下，以状态0退出，表示合法
		return 0
	else
		return 1
	fi

}

# 通过一个for循环语句重复地调用checkip函数检测列表中的IP地址
for IP in 127.0.0.1 127.0.0.0.1 .0.0.1 ...1 172.299.38.1 172.38.1 \
	192.168.10.1 127.0.0.
do
	# 在每一次循环迭代中，都以变量IP作为参数调用函数checkip
	checkip ${IP}

	# 检查函数的退出状态，如果等于0说明是合法的IP地址
	if [ $? -eq 0 ]
	then
		# 格式化输出IP地址合法
		printf "%15s: %s\n" "${IP}" "valid ip address"
	else
		# 格式化输出IP地址非法
		printf "%15s: %s\n" "${IP}" "invalid ip address"
	fi
done

# 无限循环会不停地让用户输入IP地址进行检测
while [ true ]
do 
	echo -n "Please enter IP address you want ot check(type \"quit\" to \
finish):"
	# 读取用户输入的IP地址
	read IPADDRESS
	if [ "$IPADDRESS" == "quit" ]
	then
		# 如果用户输入的字符串为quit，则退出循环，脚本结束
		break;
	fi
	
	# 以IP地址为参数调用函数checkip进行检测
	checkip ${IPADDRESS}

	# 检查函数的退出状态，如果等于0说明是合法的IP地址
	if [ $? -eq 0 ]
	then
		# 格式化输出IP地址合法
		printf "%15s: %s\n" "${IPADDRESS}" "valid ip address"
	else
		# 格式化输出IP地址非法
		printf "%15s: %s\n" "${IPADDRESS}" "invalid ip address"
	fi
done

exit 0

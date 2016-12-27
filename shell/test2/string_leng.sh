#!/bin/bash

echo 
echo "This script will tell you the length of your input!"

# 使用空命令冒号构建无限循环
while :
do
	read -p "Please Enter a String(or quit):"
	case "$REPLY" in 
		[Qq] | [Qq][Uu][Ii][Tt])
			echo " Bye."
			# 在输入大、小写quit时，退出
			exit 0
			;;
		*)
			# 使用expr命令计算字符串的长度
			LEN=$(expr length "$REPLY")
			echo "Length:$LEN"
			;;
	esac
done

exit 0

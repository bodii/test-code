#!/bin/bash

echo
echo "This script will tell you the length of you input!"

# 使用空命令冒号构建无限循环
while :
do
	read -p "Please Enter a String(or quit):"
	case "$REPLY" in 
		[Qq] | [Qq][Uu][Ii][Tt])
			echo "Bye."
			# 如果输入的是大、小写的quit，则退出
			exit 0
			;;
		*)
			# 统计用户输入的字符串长度
			LEN=$(expr length "$REPLY")
			echo "length: $LEN"
			;;
	esac
done

exit 0


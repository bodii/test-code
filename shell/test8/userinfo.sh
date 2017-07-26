#!/bin/bash

echo
echo "This script will find user's home directory"

read -p "Enter the user name(or quit):"

until [ "$REPLY" = "quit" ] || [ "$REPLY" = "Quit" ] || [ "$REPLY" = "q" ]
do
	echo
	awk '
		BEGIN {
			# 修改分隔符为冒号
			FS=":";
		}
		# 当第一个字段和awk变量USERNAME 相同时
		$1 == USERNAME {
			# 打印主目录
			printf ("User: %-8s UID: %-5s GID: %-5s Home Dir: %-15s shell:%-10s\n",$1,$3,$4,$6,$7);
		}

		# 定义awk 变量USERNAME
		' USERNAME="$REPLY" /etc/passwd
		
		echo
		read -p "Enter the user name(or quit):"
	done

	echo "Bye."
	exit 0

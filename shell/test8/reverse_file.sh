#!/bin/bash

echo
read -p "Enter the file you want to reverse(or quit):"

until [ "$REPLY" = "quit" ] || [ "$REPLY" = "Quit" ] || [ "$REPLY" = "q" ]
do
	# 文件必须是可读的普通文件，并且其大小大于0
	if [ -f "$REPLY" ] && [ -r "$REPLY" ] && [ -s "$REPLY" ]
	then
		awk '{
			# 使用for循环对一行中的单词进行翻转
			# 从每一行的最后一个字段$NF开始打印，一直到第一个字段$1
			for (i = NF;i>0;i--){
				printf "%s ",$i;
			}
			printf "\n";
		}' "$REPLY"
	else
		echo "File "$REPLY" can not be reserved,Retype another file!"
	fi

	echo
	read -p "Enter the file you want to reverse(or quit):"
done

echo "Bye."

exit 0

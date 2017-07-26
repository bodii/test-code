#!/bin/bash

echo
read -p "请输入一个用于反转内容的文件(or quit):"
echo "$REPLY"

until [ "$REPLY" = "quit" ] || [ "$REPLY" = "Quit" ] || [ "$REPLY" = "q" ]
do
	if [ -f "$REPLY" ] && [ -r "$REPLY" ] && [ -s "$REPLY" ]
	then
		awk '
			{
				i=NF;
				while(i>0){
					printf "%s ", $i;
					i--;
				}
				printf "\n"
			}
		' "$REPLY"
	else
		echo "你输入的内容不能用于反转内容的文件"
	fi

	echo
	read -p "请输入一个用于反转内容的文件(or quit):"
done

echo "Bye."

exit 0

#!/bin/bash

echo
read -p "请输入一个文件用于反转内容（或是quit退出):"
echo "你输入的是:$REPLY"

until [ "$REPLY" = "quit" ] || [ "$REPLY" = "Quit" ] || [ "$REPLY" = "q" ]
do
	if [ -f "$REPLY" ] && [ -r "$REPLY" ] && [ -s "$REPLY" ]
	then
		awk '
		$1 ~/^.*$/{
				i=NF;
				do{
					printf "%s ", $i;
					i--;
				}
				while(i>0)
				printf "\n";
			}' "$REPLY"
	else
		echo "这不是一个可操作反转的文件"
	fi

	echo
	read -p "请输入一个文件用于反转内容（或是quit退出):"
done

echo "Bye."

exit 0

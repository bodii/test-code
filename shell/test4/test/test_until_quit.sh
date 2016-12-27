#!/bin/bash

QUIT_COMMAND=quit

# 直到用户输入的字符串是quit时，until循环才会退出
until [ "$USER_INPUT" = "$QUIT_COMMAND" ]
do
	echo "Please input command:"
	echo "(type command $QUIT_COMMAND to exit)"
	read USER_INPUT
	echo
	echo ">>your command is $USER_INPUT"

	# 对用户的输入进行匹配
	case $USER_INPUT in
		"open")
			echo ">>opening..."
			;;
		"close")
			echo ">>closed."
			;;
		*)
			echo ">>Unknown command."
			;;
	esac

	echo
done

echo "Bye."
exit 0

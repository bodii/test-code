#!/bin/bash

recursive ()
{
	if [ $1 -eq 1 ]
	then
		echo 1

		return 0

	fi

	# 定义本地变量TEMP，其中保存了参数值减1后的值
	local TEMP=`expr $1 - 1`

	local PREVIOUS=`recursive $TEMP`

	# 把数值（$1 减 1） 的阶乘计算结果和参数值$1相乘
	# 特殊字符星号×需被引用，才能正确地传递给expr命令
	RESULT=`expr $1 \* $PREVIOUS`

	echo $RESULT

	return 0
}

echo 
echo -n " Enter a number(<20) to do n! operation(type "quit"to exit):"
read INPUT

# 当用户输入字符串quit时，unti循环结束，脚本退出
until [ "$INPUT" = "quit" ]
do
	case $INPUT in
		[1-9]|[1][0-9]|20)
			OUTCOME=`recursive "$INPUT"`
			echo -e "\tAfter computing,$INPUT!=$OUTCOME"
			echo
			;;
		*)
			echo
			echo -e "\tNot a valid number,Enter again!"
			;;
	esac

	echo -en "\tEnter a number(<20) to do n! operation(type "quit" to exit):"
	read INPUT
done
echo
echo "	Bye."

exit 0


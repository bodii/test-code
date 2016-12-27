#!/bin/bash


# 首先清除屏幕
clear

# 显示带单项，从而使用户可以选择期望的操作
echo -e "\t\t\tFile Operation List"
echo -e "\t\t\t----- ------------ -----"


echo "Choose one of the following operations:"
echo

echo "[O]pen File"
echo "[D]elete File"
echo "[R]ename File"
echo "[M]ove File"
echo "[C]opy File"
echo "[P]aste File"
echo

while :
do

	# 等待用户的输入
	echo -n "Please enter your operation:"
	read operation
	
	# 根据用户输入的操作执行相应的操作
	case "$operation" in
		# 同时接受大小写字母
	"O"|"o")
		# 用户选择了打开文件的操作
		echo
		echo "Opening File..."
		echo "Successed!"
		;;    # Note double semicolon to terminate each option.
	"D"|"d")
		# 用户选择了删除文件的操作
		echo
		echo "Deleteing File..."
		echo "Successed!"
		;;
	"R"|"r")
		# 用户选择了读取文件的操作
		echo
		echo "Reading File..."
		echo "Successed!"
		;;
	"M"|"m")
		# 用户选择了移动文件的操作
		echo
		echo "Moveing File..."
		echo "Successed!"
		;;
	"C"|"c")
		# 用户选择了复制文件的操作
		echo
		echo "Copying File..."
		echo "Successed!"
		;;
	"P"|"p")
		# 用户选择了粘贴文件的操作
		echo
		echo "Paste File..."
		echo "Successed!"
		;;
	[Qq] | [Qq][Uu][Ii][Tt])
		echo 
		echo "Bye."
		exit 0
		;;
	*)
		# 用户的输入不匹配上面的任何一种操作
		# 因此是未知操作，以错误状态1退出脚本
		echo
		echo "Error,Unknown operation."
		echo "Program exit!"
		exit 1
		;;
	esac
done

echo 

exit 0


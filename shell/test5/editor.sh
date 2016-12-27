#!/bin/bash

# 定义脚本使用函数
help()
{
	echo 
	echo "Usage:`basename $0` [OPTION] -f File "
	echo "Edit a Editable File using specified Editor and Pager."
	echo
	echo -e "  -e[Editor]\tEditor to edit the file: vim,pico,emacs,nano..."
	echo -e "  -p[Pager]\tPager to view the file:less,more,cat,head..."
	echo
	exit 1
}

# 关闭诊断信息
#OPTERR=0
#  处理命令行选项和参数
while getopts f:e:p: OPTION
do
	case "$OPTION" in
		f)
			#每一个选项的参数可以使用变量$OPTARG来访问
			TARGET_FILE="$OPTARG";;
		e)
			EDITOR="$OPTIARG";;
		p)
			PAGER="$OPTARG";;
		\?)
			# 当遇到未知选项时，显示脚本的使用方法
			help;;
	esac
done

# 文件参数必须被设置
if [ -z "$TARGET_FILE" ]
then
	help
fi
echo

# 确保文件是可读、可写的
if [ -r "$TARGET_FILE" -a -w "$TARGET_FILE" ]
then
	# 询问用户是否想编辑文件
	read -p "Do you want to editing file '$TARGET_FILE'(y/n)?"
	if [ "$REPLY" = 'y' -o "$REPLY" = 'Y' -o "$REPLY" = 'Yes' -o "$REPLY" = 'yes' ]
	then
		echo
		echo "Opening file '$TARGET_FILE' using ${EDITOR:-vim}..."
		sleep 3

		# 使用编辑器打开文件，如果变量EDITOR为空使用vim
		${EDITOR:-vim} "$TARGET_FILE"
		echo "Thank you for editing the file."
	fi

	echo
	# 询问用户是否想查看文件内容
	read -p "Do you want to see the file just edited (y/n)?"
	if [ "$REPLY" = 'y' -o "$REPLY" = 'Y' -o "$REPLY" = 'Yes' -o "$REPLY" = 'yes' ]
	then
		echo
		echo "Opening the file using '${PAGER:-less}'..."
		sleep 3

		# 显示文件内容，如果变量PAGER为空就使用less
		${PAGER:-less} "$TARGET_FILE"
	fi

	echo
	echo "Bye!"
	echo
else
	echo
	echo "File $TARGET_FILE can not be accessed,Please specify a Readable and Writable file."
	echo
	echo "Bye!"
	echo
	exit 1
fi

exit 0

#!/bin/bash

# 定义函数count,用来统计一个文件中的字数
count(){
	# 函数需要一个参数才可以被正确调用
	if [ $# != 1 ]
	then
		echo "Need on file parameter to  work!"
		exit 1;
	fi

	# 删除标点符号和特殊字符
	# 构建一个很长的管道命令，每一段都单独写在一行中，增加可读性
	tr '+-=*.,;:{}()#!?<>"\n\t' ' ' <$1 |\
	
	#把所有大写字母转换为小写字母
	tr 'A-Z' 'a-z' |\
	
	# 把连续重复的空格符替换为一个空格符
	tr -s ' ' |\
	
	# 把空格符转换为换行符
	tr ' ' '\n' |\

	# 把相同的单词放到一起
	sort |\
	
	# 删除重复单词，并进行统计
	uniq -c |\
	
	# 根据重复的次数进行排序
	sort -rn
}

echo 
echo "This script can count words of a specified file."

# 使用空命令冒号构建无限循环
while :
do
	read -p "Enter the file path(or quit):"
	case "$REPLY" in 
		[Qq] | [Qq][Uu][Ii][Tt])
			echo "Bye."
			# 在输入大、小写quit时，退出
			exit 0
			;;
		*)
			# 判断输入的是一个可读的普通文件，并且内容不为空
			if [ -f "$REPLY" ] #&& [ -r "$PEPLY" ] #&& [ -s "$PEPLY" ]
			then
				# 当用户输入一个合法文件时
				# 调用count函数统计文件的单词个数
				count "$REPLY"
			else
				# 如果输入了非法文件， 显示不能处理它
				echo "$REPLY can not be dealed with."
			fi
			;;
	esac
done

exit 0

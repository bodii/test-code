#!/bin/bash

echo
echo "Which directory do you want to list:"
echo "(Press\"Enter\" directory show menu again)"
echo

# 把要显示的菜单项写在in后面的列表里
select dir in /bin/ /usr/bin/ /usr/local/bin/ /sbin/ /usr/sbin/ /usr/local/sbin/ /usr/share/ quit
do
	if [ ! -z "$dir" ]
	then
		if [ "$dir" = "quit" ]
		then
			# 如果用户选择quit，则退出脚本
			echo
			echo "You entered quit. Byebye!"
			break;
		else
			# 用户选择了某一个目录
			echo
			echo "You selected directory \"$dir\", it contains following fles:"
			echo

			# 显示所选目录中的内容
			ls -l $dir
		fi
	else
		echo
		echo "Error, invalid selection \"$PERLY\", chose again!"
	fi
	

	echo 
	echo "Which directory do you want to list:"
	echo "(Press \"Enter\" directly to show menu again)"
done

exit 0

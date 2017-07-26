#!/bin/bash 

echo
echo "Which directory do you want to list:"
echo "(Press \"Enter\" directory to show menu again)"
echo

# 把要显示的菜单项写在in 后面的列表里
select dir in /bin/ /user/bin/ /usr/local/bin/ /sbin/ /usr/sbin/ /usr/local/sbin/ /usr/share/ quit
do
	case $dir in
		quit)
			# 如果用户选择quit,则退出脚本
			echo
			echo "You entered quit, Byebye!"
			break
			;;
		/*)
			# 匹配所选择目录，$dir以/开始说明用户选择了某个目录
			echo
			echo "You selected directory \"$dir\", it contains following files."
			echo


			# 显示所选目录中的内容
			ls -l $dir

			;;
		*)
			#所有其他的输入都为非法
			echo
			echo "Error invalid selection \"$REPLY\" chose again!"
			;;
	esac

	echo 
	echo "Which directory do you want to list:"
	echo "(Press \"Enter\" directly to show menu again)"
done

exit 0


#!/bin/bash

# 检查是否指定了命令行选项--export
if [ "$1" = "--export" ]
then

	# 在父进程中定义全局变量
	export LOCATION=USA

elif [ "$1" = "--no-export" ]
then

	# 在父进程中定义一个本地变量
	LOCATION=USA
else

	# 显示本脚要的使用方法
	echo
	echo -e "`basename $0` --export\texport parent process's variable to child process"
	echo -e "`basename $0` --no-export\tdon't export parent process's variable to child process"
	echo

	exit 0

fi


echo
echo "Your parent is at '$LOCATION'."

#创建子进程，子进程是否能够得到LOCATION变量的拷贝依赖于export关键字
./child.sh


# 子进程退出，父进程继续运行
echo
echo -e "Your parent is at '$LOCATION'."
echo -e "Child process and parent process have different LOCATION."
echo

exit 0


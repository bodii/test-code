#!/bin/bash

# 定义debug 函数
# exprot DEBUG=true 定义环境变量
# unset DEBUG 册除环境变量
debug(){
	# 环境变量DEBUG=ture 时将打开debug模式
	if [ "$DEBUG" = "true" ]
	then
		if [ "$1" = "off" -o "$1" = "OFF" ]
		then
			# 关闭debug trace模式
			set +x
		elif [ "$1" = "on" -o "$1" = "ON" ]
		then
			# 打开debug trace 模式
			set -x
		fi
	fi
}

# 定义变量保存profile文件
PROFILE_FILE=

judge_profile(){
	# 打开 debug
	debug on

	if [ -e ~/.bash_profile -a -r ~/.bash_profile ]
	then
		PROFILE_FILE="~/.bash_profile"
	elif [ -e ~/.bash_login -a -r ~/.bash_login ]
	then
		PROFILE_FILE="~/.bash_login"
	elif [ -e ~/.profile -a -r ~/.profile ]
	then
		PROFILE_FILE="~/.profile"
	fi

	# 关闭 debug
	debug off

}

# 调用函数judge_profile
judge_profile

# unset 变量PROFILE_FILE来模拟所有这些文件都不存在的情况
# unset PROFILE_FILE

# 打开debug trace 模式
debug on

if [ $PROFILE_FILE != "" ]
then
	echo "Use profile file: $PROFILE_FILE"
else
	echo "No prifile files exist."
fi

# 关闭debug trace 模式
debug off

exit 0

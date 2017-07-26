#!/bin/bash

# 引入debug 调试函数
. debug.sh

PROFILE_FILE=
# 定义judge_profile函数
judge_profile(){
	
	# 打开debug trace 调试
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

	# 关闭debug trace 调试
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

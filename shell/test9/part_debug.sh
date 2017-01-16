#!/bin/bash

# 定义变量保存profile文件
PROFILE_FILE=

judeg_profile(){
	# 嵌套的if语句
	if [ -e ~/.bash_profile -a -r ~/.bash_profile ]
	then
		PROFILE_FILE="~/.bash_profile"
	elif [ -e ~/.bash_login -a -r ~/.bash_profile ]
	then
		PROFILE_FILE="~/.bash_login"
	elif [ -e ~/.profile -a -r ~/.profile ]
	then
		PROFILE_FILE= "~/.profile"
	fi

}


# 调用函数judge_profile()
set -x
judeg_profile
set +x

# 打开trace debug模式
set -x

# unset 变量PRO_FILE_FILE来模拟所有这些文件都不存在的情况
unset PROFILE_FILE

if [ $PROFILE_FILE != "" ]
then
	echo "Use profile file: $PROFILE_FILE"
else
	echo "No prifile files exist."
fi

# 关闭trace debug 模式
set +x

exit 0

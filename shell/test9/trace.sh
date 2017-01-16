#!/bin/bash -x

# 定义一个空变量
PROFILE_FILE=

# nested if clause 
if [ -e ~/.bash_profile -a -r ~/.bash_profile ]
then
	PROFILE_FILE="~/.bash_profile"
elif [ -e ~/.bash_login -a -r ~/.bash_login ]
then
	PROFILE_FILE="~/.bash_login"
elif [ -e ~/.profile -a -r ~/.profile ]
then
	PROFILE_FILE="./.profile"
fi

# unset 这个变量用来模拟所有以上文件都不存在
unset PROFILE_FILE
if [ $PROFILE_FILE != "" ]
then
	echo "Use profile file: $PROFILE_FILE"
else
	echo "No prifile files exist."
fi

exit 0

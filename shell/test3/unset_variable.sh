#!/bin/bash

# 使用unset命令删除一个定义过的变量

echo
read -p "Please input your name:" NAME

echo -e "Welcome,$NAME!\n Your name is saved in variable 'NAME'!"

echo "-------------------------"
echo 
echo "Clear variable 'NAME' using 'unset' command."

# 清除变量的值
unset NAME

echo 
echo "Now variable 'Name' is: "

# 判断变量是否被清除
if [ -z $NAME ]
then
	echo
	echo "NAME:$NAME"
	echo "Variable NAME is null."
else
	echo
	echo "NAME:$NAME"
	echo "Variable Name is not null."
fi

echo 
echo ' quit'
echo


exit 0


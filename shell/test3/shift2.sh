#!/bin/bash

# 如果没有指定文件
# 使用if[ -z $1 ](指定的参数文件为空)
if [ -z $1 ]
then
	echo 
	echo "Usage:`basename $0` [filename] ..."
	echo "At least specify one filename to remove execute permisson
for group and others."
	echo 
	exit 0
fi
echo 

# 检查是否有其他的参数文件
until [ -z $1 ]
do
	echo "Processing file '$1'"
	chmod go-x $1

	# 删除第一个参数
	shift

done

echo "Done."
echo

exit

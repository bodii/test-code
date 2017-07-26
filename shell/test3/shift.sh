#!/bin/bash

# 如果没有指定文件
# 使用 if[ -z $1 ]也可以过到同样的效果
if [ $# -eq 0 ]
then
	echo 
	echo "Usage:`basename $0` [filename]..."
	echo "At least specify one filename to remove execute permisson for
	group and others."
	echo 
	exit 0
fi

echo

# 检查是否有其他的参数
# 使用until [ -z $1 ] 可以达到同样的效果
while [ $# -gt 0 ]
do
	echo "Processing file '$1'."
	chmod go-x $1

	# 删除第一个参数
	shift
done

echo "Done."
echo 

exit 0

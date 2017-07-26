#!/bin/bash

# 如果没有指定文件
# 使用if [ -z $1 ]也可以达到同样效果
if [ $# -eq 0 ]
then
	echo 
	echo "Usage:`basename $0` [filename]..."
	echo "At least specify one filename to remove execute
permisson for group and others."
	echo
	exit 1
fi

echo 
echo "You specified the following parameters:"
echo " << $* >>"
# = echo "<< $@ >>"
echo


# 遍历所有的参数
for PARAMETER in $@
# = for PARAMETER in $*
# = for PARAMETER in "$@"
do
	echo "Processing file '$PARAMETER'."
	# 删除权限
	#chmod go-x $PARAMETER
done

echo "Done."
echo

exit 0
